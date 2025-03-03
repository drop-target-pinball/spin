local check = require("check")
local std = require("std")

local queue = {}


local function to_vars_table(tbl, params)
    local meta = {}
    local label = check.nv("label", params.label, "string")
    local namespace = check.nv("namespace", params.namespace, "string")
    local conf = check.nv("conf", params.conf, "table")
    local raw = check.nv("raw", params.raw, "table")
    local player = params.player

    function meta.__index(_, name)
        local val = raw[name]
        if val == nil then
            error("no such " .. label .. ": " .. name)
        end
        return val
    end

    function meta.__newindex(_, name, val)
        local def = conf[name]
        if def == nil then
            error("no such " .. label .. ": " .. name)
        end
        local u_val = nil
        if def.kind.int ~= nil then
            u_val = { int = math.tointeger(val) }
        elseif def.kind.float ~= nil then
            u_val = { float = tonumber(val) }
        elseif def.kind.string ~= nil then
            u_val = { string = tostring(val) }
        elseif def.kind.bool ~= nil then
            if type(val) ~= "boolean" then
                error("not a boolean: " .. name .. " = " ..val)
            end
            u_val = { bool = val }
        else
            error("unexpected type: " .. name .. " = " .. val)
        end

        local ns_val = {}
        if namespace == "player" then
            ns_val = { player = player }
        else
            ns_val = namespace
        end
        local msg = {
            set = {
                namespace = ns_val,
                name = name,
                value = u_val,
            }
        }
        table.insert(queue, msg)
    end

    function meta.__pairs(_)
        return next, raw, nil
    end

    setmetatable(tbl, meta)
end

local pub = {
    conf = {},
    gfx = require("_render").gfx,
    elapsed = 0,
    raw_vars = {},
    raw_settings = {},
    raw_players = {},
    vars = {},
    settings = {},
    players = {},
}

local script_defs = {}
local scripts = {}
local alive = {}

-- Colors
pub.BLACK       = { r = 0,   g = 0,   b = 0,   a = 255 }
pub.CLEAR       = { r = 0,   g = 0,   b = 0,   a = 0   }
pub.FULL        = { r = 255, g = 255, b = 255, a = 255 }
pub.OFF         = pub.BLACK
pub.ON          = pub.FULL

-------------------------------------------------------------------------------
local function halt()
    alive = {}
end

local function reset()
    alive = {}
end

function pub._init()
    for name, def in pairs(pub.conf.scripts) do
        local mod = require(def.module)
        if type(mod) ~= "table" then
            error("module '" .. def.module .. "' did not return a table")
        end
        local script = mod[name]
        if script == nil then
            error("script '" .. name .. "' not found in module '" .. def.module .. "'")
        end
        script_defs[name] = def
        scripts[name] = mod[name]
    end

   to_vars_table(pub.vars, {
        label="var",
        namespace="var",
        conf=pub.conf.vars,
        raw=pub.raw_vars
    })
    to_vars_table(pub.settings, {
        label="setting",
        namespace="setting",
        conf=pub.conf.settings,
        raw=pub.raw_settings,
    })
    for i=1,pub.conf.max_players do
        pub.raw_players[i] = {}
        pub.players[i] = {}
        to_vars_table(pub.players[i], {
            label="player var",
            namespace="player",
            player=i,
            conf=pub.conf.player,
            raw=pub.raw_players[i],
        })
    end
    return true
end

local function kill(name)
    if alive[name] == nil then
        return
    end
    table.insert(queue, { script_killed = {name = name} })
    alive[name] = nil
end

local function kill_group(group)
    for name, def in pairs(pub.conf.scripts) do
        if def.group == group and alive[name] ~= nil then
            kill(name)
        end
    end
    for name, def in pairs(pub.conf.run_groups) do
        if group == def.parent then
            kill_group(name)
        end
    end
end

local function run(name)
    local script = scripts[name]
    if script == nil then
        error("no such script: " .. name)
    end

    -- See if this script, when run, replaces all scripts in the group
    local this_def = script_defs[name]
    if this_def.replace and this_def.group ~= nil then
        kill_group(this_def.group)
    end

    -- Create the coroutine and place it in the alive table. Set the wait
    -- condition to ready so that it will execute on the next tick
    local co = coroutine.create(script)
    alive[name] = {
        co = co,
        can_resume = pub.ready
    }
end

local function service_coroutines(kind, msg)
    for name, script in pairs(alive) do
        if coroutine.status(script.co) == "dead" then
            table.insert(queue, { script_ended = {
                name = name
            }})
            alive[name] = nil
        else
            local yes, r_kind, r_msg = script.can_resume(kind, msg)
            if yes then
                local running, result = coroutine.resume(script.co, r_kind, r_msg)
                if not running and result ~= nil then
                    error(debug.traceback(script.co, "in script '" .. name .. "': " .. result))
                end
                if running then
                    script.can_resume = result
                end
            end
        end
    end
end

function pub.post(elapsed, msg)
    pub.elapsed = elapsed
    local kind = ""
    local body = nil
    if type(msg) == "string" then
        kind = msg
        body = {}
    else
        for key, value in pairs(msg) do
            if kind ~= "" then
                error("table should only have one entry")
            end
            kind = key
            body = value
        end
    end

    if kind == 'halt' then
        halt()
    elseif kind == 'reset' then
        reset()
    elseif kind == 'kill' then
        kill(body.name)
    elseif kind == 'kill_group' then
        kill_group(body.name)
    elseif kind == 'run' then
        run(body.name)
    end

    service_coroutines(kind, body)

    if next(queue) == nil then
        return nil
    else
        local ret = queue
        queue = {}
        return ret
    end
end

-------------------------------------------------------------------------------
local function copy_opts(src, dest, ...)
    local arg = {...}
    if src == nil then
        return
    end
    if arg == nil then
        error("field names to copy are required")
    end
    for i, name in ipairs(arg) do
        if src[name] ~= nil then
            dest[name] = src[name]
        end
    end
end

function pub.video(name)
    check.nv("name", name)
    local v = spin.conf.video[name]
    if v == nil then
        error("no such video: " .. name)
    end
    return v
end

-------------------------------------------------------------------------------
function pub.format_score(score)
    check.nv("score", score)
    if score == 0 then
        return "00"
    end
    local s_score = tostring(score)
    local f_score = ""
    local n_digits = 0
    for i=#s_score,1,-1 do
        f_score = string.sub(s_score, i, i) .. f_score
        n_digits = n_digits + 1
        if n_digits % 3 == 0 and i > 1 then
            f_score = "," .. f_score
        end
    end
    return f_score
end

-------------------------------------------------------------------------------
local function extract_var(msg)
    local kind, value
    for k, v in pairs(msg.value) do
        kind = k
        value = v
    end
    return msg.name, kind, value
end

function pub.player()
    return pub.players[pub.vars.player]
end

-------------------------------------------------------------------------------
function pub.ready()
    return true
end

function pub.sleep(secs)
    local millis = secs * 1000
    local wake_at = pub.elapsed + millis
    coroutine.yield(function ()
        return pub.elapsed >= wake_at, 'wake'
    end)
end

function pub.wait(...)
    local conds = {...}
    return coroutine.yield(function(kind, msg)
        for i, cond in ipairs(conds) do
            local result, r_kind, r_msg = cond(kind, msg)
            if result then
                return true, r_kind, r_msg
            end
        end
        return false
    end)
end

function pub.for_any(name)
    check.nv("name", name)
    return function(kind)
        return kind == name, kind, msg
    end
end

function pub.for_ball(name, time)
    check.nv("name", name, "string")
    check.nv("time", time, "number")

    local here = false
    local expires = 0
    local time_ms = time * 1000
    return function(kind, msg)
        local now = spin.elapsed
        if kind == std.SWITCH_UPDATED and msg.name == name and msg.active then
            here = true
            expires = now + time_ms
        elseif kind == std.SWITCH_UPDATED and msg.name == name and not msg.active then
            here = false
            expires = 0
        elseif kind == std.TICK and here and now >= expires then
            return true, std.BALL_ARRIVED, { name = name }
        end
        return false, kind, msg
    end

end

function pub.for_switch(name, active)
    check.nv("name", name)
    if active == nil then
        active = true
    end
    return function (kind, msg)
        return kind == "switch_updated" and msg.name == name and msg.active == active, kind, msg
    end
end

function pub.for_eq(name, value)
    check.nv("name", name)
    check.nv("value", value)
    return function (kind, msg)
        if kind == "updated" then
            local var_name, _, var_value = extract_var(msg)
            return var_name == name and var_value == value, kind, msg
        else
            return false
        end
    end
end

function pub.for_time(secs)
    check.nv("secs", secs)
    local millis = secs * 1000
    local wake_at = pub.elapsed + millis
    return function (kind, msg)
        return pub.elapsed >= wake_at, 'wake'
    end
end

function pub.for_script(name)
    check.nv("name", name)
    return function (kind, msg)
        return (kind == "script_ended" or kind == "script_killed") and msg.name == name
    end
end

function pub.forever()
    return function()
        return false
    end
end

-------------------------------------------------------------------------------
function pub.alert(message)
    table.insert(queue, { note = {
        kind = 'alert',
        message = message,
    }})
end

function pub.diag(message)
    table.insert(queue, { note = {
        kind = 'diag',
        message = message,
    }})
end

function pub.fault(message)
    table.insert(queue, { note = {
        kind = 'fault',
        message = message,
    }})
end

-------------------------------------------------------------------------------
function pub.halt()
    table.insert(queue, "halt")
end

function pub.kill(name)
    check.nv("name", name)
    table.insert(queue, { kill = { name = name } })
end

function pub.kill_group(name)
    check.nv("name", name)
    table.insert(queue, { kill_group = { name = name } })
end

function pub.info(message)
    table.insert(queue, { note = {
        kind = 'info',
        message = message,
    }})
end

function pub.blink_driver(name)
    check.nv("name", name)
    table.insert(queue, { schedule_driver = {
        name = name,
        cycle_time = 1000,
        schedule = {
            {true, 0.125}, {false, 0.125},
            {true, 0.125}, {false, 0.125},
            {true, 0.125}, {false, 0.125},
            {true, 0.125}, {false, 0.125},
        }
    }})
end

function pub.play_music(name, opts)
    check.nv("name", name)
    local msg = {
        name = name
    }
    copy_opts(opts, msg,
        'loops',
        'no_restart',
        'notify'
    )
    table.insert(queue, { play_music = msg })
end

function pub.play_sound(name, opts)
    check.nv("name", name)
    local msg = { name = name }
    copy_opts(opts, msg,
        'loops',
        'notify'
    )
    table.insert(queue, { play_sound = msg })
end

function pub.play_vocal(name, opts)
    check.nv("name", name)
    local msg = { name = name }
    copy_opts(opts, msg,
        'loops',
        'notify'
    )
    table.insert(queue, { play_vocal = msg })
end

function pub.pulse_driver(name, time)
    check.nv("name", name, "string")
    table.insert(queue, { pulse_driver = {
        name = name,
        time = time,
    }})
end

function pub.pwm_driver(name, time_on, time_off)
    check.nv("name", name, "string")
    check.nv("time_on", time_on, "number")
    time_off = check.default(time_off, time_on)
    table.insert(queue, { pwm_driver = {
        name = name,
        time_on = time_on,
        time_off = time_off,
    }})
end

function pub.rejected(reason)
    check.nv("reason", reason)
    table.insert(queue, { rejected = {reason=reason}})
end

function pub.reset()
    table.insert(queue, "reset")
end

function pub.reset_lights()
    table.insert(queue, "reset_lights")
end

function pub.reset_timer(name)
    check.nv("name", name, "string")
    table.insert(queue, { reset_timer = {
        name = name
    }})
end

function pub.run(name)
    check.nv("name", name)
    table.insert(queue, { run = {
        name = name
    }})
end

function pub.schedule_driver(name, schedule)
    check.nv("name", name)
    table.insert(queue, { schedule_driver = {
        name = name,
        cycle_time = 1000,
        schedule = schedule,
    }})
end

-- function pub.set(name, value)
--     check.nv("name", name, "string")
--     check.nv("value", value)
--     table.insert(queue, { set = {
--         vars = {
--             [name] = set_nv(name, value)
--         }
--     }})
-- end

function pub.silence()
    table.insert(queue, "silence")
end

function pub.start_driver(name)
    check.nv("name", name, "string")
    table.insert(queue, { start_driver = {
        name = name
    }})
end

function pub.start_timer(name)
    check.nv("name", name, "string")
    table.insert(queue, { start_timer = {
        name = name
    }})
end

function pub.stop_driver(name)
    check.nv("name", name, "string")
    table.insert(queue, { stop_driver = {
        name = name
    }})
end

function pub.stop_music(name)
    if name == nil then
        name = ""
    end
    table.insert(queue, { stop_music = {
        name = name
    }})
end

function pub.stop_timer(name)
    check.nv("name", name, "string")
    table.insert(queue, { stop_timer = {
        name = name
    }})
end

function pub.stop_vocal(name)
    if name == nil then
        name = ""
    end
    table.insert(queue, { stop_vocal = {
        name = name
    }})
end

function pub.switch_updated(name, active)
    check.nv("name", name)
    if active == nil then
        active = true
    end
    table.insert(queue, {switch_updated = {name=name, active=active}})
end

-------------------------------------------------------------------------------

package.loaded["spin"] = pub
spin = pub

return pub
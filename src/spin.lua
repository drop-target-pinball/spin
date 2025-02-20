local check = require("check")
local std = require("std")

local pub = {
    conf = {},
    vars = {},
    gfx = require("_render").gfx,
}

local script_defs = {}
local scripts = {}
local alive = {}
local queue = {}

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

function pub._init()
    for name, def in pairs(pub.conf.scripts) do
        local mod = require(def.module)
        if type(mod) ~= "table" then
            error("module '" .. def.module .. "' did not return a table")
        end
        script = mod[name]
        if script == nil then
            error("script '" .. name .. "' not found in module '" .. def.module .. "'")
        end
        script_defs[name] = def
        scripts[name] = mod[name]
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

function pub.post(msg)
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

local function set_nv(name, value)
    check.nv('name', name)
    check.nv('value', value)

    if type(value) == "number" then
        if tonumber(tostring(value), 10) then
            return { int = value }
        else
            return { float = value }
        end
    elseif type(value) == "boolean" then
        return { bool = value }
    elseif type(value) == "string" then
        return { string = value }
    end

    error("unsupported type: " .. value)
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

function pub.ns(ns_name)
    local vars = nil
    if ns_name == nil then
        vars = pub.vars
    else
        local ns = pub.vars[ns_name]
        if ns == nil or ns.vars == nil then
            error("not a namespace: " .. ns_name)
        end
        vars = ns.vars
    end

    local ns = {}

    function ns.add_int(name, value)
        check.nv("name", name, "string")
        check.nv("value", value, "number")
        local old = ns.int(name)
        ns.set(name, old + value)
    end

    function ns.bool(name)
        check.nv('name', name)
        local v = vars[name]
        if v == nil then
            error("undefined variable: " .. name)
        end
        if v["bool"] == nil then
            error("variable is not a bool: " .. name)
        end
        return v["bool"]
    end

    function ns.int(name)
        check.nv('name', name)
        local v = vars[name]
        if v == nil then
            error("undefined variable: " .. name)
        end
        if v["int"] == nil then
            error("variable is not an int: " .. name)
        end
        return v["int"]
    end

    function ns.set(name, value)
        check.nv("name", name, "string")
        check.nv("value", value)
        table.insert(queue, { set = {
            ns = ns_name,
            vars = {
                [name] = set_nv(name, value)
            }
        }})
    end

    return ns
end

function pub.bool(name)
    return pub.ns().bool(name)
end

function pub.int(name)
    return pub.ns().int(name)
end

function pub.player()
    return pub.ns("player_" .. pub.int("player"))
end

-------------------------------------------------------------------------------
function pub.ready()
    return true
end

function pub.sleep(secs)
    local millis = secs * 1000
    local wake_at = pub.int('elapsed') + millis
    coroutine.yield(function ()
        return pub.int('elapsed') >= wake_at, 'wake'
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
        local now = spin.int('elapsed')
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
    local wake_at = pub.int('elapsed') + millis
    return function (kind, msg)
        return pub.int('elapsed') >= wake_at, 'wake'
    end
end

function pub.for_script(name)
    check.nv("name", name)
    return function (kind, msg)
        return kind == "script_ended" and msg.name == name
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
            {true, 125}, {false, 125},
            {true, 125}, {false, 125},
            {true, 125}, {false, 125},
            {true, 125}, {false, 125},
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

function pub.set(name, value)
    check.nv("name", name, "string")
    check.nv("value", value)
    table.insert(queue, { set = {
        vars = {
            [name] = set_nv(name, value)
        }
    }})
end

function pub.set_ns(ns, name, value)
    check.nv(ns, "ns")
    check.nv(name, "name")
    check.nv(value, "value")
    table.insert(queue, { set = {
        ns = ns,
        vars = {
            [name] = set_nv(name, value)
        }
    }})
end

function pub.set_multi(vars)
    check.nv("vars", vars)
    local msg = {}
    for name, value in pairs(vars) do
        msg[name] = set_nv(name, value)
    end
    table.insert(queue, { set = {vars=msg} })
end

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
local pub = {
    conf = {},
    vars = {},
}

local script_defs = {}
local scripts = {}
local alive = {}
local queue = {}

-------------------------------------------------------------------------------
local function halt()
    alive = {}
end

local function init()
    for name, def in pairs(pub.conf.scripts) do
        local mod = require(def.module)
        if type(mod) ~= "table" then
            error("module '" .. def.module .. "' did not return a table")
        end
        script_defs[name] = def
        scripts[name] = mod[name]
    end
end

local function kill(name)
    if alive[name] == nil then
        return
    end
    alive[name] = nil
end

local function kill_group(group)
    for name, def in pairs(pub.conf.scripts) do
        if def.group == group then
            kill(name)
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
    if this_def.replace and this_def.group ~= "" then
        for name, def in pairs(pub.conf.scripts) do
            if def.group == this_def.group then
                kill(name)
            end
        end
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
            if script.can_resume(kind, msg) then
                local running, result = coroutine.resume(script.co)
                if not running and result ~= nil then
                    error(result)
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
    elseif kind == 'init' then
        init()
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
local function must_have(name, val)
    if val == nil then
        error(name .. " is required")
    end
end

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

-------------------------------------------------------------------------------
local function extract_var(msg)
    local kind, value
    for k, v in pairs(msg.value) do
        kind = k
        value = v
    end
    return msg.name, kind, value
end

function pub.bool(name)
    must_have('name', name)
    local v = pub.vars[name]
    if v == nil then
        error("undefined variable: " .. name)
    end
    if v["bool"] == nil then
        error("variable is not a bool: " .. name)
    end
    return v["bool"]
end

function pub.int(name)
    must_have('name', name)
    local v = pub.vars[name]
    if v == nil then
        error("undefined variable: " .. name)
    end
    if v["int"] == nil then
        error("variable is not an int: " .. name)
    end
    return v["int"]
end

-------------------------------------------------------------------------------
function pub.ready()
    return true
end

function pub.sleep(secs)
    local millis = secs * 1000
    local wake_at = pub.int('elapsed') + millis
    coroutine.yield(function ()
        return pub.int('elapsed') >= wake_at
    end)
end

function pub.wait(...)
    local conds = {...}
    coroutine.yield(function(kind, msg)
        for i, cond in ipairs(conds) do
            if cond(kind, msg) then
                return true
            end
        end
        return false
    end)
end

function pub.for_any(name)
    must_have("name", name)
    return function(kind)
        return kind == name
    end
end

function pub.for_switch(name, active)
    must_have("name", name)
    if active == nil then
        active = true
    end
    return function (kind, msg)
        return kind == "switch_updated" and msg.name == name and msg.active == active
    end
end

function pub.for_eq(name, value)
    must_have("name", name)
    must_have("value", value)
    return function (kind, msg)
        if kind == "updated" then
            local var_name, _, var_value = extract_var(msg)
            return var_name == name and var_value == value
        else
            return false
        end
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
    if name == nil then
        error('name is required')
    end
    table.insert(queue, { kill = { name = name } })
end

function pub.kill_group(name)
    if name == nil then
        error('name is required')
    end
    table.insert(queue, { kill_group = { name = name } })
end

function pub.info(message)
    table.insert(queue, { note = {
        kind = 'info',
        message = message,
    }})
end

function pub.play_music(name, opts)
    if name == nil then
        error("name is required")
    end
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
    must_have("name", name)
    local msg = { name = name }
    copy_opts(opts, msg,
        'loops',
        'notify'
    )
    table.insert(queue, { play_sound = msg })
end

function pub.play_vocal(name, opts)
    if name == nil then
        error("name is required")
    end
    local msg = { name = name }
    copy_opts(opts, msg,
        'loops',
        'notify'
    )
    table.insert(queue, { play_vocal = msg })
end

function pub.rejected(reason)
    must_have("reason", reason);
    table.insert(queue, { rejected = {reason=reason}})
end

function pub.run(name)
    must_have("name", name);
    table.insert(queue, { run = {
        name = name
    }})
end

local function set_nv(name, value)
    must_have('name', name)
    must_have('value', value)

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

function pub.set(name, value)
    table.insert(queue, { set = {
        vars = {
            [name] = set_nv(name, value)
        }
    }})
end

function pub.set_multi(vars)
    must_have("vars", vars)
    local msg = {}
    for name, value in pairs(vars) do
        msg[name] = set_nv(name, value)
    end
    table.insert(queue, { set = {vars=msg} })
end

function pub.silence()
    table.insert(queue, "silence")
end

function pub.stop_music(name)
    if name == nil then
        name = ""
    end
    table.insert(queue, { stop_music = {
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
    must_have("name", name)
    if active == nil then
        active = true
    end
    table.insert(queue, {switch_updated = {name=name, active=active}})
end

package.loaded["spin"] = pub
spin = pub

return pub
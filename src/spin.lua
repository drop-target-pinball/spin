spin = {
    conf = {},
    vars = {},
}

local script_defs = {}
local scripts = {}
local running = {}
local queue = {}

-------------------------------------------------------------------------------
local function halt()
    running = {}
end

local function init()
    for i, def in ipairs(spin.conf.scripts) do
        local mod = require(def.module)
        if type(mod) ~= "table" then
            error("module '" .. def.module .. "' did not return a table")
        end
        script_defs[def.name] = def
        scripts[def.name] = mod[def.name]
    end
end

local function kill(name)
    if running[name] == nil then
        return
    end
    running[name] = nil
end

local function kill_group(group)
    for i, def in ipairs(spin.conf.scripts) do
        if def.group == group then
            kill(def.name)
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
        for i, def in ipairs(spin.conf.scripts) do
            if def.group == this_def.group then
                kill(def.name)
            end
        end
    end

    -- Create the coroutine and place it in the running table. Set the wait
    -- condition to ready so that it will execute on the next tick
    local co = coroutine.create(script)
    running[name] = {
        co = co,
        can_resume = spin.ready
    }
end

local function service_coroutines(kind, msg)
    for name, script in pairs(running) do
        if coroutine.status(script.co) == "dead" then
            table.insert(queue, { script_ended = {
                name = name
            }})
            running[name] = nil
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

function spin.post(msg)
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

    service_coroutines(kind, msg)

    if next(queue) == nil then
        return nil
    else
        local ret = queue
        queue = {}
        return ret
    end
end

-------------------------------------------------------------------------------
function spin.ready()
    return true
end

function spin.sleep(secs)
    local millis = secs * 1000
    local wake_at = spin.int('elapsed') + millis
    coroutine.yield(function ()
        return spin.int('elapsed') >= wake_at
    end)
end

function spin.wait_for(name)
    coroutine.yield(function (kind)
        return kind == name
    end)
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

function spin.int(name)
    must_have('name', name)
    return spin.vars[name]["int"]
end

-------------------------------------------------------------------------------
function spin.alert(message)
    table.insert(queue, { note = {
        kind = 'alert',
        message = message,
    }})
end

function spin.diag(message)
    table.insert(queue, { note = {
        kind = 'diag',
        message = message,
    }})
end

function spin.fault(message)
    table.insert(queue, { note = {
        kind = 'fault',
        message = message,
    }})
end

function spin.halt()
    table.insert(queue, "halt")
end

function spin.kill(name)
    if name == nil then
        error('name is required')
    end
    table.insert(queue, { kill = { name = name } })
end

function spin.kill_group(name)
    if name == nil then
        error('name is required')
    end
    table.insert(queue, { kill_group = { name = name } })
end

function spin.info(message)
    table.insert(queue, { note = {
        kind = 'info',
        message = message,
    }})
end

function spin.play_music(name, opts)
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

function spin.play_sound(name, opts)
    must_have("name", name)
    local msg = { name = name }
    copy_opts(opts, msg,
        'loops',
        'notify'
    )
    table.insert(queue, { play_sound = msg })
end

function spin.play_vocal(name, opts)
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

function spin.run(name)
    must_have("name", name);
    table.insert(queue, { run = {
        name = name
    }})
end

function spin.set_var(name, value)
    must_have('name', name)
    must_have('value', value)
    msg = {
        name = name,
        value = {},
    }
    if type(value) == "number" then
        if tonumber(tostring(value), 10) then
            msg.value = { int = value }
        else
            msg.value = { float = value }
        end
    elseif type(value) == "boolean" then
        msg.value = { bool = value }
    elseif type(value) == "string" then
        msg.value = { string = value }
    else
        error("unsupported type: " .. value)
    end
    table.insert(queue, { set_var = msg })
end

function spin.silence()
    table.insert(queue, "silence")
end

function spin.stop_music(name)
    if name == nil then
        name = ""
    end
    table.insert(queue, { stop_music = {
        name = name
    }})
end

function spin.stop_vocal(name)
    if name == nil then
        name = ""
    end
    table.insert(queue, { stop_vocal = {
        name = name
    }})
end


package.loaded['spin'] = spin


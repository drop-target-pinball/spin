spin = {
    conf = {},
    vars = {},
}

local scripts = {}
local running = {}
local queue = {}

local function init()
    for i, def in ipairs(spin.conf.scripts) do
        local mod = require(def.module)
        scripts[def.name] = mod[def.call]
    end
end

local function run(name)
    local script = scripts[name]
    if script == nil then
        error("no such script: " .. name)
    end

    -- Create the coroutine and place it in the running table. Set the wait
    -- condition to ready so that it will execute on the next tick
    local co = coroutine.create(script)
    running[name] = {
        co = co,
        can_resume = spin.ready
    }
end

local function tick()
    for name, script in pairs(running) do
        if coroutine.status(script.co) == "dead" then
            table.insert(queue, { script_ended = {
                name = name
            }})
            running[name] = nil
        else
            if script.can_resume() then
                local running, cond = coroutine.resume(script.co)
                if running then
                    script.can_resume = cond
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

    if kind == 'init' then
        init()
    elseif kind == 'run' then
        run(body.name)
    elseif kind == 'tick' then
        tick()
    end

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
    local wake_at = spin.vars.elapsed + millis
    coroutine.yield(function ()
        return spin.vars.elapsed >= wake_at
    end)
end

-------------------------------------------------------------------------------
function spin.alert(message)
    table.insert(queue, { note = {
        kind = 'alert',
        message = message,
    }})
end

function spin.fault(message)
    table.insert(queue, { note = {
        kind = 'fault',
        message = message,
    }})
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
    msg = {
        name = name
    }
    copy_opts(opts, msg,
        'volume',
        'loops',
        'notify'
    )
    table.insert(queue, { play_music = msg })
end

function spin.play_sound(name, opts)
    table.insert(queue, { play_sound = {
        name = name
    }})
end

function spin.stop_music(name)
    if name == nil then
        name = ""
    end
    table.insert(queue, { stop_music = {
        name = name
    }})
end

function spin.run(name)
    table.insert(queue, { run = {
        name = name
    }})
end

-------------------------------------------------------------------------------
function copy_opts(src, dest, ...)
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

package.loaded['spin'] = spin


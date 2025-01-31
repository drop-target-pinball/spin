spin = {
    conf = {}
}

local procs = {}
local running = {}
local queue = {}

local function init()
    for i, def in ipairs(spin.conf.procs) do
        local mod = require(def.module)
        procs[def.name] = mod[def.call]
    end
end

local function run(name)
    local proc = procs[name]
    if proc == nil then
        error("no such procedure: " .. name)
    end

    -- Create the coroutine and place it in the running table. Set the wait
    -- condition to ready so that it will execute on the next tick
    local co = coroutine.create(proc)
    running[name] = {
        co = co,
        resume_when = ready
    }
end

local function tick()
    for name, proc in pairs(running) do
        if coroutine.status(proc.co) == "dead" then
            table.insert(queue, { proc_ended = {
                name = name
            }})
            running[name] = nil
        else
            if proc.resume_when() then
                coroutine.resume(proc.co)
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

function ready()
    return true
end

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

function spin.play_sound(name, opts)
    table.insert(queue, { play_sound = {
        name = name
    }})
end

function spin.run(name)
    table.insert(queue, { run = {
        name = name
    }})
end

package.loaded['spin'] = spin


local spin = require("spin")

engine = {}

local function hello()
    spin.info("hello world")
    spin.play_sound("foo")
end

local procs = {
    hello = hello
}

local running = {}

local function run(name)
    local proc = procs[name]
    if proc == nil then
        error("no such procedure: " .. name)
    end
    local co = coroutine.create(proc)
    coroutine.resume(co)
end

function engine.process(msg)
    spin.out = {}
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
    if kind == 'run' then
        run(body.name)
    end

    if next(spin.out) == nil then
        return nil
    else
        return spin.out
    end
end



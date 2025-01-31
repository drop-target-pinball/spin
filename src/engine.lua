spin = require("spin")
engine = {}

local procs = {}

local running = {}

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
    local co = coroutine.create(proc)
    coroutine.resume(co)
end

function engine.process(msg)
    -- spin.out = {}
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
    end

    if next(spin.out) == nil then
        return nil
    else
        -- FIXME: hack
        local ret = spin.out
        spin.out = {}
        return ret
    end
end

package.loaded['engine'] = engine



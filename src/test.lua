local spin = require("spin")
local std = require("std")
local check = require("check")

local pub = {}

function pub.press(name)
    check.nv("name", name)
    spin.switch_updated(name)
    spin.switch_updated(name, false)
end

--function pub.ok()
    -- spin.post(std.TEST_OK)
--end

function pub.wait(timeout, desc, ...)
    check.nv("timeout", timeout, "number")
    check.nv("desc", desc, "string")
    local conds = {...}
    table.insert(conds, spin.for_time(timeout))
    kind, msg = spin.wait(table.unpack(conds))
    if kind == std.WAKE then
        error("test timeout waiting for: " .. desc)
    end
    return kind, msg
end


package.loaded["test"] = pub
return pub

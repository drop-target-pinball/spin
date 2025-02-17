local spin = require("spin")
local std = require("std")
local check = require("check")

local pub = {}

function pub.press(name)
    check.nv("name", name)
    spin.switch_updated(name)
    spin.switch_updated(name, false)
end

function pub.ok()
    spin.post(std.TEST_OK)
end


package.loaded["test"] = pub
return pub

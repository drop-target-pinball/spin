local spin = require("spin")
local std = require("std")
local test = require("test")

local pub = {}

function pub.test_start_service_no_credits()
    spin.set(std.FREE_PLAY, false)
    spin.run(std.START_SERVICE)
    test.press(std.START_BUTTON)
    local kind = spin.wait(spin.for_any(std.REJECTED, spin.for_time(1)))
    if kind == std.WAKE then
        error("expected credits required")
    end
end

package.loaded["_game_test"] = pub
return pub
local spin = require("spin")
local conf, settings = spin.conf, spin.settings
local std = require("std")
local test = require("test")

local pub = {}

function pub.test_start_service_no_credits()
    settings.free_play = false
    spin.run(std.START_SERVICE)
    test.press(std.START_BUTTON)
    test.wait(1, "rejected", spin.for_any(std.REJECTED))
end

function pub.test_start_service_player_4()
    settings.free_play = true
    conf.max_players = 4
    spin.run(std.START_SERVICE)

    test.press(std.START_BUTTON)
    test.wait(1, "game_active=true", spin.for_eq(std.GAME_ACTIVE, true))
    test.press(std.START_BUTTON)
    test.wait(1, "player_count=2", spin.for_eq(std.PLAYER_COUNT, 2))
    test.press(std.START_BUTTON)
    test.wait(1, "player_count=3", spin.for_eq(std.PLAYER_COUNT, 3))
    test.press(std.START_BUTTON)
    test.wait(1, "player_count=4", spin.for_eq(std.PLAYER_COUNT, 4))
    test.press(std.START_BUTTON)
    test.wait(1, "rejected", spin.for_any(std.REJECTED))
end


package.loaded["_game_test"] = pub
return pub
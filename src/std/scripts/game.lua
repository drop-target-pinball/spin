local spin = require("spin")
local std = require("std")

local pub = {}

local function assert_open_spot()
    if spin.int(std.PLAYER_COUNT) >= spin.int(std.MAX_PLAYERS) then
        spin.rejected(std.GAME_FULL)
        return false
    end
    return true
end

local function accept_payment()
    if spin.bool(std.FREE_PLAY) then
        return true
    end
    local credits = spin.int(std.CREDITS)
    if credits == 0 then
        spin.rejected(std.CREDITS_REQUIRED)
        return false
    end
    spin.set(std.CREDITS, credits - 1)
    return true
end

function pub.start_service()
    while true do
        spin.wait(spin.for_switch(std.START_BUTTON))
        if assert_open_spot() and accept_payment() then
            if spin.bool(std.GAME_ACTIVE) then
                spin.run(std.ADD_PLAYER)
            else
                spin.run(std.START_GAME)
            end
        end
    end
end

function pub.start_game()
    spin.set_multi({
        [std.PLAYER_COUNT] = 1,
        [std.PLAYER] = 1,
        [std.GAME_ACTIVE] = true
    })
end

function pub.add_player()
    if not assert_open_spot() then
        return
    end
    local new_count = spin.int(std.PLAYER_COUNT) + 1
    spin.set(std.PLAYER_COUNT, new_count)
end

package.loaded["game"] = pub

return pub





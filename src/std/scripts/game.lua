local spin = require("spin")
local conf, vars, settings = spin.conf, spin.vars, spin.settings
local std = require("std")

local pub = {}

local function assert_open_spot()
    if vars.player_count >= conf.max_players then
        spin.rejected(std.GAME_FULL)
        return false
    end
    return true
end

local function accept_payment()
    if settings.free_play then
        return true
    end
    if settings.credits == 0 then
        spin.rejected(std.CREDITS_REQUIRED)
        return false
    end
    settings.credits = settings.credits - 1
    return true
end

function pub.start_service()
    while true do
        spin.wait(spin.for_switch(std.START_BUTTON))
        if assert_open_spot() and accept_payment() then
            if vars.game_active then
                spin.run(std.ADD_PLAYER)
            else
                spin.run(std.START_GAME)
            end
        end
    end
end

function pub.start_game()
    vars.player_count = 1
    vars.player = 1
    vars.game_active = true
end

function pub.add_player()
    if not assert_open_spot() then
        return
    end
    vars.player_count = vars.player_count + 1
end

package.loaded["_game"] = pub

return pub





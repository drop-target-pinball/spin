local spin = require("spin")
local std = require("std")

local pub = {}

function pub.start_button_credit()
    while true do
        spin.wait(spin.for_switch(std.START_BUTTON))

        if not spin.bool(std.GAME_ON) then
            if spin.bool(std.FREE_PLAY) then
                run(std.START_GAME)
            else
                local credits = spin.int(std.CREDITS)
                if credits == 0 then
                    spin.credits_required()
                else
                    spin.set_int(std.CREDITS, credits - 1)
                    run(std.START_GAME)
                end
            end
        else
            -- Game On
            if spin.int(std.PLAYER_NUM) >= spin.int(std.MAX_PLAYERS) then
                spin.game_full()
            else
                run(std.ADD_PLAYER)
            end
        end
    end
end

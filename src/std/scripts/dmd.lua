local spin = require("spin")
local std = require("std")

local pub = {}

local function score_footer(gfx)
    local ball = spin.int("ball")
    local y = spin.video(std.DMD).height - 5

    gfx.font = std.DMD_04B_03_7PX
    gfx.draw_text(24, y, "BALL " .. ball)
    gfx.draw_text(75, y, "FREE PLAY")
end

local function score_single_draw(gfx)
    local score = spin.player().int(std.SCORE)
    local font = std.DMD_18X10
    if score < 10^9 then
        font = std.DMD_18X12
    elseif score < 10^12 then
        font = std.DMD_18X11
    end

    gfx.font = font
    gfx.draw_text_y(3, spin.format_score(score))
    score_footer(gfx)
end

local function score_multi_draw(gfx)
    local function font_for(player, score)
        local active = player == spin.int("player")
        if active and score < 10^7 then
            return std.DMD_14X10
        elseif active and score < 10^10 then
            return std.DMD_14X9
        elseif active then
            return std.DMD_14X8
        elseif score < 10^7 then
            return std.DMD_09X7
        elseif score < 10^10 then
            return std.DMD_09X6
        else
            return std.DMD_09X5
        end
    end

    local score_1 = spin.ns(std.PLAYER_1).int(std.SCORE)
    gfx.font = font_for(1, score_1)
    gfx.draw_text(0, 0, spin.format_score(score_1))
    score_footer(gfx)
end

function pub.score_draw()
    while true do
        local gfx = spin.gfx(std.DMD)
        gfx.new(gfx.BLACK)
        if spin.int(std.PLAYER_COUNT) == 1 then
            score_single_draw(gfx)
        else
            score_multi_draw(gfx)
        end
        spin.wait(spin.for_any(std.TICK))
    end
end

package.loaded["_dmd"] = pub

return pub


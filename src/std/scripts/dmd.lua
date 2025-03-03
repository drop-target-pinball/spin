local spin = require("spin")
local vars = spin.vars
local std = require("std")

local pub = {}

local function score_footer(gfx)
    local y = spin.video(std.DMD).height - 5

    gfx.font = std.DMD_04B_03_7PX
    gfx.draw_text(24, y, "BALL " .. vars.ball)
    gfx.draw_text(75, y, "FREE PLAY")
end

local function score_single_draw(gfx)
    local score = spin.player().score
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
        local active = player == vars.player
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

    local dmd = spin.video(std.DMD)

    local score_1 = spin.players[1].score
    gfx.font = font_for(1, score_1)
    gfx.draw_text(0, 0, spin.format_score(score_1))

    local score_2 = spin.players[2].score
    gfx.font = font_for(2, score_2)
    gfx.draw_text(dmd.width + 1, 0, spin.format_score(score_2), {right=true})

    if vars.player_count >= 3 then
        local score_3 = spin.players[3].score
        gfx.font = font_for(3, score_3)
        gfx.draw_text(0, dmd.height - 6, spin.format_score(score_3), {bottom=true})
    end

    if vars.player_count >= 4 then
        local score_4 = spin.players[4].score
        gfx.font = font_for(4, score_4)
        gfx.draw_text(dmd.width + 1, dmd.height -6, spin.format_score(score_4), {right=true, bottom=true})
    end

    score_footer(gfx)
end

function pub.score_draw()
    while true do
        local gfx = spin.gfx(std.DMD)
        gfx.new(spin.OFF)
        if vars.player_count == 1 then
            score_single_draw(gfx)
        else
            score_multi_draw(gfx)
        end
        spin.wait(spin.for_any(std.TICK))
    end
end

package.loaded["std_dmd"] = pub

return pub


local spin = require("spin")
local std = require("std")

local pub = {}

function pub.single_score_draw()
    while true do
        spin.wait(spin.for_any(std.TICK))
    end
end

-- switch {
-- 	case player.Score < 1_000_000_000:
-- 		g.Font = Font18x12
-- 	case player.Score < 10_000_000_000:
-- 		g.Font = Font18x11
-- 	default:
-- 		g.Font = Font18x10
-- 	}
-- 	g.Y = 3
-- 	r.Print(g, FormatScore("%d", player.Score))
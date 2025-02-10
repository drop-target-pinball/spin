local std = require("std")
local pub = {}

function pub.dmd_gradient()
    local dmd = spin.video(std.DMD)
    local gfx = spin.gfx(std.DMD)
    local w = dmd.width / 16
    for i=0, 15 do
        gfx.dot_color(i)
        gfx.fill_rect(i * w, 0, w, dmd.height)
    end
end

package.loaded["_service"] = pub

return pub

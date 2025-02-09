local pub = {}

function pub.dmd_gradient()
    local gfx = spin.gfx("dmd")
    local w = 128 / 16
    for i=0, 15 do
        gfx.dot_color(i)
        gfx.fill_rect(i * w, 0, w, 32)
    end
end

package.loaded["_diag"] = pub

return pub

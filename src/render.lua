local check = require("check")
local math = require("math")

local pub = {}

pub.ops = {}

local dots = {
    0,
    17, -- 0x11
    34, -- 0x22
    51, -- 0x33
    68, -- etc.
    85,
    102,
    119,
    136,
    153,
    170,
    187,
    204,
    221,
    238,
    255,
}

function pub.gfx(device)
    local gfx = {
        device = device,
        layer = 0,
        priority = 0,
    }

    local function insert_op(op_name, args)
        table.insert(pub.ops, {
            device = gfx.device,
            layer = gfx.layer,
            priority = gfx.priority,
            op = {
                [op_name] = args
            }
        })
    end

    function gfx.color(r, g, b, a)
        check.nv("r", r)
        check.nv("g", g)
        check.nv("b", b)
        if a == nil then
            a = 255
        end
        insert_op('color', {
            r=math.floor(r),
            g=math.floor(g),
            b=math.floor(b),
            a=math.floor(a)
        })
    end

    function gfx.dot_color(dot)
        check.nv("dot", dot)
        if dot < 0 then
            dot = 0
        elseif dot > 15 then
            dot = 15
        end
        local v = dots[dot + 1]
        gfx.color(v, v, v)
    end

    function gfx.fill_rect(x, y, w, h)
        check.nv("x", x)
        check.nv("y", y)
        check.nv("w", w)
        check.nv("h", h)
        insert_op('fill_rect', {
            x=math.floor(x),
            y=math.floor(y),
            w=math.floor(w),
            h=math.floor(h)
        })
    end

    return gfx
end

package.loaded["_render"] = pub
_render = pub

return pub

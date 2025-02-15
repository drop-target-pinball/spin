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

local function unpack_color(color)
    check.nv("color", color)
    local r = math.floor(check.default(color.r, 0))
    local g = math.floor(check.default(color.g, 0))
    local b = math.floor(check.default(color.b, 0))
    local a = math.floor(check.default(color.a, 255))
    return r, g, b, a
end

function pub.gfx(device, layer, priority)
    check.nv("device", device)

    layer = check.default(layer, 0)
    priority = check.default(priority, 0)

    local gfx = {
        device = device,
        layer = layer,
        priority = priority,

        color = {r=255, g=255, b=255, a=255},
        font = "",
    }

    gfx.BLACK       = { r = 0,   g = 0,   b = 0,   a = 255 }
    gfx.CLEAR       = { r = 0,   g = 0,   b = 0,   a = 0   }
    gfx.FULL        = { r = 255, g = 255, b = 255, a = 255 }

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

    function gfx.new(color)
        local r, g, b, a = unpack_color(color)
        insert_op('new', {
            r=r,
            g=g,
            b=b,
            a=a
        })
    end

    function gfx.draw_text(x, y, text, opts)
        check.nv("x", x)
        check.nv("y", y)
        check.nv("text", text)

        local msg = {
            x=math.tointeger(x),
            y=math.tointeger(y),
            text=text,
            color=gfx.color,
            font=gfx.font,
        }
        check.copy_opts(opts, msg,
            {"right", "bool"},
            {"bottom", "bool"}
        )
        insert_op("draw_text", msg)
    end

    function gfx.draw_text_x(x, text, opts)
        check.nv("x", x, "number")
        check.nv("text", text, "string")

        local msg = {
            x=math.tointeger(x),
            center_y=true,
            text=tostring(text),
            color=check.nv("gfx.color", gfx.color),
            font=check.nv("gfx.font", gfx.font),
        }
        check.copy_opts(opts, msg, {
            {"right", "bool"},
        })
        insert_op("draw_text", msg)
    end

    function gfx.draw_text_y(y, text, opts)
        check.nv("y", y, "number")
        check.nv("text", text, "string")

        local msg = {
            y=math.tointeger(y),
            center_x=true,
            text=tostring(text),
            color=check.nv("gfx.color", gfx.color),
            font=check.nv("gfx.font", gfx.font),
        }
        check.copy_opts(opts, msg, {
            {"bottom", "bool"}
        })
        insert_op("draw_text", msg)
    end

    function gfx.draw_centered_text(text)
        check.nv("text", text, "string")

        local msg = {
            center_x=true,
            center_y=true,
            text=text,
            color=gfx.color,
            font=gfx.font,
        }
        insert_op("draw_text", msg)
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
            h=math.floor(h),
            color=gfx.color,
        })
    end
    return gfx
end

package.loaded["_render"] = pub
_render = pub

return pub

local pub = {}

local ops = {}

function pub.gfx(device)
    local gfx = {
        device = device,
        layer = 0,
        priority = 0,
    }

    local function insert_op(op_name, args)
        table.insert(ops, {
            device = gfx.device,
            layer = gfx.layer,
            priority = gfx.priority,
            [op_name] = args
        })
    end

    function gfx.color(r, g, b, a)
        if a == nil then
            a = 255
        end
        insert_op('color', {r=r, g=g, b=b, a=a})
    end

    function gfx.fill_rect(x, y, w, h)
        insert_op('fill_rect', {x=x, y=y, w=w, h=h})
    end

    return gfx
end

package.loaded["_render"] = pub

return pub

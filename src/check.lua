local pub = {}

function pub.nv(name, value)
    if value == nil then
        error("value required for '" .. name .. "'")
    end
end

package.loaded["check"] = pub

return pub

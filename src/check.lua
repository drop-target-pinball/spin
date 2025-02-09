local pub = {}

function pub.nv(name, value)
    if value == nil then
        error("value required for '" .. name .. "'")
    end
end

function pub.default(value, default)
    if value == nil then
        return default
    end
    return value
end

package.loaded["check"] = pub

return pub

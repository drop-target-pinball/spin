local pub = {}

function pub.type(name, value, value_type)
    if value_type == nil then
        return value
    end
    if type(value) ~= value_type then
        error("invalid type for '" .. name .. "', expected " .. value_type .. ", got " .. type(value))
    end
    return value
end

function pub.nv(name, value, value_type)
    if type(name) ~= "string" then
        error("expected string for 'name', got " .. type(name))
    end
    if value == nil then
        error("value required for '" .. name .. "'")
    end
    return pub.type(name, value, value_type)
end

function pub.default(value, default)
    if value == nil then
        return default
    end
    return value
end

function pub.copy_opts(src, dest, ...)
    if src == nil then
        return
    end
    local args = {...}
    pub.nv("dest", dest)
    pub.nv("args", args)

    for i, arg in ipairs(args) do
        local name, value_type
        if type(arg) == "string" then
            name = arg
        elseif type(arg) == "table" then
            name, value_type = table.unpack(arg)
        else
            error("invalid type for args: " .. type(arg))
        end
        if src[name] ~= nil then
            pub.nv(name, value_type)
            dest[name] = src[name]
        end
    end
    return dest
end

package.loaded["check"] = pub

return pub

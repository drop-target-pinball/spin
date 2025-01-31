local spin = require("spin")

local hello = {}

function hello.hello()
    spin.info("hello world")
    spin.play_sound("foo")
end

return hello


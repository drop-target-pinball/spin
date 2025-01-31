local spin = require("spin")

local hello = {}

function hello.hello()
    spin.info("hello world")
    spin.sleep(1)
    spin.play_sound("foo")
    spin.sleep(1)
    spin.info("done")
end

return hello


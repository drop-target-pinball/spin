local spin = require("spin")
local std = require("std")

alert = spin.alert
fault = spin.fault
halt = spin.halt
info = spin.info
kill = spin.kill
kill_group = spin.kill_group
play_music = spin.play_music
play_sound = spin.play_sound
play_vocal = spin.play_vocal
run = spin.run
set_var = spin.set_var
silence = spin.silence
stop_music = spin.stop_music
stop_vocal = spin.stop_vocal

function credit(num)
    if num == nil then
        num = 1
    end
    local credits = spin.int(std.CREDITS)
    spin.set(std.CREDITS, credits + num)
end

function press(name)
    spin.switch_updated(name)
    spin.switch_updated(name, false)
end

function start()
    press(std.START_BUTTON)
end


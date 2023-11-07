defaults {
    module = "service"
}

audio "service_enter" {
    type = "sound"
    file = "lib/sound/proc-shared/menu_in.wav"
}

audio "service_exit" {
    type = "sound"
    file = "lib/sound/proc-shared/menu_out.wav"
}

audio "service_up" {
    type = "sound"
    file = "lib/sound/proc-shared/next_item.wav"
}

audio "service_down" {
    type = "sound"
    file = "lib/sound/proc-shared/previous_item.wav"
}
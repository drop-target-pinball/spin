include = [
    "lib/config/defaults.hcl",
    "lib/module/service/service.hcl"
]

load = [
    "service"
]

audio_device "" {
    handler = "sdl_mixer"
}

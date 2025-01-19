use sdl2;
use sdl2::mixer;

struct Audio {

}

impl Audio {
    pub fn new() -> Self {
        let sdl_context = sdl2::init().unwrap();
        let audio_subsystem = sdl_context.audio().unwrap();
        Self{}
    }
}
use crate::prelude::*;

use sdl2;
use sdl2::mixer;
use std::collections::HashMap;
use crate::sdl::System;

struct Audio {
    sounds: HashMap<String,mixer::Chunk>,
}

impl Audio {
    pub fn new() -> Self {
        Self {
            sounds: HashMap::new(),
        }
    }

    fn init(&mut self, sdl: &sdl2::Sdl, ctx: &mut Context) {
        println!("************************ got init!");
    }
}

impl System for Audio {
    fn topic(&self) -> Topic {
        Topic::Audio
    }

    fn process(&mut self, sdl: &sdl2::Sdl, ctx: &mut Context, msg: &Message) {
        match msg {
            Message::Init => self.init(sdl, ctx),
            _ => (),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::sdl;

    #[test]
    pub fn test_init() {
        let mut audio = Audio::new();

        let mut dev_sdl = sdl::Device::new(0).unwrap();
        dev_sdl.add_system(&mut audio);

        let mut e = Engine::default();
        e.add_device(&mut dev_sdl);

        e.init();
        assert!(false);
    }
}
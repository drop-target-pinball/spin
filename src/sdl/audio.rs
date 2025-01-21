use crate::prelude::*;

use sdl2;
use sdl2::mixer;
use std::collections::HashMap;
use crate::sdl::System;

pub struct Audio {
    enabled: bool,
    sounds: HashMap<String,mixer::Chunk>,
    audio: Option<sdl2::AudioSubsystem>,
}

impl Audio {
    pub fn new() -> Self {
        Self {
            enabled: true,
            sounds: HashMap::new(),
            audio: None,
        }
    }

    fn init(&mut self, sdl: &sdl2::Sdl, ctx: &mut Context) {
        self.audio = match sdl.audio() {
            Ok(a) => Some(a),
            Err(e) => {
                fault!(ctx.queue, "unable to open SDL audio: {}", e);
                self.enabled = false;
                return
            }
        };

        let frequency = 44_100;
        let format = mixer::AUDIO_S16LSB; // signed 16 bit samples, in little-endian byte order
        let channels = mixer::DEFAULT_CHANNELS; // Stereo
        let chunk_size = 1_024;

        if let Err(e) = mixer::open_audio(frequency, format, channels, chunk_size) {
            fault!(ctx.queue, "unable to open mixer: {}", e);
            self.enabled = false;
            return
        }

        if let Err(e) = mixer::init(mixer::InitFlag::MP3 | mixer::InitFlag::OGG) {
            fault!(ctx.queue, "unable to open mixer: {}", e);
            self.enabled = false;
            return
        }

        for sound in &ctx.conf.sounds {
            let path = path_to(&sound.path);
            match mixer::Chunk::from_file(&path) {
                Err(e) => fault!(ctx.queue, "unable to load sound: {}", &e),
                Ok(chunk) => match self.sounds.insert(sound.id.clone(), chunk) {
                    Some(_) => fault!(ctx.queue, "sound already loaded: {}", &path.to_string_lossy()),
                    None => (),
                }
            };
        }
    }
}

impl System for Audio {
    fn topic(&self) -> Topic {
        Topic::Audio
    }

    fn process(&mut self, sdl: &sdl2::Sdl, ctx: &mut Context, msg: &Message) {
        if !self.enabled {
            return
        }
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
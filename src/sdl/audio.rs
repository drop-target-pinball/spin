use crate::prelude::*;

use sdl2;
use sdl2::mixer;
use std::collections::HashMap;
use crate::sdl::device_error;

pub struct AudioOptions {
    pub freq: i32,
    pub format: u16,
    pub channels: i32,
    pub chunk_size: i32,
    pub flags: mixer::InitFlag,
}

impl Default for AudioOptions {
    fn default()-> Self {
        Self {
            freq: 44_100,
            format: mixer::AUDIO_S16LSB,
            channels: mixer::DEFAULT_CHANNELS,
            chunk_size: 1_024,
            flags: mixer::InitFlag::MP3 | mixer::InitFlag::OGG,
        }
    }
}


pub struct Audio {
    sounds: HashMap<String,mixer::Chunk>,
    audio: sdl2::AudioSubsystem,
}

impl Audio {
    pub fn new(sdl: &sdl2::Sdl, opt: AudioOptions) -> Result<Self> {
        let audio = match sdl.audio() {
            Ok(a) => a,
            Err(reason) => {
                return device_error("unable to open SDL audio", reason);
            }
        };

        if let Err(reason) = mixer::open_audio(opt.freq, opt.format, opt.channels, opt.chunk_size) {
            return device_error("unable to open mixer", reason);
        }

        if let Err(reason) = mixer::init(opt.flags) {
            return device_error("unable to open mixer", reason);
        }

        Ok(Self {
            sounds: HashMap::new(),
            audio,
        })
    }

    fn init(&mut self, sdl: &sdl2::Sdl, ctx: &mut Context) {
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

    pub fn process(&mut self, sdl: &sdl2::Sdl, ctx: &mut Context, msg: &Message) -> bool {
        match msg {
            Message::Init => self.init(sdl, ctx),
            _ => return false,
        }
        true
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::sdl;

    #[test]
    pub fn test_init() -> Result<()> {
        // let mut dev = sdl::Device
        // let mut dev_sdl = sdl::Device::new(0).unwrap();
        // dev_sdl.add_system(&mut audio);

        // let mut e = Engine::default();
        // e.add_device(&mut dev_sdl);

        // e.init();
        // assert!(false);
        Ok(())
    }
}
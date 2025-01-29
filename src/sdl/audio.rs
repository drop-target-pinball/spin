use crate::prelude::*;

use sdl2;
use sdl2::mixer;
use std::collections::HashMap;

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
    id: u8,
    sounds: HashMap<String,mixer::Chunk>,
    _audio: sdl2::AudioSubsystem,
}

impl Audio {
    pub fn new(ctx: &sdl2::Sdl, id: u8, opt: AudioOptions) -> Self {
        let audio = match ctx.audio() {
            Ok(a) => a,
            Err(reason) => {
                panic!("unable to open SDL audio: {}", reason);
            }
        };

        if let Err(reason) = mixer::open_audio(opt.freq, opt.format, opt.channels, opt.chunk_size) {
            panic!("unable to open mixer: {}", reason);
        }

        if let Err(reason) = mixer::init(opt.flags) {
            panic!("unable to initialize mixer: {}", reason);
        }

        Self {
            id,
            sounds: HashMap::new(),
            _audio: audio,
        }
    }

    fn init(&mut self, env: &mut Env, q: &mut Queue) {
        for sound in &env.conf.sounds {
            if sound.device_id != self.id {
                continue
            }
            let path = env.conf.app_dir.join(&sound.path);
            match mixer::Chunk::from_file(&path) {
                Err(e) => fault!(q, "unable to load sound: {}", &e),
                Ok(chunk) => {
                    if self.sounds.insert(sound.name.clone(), chunk).is_some() {
                        fault!(q, "sound already loaded: {}", &path.to_string_lossy());
                    }
                }
            };
        }
    }

    fn play_sound(&mut self, q: &mut Queue, cmd: &PlayAudio)  {
        let chunk = match self.sounds.get(&cmd.name) {
            Some(c) => c,
            None => {
                fault!(q, "unregistered sound: {}", cmd.name);
                return;
            }
        };

        if let Err(e) = sdl2::mixer::Channel::all().play(chunk, 0) {
            fault!(q, "cannot play sound: {}", e);
        }
    }

    pub fn process(&mut self, _: &sdl2::Sdl, env: &mut Env, q: &mut Queue, msg: &Message)  {
        match msg {
            Message::Init => self.init(env, q),
            Message::PlaySound(a) => self.play_sound(q, a),
            _ => ()
        }
    }
}

// #[cfg(test)]
// mod tests {
//     use super::*;
//     use crate::sdl;

//     #[test]
//     pub fn test_init() -> Result<()> {
//         // let mut dev = sdl::Device
//         // let mut dev_sdl = sdl::Device::new(0).unwrap();
//         // dev_sdl.add_system(&mut audio);

//         // let mut e = Engine::default();
//         // e.add_device(&mut dev_sdl);

//         // e.init();
//         // assert!(false);
//         Ok(())
//     }
// }
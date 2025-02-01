use crate::prelude::*;

use sdl2;
use sdl2::mixer;
use std::collections::HashMap;
use std::sync::atomic::{AtomicBool, Ordering};

static MUSIC_FINISHED: AtomicBool = AtomicBool::new(false);

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


pub struct Audio<'a> {
    id: u8,
    music: HashMap<String,mixer::Music<'a>>,
    sounds: HashMap<String,mixer::Chunk>,
    _audio: sdl2::AudioSubsystem,

    music_playing: Option<String>,
    music_notify: bool,
}

impl<'a> Audio<'a> {
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
            music: HashMap::new(),
            sounds: HashMap::new(),
            _audio: audio,
            music_playing: None,
            music_notify: false,
        }
    }

    fn init(&mut self, env: &mut Env) {
        for music in &env.conf.music {
            if music.device_id != self.id {
                continue
            }
            let path = env.conf.data_dir.join(&music.path);
            match mixer::Music::from_file(&path) {
                Err(e) => fault!(env.queue, "unable to load sound: {}", &e),
                Ok(m) => {
                    if self.music.insert(music.name.clone(), m).is_some() {
                        fault!(env.queue, "music already loaded: {}", &path.to_string_lossy());
                    }
                }
            }
        }

        for sound in &env.conf.sounds {
            if sound.device_id != self.id {
                continue
            }
            let path = env.conf.data_dir.join(&sound.path);
            match mixer::Chunk::from_file(&path) {
                Err(e) => fault!(env.queue, "unable to load sound: {}", &e),
                Ok(chunk) => {
                    if self.sounds.insert(sound.name.clone(), chunk).is_some() {
                        fault!(env.queue, "sound already loaded: {}", &path.to_string_lossy());
                    }
                }
            };
        }

        fn music_finished() {
            MUSIC_FINISHED.store(true, Ordering::SeqCst);
        }
        mixer::Music::hook_finished(music_finished);
    }

    fn play_sound(&mut self, env: &mut Env, cmd: &PlayAudio)  {
        let chunk = match self.sounds.get(&cmd.name) {
            Some(c) => c,
            None => {
                fault!(env.queue, "unregistered sound: {}", cmd.name);
                return;
            }
        };

        if let Err(e) = sdl2::mixer::Channel::all().play(chunk, 0) {
            fault!(env.queue, "cannot play sound: {}", e);
        }
    }

    fn play_music(&mut self, env: &mut Env, cmd: &PlayAudio) {
        let Some(music) = self.music.get(&cmd.name) else { return };

        let mut loops = cmd.loops;
        if loops == 0 {
            loops = -1;
        }
        if let Err(e) = music.play(loops) {
            fault!(env.queue, "cannot play music: {}", e);
        }
        self.music_playing = Some(cmd.name.clone());
        self.music_notify = cmd.notify;
    }

    fn stop_music(&mut self, env: &mut Env, cmd: &Name) {
        let Some(playing) = &self.music_playing else { return };
        if cmd.name != "" && cmd.name != *playing {
            return
        }
        sdl2::mixer::Music::halt();
        self.music_playing = None;
        self.music_notify = false;
    }

    pub fn process(&mut self, _: &sdl2::Sdl, env: &mut Env, msg: &Message) {
        if MUSIC_FINISHED.load(Ordering::SeqCst) {
            MUSIC_FINISHED.store(false, Ordering::SeqCst);
            if let Some(name) = &self.music_playing {
                if self.music_notify {
                    env.queue.post(Message::MusicEnded(Name{name: name.to_string()}));
                }
                self.music_playing = None;
            }
        }

        match msg {
            Message::Init => self.init(env),
            Message::PlayMusic(m) => self.play_music(env, m),
            Message::PlaySound(m) => self.play_sound(env, m),
            Message::StopMusic(m) => self.stop_music(env, m),
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
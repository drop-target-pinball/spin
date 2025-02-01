use crate::prelude::*;

use sdl2;
use sdl2::mixer;
use std::collections::HashMap;
use std::sync::atomic::{AtomicBool, Ordering};

static MUSIC_FINISHED: AtomicBool = AtomicBool::new(false);

struct ActiveAudio {
    name: String,
    priority: i8,
    notify: bool,
}

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
    vocals: HashMap<String,mixer::Chunk>,
    _audio: sdl2::AudioSubsystem,

    music_playing: Option<ActiveAudio>,
    vocal_playing: Option<ActiveAudio>,
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

        // Reserve one channel for vocals
        sdl2::mixer::reserve_channels(1);

        // Allocate the maximum number as specified in the AudioOptions
        sdl2::mixer::allocate_channels(-1);

        Self {
            id,
            music: HashMap::new(),
            sounds: HashMap::new(),
            vocals: HashMap::new(),
            _audio: audio,
            music_playing: None,
            vocal_playing: None,
        }
    }

    fn init(&mut self, env: &mut Env) {
        for music in &env.conf.music {
            if music.device_id != self.id {
                continue
            }
            let path = env.conf.data_dir.join(&music.path);
            match mixer::Music::from_file(&path) {
                Err(e) => fault!(env.queue, "unable to load music '{}': {}", &music.name, &e),
                Ok(m) => {
                    if self.music.insert(music.name.clone(), m).is_some() {
                        fault!(env.queue, "music already loaded '{}': {}", &music.name, &path.to_string_lossy());
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
                Err(e) => fault!(env.queue, "unable to load sound '{}': {}", &sound.name, &e),
                Ok(chunk) => {
                    if self.sounds.insert(sound.name.clone(), chunk).is_some() {
                        fault!(env.queue, "sound already loaded '{}': {}", &sound.name, &path.to_string_lossy());
                    }
                }
            };
        }

        for vocal in &env.conf.vocals {
            if vocal.device_id != self.id {
                continue
            }
            let path = env.conf.data_dir.join(&vocal.path);
            match mixer::Chunk::from_file(&path) {
                Err(e) => fault!(env.queue, "unable to load vocal '{}': {}", &vocal.name, &e),
                Ok(chunk) => {
                    if self.vocals.insert(vocal.name.clone(), chunk).is_some() {
                        fault!(env.queue, "sound already loaded '{}': {}", &vocal.name, &path.to_string_lossy());
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
        let Some(chunk) = self.sounds.get(&cmd.name) else { return };

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
        self.music_playing = Some(ActiveAudio{
            name: cmd.name.clone(),
            priority: 0,
            notify: cmd.notify
        });
    }

    fn play_vocal(&mut self, env: &mut Env, cmd: &PlayAudio) {
        let Some(vocal) = self.vocals.get(&cmd.name) else { return };

        match sdl2::mixer::Channel(0).play(vocal, 0) {
            Ok(_) => sdl2::mixer::Music::set_volume(64),
            Err(e) => fault!(env.queue, "cannot play vocal: {}", e),
        }
    }

    fn stop_music(&mut self, _: &mut Env, cmd: &Name) {
        let Some(playing) = &self.music_playing else { return };
        if cmd.name != "" && cmd.name != *playing.name {
            return
        }
        sdl2::mixer::Music::halt();
        self.music_playing = None;
    }

    pub fn process(&mut self, _: &sdl2::Sdl, env: &mut Env, msg: &Message) {
        if MUSIC_FINISHED.load(Ordering::SeqCst) {
            MUSIC_FINISHED.store(false, Ordering::SeqCst);
            if let Some(playing) = &self.music_playing {
                if playing.notify {
                    env.queue.post(Message::MusicEnded(Name{name: playing.name.to_string()}));
                }
                self.music_playing = None;
            }
        }

        if let Some(playing) = &self.vocal_playing {
            if !mixer::Channel(0).is_playing() {
                if playing.notify {
                    env.queue.post(Message::VocalEnded(Name{name: playing.name.to_string()}));
                }
                sdl2::mixer::Music::set_volume(128);
                self.vocal_playing = None;
            }
        }

        match msg {
            Message::Init => self.init(env),
            Message::PlayMusic(m) => self.play_music(env, m),
            Message::PlaySound(m) => self.play_sound(env, m),
            Message::PlayVocal(m) => self.play_vocal(env, m),
            Message::StopMusic(m) => self.stop_music(env, m),
            _ => ()
        }
    }
}

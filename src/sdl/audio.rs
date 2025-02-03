use crate::prelude::*;

use sdl2;
use sdl2::mixer;
use sdl2::mixer::Channel;
use std::collections::HashMap;
use std::sync::atomic::{AtomicBool, Ordering};

static MUSIC_FINISHED: AtomicBool = AtomicBool::new(false);
const MAX_VOLUME: i32 = 128;

struct ActiveAudio {
    name: String,
    chan: i32,
    duck: i32,
    start_time: u64,
    priority: i32,
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

struct Sound {
    conf: config::Sound,
    chunk: mixer::Chunk,
}

struct Vocal {
    conf: config::Vocal,
    chunk: mixer::Chunk
}

pub struct Audio<'a> {
    id: u8,
    music: HashMap<String,mixer::Music<'a>>,
    sounds: HashMap<String,Sound>,
    vocals: HashMap<String,Vocal>,
    _audio: sdl2::AudioSubsystem,

    music_playing: Option<ActiveAudio>,
    active: Vec<Option<ActiveAudio>>,
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

        let mut channels: Vec<Option<ActiveAudio>> = Vec::new();
        for _ in 0..opt.channels {
            channels.push(None);
        }

        Self {
            id,
            music: HashMap::new(),
            sounds: HashMap::new(),
            vocals: HashMap::new(),
            _audio: audio,
            music_playing: None,
            active: channels
        }
    }

    fn find_available_channel(&self, env: &mut Env, priority: i32) -> i32 {
        let mut opt_candidate: Option<&ActiveAudio> = None;

        // First check if there are any open slots
        for i in 1..self.active.len() {
            match &self.active[i] {
                None => return i as i32,
                // Otherwise, see if this is a candidate that could be
                // freed.
                Some(aa) => {
                    if aa.priority >= priority {
                        continue
                    }
                    match opt_candidate {
                        // If it has a lower priority and we don't have a
                        // candidate yet, this is it.
                        None => opt_candidate = Some(&aa),
                        // If we do have a candidate already, take it if
                        // it has a lower priority or an earlier start
                        // time with the same priority.
                        Some(cand) => {
                            if aa.priority < cand.priority {
                                opt_candidate = Some(&aa);
                                continue;
                            }
                            if aa.priority == cand.priority && aa.start_time < cand.start_time {
                                opt_candidate = Some(&aa);
                            }
                        }
                    }
                }
            }
        }

        // No channels are open. If we have a candidate, we will free that
        // channel and return that.
        match opt_candidate {
            Some(aa) => {
                diag!(env.queue, "a lower priority channel was freed");
                let chan = aa.chan;
                chan
            },
            None => 0
        }
    }

    fn free_channel(&self, env: &mut Env, num: i32) {
        let chan: Channel = Channel(num);
        if chan.is_playing() {
            chan.halt();
        }

        let Some(aa) =  &self.active[num as usize] else { return };
        if aa.notify {
            let name = Name{name: aa.name.clone() };
            if num == 0 {
                env.queue.post(Message::VocalEnded(name));
            } else {
                env.queue.post(Message::SoundEnded(name));
            }
        }
    }

    fn duck(&mut self) {
        // Music shouldn't be ducking when fading in or out
        match mixer::Music::get_fading() {
            mixer::Fading::NoFading => (),
            mixer::Fading::FadingIn => return,
            mixer::Fading::FadingOut => return,
        }

        let curr_vol = mixer::Music::get_volume();

        // Volume is going to be full (128) unless there are requests to
        // duck the music. If so, take the lowest volume found.
        let mut requested_vol: i32 = 128;
        for opt_aa in &self.active {
            match opt_aa {
                None => (),
                Some(aa) => {
                    if aa.duck != 0 && aa.duck < requested_vol {
                        requested_vol = aa.duck;
                    }
                }
            }
        }

        if requested_vol != curr_vol {
            mixer::Music::set_volume(requested_vol);
        }
    }

    fn reap_channels(&mut self, env: &mut Env) {
        for i in 0..self.active.len() {
            let chan = Channel(i as i32);
            match &self.active[i] {
                None => {
                    if chan.is_playing() {
                        chan.halt();
                    }
                }
                Some(_) => {
                    if !chan.is_playing() {
                        //diag!(env.queue, "reap: {} {}", aa.chan, aa.name);
                        self.free_channel(env, i as i32);
                        self.active[i] = None;
                    }
                }
            }
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
                    let s= Sound{conf: sound.clone(), chunk };
                    if self.sounds.insert(sound.name.clone(), s).is_some() {
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
                    let entry = Vocal{conf: vocal.clone(), chunk };
                    if self.vocals.insert(vocal.name.clone(), entry).is_some() {
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

    fn play_music(&mut self, env: &mut Env, cmd: &PlayMusic) {
        let Some(music) = self.music.get(&cmd.name) else { return };

        if let Some(playing) = &self.music_playing {
            if playing.name == cmd.name && cmd.no_restart {
                return
            }
        }

        let mut loops = cmd.loops;
        if loops == 0 {
            loops = -1;
        }
        if let Err(e) = music.play(loops) {
            fault!(env.queue, "cannot play music: {}", e);
        }

        let active = ActiveAudio{
            name: cmd.name.clone(),
            chan: -1,
            duck: 0,
            priority: 0,
            start_time: env.vars.elapsed,
            notify: cmd.notify
        };
        self.music_playing = Some(active);
    }

    fn play_sound(&mut self, env: &mut Env, cmd: &PlaySound)  {
        let Some(sound) = self.sounds.get(&cmd.name) else { return };

        let mut chan_num: i32 = 0;
        // See if the same sound is already playing
        for i in 1..self.active.len() {
            let Some(aa) = &self.active[i] else { continue };
            if aa.name == cmd.name {
                // If this sound is already active, do nothing if within the
                // debounce period
                let delta = env.vars.elapsed - aa.start_time;
                let debounce = sec_to_millis(sound.conf.debounce);
                if debounce > 0 && debounce > delta {
                    diag!(env.queue, "debounce: {}", cmd.name);
                    return
                }
                // Since the sound is already playing, this channel will
                // be reused.
                chan_num = i as i32;
                break;
            }
        }

        if chan_num > 0 {
            self.free_channel(env, chan_num);
            self.active[chan_num as usize] = None;
        } else {
            chan_num = self.find_available_channel(env, sound.conf.priority);
            if chan_num == 0 {
                diag!(env.queue, "no sound channels are available");
                return;
            }
            self.free_channel(env, chan_num);
            self.active[chan_num as usize] = None;
        }

        match sdl2::mixer::Channel(chan_num).play(&sound.chunk, cmd.loops) {
            Ok(_) => {
                let active = ActiveAudio{
                    name: cmd.name.clone(),
                    chan: chan_num,
                    duck: scale_volume(MAX_VOLUME, sound.conf.duck),
                    priority: sound.conf.priority,
                    start_time: env.vars.elapsed,
                    notify: cmd.notify
                };
                self.active[chan_num as usize] = Some(active);
            }
            Err(e) =>  {
                fault!(env.queue, "cannot play sound: {}", e);
                return
            }
        };
    }

    fn play_vocal(&mut self, env: &mut Env, cmd: &PlayVocal) {
        let Some(vocal) = self.vocals.get(&cmd.name) else { return };

        // Is a vocal already playing?
        if let Some(aa) = &self.active[0] {
            // Don't interrupt if it has a higher priority
            if aa.priority > vocal.conf.priority {
                diag!(env.queue, "a higher priority vocal is already playing");
                return
            }
            self.free_channel(env, 0);
            self.active[0] = None
        }

        match sdl2::mixer::Channel(0).play(&vocal.chunk, 0) {
            Ok(_) => {
                let active = ActiveAudio{
                    name: cmd.name.clone(),
                    chan: 0,
                    duck: scale_volume(MAX_VOLUME, vocal.conf.duck),
                    priority: vocal.conf.priority,
                    start_time: env.vars.elapsed,
                    notify: cmd.notify
                };
                self.active[0] = Some(active);
            },
            Err(e) => fault!(env.queue, "cannot play vocal: {}", e),
        }
    }

    fn silence(&mut self, env: &mut Env) {
        self.stop_music(env, &Name{name: "".to_string()});
        self.stop_vocal(env, &Name{name: "".to_string()});
    }

    fn stop_music(&mut self, _: &mut Env, cmd: &Name) {
        let Some(playing) = &self.music_playing else { return };
        if cmd.name != "" && cmd.name != *playing.name {
            return
        }
        sdl2::mixer::Music::halt();
        self.music_playing = None;
    }

    fn stop_vocal(&mut self, env: &mut Env, cmd: &Name) {
        let Some(playing) = &self.active[0] else { return };
        if cmd.name != "" && cmd.name != *playing.name {
            return
        }
        self.free_channel(env, 0);
        self.active[0] = None;
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
        self.reap_channels(env);
        match msg {
            Message::Halt => self.silence(env),
            Message::Init => self.init(env),
            Message::PlayMusic(m) => self.play_music(env, m),
            Message::PlaySound(m) => self.play_sound(env, m),
            Message::PlayVocal(m) => self.play_vocal(env, m),
            Message::Silence => self.silence(env),
            Message::StopMusic(m) => self.stop_music(env, m),
            Message::StopVocal(m) => self.stop_vocal(env, m),
            _ => ()
        }
        self.duck();
    }
}

fn scale_volume(vol: i32, scale: f64) -> i32 {
    (scale * vol as f64) as i32
}
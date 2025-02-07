use crate::prelude::*;

use serde::{Serialize, Deserialize};
use sdl2;
use sdl2::mixer;
use sdl2::mixer::Channel;
use std::collections::HashMap;
use std::sync::atomic::{AtomicBool, Ordering};

static MUSIC_FINISHED: AtomicBool = AtomicBool::new(false);

const MAX_VOLUME: i32 = 128;
const FREQUENCY: i32 = 44_100;
const FORMAT: u16 = mixer::AUDIO_S16LSB;
const CHUNK_SIZE: i32 = 1024;

#[cfg(feature = "debug_audio")]
macro_rules! debug_audio {
    ($q:expr, $($args:expr),+) => {
        $q.post(Message::Note(Note{kind: NoteKind::Diag, message: format!($($args),+)}))
    }
}
#[cfg(not(feature = "debug_audio"))]
macro_rules! debug_audio {
    ($q:expr, $($args:expr),+) => {
    }
}

#[derive(Serialize, Deserialize, Debug, Default, Clone)]
#[serde(rename_all = "snake_case")]
pub enum Output {
    #[default]
    Mono,
    Stereo,
}

#[derive(Serialize, Deserialize, Debug, Default, Clone)]
#[serde(deny_unknown_fields)]
pub struct AudioConfig {
    pub output: Output,
    pub channels: i32,
    pub with_mp3: bool,
    pub with_ogg: bool,
}

struct ActiveChan {
    name: String,
    chan: i32,
    duck: i32,
    start_time: i64,
    priority: i32,
    notify: bool,
}

struct Sound {
    def: SoundDef,
    chunk: mixer::Chunk,
}

struct Vocal {
    def: VocalDef,
    chunk: mixer::Chunk
}

pub struct Audio<'a> {
    music: HashMap<String,mixer::Music<'a>>,
    sounds: HashMap<String,Sound>,
    vocals: HashMap<String,Vocal>,

    music_playing: Option<ActiveChan>,
    active: Vec<Option<ActiveChan>>,
}

impl Audio<'_> {
    pub fn new(conf: &AudioConfig) -> Self {
        let output_chans = match conf.output {
            Output::Mono => 1,
            Output::Stereo => 2
        };

        if let Err(reason) = mixer::open_audio(FREQUENCY, FORMAT, output_chans, CHUNK_SIZE) {
            panic!("unable to open mixer: {}", reason);
        }

        let mut flags = mixer::InitFlag::empty();
        if conf.with_mp3 {
            flags |= mixer::InitFlag::MP3;
        }
        if conf.with_ogg {
            flags |= mixer::InitFlag::OGG;
        }
        if let Err(reason) = mixer::init(flags) {
            panic!("unable to initialize mixer: {}", reason);
        }

        // Config specifies the number of sound channels only. Another one
        // is then reserved for vocals hence the plus 1.
        let sound_chans = conf.channels + 1;
        let mut channels: Vec<Option<ActiveChan>> = Vec::new();
        sdl2::mixer::allocate_channels(sound_chans);
        for _ in 0..sound_chans {
            channels.push(None);
        }

        Self {
            music: HashMap::new(),
            sounds: HashMap::new(),
            vocals: HashMap::new(),
            music_playing: None,
            active: channels
        }
    }

    fn find_available_channel(&self, env: &mut Env, priority: i32) -> Option<i32> {
        let mut opt_candidate: Option<&ActiveChan> = None;

        // First check if there are any open slots
        for i in 1..self.active.len() {
            match &self.active[i] {
                None => return Some(i as i32),
                // Otherwise, see if this is a candidate that could be
                // freed.
                Some(aa) => {
                    if aa.priority > priority {
                        continue
                    }
                    match opt_candidate {
                        // If it has a lower priority and we don't have a
                        // candidate yet, this is it.
                        None => opt_candidate = Some(aa),
                        // If we do have a candidate already, take it if
                        // it has a lower priority or an earlier start
                        // time with the same priority.
                        Some(cand) => {
                            if aa.priority < cand.priority {
                                opt_candidate = Some(aa);
                                continue;
                            }
                            if aa.priority == cand.priority && aa.start_time < cand.start_time {
                                opt_candidate = Some(aa);
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
                diag!(env.queue, "a sound channel was interrupted to satisfy request");
                Some(aa.chan)
            },
            None => None
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
                Some(_aa) => {
                    if !chan.is_playing() {
                        debug_audio!(env.queue, "channel {} reap: {}", _aa.chan, _aa.name);
                        self.free_channel(env, i as i32);
                        self.active[i] = None;
                    }
                }
            }
        }
    }

    fn init(&mut self, env: &mut Env) {
        for (name, music) in &env.conf.music {
            let path = env.conf.data_dir.join(&music.path);
            match mixer::Music::from_file(&path) {
                Err(e) => fault!(env.queue, "unable to load music '{}': {}", &name, &e),
                Ok(m) => {
                    self.music.insert(name.clone(), m);
                }
            }
        }

        for (name, sound) in &env.conf.sounds {
            let path = env.conf.data_dir.join(&sound.path);
            match mixer::Chunk::from_file(&path) {
                Err(e) => fault!(env.queue, "unable to load sound '{}': {}", &name, &e),
                Ok(chunk) => {
                    let s= Sound{def: sound.clone(), chunk };
                    self.sounds.insert(name.clone(), s);
                }
            };
        }

        for (name, vocal) in &env.conf.vocals {
            let path = env.conf.data_dir.join(&vocal.path);
            match mixer::Chunk::from_file(&path) {
                Err(e) => fault!(env.queue, "unable to load vocal '{}': {}", &name, &e),
                Ok(chunk) => {
                    let entry = Vocal{def: vocal.clone(), chunk };
                    self.vocals.insert(name.clone(), entry);
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

        let active = ActiveChan{
            name: cmd.name.clone(),
            chan: -1,
            duck: 0,
            priority: 0,
            start_time: env.vars["elapsed"].as_int(),
            notify: cmd.notify
        };
        self.music_playing = Some(active);
    }

    fn play_sound(&mut self, env: &mut Env, cmd: &PlaySound)  {
        let Some(sound) = self.sounds.get(&cmd.name) else { return };

        let mut opt_chan_num: Option<i32> = None;
        // See if the same sound is already playing
        for i in 1..self.active.len() {
            let Some(aa) = &self.active[i] else { continue };
            if aa.name == cmd.name {
                // If this sound is already active, do nothing if within the
                // debounce period
                let delta = env.vars["elapsed"].as_int() - aa.start_time;
                let debounce = sec_to_millis(sound.def.debounce);
                if debounce > 0 && debounce > delta {
                    diag!(env.queue, "debounce: {}", cmd.name);
                    return
                }
                // Since the sound is already playing, this channel will
                // be reused.
                opt_chan_num = Some(i as i32);
                break;
            }
        }

        if opt_chan_num.is_none() {
            opt_chan_num = match self.find_available_channel(env, sound.def.priority) {
                Some(chan_num) => {
                    self.free_channel(env, chan_num);
                    self.active[chan_num as usize] = None;
                    Some(chan_num)
                }
                None => {
                    diag!(env.queue, "no sound channels are available");
                    return;
                }
            }
        }

        // There must be a channel number assigned by this point
        let chan_num = opt_chan_num.unwrap();
        self.free_channel(env, chan_num);
        self.active[chan_num as usize] = None;

        debug_audio!(env.queue, "channel {} allocate: {}", chan_num, sound.conf.name);
        match sdl2::mixer::Channel(chan_num).play(&sound.chunk, cmd.loops) {
            Ok(_) => {
                let active = ActiveChan{
                    name: cmd.name.clone(),
                    chan: chan_num,
                    duck: scale_volume(MAX_VOLUME, sound.def.duck),
                    priority: sound.def.priority,
                    start_time: env.vars["elapsed"].as_int(),
                    notify: cmd.notify
                };
                self.active[chan_num as usize] = Some(active);
            }
            Err(e) =>  {
                fault!(env.queue, "cannot play sound: {}", e);
            }
        };
    }

    fn play_vocal(&mut self, env: &mut Env, cmd: &PlayVocal) {
        let Some(vocal) = self.vocals.get(&cmd.name) else { return };

        // Is a vocal already playing?
        if let Some(aa) = &self.active[0] {
            // Don't interrupt if it has a higher priority
            if aa.priority > vocal.def.priority {
                diag!(env.queue, "a higher priority vocal is already playing");
                return
            }
            self.free_channel(env, 0);
            self.active[0] = None
        }

        match sdl2::mixer::Channel(0).play(&vocal.chunk, 0) {
            Ok(_) => {
                let active = ActiveChan{
                    name: cmd.name.clone(),
                    chan: 0,
                    duck: scale_volume(MAX_VOLUME, vocal.def.duck),
                    priority: vocal.def.priority,
                    start_time: env.vars["elapsed"].as_int(),
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
        if !cmd.name.is_empty() && cmd.name != *playing.name {
            return
        }
        sdl2::mixer::Music::halt();
        self.music_playing = None;
    }

    fn stop_vocal(&mut self, env: &mut Env, cmd: &Name) {
        let Some(playing) = &self.active[0] else { return };
        if !cmd.name.is_empty() && cmd.name != *playing.name {
            return
        }
        self.free_channel(env, 0);
        self.active[0] = None;
    }

    pub fn process(&mut self, env: &mut Env, msg: &Message) {
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
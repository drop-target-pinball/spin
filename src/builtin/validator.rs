use crate::prelude::*;
use std::collections::HashSet;

pub struct Validator {
    music: HashSet<String>,
    sounds: HashSet<String>,
    vocals: HashSet<String>,
}

impl Validator {
    pub fn new() -> Self {
        Self{
            music: HashSet::new(),
            sounds: HashSet::new(),
            vocals: HashSet::new(),
        }
    }

    fn init(&mut self, env: &Env) {
        for v in &env.conf.music  { self.music.insert(v.name.clone()); }
        for v in &env.conf.sounds { self.sounds.insert(v.name.clone()); }
        for v in &env.conf.vocals { self.vocals.insert(v.name.clone()); }
    }

    fn validate_music(&self, env: &mut Env, msg: &PlayMusic) {
        if !self.music.contains(&msg.name) {
            fault!(env.queue, "no such music: {}", msg.name);
        }
    }

    fn validate_sound(&self, env: &mut Env, msg: &PlaySound) {
        if !self.sounds.contains(&msg.name) {
            fault!(env.queue, "no such sound: {}", msg.name);
        }
    }

    fn validate_vocal(&self, env: &mut Env, msg: &PlayVocal) {
        if !self.vocals.contains(&msg.name) {
            fault!(env.queue, "no such vocal: {}", msg.name);
        }
    }
}

impl Device for Validator {
    fn process(&mut self, env: &mut Env, msg: &Message) {
        match msg {
            Message::Init => self.init(env),
            Message::PlayMusic(m) => self.validate_music(env, m),
            Message::PlaySound(m) => self.validate_sound(env, m),
            Message::PlayVocal(m) => self.validate_vocal(env, m),
            _ => ()
        }
    }
}
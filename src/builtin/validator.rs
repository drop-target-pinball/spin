use crate::prelude::*;

#[derive(Default)]
pub struct Validator {
}

impl Validator {
    fn validate_music(&self, env: &mut Env, msg: &PlayMusic) {
        if !env.conf.music.contains_key(&msg.name) {
            fault!(env.queue, "no such music: {}", msg.name);
        }
    }

    fn validate_sound(&self, env: &mut Env, msg: &PlaySound) {
        if !env.conf.sounds.contains_key(&msg.name) {
            fault!(env.queue, "no such sound: {}", msg.name);
        }
    }

    fn validate_vocal(&self, env: &mut Env, msg: &PlayVocal) {
        if !env.conf.vocals.contains_key(&msg.name) {
            fault!(env.queue, "no such vocal: {}", msg.name);
        }
    }
}

impl Device for Validator {
    fn process(&mut self, env: &mut Env, msg: &Message) {
        match msg {
            Message::PlayMusic(m) => self.validate_music(env, m),
            Message::PlaySound(m) => self.validate_sound(env, m),
            Message::PlayVocal(m) => self.validate_vocal(env, m),
            _ => ()
        }
    }
}
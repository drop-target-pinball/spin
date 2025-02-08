use crate::prelude::*;

#[derive(Default)]
pub struct Validator {
}

impl Validator {
    fn validate_music(&self, s: &mut State, msg: &PlayMusic) {
        if !s.conf.music.contains_key(&msg.name) {
            fault!(s.queue, "no such music: {}", msg.name);
        }
    }

    fn validate_sound(&self, s: &mut State, msg: &PlaySound) {
        if !s.conf.sounds.contains_key(&msg.name) {
            fault!(s.queue, "no such sound: {}", msg.name);
        }
    }

    fn validate_vocal(&self, s: &mut State, msg: &PlayVocal) {
        if !s.conf.vocals.contains_key(&msg.name) {
            fault!(s.queue, "no such vocal: {}", msg.name);
        }
    }
}

impl Device for Validator {
    fn process(&mut self, s: &mut State, msg: &Message) {
        match msg {
            Message::PlayMusic(m) => self.validate_music(s, m),
            Message::PlaySound(m) => self.validate_sound(s, m),
            Message::PlayVocal(m) => self.validate_vocal(s, m),
            _ => ()
        }
    }
}
use crate::prelude::*;
use sdl2::keyboard::{Keycode, Mod};
use sdl2::event::{Event, WindowEvent};
use std::collections::HashMap;


#[derive(PartialEq, Eq, Copy, Clone, Debug, Hash)]
struct Key {
    pub code: Keycode,
    pub mods: Mod,
}

pub struct Input {
    keys: HashMap<Key, KeyDef>
}

impl Input {
    pub fn new(conf: &AppConfig) -> Input {
        let mut keys = HashMap::new();
        for def in &conf.keyboard {
            let code = match Keycode::from_name(&def.key) {
                Some(c) => c,
                None => panic!("unknown keycode: {}", &def.key),
            };
            let mut mods = Mod::empty();
            if def.left_shift {
                mods.insert(Mod::LSHIFTMOD);
            }
            let key =Key{code, mods};
            if keys.insert(Key{code, mods}, def.clone()).is_some() {
                panic!("duplicate key def: {:?}", key);
            }
        }
        Input { keys }
    }

    pub fn key(&self, s: &mut State, opt_code: &Option<Keycode>, keymod: &Mod, repeat: bool, down: bool) {
        if repeat { return };
        let Some(code) = opt_code else { return };
        let Some(def) = self.keys.get(&Key{code: *code, mods: *keymod}) else { return };

        let opt_msg = if down { &def.down } else { &def.up };
        if let Some(msg) = &opt_msg {
            s.queue.post(msg.clone());
        }
    }

    pub fn event(&self, s: &mut State, evt: &Event) {
        match evt {
            Event::KeyDown { timestamp: _, window_id: _, keycode, scancode: _, keymod, repeat } => {
                self.key(s, keycode, keymod, *repeat, true);
            },
            Event::KeyUp { timestamp: _, window_id: _, keycode, scancode: _, keymod, repeat } => {
                self.key(s, keycode, keymod, *repeat, false);
            }
            Event::Window { timestamp: _, window_id: _, win_event: WindowEvent::Close } => {
                s.queue.post(Message::Shutdown);
                if s.runtime.is_auto_test() {
                    std::process::exit(1);
                }
            },
            Event::Quit { .. } => {
                s.queue.post(Message::Shutdown);
                if s.runtime.is_auto_test() {
                    std::process::exit(1);
                }
            }
            _ => ()
        }
    }

}
use crate::prelude::*;

#[derive(Default)]
pub struct Store {
}

impl Store {
    pub fn new() -> Self {
        Self{}
    }

    fn set(&self, s: &mut State, msg: &Vars) {
        for (name, value) in &msg.vars {
            match vars::set(&mut s.vars, &msg.ns, &name, &value) {
                Ok(msg) => s.queue.post(Message::Updated(msg)),
                Err(e) => fault!(s.queue, "{}", e),
            }
        }
    }
}

impl Device for Store {
    fn init(&mut self, s: &mut State, _: &mut render::State) {
        for (name, v) in &s.conf.vars {
            vars::define(&mut s.queue, &mut s.vars, &s.conf.namespaces, &name, &v.kind);
        }
    }

    fn process(&mut self, s: &mut State, msg: &Message) {
        match msg {
            Message::Set(m) => self.set(s, m),
            _ => (),
        }
    }

    fn render(&mut self, _: &mut render::State) {}
    fn present(&mut self, _: &render::State) {}
}
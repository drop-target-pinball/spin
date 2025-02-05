use crate::prelude::*;

#[derive(Default)]
pub struct Store {
}

impl Store {
    pub fn new() -> Self {
        Self{}
    }

    fn init(&mut self, env: &mut Env) {
        for v in &env.conf.vars {
            vars::define(&mut env.queue, env.vars, &env.conf.namespaces, &v.name, &v.kind);
        }
    }

    fn set(&self, env: &mut Env, msg: &Vars) {
        for (name, value) in &msg.vars {
            vars::set(env, &name, &value);
        }
    }
}

impl Device for Store {
    fn process(&mut self, env: &mut Env, msg: &Message) {
        match msg {
            Message::Init => self.init(env),
            Message::Set(m) => self.set(env, m),
            _ => (),
        }
    }
}
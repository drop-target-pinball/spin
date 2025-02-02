use crate::prelude::*;

pub struct Store {
}

impl Store {
    pub fn new() -> Self {
        Self{}
    }

    fn init(&mut self, env: &mut Env) {
        for v in &env.conf.vars {
            env.vars.define(&mut env.queue, &v.name, &v.value);
        }
    }

    fn set_var(&self, env: &mut Env, msg: &Var) {
        env.vars.set(&mut env.queue, &msg.name, &msg.value);
    }
}

impl Device for Store {
    fn process(&mut self, env: &mut Env, msg: &Message) {
        match msg {
            Message::Init => self.init(env),
            Message::SetVar(m) => self.set_var(env, m),
            _ => (),
        }
    }
}
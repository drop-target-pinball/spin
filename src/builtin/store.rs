use crate::prelude::*;

#[derive(Default)]
pub struct Store {
}

impl Store {
    pub fn new() -> Self {
        Self{}
    }

    fn init(&mut self, env: &mut Env) {
        let mut spaces = vars::Namespaces::new();
        for ns in &env.conf.namespaces {
            spaces.insert(ns.name.clone(), ns.vars.clone());
        }
        for v in &env.conf.vars {
            vars::define(&mut env.queue, env.vars, &spaces, &v.name, &v.kind);
        }
    }

    fn set_var(&self, env: &mut Env, msg: &Var) {
        vars::set(env, &msg.name, &msg.value);
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
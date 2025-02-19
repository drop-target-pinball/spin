use crate::prelude::*;
use crate::Result;
use std::collections::HashMap;

struct Timer {
    def: TimerDef,
    last_update: i64,
}

#[derive(Default)]
pub struct Store {
    timers: HashMap<String, Timer>
}

impl Store {
    pub fn new() -> Self {
        Self{
            timers: HashMap::new(),
        }
    }

    fn set_var(&self, s: &mut State, ns: &Option<String>, name: &str, value: &vars::Value) {
        match vars::set(&mut s.vars, &ns, &name, &value) {
            Ok(msg) => s.queue.post(Message::Updated(msg)),
            Err(e) => fault!(s.queue, "{}", e),
        }
    }

    fn set_vars(&self, s: &mut State, msg: &Vars) {
        for (name, value) in &msg.vars {
            self.set_var(s, &msg.ns, name, value);
        }
    }

    fn start_timer(&mut self, s: &mut State, msg: &Name) {
        let Some(def) = s.conf.timers.get(&msg.name) else { return };
        let elapsed = s.vars["elapsed"].as_i64();
        let timer = Timer{
            def: def.clone(),
            last_update: elapsed,
        };
        self.timers.insert(msg.name.clone(), timer);
        self.set_var(s, &None, &def.var.clone(), &vars::Value::Int(def.start));
    }

    fn stop_timer(&mut self, msg: &Name) {
        self.timers.remove(&msg.name);
    }

    fn reset_timer(&mut self, s: &mut State, msg: &Name) {
        let Some(def) = s.conf.timers.get(&msg.name) else { return };
        self.set_var(s, &None, &def.var.clone(), &vars::Value::Int(def.start));
    }

    fn halt(&mut self) {
        self.timers.clear();
    }

    fn tick(&mut self, s: &mut State) {
        let now = s.vars["elapsed"].as_i64();

        let mut updates: Vec<(String, vars::Value)> = Vec::new();
        let mut expired: Vec<String> = Vec::new();

        for (name, timer) in &mut self.timers {
            let tick_ms = (timer.def.tick * 1000.0) as i64;
            let mut delta = now - timer.last_update;
            let orig =  s.vars.get(&timer.def.var).unwrap().as_i64();
            let mut curr = orig;
            let dir = if timer.def.step < 0 { -1 } else { 1 };
            while delta > tick_ms {
                curr += timer.def.step;
                delta -= tick_ms;
                if dir == 1 && curr > timer.def.end {
                    curr = timer.def.end;
                }
                if dir == -1 && curr < timer.def.end {
                    curr = timer.def.end;
                }
            }
            if curr != orig {
                updates.push((timer.def.var.clone(), vars::Value::Int(curr)));
                timer.last_update = now;
            }
            if curr == timer.def.end {
                expired.push(name.clone());
            }
        }
        for (name, val) in updates {
            self.set_var(s, &None, &name, &val);
        }
        for name in expired {
            self.timers.remove(&name);
            s.queue.post(Message::TimerExpired(Name{name: name.clone()}))
        }
    }
}

impl Device for Store {
    fn init(&mut self, s: &mut State, _: &mut render::State) {
        for (name, v) in &s.conf.vars {
            vars::define(&mut s.queue, &mut s.vars, &s.conf.namespaces, &name, &v.kind);
        }
    }

    fn poll(&mut self, _: &mut State) -> Result<()> { Ok(()) }

    fn process(&mut self, s: &mut State, msg: &Message) {
        match msg {
            Message::Halt => self.halt(),
            Message::ResetTimer(m) => self.reset_timer(s, m),
            Message::Set(m) => self.set_vars(s, m),
            Message::StartTimer(m) => self.start_timer(s, m),
            Message::StopTimer(m) => self.stop_timer(m),
            Message::Tick => self.tick(s),
            _ => (),
        }
    }

    fn render(&mut self, _: &mut render::State) {}
    fn present(&mut self, _: &render::State) {}
}
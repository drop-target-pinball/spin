use crate::prelude::*;
use crate::Result;
use std::collections::HashMap;

struct Timer {
    def: TimerDef,
    last_update: i64,
    expire_at: Option<i64>,
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

    fn set_var(&self, s: &mut State, ns: Namespace, name: &str, value: &vars::Value) {
        let result = match ns {
            Namespace::Player(i) => vars::set(&mut s.players[i-1], ns, name, value),
            Namespace::Setting => vars::set(&mut s.settings, ns, name, value),
            Namespace::Var => vars::set(&mut s.vars, ns, name, value),
        };
        match result {
            Ok(maybe_msg) => {
                if let Some(msg) = maybe_msg {
                    s.queue.post(Message::Updated(msg));
                }
            }
            Err(e) => fault!(s.queue, "{}", e),
        }
    }

    fn start_timer(&mut self, s: &mut State, msg: &Name) {
        let Some(def) = s.conf.timers.get(&msg.name) else { return };
        let timer = Timer{
            def: def.clone(),
            last_update: s.elapsed,
            expire_at: None,
        };
        self.timers.insert(msg.name.clone(), timer);
        self.set_var(s, Namespace::Var, &def.var.clone(), &vars::Value::Int(def.start));
    }

    fn stop_timer(&mut self, msg: &Name) {
        self.timers.remove(&msg.name);
    }

    fn reset_timer(&mut self, s: &mut State, msg: &Name) {
        let Some(def) = s.conf.timers.get(&msg.name) else { return };
        self.set_var(s, Namespace::Var, &def.var.clone(), &vars::Value::Int(def.start));
    }

    fn kill_group(&mut self, s: &mut State, msg: &Name) {
        let groups = timers_for_group(&s.conf, &msg.name);
        self.timers.retain(|k, _| !groups.contains(k));
    }

    fn halt(&mut self) {
        self.timers.clear();
    }

    fn reset(&mut self, s: &mut State) {
        self.halt();
        s.vars = HashMap::new();
        for (name, v) in &s.conf.vars {
            vars::define(&mut s.queue, &mut s.vars, &name, &v.kind);
        }
        s.settings = HashMap::new();
        for (name, v) in &s.conf.settings {
            vars::define(&mut s.queue, &mut s.settings, &name, &v.kind);
        }
        let max_players = s.conf.max_players;
        for _ in 0..max_players {
            let mut vars = vars::Vars::new();
            for (name, v) in &s.conf.player {
                vars::define(&mut s.queue, &mut vars, name, &v.kind);
            }
            s.players.push(vars);
        }
    }

    fn tick(&mut self, s: &mut State) {
        let now = s.elapsed;

        let mut updates: Vec<(String, vars::Value)> = Vec::new();
        let mut expired: Vec<String> = Vec::new();

        for (name, timer) in &mut self.timers {
            if let Some(expire_at) = timer.expire_at {
                if now >= expire_at {
                    expired.push(name.clone());
                }
                continue;
            }
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
                if let Some(delay) = timer.def.expire_delay {
                    timer.expire_at = Some(now + (delay * 1000.0) as i64);
                } else {
                    expired.push(name.clone());
                }
            }
        }
        for (name, val) in updates {
            self.set_var(s, Namespace::Var, &name, &val);
        }
        for name in expired {
            self.timers.remove(&name);
            s.queue.post(Message::TimerExpired(Name{name: name.clone()}))
        }
    }
}

impl Device for Store {
    fn init(&mut self, s: &mut State, _: &mut render::State) {
        self.reset(s);
    }

    fn poll(&mut self, _: &mut State) -> Result<()> { Ok(()) }

    fn process(&mut self, s: &mut State, msg: &Message) {
        match msg {
            Message::Halt => self.halt(),
            Message::KillGroup(m) => self.kill_group(s, m),
            Message::Reset => self.reset(s),
            Message::ResetTimer(m) => self.reset_timer(s, m),
            Message::Set(m) => self.set_var(s, m.namespace, &m.name, &m.value),
            Message::StartTimer(m) => self.start_timer(s, m),
            Message::StopTimer(m) => self.stop_timer(m),
            Message::Tick => self.tick(s),
            _ => (),
        }
    }

    fn render(&mut self, _: &mut State,  _: &mut render::State) {}
    fn present(&mut self, _: &mut State, _: &render::State) {}
}

fn timers_for_group(conf: &AppConfig, kill_group: &str) -> Vec<String> {
    let mut timer_names = Vec::new();
    for (name, def) in &conf.timers {
        if let Some(group) = &def.group {
            if group == kill_group {
                timer_names.push(name.clone());
            }
        }
    }
    for (rg_name, def) in &conf.run_groups {
        if let Some(parent) = &def.parent {
            if parent == kill_group {
                timer_names.extend(timers_for_group(conf, &rg_name));
            }
        }
    }
    timer_names
}
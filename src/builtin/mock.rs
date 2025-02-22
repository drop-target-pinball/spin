use crate::prelude::*;
use crate::Result;
use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct LockDef {
    pub switches: Vec<String>,
    pub eject: String,
    pub eject_to: Option<String>,
    #[serde(default)]
    pub default: usize,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct Config {
    #[serde(default)]
    pub locks: Vec<LockDef>
}

pub struct Lock {
    pub def: LockDef,
    pub count: usize,
}

pub struct Device {
    locks: Vec<Lock>
}

impl Device {
    pub fn new(conf: &Config) -> Device {
        let mut locks = Vec::new();
        for def in &conf.locks {
            locks.push(Lock{
                def: def.clone(),
                count: def.default,
            });
        }
        Device { locks }
    }

    fn lock_arrival(&mut self, msg: &SwitchUpdated) {
        // Only interested in balls arriving to the ball lock.
        if !msg.active {
            return
        }

        // Is this switch a member of a lock? Return if not
        let lock = 'lock: {
            for lock in &mut self.locks {
                if lock.def.switches.contains(&msg.name) {
                    break 'lock lock
                }
            }
            return
        };

        // What position in the ball queue is this switch?
        let idx = lock.def.switches
            .iter()
            .position(|e| e == &msg.name)
            .expect("must exist, found in previous search");

        // Only interested If this is the top-most empty ball slot
        if idx != lock.count {
            return
        }

        // A ball has been locked
        lock.count = lock.count + 1
    }

    fn lock_departure(&mut self, s: &mut State, msg: &PulseDriver) {
        // Is this an eject solenoid for a lock? Return if not
        let Some(lock) = self.locks
            .iter_mut()
            .find(|e| e.def.eject == msg.name)
            else { return };

        // Return if the lock is empty
        if lock.count == 0 {
            return
        }

        // The top-most switch is going to open up
        lock.count -= 1;
        let su_msg = SwitchUpdated{
            name: lock.def.switches[lock.count].clone(),
            active: false,
        };
        s.queue.post(Message::SwitchUpdated(su_msg));
    }

    fn switch_updated(&mut self, s: &mut State, msg: &SwitchUpdated) {
        let Some(sw) = s.switches.get_mut(&msg.name) else { return };
        sw.active = msg.active;
        sw.last_update = s.elapsed;
        self.lock_arrival(msg);
    }

    fn pulse_driver(&mut self, s: &mut State, msg: &PulseDriver) {
        self.lock_departure(s, msg);
    }
}

impl crate::Device for Device {
    fn init(&mut self, s: &mut State, _: &mut render::State) {
        // Post switch active events the ball starting positions
        for lock in &self.locks {
            if lock.count == 0 {
                continue
            }
            for i in 0..lock.count {
                let su_msg = SwitchUpdated{
                    name: lock.def.switches[i].clone(),
                    active: true,
                };
                s.queue.post(Message::SwitchUpdated(su_msg))
            }
        }
    }

    fn poll(&mut self, _: &mut State) -> Result<()> { Ok(()) }
    fn process(&mut self, s: &mut State, msg: &Message) {
        match msg {
            Message::SwitchUpdated(m) => self.switch_updated(s, m),
            Message::PulseDriver(m) => self.pulse_driver(s, m),
            _ => (),
        }
    }
    fn render(&mut self, _: &mut State, _: &mut render::State) {}
    fn present(&mut self, _: &mut State, _: &render::State) {}
}
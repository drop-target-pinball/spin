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
    conf: Config,
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
        Device {
            conf: conf.clone(),
            locks,
        }
    }

    fn lock_arrival(&mut self, s: &mut State, msg: &SwitchUpdated) {
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

        // Ignore if the lock is full
        if lock.def.switches.len() == lock.count {
            return
        }

        // What position in the ball queue is this switch?
        let idx = lock.def.switches
            .iter()
            .position(|e| e == &msg.name)
            .expect("must exist, found in previous search");

        // See if the ball should slide to the next open slot
        if idx > lock.count {
            let su_msg: SwitchUpdated = SwitchUpdated{
                name: lock.def.switches[idx].clone(),
                active: false,
            };
            s.queue.post(Message::SwitchUpdated(su_msg));

            let su_msg: SwitchUpdated = SwitchUpdated{
                name: lock.def.switches[idx-1].clone(),
                active: true,
            };
            s.queue.post(Message::SwitchUpdated(su_msg));
        } else {
            // Otherwise, a ball has been locked
            lock.count = lock.count + 1
        }
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

        // If the ball is going to be ejected to a another switch, check to
        // see if that is available.
        let maybe_eject_to = &lock.def.eject_to;
        if let Some(eject_to) = maybe_eject_to {
            let sw = &s.switches[eject_to];
            // Switch is occupied, do nothing
            if sw.active {
                return
            }
        }

        // The top-most switch is going to open up
        lock.count -= 1;
        let su_msg: SwitchUpdated = SwitchUpdated{
            name: lock.def.switches[lock.count].clone(),
            active: false,
        };
        s.queue.post(Message::SwitchUpdated(su_msg));

        // Eject the ball to the other switch if defined
        if let Some(eject_to) = maybe_eject_to {
            let su_msg = SwitchUpdated{
                name: eject_to.clone(),
                active: true,
            };
            s.queue.post(Message::SwitchUpdated(su_msg));
        }
    }

    fn switch_updated(&mut self, s: &mut State, msg: &SwitchUpdated) {
        let Some(sw) = s.switches.get_mut(&msg.name) else { return };
        sw.active = msg.active;
        sw.last_update = s.elapsed;
        self.lock_arrival(s, msg);
    }

    fn pulse_driver(&mut self, s: &mut State, msg: &PulseDriver) {
        self.lock_departure(s, msg);
    }

    fn reset(&mut self, s: &mut State) {
        for sw in s.switches.values_mut() {
            sw.active = false;
            sw.last_update = s.elapsed;
        }

        self.locks.clear();
        for def in &self.conf.locks {
            self.locks.push(Lock{
                def: def.clone(),
                count: def.default,
            });
        }

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
}

impl crate::Device for Device {
    fn init(&mut self, s: &mut State, _: &mut render::State) {
        self.reset(s);
    }

    fn poll(&mut self, _: &mut State) -> Result<()> { Ok(()) }
    fn process(&mut self, s: &mut State, msg: &Message) {
        match msg {
            Message::PulseDriver(m) => self.pulse_driver(s, m),
            Message::Reset => self.reset(s),
            Message::SwitchUpdated(m) => self.switch_updated(s, m),
            _ => (),
        }
    }
    fn render(&mut self, _: &mut State, _: &mut render::State) {}
    fn present(&mut self, _: &mut State, _: &render::State) {}
}
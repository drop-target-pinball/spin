use std::time;

pub struct Vars {
    pub uptime: time::Instant,
    pub now: time::Instant,
}

impl Default for Vars {
    fn default() -> Vars {
        let now = time::Instant::now();
        Vars {
            uptime: now,
            now,
        }
    }
}
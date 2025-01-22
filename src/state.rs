use std::time;

pub struct Vars {
    pub uptime: time::Instant,
    pub now: time::Instant,
}

impl Vars {
    pub fn new() -> Vars {
        let now = time::Instant::now();
        Vars {
            uptime: now,
            now,
        }
    }


}
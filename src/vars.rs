use std::time;

pub struct Vars {
    pub elapsed: time::Duration
}

impl Default for Vars {
    fn default() -> Vars {
        Vars {
            elapsed: time::Duration::ZERO,
        }
    }
}
use std::time;

pub struct State {
    pub uptime: time::Instant,
    pub now: time::Instant,
}

impl State {
    pub fn new() -> State {
        let now = time::Instant::now();
        State {
            uptime: now,
            now,
        }
    }


}
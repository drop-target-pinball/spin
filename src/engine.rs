use crate::prelude::*;
use std::time;

pub struct Env {
    pub conf: Config,
    pub vars: Vars,
}

impl Env  {
    pub fn new(conf: Config, vars: Vars) -> Self {
        Self {conf, vars}
    }
}

impl Default for Env {
    fn default() -> Self {
        Self {
            conf: Config::new(RunMode::Develop),
            vars: Vars::new(),
        }
    }
}

pub trait Device {
    fn process(&mut self, e: &mut Env, q: &mut Queue, msg: &Message) -> bool;
}

pub struct Engine {
    pub env: Env,
    pub queue: Queue,
    devices: Vec<Box<dyn Device>>,
}


impl Engine {
    pub fn new(conf: Config) -> Self {
        let env = Env::new(conf, Vars::new());
        Engine {
            env,
            queue: Queue::new(),
            devices: Vec::new(),
        }
    }

    pub fn add_device(&mut self, d: Box<dyn Device>) {
        self.devices.push(d);
    }

    pub fn init(&mut self) {
        self.queue.push(Message::Init);
        self.process_queue();
    }

    pub fn tick(&mut self) {
        self.env.vars.now = time::Instant::now();
        self.process_queue();
    }

    fn process_queue(&mut self) {
        loop {
            // As messages are being processed, the systems may want to
            // generate additional messages. Collect those here.
            let mut out_queue = Queue::new();

            // Send each message in the queue to every system for processing.
            match self.queue.pop() {
                None => break,
                Some(msg) => {
                    let mut handled = false;
                    for dev in &mut self.devices {
                            handled |= dev.process(&mut self.env, &mut out_queue, &msg);
                    }
                    handled |= Self::process(&mut self.env, &mut out_queue, &msg);
                    if !handled && self.env.conf.is_develop() {
                        panic!("message not handled: {:?}", &msg);
                    }
                },
            }

            // If systems generated additional messages, place these on the
            // main queue.
            while let Some(msg) = out_queue.pop() {
                self.queue.push(msg);
            }

            // Continue processing until the main queue is empty
            if self.queue.is_empty() {
                break
            }
        }
    }

    fn process(e: &mut Env, queue: &mut Queue, msg: &Message) -> bool {
        match msg {
            Message::Note(n) => {
                if e.conf.is_develop() && n.kind == NoteKind::Fault {
                    panic!("fault: {}", n.message);
                }
            },
            _ => return false
        }
        true
    }

}

impl Default for Engine {
    fn default() -> Self {
        Engine::new(Config::default())
    }
}


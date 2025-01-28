use crate::prelude::*;
use std::sync::atomic::{AtomicBool, Ordering};
use std::sync::mpsc::Receiver;
use std::sync::Arc;
use std::{
    sync::mpsc::{self, TryRecvError},
    thread,
    time::{self, Duration},
};

pub struct Env {
    pub conf: Config,
    pub vars: Vars,
}

impl Env {
    pub fn new(conf: Config, vars: Vars) -> Self {
        Self { conf, vars }
    }
}

impl Default for Env {
    fn default() -> Self {
        Self {
            conf: Config::new(RunMode::Develop),
            vars: Vars::default(),
        }
    }
}

pub trait Device {
    fn process(&mut self, e: &mut Env, q: &mut Queue, msg: &Message) -> bool;
}

pub struct Engine<'e> {
    pub env: Env,
    pub queue: Queue,
    pub rx: Receiver<Message>,
    devices: Vec<&'e mut dyn Device>,
}

impl<'e> Engine<'e> {
    pub fn new(conf: Config) -> Self {
        let env = Env::new(conf, Vars::default());
        let (tx, rx) = mpsc::channel();
        Engine {
            env,
            queue: Queue::new(tx),
            rx,
            devices: Vec::new(),
        }
    }

    pub fn add_device(&mut self, d: &'e mut dyn Device) {
        self.devices.push(d);
    }

    pub fn queue(&self) -> Queue {
        self.queue.clone()
    }

    pub fn init(&mut self) {
        self.queue.post(Message::Init);
        self.process_queue();
        info!(self.queue, "ready");
    }

    pub fn tick(&mut self) {
        self.env.vars.now = time::Instant::now();
        self.process_queue();
    }

    pub fn run(&mut self) -> Result<(), String> {
        let rate = Duration::from_micros(16670);

        self.init();

        let running = Arc::new(AtomicBool::new(true));
        let running_2 = running.clone();

        let result = ctrlc::set_handler(move || {
            running_2.store(false, Ordering::SeqCst);
        });
        if let Err(e) = result {
            return Err(format!("unable to set signal handler: {}", e));
        }

        while running.load(Ordering::SeqCst) {
            let t0 = time::Instant::now();
            self.tick();

            let elapsed = t0.elapsed();
            let remaining = rate - elapsed;
            if remaining > Duration::ZERO {
                thread::sleep(remaining);
            }
        }

        info!(self.queue, "stop requested");
        self.tick();

        Ok(())
    }

    fn process_queue(&mut self) {
        // Send each message in the queue to every system for processing.
        loop {
            // As messages are being processed, the systems may want to
            // generate additional messages. Create a new sender for those.
            let mut q = self.queue.clone();

            match self.rx.try_recv() {
                Err(TryRecvError::Empty) => break,
                Err(TryRecvError::Disconnected) => panic!("channel closed"),
                Ok(msg) => {
                    let mut handled = false;
                    for dev in &mut self.devices {
                        handled |= dev.process(&mut self.env, &mut q, &msg);
                    }
                    handled |= Self::process(&mut self.env, &mut q, &msg);
                    if !handled && self.env.conf.is_develop() {
                        panic!("message not handled: {:?}", &msg);
                    }
                }
            }
        }
    }

    fn process(e: &mut Env, _: &mut Queue, msg: &Message) -> bool {
        match msg {
            Message::Note(n) => {
                if e.conf.is_develop() && n.kind == NoteKind::Fault {
                    panic!("fault: {}", n.message);
                }
            }
            _ => return false,
        }
        true
    }
}

impl Default for Engine<'_> {
    fn default() -> Self {
        Engine::new(Config::default())
    }
}

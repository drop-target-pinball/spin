use crate::prelude::*;
use std::time;

pub struct Context<'s> {
    pub state: &'s mut State,
    pub queue: &'s mut Queue,
    pub conf: &'s Config,
}

impl<'s> Context<'s> {
    pub fn new(state: &'s mut State, queue: &'s mut Queue, conf: &'s Config) -> Self {
        Self { state, queue, conf }
    }
}

pub trait Device {
    fn process(&mut self, ctx: &mut Context, msg: &Message) -> bool;
}

pub struct Engine<'e>  {
    pub queue: Queue,
    pub state: State,
    devices: Vec<&'e mut dyn Device>,
    conf: Config,
}


impl<'e> Engine<'e> {
    pub fn new(conf: Config) -> Self {
        Engine {
            state: State::new(),
            queue: Queue::new(),
            devices: Vec::new(),
            conf,
        }
    }

    pub fn add_device(&mut self, s: &'e mut dyn Device) {
        self.devices.push(s);
    }

    pub fn init(&mut self) {
        self.queue.push(Message::Init);
        self.process_queue();
    }

    pub fn tick(&mut self) {
        self.state.now = time::Instant::now();
        self.process_queue();
    }

    fn process_queue(&mut self) {
        loop {
            // As messages are being processed, the systems may want to
            // generate additional messages. Collect those here.
            let mut out_queue = Queue::new();
            let mut ctx = Context::new(&mut self.state, &mut out_queue, &self.conf);

            // Send each message in the queue to every system for processing.
            match self.queue.pop() {
                None => break,
                Some(msg) => {
                    let mut handled = false;
                    for dev in &mut self.devices {
                            handled |= dev.process(&mut ctx, &msg);
                    }
                    if ctx.conf.is_develop() && !handled {
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

    fn process(&mut self, ctx: &mut Context, msg: &Message) -> bool {
        match msg {
            Message::Note(n) => {
                if ctx.conf.is_develop() && n.kind == NoteKind::Fault {
                    panic!("fault: {}", n.message);
                }
            },
            _ => return false
        }
        true
    }

}

impl<'e> Default for Engine<'e> {
    fn default() -> Self {
        Engine::new(Config::default())
    }
}


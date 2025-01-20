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
    fn id(&self) -> u8;
    fn topic(&self) -> Topic;
    fn process(&mut self, ctx: &mut Context, topic: Topic, msg: &Message);

    fn is_addressed_to(&self, pkt: &Packet) -> bool {
        if self.topic() == Topic::All {
            true
        } else if pkt.topic == Topic::All {
            true
        } else {
            self.topic() == pkt.topic && self.id() == pkt.device_id
        }
    }
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
        self.queue.post(Topic::All, 0, Message::Init);
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
            match self.queue.process() {
                None => break,
                Some(pkt) => {
                    for dev in &mut self.devices {
                        if dev.is_addressed_to(&pkt) {
                            dev.process(&mut ctx, pkt.topic, &pkt.message);
                        }
                    }
                },
            }

            // If systems generated additional messages, place these on the
            // main queue.
            while let Some(pkt) = out_queue.process() {
                self.queue.post(pkt.topic, pkt.device_id, pkt.message);
            }

            // Continue processing until the main queue is empty
            if self.queue.is_empty() {
                break
            }
        }
    }


}

impl<'e> Default for Engine<'e> {
    fn default() -> Self {
        Engine::new(Config::default())
    }
}


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

pub trait System {
    fn process(&mut self, ctx: &mut Context, evt: &Event);
}

pub struct Engine<'e>  {
    pub queue: Queue,
    state: State,
    systems: Vec<&'e mut dyn System>,
    conf: Config,
}


impl<'e> Engine<'e> {
    pub fn new(conf: Config) -> Self {
        Engine {
            state: State::new(),
            queue: Queue::new(),
            systems: Vec::new(),
            conf,
        }
    }

    pub fn add_system(&mut self, s: &'e mut dyn System) {
        self.systems.push(s);
    }

    pub fn tick(&mut self) {
        self.state.now = time::Instant::now();
        self.process_queue();
    }

    fn process_queue(&mut self) {
        let mut out = Queue::new();

        // As events are being processed, the systems may want to generate
        // additional events. Collect those here.
        let mut out_queue = Queue::new();

        let mut ctx = Context::new(&mut self.state, &mut out_queue, &self.conf);
        loop {
            // Send each event in the queue to every system for processing.
            match self.queue.process() {
                None => break,
                Some(event) => {
                    for s in &mut self.systems {
                        s.process(&mut ctx, &event);
                    }
                },
            }

            // If systems generated additional events, place these on the
            // main queue.
            while let Some(event) = out.process() {
                self.queue.post(event)
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


use crate::prelude::*;
use sdl2;

pub struct Device<'d> {
    id: u8,
    sdl: sdl2::Sdl,
    systems: Vec<&'d mut dyn System>,
}

pub trait System {
    fn topic(&self) -> Topic;
    fn process(&mut self, sdl: &sdl2::Sdl, ctx: &mut Context, msg: &Message);
}


impl<'d> Device<'d> {
    pub fn new(id: u8) -> Result<Self, String> {
        let sdl = sdl2::init()?;
        Ok(Self{id, sdl, systems: Vec::new()})
    }

    pub fn add_system(&mut self, s: &'d mut dyn System) {
        self.systems.push(s);
    }
}

impl<'d> crate::engine::Device for Device<'d> {
    fn topic(&self) -> Topic {
        Topic::All
    }

    fn id(&self) -> u8 {
        self.id
    }

    fn process(&mut self, ctx: &mut Context, topic: Topic, msg: &Message) {
        for s in &mut self.systems {
            if topic == Topic::All || s.topic() == topic {
                s.process(&self.sdl, ctx, msg);
            }
        }
    }
}
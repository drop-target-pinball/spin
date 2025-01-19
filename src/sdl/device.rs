use crate::prelude::*;
use sdl2;

pub struct Device<'d> {
    ctx: sdl2::Sdl,
    systems: Vec<&'d mut dyn SubSystem>,
}

pub trait SubSystem {
    fn process(&mut self, sdl: &sdl2::Sdl, state: &mut State, queue: &mut Queue, event: &Event);
}


impl<'d> Device<'d> {
    pub fn new() -> Result<Self, String> {
        let ctx = sdl2::init()?;
        Ok(Self{ctx, systems: Vec::new()})
    }
}

// impl<'d> System for Device<'d> {
//     fn process(&mut self, state: &mut State, queue: &mut Queue, event: &Event) {
//         for s in &mut self.systems {
//             s.process(&self.ctx, state, queue, event);
//         }
//     }
// }
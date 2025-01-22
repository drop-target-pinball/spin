use crate::prelude::*;
use crate::sdl::audio::{Audio, AudioOptions};
use sdl2;

pub struct Device {
    id: u8,
    sdl: sdl2::Sdl,
    audio: Option<Audio>
}

impl Device {
    pub fn new(id: u8) -> Result<Self> {
        match sdl2::init() {
            Ok(sdl) =>  Ok(Self{
                id,
                sdl,
                audio: None,
            }),
            Err(reason) => device_error("unable to initialize SDL", reason)
        }
    }

    pub fn with_audio(mut self, options: AudioOptions) -> Result<Self> {
        self.audio = Some(Audio::new(&self.sdl, options)?);
        Ok(self)
    }
}

impl crate::engine::Device for Device {
    fn process(&mut self, ctx: &mut Context, msg: &Message) -> bool {
        let mut handled = false;
        if let Some(audio) = &mut self.audio {
            handled |= audio.process(&self.sdl, ctx, &msg);
        }
        handled
    }
}

// impl<'d> crate::engine::Device for Device<'d> {
//     fn process(&mut self, ctx: &mut Context, topic: Topic, msg: &Message) {
//         for s in &mut self.systems {
//             if topic == Topic::All || s.topic() == topic {
//                 s.process(&self.sdl, ctx, msg);
//             }
//         }
//     }
// }

pub fn device_error<T>(what: &str, reason: String) -> Result<T> {
    Err(Error::Device(what.to_string(), reason))
}
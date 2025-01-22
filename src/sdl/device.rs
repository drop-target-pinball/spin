use crate::prelude::*;
use crate::sdl::audio::{Audio, AudioOptions};
use sdl2;

pub struct SdlDevice {
    ctx: sdl2::Sdl,
    audio: Option<Audio>
}

impl SdlDevice {
    pub fn new() -> Result<Self> {
        match sdl2::init() {
            Ok(ctx) =>  Ok(Self{
                ctx,
                audio: None,
            }),
            Err(reason) => device_error("unable to initialize SDL", reason)
        }
    }

    pub fn with_audio(mut self, id: u8, options: AudioOptions) -> Result<Self> {
        self.audio = Some(Audio::new(&self.ctx, id, options)?);
        Ok(self)
    }
}

impl Device for SdlDevice {
    fn process(&mut self, env: &mut Env, queue: &mut Queue, msg: &Message) -> bool {
        let mut handled = false;
        if let Some(audio) = &mut self.audio {
            handled |= audio.process(&self.ctx, env, queue, &msg);
        }
        handled
    }
}

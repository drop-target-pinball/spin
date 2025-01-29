use crate::prelude::*;
use crate::sdl::audio::{Audio, AudioOptions};
use sdl2;

pub struct SdlDevice {
    ctx: sdl2::Sdl,
    audio: Option<Audio>
}

impl Default for SdlDevice {
    fn default() -> Self {
        match sdl2::init() {
            Ok(ctx) => Self {
                ctx,
                audio: None,
            },
            Err(reason) => panic!("unable to initialize SDL: {}", reason),
        }
    }
}

impl SdlDevice {
    pub fn with_audio(mut self, id: u8, options: AudioOptions) -> Self {
        self.audio = Some(Audio::new(&self.ctx, id, options));
        self
    }
}

impl Device for SdlDevice {
    fn process(&mut self, env: &mut Env, queue: &mut Queue, msg: &Message)  {
        if let Some(audio) = &mut self.audio {
            audio.process(&self.ctx, env, queue, msg);
        }
    }
}

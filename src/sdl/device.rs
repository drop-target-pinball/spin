use crate::prelude::*;
use crate::sdl::audio::{Audio, AudioOptions};
use sdl2;

pub struct SdlDevice<'a> {
    ctx: sdl2::Sdl,
    audio: Option<Audio<'a>>,
}

impl Default for SdlDevice<'_> {
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

impl SdlDevice<'_> {
    pub fn with_audio(mut self, id: u8, options: AudioOptions) -> Self {
        self.audio = Some(Audio::new(&self.ctx, id, options));
        self
    }
}

impl Device for SdlDevice<'_> {
    fn process(&mut self, env: &mut Env, msg: &Message)  {
        if let Some(audio) = &mut self.audio {
            audio.process(&self.ctx, env, msg);
        }
    }
}

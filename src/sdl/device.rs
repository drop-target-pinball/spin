use crate::prelude::*;
use crate::sdl::audio::{Audio, AudioConfig};
use crate::sdl::dmd::DmdConfig;
use sdl2;
use sdl2::pixels::PixelFormatEnum;
use serde::{Serialize, Deserialize};
use sdl2::surface::Surface;

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct Config {
    pub audio: Option<AudioConfig>,
    pub dmd: Option<DmdConfig>,
}

pub struct Device<'a> {
    ctx: sdl2::Sdl,
    audio: Option<Audio<'a>>,
}

impl Default for Device<'_> {
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

impl<'a> Device<'a> {
    pub fn new(conf: &Config) -> Self {
        let ctx = expect!(sdl2::init(), "unable to initialize SDL");

        let audio : Option<Audio<'a>> = match &conf.audio {
            Some(c) => Some(Audio::new(&ctx,&c)),
            None => None,
        };

        Self { ctx, audio }
    }
}

impl crate::engine::Device for Device<'_> {
    fn process(&mut self, env: &mut Env, msg: &Message)  {
        if let Some(audio) = &mut self.audio {
            audio.process(&self.ctx, env, msg);
        }
    }
}

pub fn new_layer<'a>(conf: crate::config::Display) -> Layer<'a> {
    Layer::Sdl(expect!(Surface::new(
        conf.width,
        conf.height,
        PixelFormatEnum::RGBA8888,
    ), "unable to create rendering surface"))
}
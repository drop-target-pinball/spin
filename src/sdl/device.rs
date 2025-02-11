use crate::prelude::*;
use crate::sdl::audio::{Audio, AudioConfig};
use crate::sdl::dmd::{Dmd, DmdConfig};
use crate::sdl::video::Renderer;
use sdl2::{self, AudioSubsystem, VideoSubsystem};
use sdl2::ttf;
use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct Config {
    pub audio: Option<AudioConfig>,
    pub dmd: Option<DmdConfig>,
}

pub struct Context {
    pub sdl: sdl2::Sdl,
    pub audio: AudioSubsystem,
    pub ttf: ttf::Sdl2TtfContext,
    pub video: VideoSubsystem,
}

impl<'a> Default for Context {
    fn default() -> Self {
        let sdl = expect!(sdl2::init(), "unable to initialize SDL");
        let audio = expect!(sdl.audio(), "unable to initialize audio");
        let ttf = expect!(ttf::init(), "unable to initialize TTF");
        let video = expect!(sdl.video(), "unable to initialize video");

        Self { sdl, audio, ttf, video }
    }
}

pub struct Device {
    ctx: Context,
    audio: Option<Audio>,
    dmd: Option<Dmd>,
    renderer: Renderer,
}

impl Device {
    pub fn new(app_conf: &AppConfig, device_conf: &Config) -> Self {
        let ctx = Context::default();
        let audio : Option<Audio> = match &device_conf.audio {
            Some(c) => Some(Audio::new(&c)),
            None => None,
        };

        let dmd = match &device_conf.dmd {
            Some(c) => Some(Dmd::new(&ctx, &app_conf.video, &c)),
            None => None,
        };

        Self { ctx, audio, dmd, renderer: Renderer::default() }
    }

    fn poll(&mut self, _: &mut State) {
        let mut pump = expect!(self.ctx.sdl.event_pump(), "unable to obtain SDL event pump");
        for _ in pump.poll_iter() {
            // do nothing for now
        }
    }
}

impl crate::Device for Device {
    fn init(&mut self, s: &mut State, _: &mut render::State) {
        if let Some(audio) = &mut self.audio {
            audio.init(s);
        }
        self.renderer.init(s);
    }

    fn process(&mut self, s: &mut State, msg: &Message)  {
        match msg {
            Message::Poll => self.poll(s),
            _ => (),
        }
        if let Some(audio) = &mut self.audio {
            audio.process(s, msg);
        }
    }

    fn render(&mut self, state: &mut render::State) {
        self.renderer.render(state);
    }

    fn present(&mut self, state: &render::State) {
        if let Some(dmd) = &mut self.dmd {
            unwrap!(dmd.present(state));
        }
    }
}


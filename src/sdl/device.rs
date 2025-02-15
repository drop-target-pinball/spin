use crate::prelude::*;
use crate::sdl::audio::{Audio, AudioConfig};
use crate::sdl::dmd::{Dmd, DmdConfig};
use crate::sdl::video::Renderer;
use crate::sdl::image::{Image, ImageConfig};
use crate::sdl::monitor::{Monitor, MonitorConfig};
use super::input::Input;
use sdl2::{self, AudioSubsystem, VideoSubsystem};
use sdl2::ttf;
use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct Config {
    pub audio: Option<AudioConfig>,
    pub dmd: Option<DmdConfig>,
    pub image: Option<ImageConfig>,
    pub monitor: Option<MonitorConfig>,
}

pub struct Context {
    pub sdl: sdl2::Sdl,
    pub audio: AudioSubsystem,
    pub ttf: &'static ttf::Sdl2TtfContext,
    pub video: VideoSubsystem,
}

impl Context {
    pub fn new() -> Context {
        let sdl = expect!(sdl2::init(), "unable to initialize SDL");
        let audio = expect!(sdl.audio(), "unable to initialize SDL audio");
        let ttf = Box::new(expect!(sdl2::ttf::init(), "unable to initialize SDL truetype"));
        let video = expect!(sdl.video(), "unable to initialize SDL video");

        Context { sdl, audio, ttf: Box::leak(ttf), video }
    }
}

pub struct Device {
    ctx: Context,
    audio: Option<Audio>,
    dmd: Option<Dmd>,
    input: Input,
    image: Option<Image>,
    monitor: Option<Monitor>,
    renderer: Renderer<'static>,
}

impl Device {
    pub fn new(app_conf: &AppConfig, runtime: &Runtime) -> Self {
        let ctx = Context::new();
        let device_conf = app_conf.sdl.as_ref().unwrap();

        let audio = match &device_conf.audio {
            Some(conf) => Some(Audio::new(&conf)),
            None => None,
        };

        let image = match &device_conf.image {
            Some(c) => Some(unwrap!(Image::new(&c))),
            None => None
        };
        let dmd = match &device_conf.dmd {
            Some(c) => Some(Dmd::new(&ctx, &app_conf.video, &c)),
            None => None,
        };
        let monitor = match &device_conf.monitor {
            Some(c) => Some(unwrap!(Monitor::new(&ctx, &c, runtime))),
            None => None
        };
        let input = Input::new(app_conf);
        let renderer = Renderer::default();

        Self { ctx, audio, dmd, input, image, monitor, renderer }
    }

    fn poll(&mut self, s: &mut State) {
        let mut pump = expect!(self.ctx.sdl.event_pump(), "unable to obtain SDL event pump");
        for event in pump.poll_iter() {
           self.input.event(s, &event);
        }
    }
}

impl<'a> crate::Device for Device {
    fn init(&mut self, s: &mut State, _: &mut render::State) {
        if let Some(audio) = &mut self.audio {
            audio.init(s);
        }
        self.renderer.init(&self.ctx, s);
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
        if let Err(e) = self.renderer.render(state) {
            fault!(state.queue, "{}", e);
        }
    }

    fn present(&mut self, state: &render::State) {
        if let Some(dmd) = &mut self.dmd {
            if let Err(e) = dmd.present(state) {
                fault!(state.queue, "{}", e);
            }
        }
        if let Some(monitor) = &mut self.monitor {
            if let Err(e) = monitor.present(state) {
                fault!(state.queue, "{}", e);
            }
        }
    }
}


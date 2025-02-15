use crate::prelude::*;
use crate::error::{Error, Result};

use serde::{Serialize, Deserialize};
use sdl2::video::Window;
use sdl2::image::LoadSurface;
use sdl2::surface::Surface;
use sdl2::render::{Canvas, Texture, TextureCreator};
use super::Context;

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct MonitorConfig {
    pub playfield: String
}

pub struct Monitor {
    conf: MonitorConfig,
    canvas: Canvas<Window>,
    playfield: Texture<'static>,
}

impl Monitor {
    pub fn new(ctx: &Context, conf: &MonitorConfig, runtime: &Runtime) -> Result<Monitor> {
        let path = runtime.dirs.data.join(&conf.playfield);
        let pf_surface = try_init!(Surface::from_file(&path));
        let (pf_w, pf_h) = (pf_surface.width(), pf_surface.height());

        let mode = try_init!(ctx.video.current_display_mode(0));

        let win = try_init!(
            ctx.video.window("Monitor", pf_w, pf_h)
            .position(mode.w - pf_w as i32, 0)
            .hidden()
            .build());
        let mut canvas = try_init!(win.into_canvas()
            .accelerated()
            .present_vsync()
            .build());

        let texture_creator = Box::leak(Box::new(canvas.texture_creator()));
        let pf_texture = try_init!(texture_creator.create_texture_from_surface(pf_surface));

        canvas.window_mut().show();

        Ok(Monitor{ conf: conf.clone(), canvas, playfield: pf_texture })
    }

    pub fn present(&mut self, s: &render::State) -> Result<()> {
        try_present!(self.canvas.copy(&self.playfield, None, None));
        self.canvas.present();
        Ok(())
    }
}
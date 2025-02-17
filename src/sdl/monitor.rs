use crate::prelude::*;
use crate::error::{Error, Result};
use crate::vars::Value;

use std::collections::HashMap;
use sdl2::gfx::primitives::DrawRenderer;
use sdl2::pixels::Color;
use serde::{Serialize, Deserialize};
use sdl2::video::Window;
use sdl2::image::LoadSurface;
use sdl2::surface::Surface;
use sdl2::render::{BlendMode, Canvas, Texture};
use sdl2::rect::Rect;
use super::Context;

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct MonitorConfig {
    pub playfield: String
}

#[derive(Default)]
struct DriverState {
    on: bool,
    last_update: i64,
    proc_schedule: u32,
    proc_cycle_seconds: u8,
    proc_now: bool,
    pulse: bool,
    pulse_expire: i64,
}

pub struct Monitor {
    conf: MonitorConfig,
    canvas: Canvas<Window>,
    playfield: Texture<'static>,
    states: HashMap<String, DriverState>,
    layouts: HashMap<String, Vec<Layout>>
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

        Ok(Monitor{
            conf: conf.clone(),
            canvas,
            playfield: pf_texture,
            states: HashMap::new(),
            layouts: HashMap::new(),
         })
    }

    fn start_driver(&mut self, s: &mut State, elapsed: i64, msg: &Name) {
        let mut maybe_ds = self.states.get_mut(&msg.name);
        let Some(ds) = maybe_ds.as_mut() else { return };
        ds.on = true;
        ds.last_update = elapsed;
        ds.pulse = false;
    }

    fn stop_driver(&mut self, s: &mut State, elapsed: i64, msg: &Name) {
        let mut maybe_ds = self.states.get_mut(&msg.name);
        let Some(ds) = maybe_ds.as_mut() else { return };
        ds.on = false;
        ds.last_update = elapsed;
        ds.pulse = true;
    }

    fn pulse_driver(&mut self, s: &mut State, elapsed: i64, msg: &PulseDriver) {
        let mut maybe_ds = self.states.get_mut(&msg.name);
        let Some(ds) = maybe_ds.as_mut() else { return };
        ds.on = true;
        ds.last_update = elapsed;
        ds.pulse = true;
        ds.pulse_expire = elapsed + ( 4 * match msg.time {
            Some(t) => t as i64,
            None => 25,
        });
    }

    pub fn init(&mut self, s: &mut State) {
        for (name, def) in &s.conf.drivers {
            self.states.insert(name.to_string(), DriverState::default());
            self.layouts.insert(name.to_string(), def.layout.clone());

        }
    }

    pub fn process(&mut self, s: &mut State, msg: &Message) {
        let elapsed = s.vars.get("elapsed").unwrap_or(&Value::Int(0)).as_int();

        match msg {
            Message::StartDriver(m) => self.start_driver(s, elapsed, &m),
            Message::StopDriver(m) => self.stop_driver(s, elapsed, &m),
            Message::PulseDriver(m) => self.pulse_driver(s, elapsed, &m),
            _ => (),
        }
    }

    pub fn present(&mut self, s: &render::State) -> Result<()> {
        try_present!(self.canvas.copy(&self.playfield, None, None));
        for (name, ds) in &mut self.states {
            let mut alpha_pct = 1.0;
            if !ds.on {
                continue
            }
            if ds.pulse && s.elapsed >= ds.pulse_expire {
                ds.on = false;
                continue
            }
            draw_layout(&mut self.canvas, &self.layouts[name], alpha_pct)?;
        }
        self.canvas.present();
        Ok(())
    }
}


fn draw_layout(cvs: &mut Canvas<Window>, layouts: &Vec<Layout>, alpha_pct: f64) -> Result<()> {
    for layout in layouts {
        let color_name = layout.color_name.as_ref().unwrap_or(&ColorName::White);
        let mut color = color_name.to_color().to_sdl();
        color.a = (0xa0 as f64 * alpha_pct) as u8;
        match layout.shape {
            Shape::Circle => {
                let r = std::cmp::max(layout.w, layout.h) as i16;
                let hr = r / 2;
                try_present!(cvs.filled_circle(
                    layout.x as i16 + hr, layout.y as i16 + hr,
                    hr,
                    color));
                try_present!(cvs.circle(
                    layout.x as i16 + hr, layout.y as i16 + hr,
                    hr,
                    Color::BLACK));
            }
            Shape::Rect => {
                cvs.set_draw_color(color);
                cvs.set_blend_mode(BlendMode::Blend);
                try_present!(cvs.fill_rect(Rect::new(
                    layout.x, layout.y,
                    layout.w, layout.h)));
                cvs.set_draw_color(Color::BLACK);
                try_present!(cvs.draw_rect(Rect::new(
                    layout.x, layout.y,
                    layout.w, layout.h)));
            }
        }
    }
    Ok(())
}
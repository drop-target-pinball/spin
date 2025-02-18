use crate::{prelude::*, DEFAULT_PULSE_TIME};
use crate::error::{Error, Result};
use crate::vars::Value;

use std::collections::HashMap;
use sdl2::gfx::primitives::DrawRenderer;
use sdl2::libc::Elf32_Addr;
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
enum DriverMode {
    #[default]
    Off,
    On,
    Pulse,
    Schedule
}

#[derive(Default)]
struct DriverState {
    on_now: bool,
    start: i64,
    cycle_len: i64,
    mode: DriverMode,
    schedule: Vec<(bool, i64)>,
    pos: usize,
    expire: i64,
}

pub struct Monitor {
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
            canvas,
            playfield: pf_texture,
            states: HashMap::new(),
            layouts: HashMap::new(),
         })
    }

    fn start_driver(&mut self, msg: &Name) {
        let mut maybe_ds = self.states.get_mut(&msg.name);
        let Some(ds) = maybe_ds.as_mut() else { return };
        ds.mode = DriverMode::On;
        ds.on_now = true;
    }

    fn stop_driver(&mut self, msg: &Name) {
        let mut maybe_ds = self.states.get_mut(&msg.name);
        let Some(ds) = maybe_ds.as_mut() else { return };
        ds.mode = DriverMode::Off;
        ds.on_now = false;
    }

    fn pulse_driver(&mut self, elapsed: i64, msg: &PulseDriver) {
        let mut maybe_ds = self.states.get_mut(&msg.name);
        let Some(ds) = maybe_ds.as_mut() else { return };
        ds.mode = DriverMode::Pulse;
        ds.on_now = true;
        ds.start = elapsed;
        ds.expire = elapsed + match msg.time {
            Some(t) => t,
            None => DEFAULT_PULSE_TIME,
        };
    }

    fn pwm_driver(&mut self, elapsed: i64, msg: &PwmDriver) {
        let mut maybe_ds = self.states.get_mut(&msg.name);
        let Some(ds) = maybe_ds.as_mut() else { return };
        ds.mode = DriverMode::Schedule;
        ds.start = elapsed;
        ds.schedule = vec![
            (true, msg.time_on),
            (false, msg.time_off),
        ];
        ds.pos = 0;
        ds.on_now = true;
    }

    fn schedule_driver(&mut self, elapsed: i64, msg: &ScheduleDriver) {
        let mut maybe_ds = self.states.get_mut(&msg.name);
        let Some(ds) = maybe_ds.as_mut() else { return };
        ds.mode = DriverMode::Schedule;
        ds.cycle_len = msg.schedule
            .iter()
            .map(|s| s.1)
            .sum();
        ds.start = elapsed / ds.cycle_len * ds.cycle_len;
        ds.schedule = msg.schedule.clone();
        ds.pos = find_schedule_pos(elapsed % ds.cycle_len, &ds.schedule);
        ds.on_now = ds.schedule[ds.pos].0;

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
            Message::ScheduleDriver(m) => self.schedule_driver(elapsed, &m),
            Message::StartDriver(m) => self.start_driver(&m),
            Message::StopDriver(m) => self.stop_driver( &m),
            Message::PulseDriver(m) => self.pulse_driver(elapsed, &m),
            Message::PwmDriver(m) => self.pwm_driver(elapsed, &m),
            _ => (),
        }
    }

    pub fn present(&mut self, s: &render::State) -> Result<()> {
        try_present!(self.canvas.copy(&self.playfield, None, None));
        for (name, ds) in &mut self.states {
            let alpha_pct = 1.0;
            match ds.mode {
                DriverMode::Off => ds.on_now = false,
                DriverMode::On => ds.on_now = true,
                DriverMode::Pulse => ds.on_now = s.elapsed >= ds.expire,
                DriverMode::Schedule => {
                    let cycle_pos = s.elapsed & ds.cycle_len;
                    ds.pos = find_schedule_pos(cycle_pos, &ds.schedule);
                    ds.on_now = ds.schedule[ds.pos].0;
                }
            }
            if ds.on_now {
                draw_layout(&mut self.canvas, &self.layouts[name], alpha_pct)?;
            }
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

fn find_schedule_pos(cycle_pos: i64, sched: &Vec<(bool, i64)>) -> usize {
    let mut cycle_pos = cycle_pos;
    for (pos, s) in sched.iter().enumerate() {
        if cycle_pos - s.1 <= 0 {
            return pos
        }
        cycle_pos -= s.1;
    }
    0
}


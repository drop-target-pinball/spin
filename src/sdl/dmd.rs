use crate::prelude::*;

use crate::sdl::*;

use serde::{Serialize, Deserialize};
use std::collections::HashMap;
use std::slice::from_raw_parts;
use std::sync::Arc;
use sdl2::video::Window;
use sdl2::render::Canvas;
use sdl2::rect::Rect;

fn default_dot_size() -> u32 { 4 }
fn default_padding() -> u32 { 1 }
fn default_border_size() -> u32 { 20 }
fn default_title() -> String { return "Dot Matrix Display".to_string() }
fn default_panel_color() -> ColorDef { ColorDef::new(0x40, 0x40, 0x40, 0xff) }
fn default_border_color() -> ColorDef { ColorDef::new(0x80, 0x80, 0x80, 0xff) }

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct DmdConfig {
    pub video: String,
    #[serde(default = "default_dot_size")]
    pub dot_size: u32,
    #[serde(default = "default_padding")]
    pub padding: u32,
    #[serde(default = "default_panel_color")]
    pub panel_color: ColorDef,
    #[serde(default = "default_border_size")]
    pub border_size: u32,
    #[serde(default = "default_border_color")]
    pub border_color: ColorDef,
    #[serde(default = "default_title")]
    pub title: String,
}

pub struct Dmd {
    video_def: VideoDef,
    conf: DmdConfig,
    canvas: Canvas<Window>,
}

impl Dmd {
    pub fn new(ctx: &Context, video_confs: &HashMap<String, VideoDef>, dmd_conf: &DmdConfig) -> Dmd {
        let video_def = match video_confs.get(&dmd_conf.video) {
            Some(c) => c,
            None => panic!("in sdl.dmd: video device not defined: {}", dmd_conf.video),
        };

        let w =
            (dmd_conf.border_size * 2) +
            (video_def.width * dmd_conf.dot_size) +
            (video_def.width* dmd_conf.padding) +
            dmd_conf.padding;

        let h =
            (video_def.height * dmd_conf.dot_size) +
            (video_def.height * dmd_conf.padding) +
            (dmd_conf.border_size * 2) +
            dmd_conf.padding;

        let win = expect!(ctx.video.window(&dmd_conf.title, w, h)
            .position(0, 0)
            .hidden()
            .build(),
            "unable to create DMD window"
        );
        let mut canvas = expect!(win.into_canvas()
            .accelerated()
            .present_vsync()
            .build(),
            "unable to create DMD canvas"
        );
        canvas.window_mut().show();

        Self { video_def: video_def.clone(), conf: dmd_conf.clone(), canvas }
    }

    pub fn present(&mut self, s: &render::State) -> Result<(), String> {
        let c = &mut self.canvas;

        let size = self.video_def.width * self.video_def.height * 4;
        let raw_surface = s.videos[&self.conf.video].frame().surface().raw();

        let data = unsafe {
            let pixels =  (*raw_surface).pixels as *const u8 ;
            from_raw_parts(pixels, size as usize)
        };

        let (win_w, win_h) = c.window().size();
        let panel_w = self.video_def.width;
        let panel_h = self.video_def.height;

        // Background
        c.set_draw_color(self.conf.panel_color);
        c.clear();

        // Borders
        let border_size = self.conf.border_size;
        c.set_draw_color(self.conf.border_color);
        c.fill_rect(Rect::new(0, 0, win_w, border_size))?;
        c.fill_rect(Rect::new(0, (win_h - border_size) as i32, win_w, border_size))?;
        c.fill_rect(Rect::new(0, 0, border_size, win_h))?;
        c.fill_rect(Rect::new((win_w - border_size) as i32, 0, border_size, win_h))?;

        // Dots
        for y in 0..panel_h {
            for x in 0..panel_w {
                let dx = border_size + self.conf.padding + (x * self.conf.padding) + (x * self.conf.dot_size);
                let dy = border_size + self.conf.padding + (y * self.conf.padding) + (y * self.conf.dot_size);
                let offset = ((y * self.video_def.width + x) * 4) as usize;
                let dot = rgb_to_gray(data[offset+0] as u8, data[offset+1] as u8, data[offset+2] as u8) / 16;
                c.set_draw_color(palettes::ORANGE[dot as usize]);
                c.fill_rect(Rect::new(dx as i32, dy as i32, self.conf.dot_size, self.conf.dot_size))?;
            }
        }
        self.canvas.present();
        Ok(())
    }
}

mod palettes {
    use sdl2::pixels::Color;

    pub static ORANGE: [Color; 16] = [
        Color{r: 0, g: 0, b: 0, a: 0xff},
        Color{r: 15, g: 8, b: 0, a: 0xff},
        Color{r: 33, g: 17, b: 0, a: 0xff},
        Color{r: 51, g: 25, b: 0, a: 0xff},
        Color{r: 66, g: 33, b: 0, a: 0xff},
        Color{r: 84, g: 42, b: 0, a: 0xff},
        Color{r: 102, g: 51, b: 0, a: 0xff},
        Color{r: 117, g: 58, b: 0, a: 0xff},
        Color{r: 135, g: 67, b: 0, a: 0xff},
        Color{r: 153, g: 76, b: 0, a: 0xff},
        Color{r: 168, g: 84, b: 0, a: 0xff},
        Color{r: 186, g: 93, b: 0, a: 0xff},
        Color{r: 204, g: 102, b: 0, a: 0xff},
        Color{r: 219, g: 109, b: 0, a: 0xff},
        Color{r: 237, g: 118, b: 0, a: 0xff},
        Color{r: 255, g: 127, b: 0, a: 0xff},
    ];

}

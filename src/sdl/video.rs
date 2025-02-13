use crate::prelude::*;
use crate::render;
use std::collections::HashMap;
use sdl2::surface::Surface;
use sdl2::render::{BlendMode, Canvas};
use sdl2::pixels::PixelFormatEnum;
use sdl2::rect::{Point, Rect};
use sdl2::pixels::Color;
use sdl2::ttf::Font;
use super::Context;
use serde::{Serialize, Deserialize};
use std::path::Path;

const TRANSPARENT: Color = Color{r: 0, g: 0, b: 0, a: 0};

pub struct Video {
    frame: Canvas<Surface<'static>>,
    layers: Vec<Canvas<Surface<'static>>>,
    dirty: bool,
}

impl Video {
    pub fn new(conf: &VideoDef) -> Video {
        let mut layers = Vec::new();
        for _ in 0..conf.layers {
            layers.push(new_canvas(conf));
        }
        Video {
            frame: new_canvas(&conf),
            layers,
            dirty: true
        }
    }

    pub fn layer(&mut self, i: usize) -> &mut Canvas<Surface<'static>> {
        self.dirty = true;
        &mut self.layers[i]
    }

    pub fn flatten(&mut self) {
        if !self.dirty {
            return
        }
        let frame_rect = Rect::new(0, 0, self.frame.surface().width(), self.frame.surface().height());
        self.frame.set_draw_color(Color::BLACK);
        self.frame.clear();
        unwrap!(self.frame.surface_mut().set_blend_mode(BlendMode::Blend));
        for layer in &mut self.layers {
            unwrap!(layer.surface().blit(frame_rect, &mut self.frame.surface_mut(), frame_rect));
            layer.set_draw_color(TRANSPARENT);
            layer.clear();
        }
        unwrap!(self.frame.surface_mut().set_blend_mode(BlendMode::None));
        self.dirty = false;
    }

    pub fn frame(&self) -> &Canvas<Surface<'static>> {
        &self.frame
    }
}

pub fn new_canvas(conf: &VideoDef) -> Canvas<Surface<'static>> {
    let surf = expect!(Surface::new(
        conf.width,
        conf.height,
        PixelFormatEnum::ABGR8888,
    ), "unable to create rendering surface");
    expect!(surf.into_canvas(), "unable to create rendering canvas")
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
#[serde(rename_all = "PascalCase")]
struct Tile {
    x: i32,
    y: i32,
    w: u32,
    h: u32,
    offset_x: i32,
}

struct BitmapFont {
    surface: Surface<'static>,
    tile_map: HashMap<String, Tile>,
    tracking: i32,
}

impl BitmapFont {
    pub fn load(path: &Path) -> SpinResult<BitmapFont> {
        let data = match std::fs::read(path) {
            Ok(d) => d,
            Err(e) => return raise!(Error::Load, "unable to load '{}': {}", path.to_string_lossy(), e),
        };
        let mut frames = Self::decode_dmd(&data)?;

        let mut info_path = path.to_path_buf();
        info_path.set_extension(".dmd.json");
        let info_text = match std::fs::read_to_string(&info_path) {
            Ok(i) => i,
            Err(e) => return raise!(Error::Load, "unable to load '{}': {}", info_path.to_string_lossy(), e),
        };

        let tile_map: HashMap<String, Tile> = match serde_json::from_str(&info_text) {
            Ok(i) => i,
            Err(e) => return raise!(Error::Load, "unable to parse '{}': {}", info_path.to_string_lossy(), e),
        };
        Ok(BitmapFont {
            surface: frames.remove(0),
            tile_map,
            tracking: 0,
        })
    }

    const HEADER_SIZE: usize = 16;

    pub fn decode_dmd(data: &[u8]) -> SpinResult<Vec<Surface<'static>>> {
        if data.len() < Self::HEADER_SIZE {
            return raise!(Error::InvalidFormat, "invalid DMD");
        }
        // let header = u32::from_le_bytes(unwrap!(data[0..4].try_into()));
        let n_frames = u32::from_le_bytes(unwrap!(data[4..8].try_into()));
        let width = u32::from_le_bytes(unwrap!(data[8..12].try_into()));
        let height = u32::from_le_bytes(unwrap!(data[12..16].try_into()));

        let total_size = Self::HEADER_SIZE as u32 + (width * height * n_frames);
        if total_size as usize != data.len() {
            return raise!(Error::InvalidFormat, "invalid DMD size, expected {}, got {}", total_size, data.len());
        }

        let mut frames =Vec::new();
        for _ in 0..n_frames {
            let surface = chain!(Surface::new(width, height, PixelFormatEnum::RGB888), Error::RenderError);
            let mut canvas = chain!(surface.into_canvas(), Error::RenderError);
            let start = Self::HEADER_SIZE as u32 + (n_frames * width * height);
            for y in 0..width {
                for x in 0..height {
                    let idx = (x * width) + y + start;
                    let dot = data[idx as usize];
			        // Values in file are going to be between 0x0 and 0xf. Copy
                    // the lower nibble to the higher nibble.
                    let dot = dot <<4 + dot;
                    canvas.set_draw_color(Color{r: dot, g: dot, b: dot, a: 0xff});
                    chain!(canvas.draw_point(Point::new(x as i32, y as i32)), Error::RenderError);
                }
            }
            frames.push(canvas.into_surface());
        }
        Ok(frames)
    }
}

pub struct Renderer<'ttf> {
    ttf_fonts: HashMap<String, Font<'ttf, 'static>>,
    bit_fonts: HashMap<String, BitmapFont>,
    font: Option<String>,
    color: Color,
}

impl<'ttf> Default for Renderer<'ttf> {
    fn default() -> Renderer<'ttf> {
        Renderer {
            ttf_fonts: HashMap::new(),
            bit_fonts: HashMap::new(),
            font: None,
            color: Color::BLACK,
        }
    }
}

impl<'ttf> Renderer<'ttf> {
    pub fn init(&mut self, ctx: &Context, s: &State) {
        for (name, font_def) in &s.conf.fonts {
            let path = s.runtime.dirs.data.join(&font_def.path);
            let font = match ctx.ttf.load_font(&path, font_def.point_size) {
                Ok(f) => f,
                Err(e) => {
                    fault!(s.queue, "unable to load font '{}': {}", name, e);
                    return
                }
            };
            self.ttf_fonts.insert(name.clone(), font);
        }
    }

    fn draw_text(&self, cvs: &mut Canvas<Surface<'static>>, args: &render::DrawText) -> SpinResult<()> {
        let Some(name) = &self.font else {
            return raise!(Error::RenderError, "no font has been set");
        };

        if let Some(font) = self.ttf_fonts.get(name) {
            self.draw_text_ttf(font, cvs, args)
        } else if let Some(font) = self.bit_fonts.get(name)  {
            self.draw_text_bit(font, cvs, args)
        } else {
            raise!(Error::RenderError, "no such font: {}", name)
        }
    }

    fn draw_text_ttf(&self, font: &Font<'ttf, 'static>, cvs: &mut Canvas<Surface<'static>>, args: &render::DrawText) -> SpinResult<()> {
        let text = match font.render(&args.text).solid(self.color) {
            Ok(s) => s,
            Err(e) => return raise!(Error::RenderError, "{}", e)
        };

        let x = if args.center_x {
            ((cvs.surface().width() - text.width()) / 2) as i32
        } else {
            args.x
        };

        let y = if args.center_y {
            ((cvs.surface().height() - text.height()) / 2) as i32
        } else {
            args.y
        };

        match text.blit(text.rect(), cvs.surface_mut(), Rect::new(x, y, text.width(), text.height())) {
            Ok(_) => Ok(()),
            Err(e) => return raise!(Error::RenderError, "{}", e)
        }
    }

    fn draw_text_bit(&self, font: &BitmapFont, cvs: &mut Canvas<Surface<'static>>, args: &render::DrawText) -> SpinResult<()> {
        for c in args.text.chars() {
            let Some(tile) = font.tile_map.get(&c.to_string()) else { continue };
            let src_rect = Rect::new(tile.x, tile.y, tile.w, tile.h);
            let dst_rect = Rect::new(args.x + tile.offset_x, args.y, tile.w, tile.h);
            chain!(font.surface.blit(src_rect, cvs.surface_mut(), dst_rect), Error::RenderError);
        }
        Ok(())
    }


    fn fill_rect(&self, cvs: &mut Canvas<Surface<'static>>, rect: &render::Rect) {
        unwrap!(cvs.fill_rect(Rect::new(
            rect.x,
            rect.y,
            rect.w,
            rect.h,
        )));
    }

    fn set_color(&mut self, cvs: &mut Canvas<Surface<'static>>, color: &render::Color) {
        let c = Color{
            r: color.r,
            g: color.g,
            b: color.b,
            a: color.a
        };
        cvs.set_draw_color(c);
        self.color = c;
    }

    fn set_font(&mut self, name: &str) -> SpinResult<()> {
        if self.ttf_fonts.contains_key(name) {
            self.font = Some(name.to_string());
            Ok(())
        } else {
            raise!(Error::RenderError, "no such font: {}", name)
        }
    }

    pub fn render_instruction(&mut self, layer: &mut Canvas<Surface<'static>>, inst: &render::Instruction) -> SpinResult<()> {
        match &inst.op {
            render::Op::Color(color) => self.set_color(layer, color),
            render::Op::DrawText(args) => self.draw_text(layer, args)?,
            render::Op::Font(name) => self.set_font(name)?,
            render::Op::FillRect(rect) => self.fill_rect(layer, rect),
        }
        Ok(())
    }

    pub fn render(&mut self, state: &mut render::State) {
        for (name, video) in &mut state.videos {
            for inst in &state.ops {
                if inst.device != *name {
                    continue
                }
                let layer = video.layer(inst.layer);
                if let Err(e) = self.render_instruction(layer, inst) {
                    fault!(state.queue, "{}", e);
                }
            }
            video.flatten();
        }
    }
}

use crate::prelude::*;
use crate::{Error, Result};
use crate::render;
use std::collections::HashMap;
use sdl2::surface::Surface;
use sdl2::render::{BlendMode, Canvas};
use sdl2::pixels::PixelFormatEnum;
use sdl2::rect::Rect;
use sdl2::pixels::Color;
use sdl2::ttf::Font;
use serde::{Serialize, Deserialize};
use std::path::Path;
use super::Context;

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

    pub fn flatten(&mut self) -> Result<()> {
        if !self.dirty {
            return Ok(())
        }
        let frame_rect = Rect::new(0, 0, self.frame.surface().width(), self.frame.surface().height());
        self.frame.set_draw_color(Color::BLACK);
        self.frame.clear();
        try_render!(self.frame.surface_mut().set_blend_mode(BlendMode::Blend));
        for layer in &mut self.layers {
            try_render!(layer.surface().blit(frame_rect, &mut self.frame.surface_mut(), frame_rect));
            layer.set_draw_color(TRANSPARENT);
            layer.clear();
        }
        try_render!(self.frame.surface_mut().set_blend_mode(BlendMode::None));
        self.dirty = false;
        Ok(())
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
    pub fn load(path: &Path) -> Result<BitmapFont> {
        let data = match std::fs::read(path) {
            Ok(d) => d,
            Err(e) => return raise!(Error::Init, "unable to load '{}': {}", path.to_string_lossy(), e),
        };
        let mut frames = super::decode_dmd(&data)?;

        let mut info_path = path.to_path_buf();
        info_path.set_extension("dmd.json");
        let info_text = match std::fs::read_to_string(&info_path) {
            Ok(i) => i,
            Err(e) => return raise!(Error::Init, "unable to load '{}': {}", info_path.to_string_lossy(), e),
        };

        let tile_map: HashMap<String, Tile> = match serde_json::from_str(&info_text) {
            Ok(i) => i,
            Err(e) => return raise!(Error::Init, "unable to parse '{}': {}", info_path.to_string_lossy(), e),
        };
        Ok(BitmapFont {
            surface: frames.remove(0),
            tile_map,
            tracking: 0,
        })
    }
}

pub struct Renderer<'ttf> {
    ttf_fonts: HashMap<String, Font<'ttf, 'static>>,
    bmp_fonts: HashMap<String, BitmapFont>,
    font: Option<String>,
    color: Color,
}

impl<'ttf> Default for Renderer<'ttf> {
    fn default() -> Renderer<'ttf> {
        Renderer {
            ttf_fonts: HashMap::new(),
            bmp_fonts: HashMap::new(),
            font: None,
            color: Color::BLACK,
        }
    }
}

impl<'ttf> Renderer<'ttf> {
    pub fn init(&mut self, ctx: &Context, s: &State) {
        for (name, font_def) in &s.conf.fonts {
            let path = s.runtime.dirs.data.join(&font_def.path);
            let ext = path.extension().unwrap_or_default();

            let result = {
                if ext == "ttf" {
                    self.load_ttf_font(ctx, s, name, font_def)
                } else if ext == "dmd" {
                    self.load_bmp_font(s, name, font_def)
                } else {
                    raise!(Error::Init, "invalid file extension for font: '{}'", font_def.path)
                }
            };
            if let Err(e) = result {
                fault!(s.queue, "{}", e);
            }
        }
    }

    fn load_ttf_font(&mut self, ctx: &Context, s: &State, name: &str, font_def: &FontDef) -> Result<()> {
        let path = s.runtime.dirs.data.join(&font_def.path);
        match ctx.ttf.load_font(&path, font_def.point_size.unwrap_or(8)) {
            Ok(f) => {
                self.ttf_fonts.insert(name.to_string(), f);
                Ok(())
            }
            Err(e) => raise!(Error::Init, "unable to load font '{}': {}", path.to_string_lossy(), e),
        }
    }

    fn load_bmp_font(&mut self, s: &State, name: &str, font_def: &FontDef) -> Result<()> {
        let path = s.runtime.dirs.data.join(&font_def.path);
        match BitmapFont::load(&path) {
            Ok(f) => {
                self.bmp_fonts.insert(name.to_string(), f);
                Ok(())
            }
            Err(e) => raise!(Error::Init, "unable to load font '{}': {}", path.to_string_lossy(), e),
        }
    }

    fn draw_text(&self, cvs: &mut Canvas<Surface<'static>>, args: &render::DrawText) -> Result<()> {
        let Some(name) = &self.font else {
            return raise!(Error::Render, "no font has been set");
        };

        if let Some(font) = self.ttf_fonts.get(name) {
            self.draw_text_ttf(font, cvs, args)
        } else if let Some(font) = self.bmp_fonts.get(name)  {
            self.draw_text_bit(font, cvs, args)
        } else {
            raise!(Error::Render, "no such font: {}", name)
        }
    }

    fn draw_text_ttf(&self, font: &Font<'ttf, 'static>, cvs: &mut Canvas<Surface<'static>>, args: &render::DrawText) -> Result<()> {
        let text = match font.render(&args.text).solid(self.color) {
            Ok(s) => s,
            Err(e) => return raise!(Error::Render, "{}", e)
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
            Err(e) => return raise!(Error::Render, "{}", e)
        }
    }

    fn draw_text_bit(&self, font: &BitmapFont, cvs: &mut Canvas<Surface<'static>>, args: &render::DrawText) -> Result<()> {
        let mut x = args.x;
        for c in args.text.chars() {
            let Some(tile) = font.tile_map.get(&c.to_string()) else { continue };
            let src_rect = Rect::new(tile.x, tile.y, tile.w, tile.h);
            let dst_rect = Rect::new(x + tile.offset_x, args.y, tile.w, tile.h);
            chain!(font.surface.blit(src_rect, cvs.surface_mut(), dst_rect), Error::Render);
            x += tile.w as i32;
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

    fn set_font(&mut self, name: &str) -> Result<()> {
        if self.ttf_fonts.contains_key(name) {
            self.font = Some(name.to_string());
            Ok(())
        } else if self.bmp_fonts.contains_key(name) {
            self.font = Some(name.to_string());
            Ok(())
        } else {
            raise!(Error::Render, "no such font: {}", name)
        }
    }

    pub fn render_instruction(&mut self, layer: &mut Canvas<Surface<'static>>, inst: &render::Instruction) -> Result<()> {
        match &inst.op {
            render::Op::Color(color) => self.set_color(layer, color),
            render::Op::DrawText(args) => self.draw_text(layer, args)?,
            render::Op::Font(name) => self.set_font(name)?,
            render::Op::FillRect(rect) => self.fill_rect(layer, rect),
        }
        Ok(())
    }

    pub fn render(&mut self, state: &mut render::State) -> Result<()> {
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
            video.flatten()?;
        }
        Ok(())
    }
}

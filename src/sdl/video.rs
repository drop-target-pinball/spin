use crate::prelude::*;
use crate::render;
use sdl2::ttf;
use std::collections::HashMap;
use sdl2::surface::Surface;
use sdl2::render::{BlendMode, Canvas};
use sdl2::pixels::PixelFormatEnum;
use sdl2::rect::Rect;
use sdl2::pixels::Color;
use sdl2::ttf::Font;

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

pub struct Renderer<'a> {
    ttf: ttf::Sdl2TtfContext,
    fonts: HashMap<String, Font<'a, 'static>>,
    // font: Option<&Font<'_, 'static>>
}

impl<'a> Default for Renderer<'a> {
    fn default() -> Renderer<'a> {
        let ttf = expect!(ttf::init(), "unable to initialize TTF");
        Renderer {
            ttf,
            fonts: HashMap::new(),
        }
    }
}

impl<'a> Renderer<'a> {
    pub fn init(&'a mut self, s: &State) {
        for (name, font_def) in &s.conf.fonts {
            let path = s.conf.data_dir.join(&font_def.path);
            match self.ttf.load_font(path, font_def.point_size) {
                Err(e) => fault!(s.queue, "unable to load font '{}': {}", name, e),
                Ok(f) => { self.fonts.insert(name.to_string(), f); },
            }
        }
    }

    fn draw_text(&self, cvs: &mut Canvas<Surface<'static>>, args: &render::DrawText) {
        // let Some(f) = self.font else { return Ok(()) };
        // let surf = f.render(args.text)?;
    }

    fn fill_rect(&self, cvs: &mut Canvas<Surface<'static>>, rect: &render::Rect) {
        unwrap!(cvs.fill_rect(Rect::new(
            rect.x,
            rect.y,
            rect.w,
            rect.h,
        )));
    }

    fn set_color(&self, cvs: &mut Canvas<Surface<'static>>, color: &render::Color) {
        cvs.set_draw_color(Color{
            r: color.r,
            g: color.g,
            b: color.b,
            a: color.a
        });
    }

    fn set_font(&mut self, name: &str) -> SpinResult<()> {
        Ok(())
        // if self.fonts.contains_key(name) {
        //     self.font = Some(&self.font[name]);
        //     Ok(())
        // } else {
        //     raise!(Error::RenderError, "no such font: {}", name)
        // }
    }

    pub fn render_instruction(&mut self, layer: &mut Canvas<Surface<'static>>, inst: &render::Instruction) -> SpinResult<()> {
        match &inst.op {
            render::Op::Color(color) => self.set_color(layer, color),
            render::Op::DrawText(args) => self.draw_text(layer, args),
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

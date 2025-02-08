use crate::prelude::*;
use sdl2::surface::{Surface, SurfaceRef};
use sdl2::render::{BlendMode, Canvas};
use sdl2::pixels::PixelFormatEnum;
use sdl2::rect::Rect;
use sdl2::pixels::Color;

const TRANSPARENT: Color = Color{r: 0, g: 0, b: 0, a: 0};

pub struct Video {
    frame: Canvas<Surface<'static>>,
    layers: Vec<Canvas<Surface<'static>>>,
}

impl Video {
    pub fn new(conf: &VideoDef) -> Video {
        let mut layers = Vec::new();
        for _ in 0..conf.layers {
            layers.push(new_canvas(conf));
        }
        Video {
            frame: new_canvas(&conf),
            layers
        }
    }

    pub fn layer(&mut self, i: usize) -> &mut Canvas<Surface<'static>> {
        &mut self.layers[i]
    }

    pub fn flatten(&mut self) {
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
    }

    pub fn frame(&self) -> &Canvas<Surface<'static>> {
        &self.frame
    }
}

pub fn new_canvas(conf: &VideoDef) -> Canvas<Surface<'static>> {
    let surf = expect!(Surface::new(
        conf.width,
        conf.height,
        PixelFormatEnum::RGBA8888,
    ), "unable to create rendering surface");
    expect!(surf.into_canvas(), "unable to create rendering canvas")
}

fn fill_rect(cvs: &mut Canvas<Surface<'static>>, rect: &render::Rect) {
    unwrap!(cvs.fill_rect(Rect::new(
        rect.x,
        rect.y,
        rect.w,
        rect.h,
    )));
}

fn set_color(cvs: &mut Canvas<Surface<'static>>, color: &render::Color) {
    cvs.set_draw_color(Color{
        r: color.r,
        g: color.g,
        b: color.b,
        a: color.a
    });
}

pub fn render(state: &mut render::State) {
    for (name, video) in &mut state.videos {
        for inst in &state.ops {
            if inst.device != *name {
                continue
            }
            let layer = video.layer(inst.layer);
            match &inst.op {
                render::Op::Color(color) => set_color(layer, color),
                render::Op::FillRect(rect) => fill_rect(layer, rect),
            }
        }
        video.flatten();
    }
}
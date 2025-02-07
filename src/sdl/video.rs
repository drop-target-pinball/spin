use crate::prelude::*;
use sdl2::surface::{Surface, SurfaceRef};
use sdl2::render::Canvas;
use sdl2::pixels::PixelFormatEnum;
use sdl2::rect::Rect;
use sdl2::pixels::Color;

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
        for layer in &mut self.layers {
            expect!(layer.surface().blit(frame_rect, &mut self.frame.surface_mut(), frame_rect),
                "unable to flatten layers");
            layer.set_draw_color(Color::BLACK);
            layer.clear();
        }
    }

    pub fn frame(&self) -> &SurfaceRef {
        self.frame.surface()
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
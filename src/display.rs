use crate::prelude::*;

pub enum Layer<'a> {
    #[cfg(feature = "sdl")]
    Sdl(sdl2::surface::Surface<'a>),
}

pub struct Display<'a> {
    flattened: Layer<'a>,
    layers: Vec<Layer<'a>>,
}

impl<'a> Display<'a> {
    pub fn new(conf: config::Display) -> Display<'a> {
        let create_layer = match conf.renderer {
            #[cfg(feature = "sdl")]
            config::Renderer::Sdl => crate::sdl::new_layer,
            _ => panic!("unknown renderer")
        };

        let mut layers = Vec::new();
        for _ in 0..conf.layers {
            layers.push(create_layer(conf.clone()));
        }

        Display {
            flattened: create_layer(conf),
            layers
        }
    }
}
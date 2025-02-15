
use crate::prelude::*;
use crate::error::{Error, Result};

use serde::{Serialize, Deserialize};
use sdl2::image;

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct ImageConfig {
    #[serde(default)]
    pub with_jpeg: bool,
    #[serde(default)]
    pub with_png: bool,
}

pub struct Image {
}

impl Image {
    pub fn new(conf: &ImageConfig) -> Result<Self> {
        let mut flags = image::InitFlag::empty();
        if conf.with_jpeg {
            flags |= image::InitFlag::JPG;
        }
        if conf.with_png {
            flags |= image::InitFlag::PNG;
        }
        try_init!(image::init(flags));
        Ok(Self{})
    }
}
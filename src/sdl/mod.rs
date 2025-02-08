use serde::{Serialize, Deserialize};

mod audio;
mod device;
mod dmd;
mod video;

pub use crate::sdl::audio::*;
pub use crate::sdl::device::*;
pub use crate::sdl::dmd::*;
pub use crate::sdl::video::*;

use sdl2::pixels::Color;

fn default_alpha() -> u8 { 255 }

#[derive(Serialize, Deserialize, Debug, Copy, Clone)]
#[serde(deny_unknown_fields)]
pub struct ColorDef {
    #[serde(default)]
    pub r: u8,
    #[serde(default)]
    pub g: u8,
    #[serde(default)]
    pub b: u8,
    #[serde(default = "default_alpha")]
    pub a: u8,
}

impl ColorDef {
    fn new(r: u8, g: u8, b: u8, a: u8) -> ColorDef {
        ColorDef { r, g, b, a, }
    }
}

impl Into<Color> for ColorDef {
    fn into(self) -> Color {
        Color{
            r: self.r,
            g: self.g,
            b: self.b,
            a: self.a,
        }
    }
}
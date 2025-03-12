use crate::prelude::*;
use crate::{Error, Result};
use serde::{Serialize, Deserialize};
use sdl2::rect::Point;
use sdl2::pixels::PixelFormatEnum;
use sdl2::surface::Surface;

mod audio;
mod device;
mod dmd;
mod image;
mod input;
mod monitor;
pub mod video;

pub use crate::sdl::audio::*;
pub use crate::sdl::device::{Context, Config, Device};
pub use crate::sdl::dmd::*;
pub use crate::sdl::input::*;

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
    pub fn new(r: u8, g: u8, b: u8, a: u8) -> ColorDef {
        ColorDef { r, g, b, a, }
    }
}

impl From<ColorDef> for Color {
    fn from(val: ColorDef) -> Self {
        Color{
            r: val.r,
            g: val.g,
            b: val.b,
            a: val.a,    
        }
    }
}

const HEADER_SIZE: usize = 16;
const MAX_DIMENSION: u32 = 10240;

pub fn decode_dmd(data: &[u8]) -> Result<Vec<Surface<'static>>> {
    if data.len() < HEADER_SIZE {
        return raise!(Error::InvalidFormat, "invalid DMD");
    }
    let n_frames = u32::from_le_bytes(unwrap!(data[4..8].try_into()));
    let width = u32::from_le_bytes(unwrap!(data[8..12].try_into()));
    let height = u32::from_le_bytes(unwrap!(data[12..16].try_into()));

    // Sanity check to make sure the dimensions are reasonable. If not, we
    // are probably not reading a DMD file. 
    if width > MAX_DIMENSION || height > MAX_DIMENSION {
        return raise!(Error::InvalidFormat, "invalid DMD file");
    }

    let total_size = HEADER_SIZE as u32 + (width * height * n_frames);
    if total_size as usize != data.len() {
        return raise!(Error::InvalidFormat, "invalid DMD size, expected {}, got {}", total_size, data.len());
    }

    let mut frames =Vec::new();
    for i in 0..n_frames {
        let surface = chain!(Surface::new(width, height, PixelFormatEnum::RGB888), Error::Render);
        let mut canvas = chain!(surface.into_canvas(), Error::Render);
        let start = HEADER_SIZE as u32 + (i * width * height);
        for y in 0..height {
            for x in 0..width {
                let idx = (y * width) + x + start;
                let dot = data[idx as usize];
                // Values in file are going to be between 0x0 and 0xf. Copy
                // the lower nibble to the higher nibble.
                let dot = (dot << 4) | dot;
                canvas.set_draw_color(Color{r: dot, g: dot, b: dot, a: 0xff});
                chain!(canvas.draw_point(Point::new(x as i32, y as i32)), Error::Render);
            }
        }
        frames.push(canvas.into_surface());
    }
    Ok(frames)
}
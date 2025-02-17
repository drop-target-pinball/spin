pub mod builtin;
pub mod engine;
pub mod error;
pub mod font;
pub mod message;
pub mod render;
pub mod config;
pub mod vars;
pub mod script;

#[cfg(feature = "sdl")]
pub mod sdl;

pub mod prelude {
    pub use crate::builtin;
    pub use crate::config::*;
    pub use crate::engine::*;
    pub use crate::message::*;
    pub use crate::render;
    pub use crate::vars;
    pub use crate::script;
    pub use crate::{alert, diag, raise, fault, info};
    pub use crate::{unwrap, expect, chain};
    pub use crate::{try_device, try_init, try_present, try_render};

    pub use crate::{Device, Video};
    pub use crate::{rgb_to_gray, sec_to_millis};

    #[cfg(feature = "sdl")]
    pub use crate::sdl;
}

use crate::prelude::*;

pub use crate::error::{Error, Result};

#[cfg(feature = "sdl")]
pub type Video = crate::sdl::video::Video;

pub trait Device {
    fn init(&mut self, s: &mut State, r: &mut render::State);
    fn poll(&mut self, s: &mut State) -> error::Result<()>;
    fn process(&mut self, s: &mut State, msg: &Message);
    fn render(&mut self, s: &mut render::State);
    fn present(&mut self, s: &render::State);
}

// https://stackoverflow.com/questions/42516203/converting-rgba-image-to-grayscale-golang
pub fn rgb_to_gray(r: u8, g: u8, b: u8) -> u8 {
	let lum = 0.299*(r as f64) + 0.587*(g as f64) + 0.114*(b as f64);
	lum as u8
}

#[macro_export]
macro_rules! unwrap {
    ($q:expr) => {
        match $q {
            Ok(a) => a,
            Err(e) => panic!("unexpected error: {}", e),
        }
    };
}

#[macro_export]
macro_rules! expect {
    ($q:expr, $msg:expr) => {
        match $q {
            Ok(a) => a,
            Err(e) => panic!("error: {}: {}", $msg, e),
        }
    };
}


pub fn sec_to_millis(sec: f64) -> i64 {
    (sec * 1000_f64) as i64
}

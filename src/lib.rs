pub mod builtin;
pub mod engine;
pub mod error;
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
    pub use crate::error::*;
    pub use crate::message::*;
    pub use crate::render;
    pub use crate::vars;
    pub use crate::script;
    pub use crate::{alert, diag, raise, fault, info, unwrap, expect};

    pub use crate::{Device, Globals, Video};
    pub use crate::{rgb_to_gray, sec_to_millis};

    #[cfg(feature = "sdl")]
    pub use crate::sdl;
}

use crate::prelude::*;

#[cfg(feature = "sdl")]
pub type Video = crate::sdl::Video;

pub struct Globals<'a> {
    pub s: &'a mut State,
    pub r: &'a mut render::State
}

pub trait Device {
    fn init(&mut self, g: &mut Globals);
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

// #[cfg(test)]
// mod tests {
//     use super::*;

//     #[test]
//     fn test_rgb_to_gray() {
//         let have = rgb_to_gray(127, 127, 127);
//         let want = 253;
//         assert_eq!(have, want);
//     }
// }
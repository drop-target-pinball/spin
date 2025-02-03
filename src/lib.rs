pub mod builtin;
pub mod engine;
pub mod error;
pub mod message;
pub mod config;
pub mod vars;
pub mod script;

#[cfg(feature = "sdl")]
pub mod sdl;


pub mod prelude {
    pub use crate::builtin;
    pub use crate::engine::*;
    pub use crate::error::*;
    pub use crate::message::*;
    pub use crate::config;
    pub use crate::vars::*;
    pub use crate::script;
    pub use crate::{alert, diag, raise, fault, info, unwrap};

    pub use crate::sec_to_millis;

    #[cfg(feature = "sdl")]
    pub use crate::sdl;
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

pub fn sec_to_millis(sec: f64) -> u64 {
    return (sec * 1000 as f64) as u64;
}


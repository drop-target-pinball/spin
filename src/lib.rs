pub mod builtin;
pub mod engine;
pub mod error;
pub mod message;
pub mod config;
pub mod vars;
pub mod proc;

#[cfg(feature = "sdl")]
pub mod sdl;


pub mod prelude {
    pub use crate::builtin;
    pub use crate::engine::*;
    pub use crate::error::*;
    pub use crate::message::*;
    pub use crate::config;
    pub use crate::vars::*;
    pub use crate::proc;
    pub use crate::{alert, raise, fault, info, unwrap};

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



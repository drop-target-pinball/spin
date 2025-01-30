pub mod builtin;
pub mod engine;
pub mod error;
pub mod message;
pub mod logger;
pub mod config;
pub mod vars;
pub mod proc;

#[cfg(feature = "sdl")]
pub mod sdl;
#[cfg(feature = "server")]
pub mod server;

pub mod prelude {
    pub use crate::builtin;
    pub use crate::engine::*;
    pub use crate::error::*;
    pub use crate::message::*;
    pub use crate::logger::*;
    pub use crate::config;
    pub use crate::vars::*;
    pub use crate::proc;
    pub use crate::{alert, raise, fault, info};

    #[cfg(feature = "sdl")]
    pub use crate::sdl;
    #[cfg(feature = "server")]
    pub use crate::server;
}






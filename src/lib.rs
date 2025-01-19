pub mod engine;
pub mod event;
pub mod logger;
pub mod config;
pub mod sdl;
pub mod state;

pub mod prelude {
    pub use crate::engine::*;
    pub use crate::event::*;
    pub use crate::logger::*;
    pub use crate::config::*;
    pub use crate::state::*;
    pub use crate::{alert, fault, info};
}


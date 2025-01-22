pub mod engine;
pub mod error;
pub mod message;
pub mod logger;
pub mod config;
#[cfg(feature = "sdl")]
pub mod sdl;
pub mod state;

pub mod prelude {
    pub use crate::engine::*;
    pub use crate::error::*;
    pub use crate::message::*;
    pub use crate::logger::*;
    pub use crate::config::*;
    pub use crate::state::*;
    pub use crate::{alert, fault, info};

    #[cfg(feature = "sdl")]
    pub use crate::sdl;

    pub use super::spin_path;
}

use std::env;
use std::path::{Path, PathBuf};

pub fn spin_path(name: &str) -> PathBuf {
    let root = env::var_os("SPIN_DIR").unwrap_or(".".into());
    Path::new(&root).join(name)
}


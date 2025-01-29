
use std::path::{Path, PathBuf};
use std::env;
use std::ffi::OsString;

#[derive(Clone, Copy, PartialEq)]
pub enum RunMode {
    Develop,
    Test,
    Release
}

#[derive(Clone)]
pub struct Sound {
    pub name: String,
    pub device_id: u8,
    pub path: String,
}

impl Sound {
    pub fn new(name: &str, path: &str) -> Sound {
        Self {
            name: name.to_string(),
            device_id: 0,
            path: path.to_string(),
        }
    }
}

#[derive(Clone)]
pub struct Config {
    pub mode: RunMode,
    pub lib_dir: PathBuf,
    pub app_dir: PathBuf,
    pub sounds: Vec<Sound>
}

impl Config {
    pub fn new(mode: RunMode) -> Self {
        Config {
            mode,
            lib_dir: PathBuf::from(env::var_os("SPIN_LIB_DIR").unwrap_or(".".into())),
            app_dir: PathBuf::from(env::var_os("SPIN_APP_DIR").unwrap_or(".".into())),
            sounds: Vec::new(),
        }
    }

    pub fn add_sound(&mut self, s: &Sound) -> &mut Self {
        self.sounds.push(s.clone());
        self
    }

    pub fn is_develop(&self) -> bool {
        self.mode == RunMode::Develop
    }

    pub fn is_release(&self) -> bool {
        self.mode == RunMode::Release
    }

    pub fn lua_path(&self) -> OsString {
        let mut path = OsString::new();
        path.push(&self.lib_dir.clone());
        path.push("/lua/?.lua;");
        path
    }
}

impl Default for Config {
    fn default() -> Self {
        Self::new(RunMode::Develop)
    }
}

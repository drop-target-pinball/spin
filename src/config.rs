
use std::path::PathBuf;
use std::env;

use serde::{Serialize, Deserialize};

#[derive(Clone, Copy, PartialEq)]
pub enum RunMode {
    Develop,
    Test,
    Release
}

impl Default for RunMode {
    fn default() -> RunMode { RunMode::Develop }
}

#[derive(Serialize, Deserialize, Clone)]
#[serde(deny_unknown_fields)]
pub struct Proc {
    pub name: String,
    pub module: String,
    pub call: String,
    #[serde(default)]
    pub group: String,
}

#[derive(Serialize, Deserialize, Clone)]
pub struct Sound {
    pub name: String,
    #[serde(default)]
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

#[derive(Serialize, Deserialize, Clone)]
pub struct App {
    #[serde(skip)]
    pub mode: RunMode,
    #[serde(skip)]
    pub app_dir: PathBuf,

    pub procs: Vec<Proc>,
    pub sounds: Vec<Sound>,
}

pub fn new(mode: RunMode, app_dir: PathBuf) -> App {
    App {
        mode,
        app_dir,
        procs: Vec::new(),
        sounds: Vec::new(),
    }
}

impl App {
    pub fn is_develop(&self) -> bool {
        self.mode == RunMode::Develop
    }

    pub fn is_release(&self) -> bool {
        self.mode == RunMode::Release
    }
}

impl Default for App {
    fn default() -> Self {
        new(RunMode::Develop, ".".into())
    }
}

pub fn app_dir() -> PathBuf {
    PathBuf::from(env::var_os("SPIN_DIR").unwrap_or(".".into()))
}
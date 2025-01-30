
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

#[derive(Serialize, Deserialize, Clone)]
pub struct Server {
    pub enabled: bool,
    pub host: String,
    pub port: u16,
    pub log_level: Option<String>,
}

impl Default for Server {
    fn default() -> Server {
        Server {
            enabled: false,
            host: String::from("0.0.0.0"),
            port: 7746, // SPIN on telephone buttons
            log_level: None,
        }
    }
}

#[derive(Serialize, Deserialize, Clone)]
pub struct App {
    #[serde(skip)]
    pub mode: RunMode,
    #[serde(skip)]
    pub app_dir: PathBuf,
    pub sounds: Vec<Sound>,
    pub server: Server,
}

pub fn new(mode: RunMode) -> App {
    App {
        mode,
        app_dir: PathBuf::from(env::var_os("SPIN_DIR").unwrap_or(".".into())),
        sounds: Vec::new(),
        server: Server::default(),
    }
}

impl App {
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
}

impl Default for App {
    fn default() -> Self {
        new(RunMode::Develop)
    }
}

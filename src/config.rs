
use crate::prelude::*;

use std::fs;
use std::io;
use std::path::{Path, PathBuf};
use std::env;
use figment::Figment;
use figment::providers::{Format, Yaml};

use serde::{Serialize, Deserialize};

#[derive(Clone, Copy, Debug, PartialEq)]
pub enum RunMode {
    Develop,
    Test,
    Release
}

impl Default for RunMode {
    fn default() -> RunMode { RunMode::Develop }
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct Script {
    pub name: String,
    pub module: String,
    pub call: String,
    #[serde(default)]
    pub group: String,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
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

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct App {
    #[serde(skip)]
    pub mode: RunMode,
    #[serde(skip)]
    pub app_dir: PathBuf,

    pub scripts: Vec<Script>,
    pub sounds: Vec<Sound>,
}

pub fn new(mode: RunMode, app_dir: PathBuf) -> App {
    App {
        mode,
        app_dir,
        scripts: Vec::new(),
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

pub fn load(dir: &Path) -> Result<App> {
    let files = match find_files(dir) {
        Ok(f) => f,
        Err(e) => return raise!(Error::Config, "{}: {}", dir.to_string_lossy(), e)
    };

    if files.is_empty() {
        return raise!(Error::Config, "no configuration files found in '{}'", dir.to_string_lossy());
    }

    let mut builder = Figment::new();
    for file in files {
        let path = dir.join(file);
        builder = builder.merge(Yaml::file(&path));
    }
    let config: App = match builder.extract() {
        Ok(c) => c,
        Err(e) => return raise!(Error::Config, "{}", e),
    };
    Ok(config)
}

fn find_files(dir: &Path) -> io::Result<Vec<PathBuf>> {
    let mut files: Vec<PathBuf> = Vec::new();
    let listing = fs::read_dir(&dir)?;
    for result  in listing {
        let entry = result?;
        if entry.file_type()?.is_dir() {
            continue
        }

        let name = PathBuf::from(&entry.file_name());
        match name.extension() {
            None => (),
            Some(os_str) => match os_str.to_str() {
                Some("yaml") => files.push(name),
                Some("yml") => files.push(name),
                _ => (),
            }
        }
    }
    return Ok(files)
}
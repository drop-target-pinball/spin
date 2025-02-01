
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
    /// Without pinball machine
    Develop,

    /// With pinball machine
    Test,

    /// Headless via systemd
    Release
}

impl Default for RunMode {
    fn default() -> RunMode { RunMode::Develop }
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct Music {
    pub name: String,
    #[serde(default)]
    pub device_id: u8,
    pub path: String,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct Script {
    pub name: String,
    pub module: String,
    #[serde(default)]
    pub group: String,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
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
#[serde(deny_unknown_fields)]
pub struct Vocal {
    pub name: String,
    #[serde(default)]
    pub device_id: u8,
    pub path: String,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct App {
    #[serde(skip)]
    pub mode: RunMode,
    #[serde(skip)]
    pub app_dir: PathBuf,
    #[serde(skip)]
    pub data_dir: PathBuf,
    #[serde(skip)]
    pub scripts_dir: PathBuf,

    pub module_name: Option<String>,

    #[serde(default)]
    pub music: Vec<Music>,
    #[serde(default)]
    pub scripts: Vec<Script>,
    #[serde(default)]
    pub sounds: Vec<Sound>,
    #[serde(default)]
    pub vocals: Vec<Vocal>,
}

// pub fn new(mode: RunMode, app_dir: PathBuf) -> App {
//     App {
//         mode,
//         app_dir: PathBuf::new(),
//         data_dir: PathBuf::new(),
//         scripts_dir: PathBuf::new(),
//         module_name: None,
//         music: Vec::new(),
//         scripts: Vec::new(),
//         sounds: Vec::new(),
//         vocals: Vec::new(),
//     }
// }

impl App {
    pub fn is_develop(&self) -> bool {
        self.mode == RunMode::Develop
    }

    pub fn is_release(&self) -> bool {
        self.mode == RunMode::Release
    }
}

// impl Default for App {
//     fn default() -> Self {
//         new(RunMode::Develop, ".".into())
//     }
// }

// ----------------------------------------------------------------------------

pub fn app_dir() -> PathBuf {
    PathBuf::from(env::var_os("SPIN_DIR").unwrap_or(".".into()))
}

pub fn load(app_dir: &Path) -> Result<App> {
    let conf_dir = app_dir.join("config");
    let data_dir = app_dir.join("data");
    let scripts_dir = app_dir.join("scripts");

    let files = match find_files(&conf_dir) {
        Ok(f) => f,
        Err(e) => return raise!(Error::Config, "{}: {}", conf_dir.to_string_lossy(), e)
    };

    if files.is_empty() {
        return raise!(Error::Config, "no configuration files found in '{}'", conf_dir.to_string_lossy());
    }

    let mut builder = Figment::new();
    for file in files {
        builder = builder.merge(Yaml::file(&file));
    }
    let mut config: App = match builder.extract() {
        Ok(c) => c,
        Err(e) => return raise!(Error::Config, "{}", e),
    };

    config.app_dir = PathBuf::from(app_dir);
    config.data_dir = data_dir;
    config.scripts_dir = scripts_dir;

    Ok(config)
}

fn find_files(dir: &Path) -> io::Result<Vec<PathBuf>> {
    let mut files: Vec<PathBuf> = Vec::new();
    let listing = fs::read_dir(&dir)?;
    for result in listing {
        let entry = result?;
        if entry.file_type()?.is_dir() {
            files.append(&mut find_files(&dir.join(&entry.file_name()))?);
            continue
        }

        let name = PathBuf::from(&entry.file_name());
        match name.extension() {
            None => (),
            Some(os_str) => match os_str.to_str() {
                Some("yaml") => files.push(dir.join(name)),
                Some("yml") => files.push(dir.join(name)),
                _ => (),
            }
        }
    }
    return Ok(files)
}
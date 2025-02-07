
use crate::prelude::*;

use std::fs;
use std::io;
use std::path::{Path, PathBuf};
use std::env;
use std::collections::HashMap;
use figment::Figment;
use figment::providers::{Format, Yaml};

use serde::{Serialize, Deserialize};

const STD: [(&str, &str); 3] = [
    ("game", include_str!("std/config/game.yaml")),
    ("player", include_str!("std/config/player.yaml")),
    ("player_4", include_str!("std/config/player_4.yaml")),
];

#[derive(Clone, Copy, Debug, Default, PartialEq)]
pub enum RunMode {
    /// Without pinball machine
    #[default]
    Develop,

    /// With pinball machine
    Test,

    /// Headless via systemd
    Release
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct MusicDef {
    pub path: String,
    #[serde(default)]
    pub device_id: u8,
}


#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct ScriptDef {
    pub module: String,
    #[serde(default)]
    pub group: String,
    #[serde(default)]
    pub replace: bool,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct SoundDef {
    pub path: String,
    #[serde(default)]
    pub device_id: u8,
    #[serde(default)]
    pub priority: i32,
    #[serde(default)]
    pub duck: f64,
    #[serde(default)]
    /// Seconds
    pub debounce: f64,
}


#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(rename_all = "snake_case")]
pub enum VarKind {
    Int(i64),
    Float(f64),
    String(String),
    Bool(bool),
    Namespace{name: String},
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct VarDef {
    pub kind: VarKind,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct VideoDef {
    pub width: u32,
    pub height: u32,
    pub layers: u32,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct VocalDef {
    pub path: String,
    #[serde(default)]
    pub device_id: u8,
    #[serde(default)]
    pub priority: i32,
    #[serde(default)]
    pub duck: f64
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct AppConfig {
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
    pub displays: HashMap<String, VideoDef>,
    #[serde(default)]
    pub music: HashMap<String, MusicDef>,
    #[serde(default)]
    pub namespaces: HashMap<String, HashMap<String, VarDef>>,
    #[serde(default)]
    pub scripts: HashMap<String, ScriptDef>,
    #[serde(default)]
    pub sounds: HashMap<String, SoundDef>,
    #[serde(default)]
    pub std: Vec<String>,
    #[serde(default)]
    pub vocals: HashMap<String, VocalDef>,
    #[serde(default)]
    pub vars: HashMap<String, VarDef>,
    #[serde(default)]
    pub video: HashMap<String, VideoDef>,

    #[cfg(feature = "sdl")]
    pub sdl: Option<crate::sdl::Config>,

}

impl AppConfig {
    pub fn is_develop(&self) -> bool {
        self.mode == RunMode::Develop
    }

    pub fn is_release(&self) -> bool {
        self.mode == RunMode::Release
    }
}

// ----------------------------------------------------------------------------

pub fn app_dir() -> PathBuf {
    PathBuf::from(env::var_os("SPIN_DIR").unwrap_or(".".into()))
}

pub fn load_config(app_dir: &Path) -> SpinResult<AppConfig> {
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
        builder = builder.admerge(Yaml::file(&file));
    }

    let config: AppConfig = match builder.extract() {
        Ok(c) => c,
        Err(e) => return raise!(Error::Config, "{}", e),
    };

    let mut std_config: HashMap<&str, &str> = HashMap::new();
    for (name, conf) in STD {
        std_config.insert(name, conf);
    }

    for name in config.std {
        let Some(conf) = std_config.get(name.as_str()) else {
            return raise!(Error::Config, "no such standard library config: {}", name);
        };
        builder = builder.adjoin(Yaml::string(conf));
    }

    let mut config: AppConfig = match builder.extract() {
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
    let listing = fs::read_dir(dir)?;
    for result in listing {
        let entry = result?;
        if entry.file_type()?.is_dir() {
            files.append(&mut find_files(&dir.join(entry.file_name()))?);
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
    Ok(files)
}

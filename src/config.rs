
use crate::prelude::*;

use std::fs;
use std::io;
use std::path::{Path, PathBuf};
use std::env;
use std::collections::HashMap;
use figment::Figment;
use figment::providers::{Format, Yaml};
use crate::{Error, Result};

use serde::{Serialize, Deserialize};

const STD: [(&str, &str); 4] = [
    ("dmd", include_str!("std/config/dmd.yaml")),
    ("service", include_str!("std/config/service.yaml")),
    ("game", include_str!("std/config/game.yaml")),
    ("player_4", include_str!("std/config/player_4.yaml")),
];

#[derive(Serialize, Deserialize, Debug, Default, Clone, Copy, PartialEq)]
pub enum RunMode {
    /// Without pinball machine
    #[default]
    Develop,

    /// With pinball machine
    PlayTest,

    /// Headless via systemd
    Release,

    /// Integration tests
    AutoTest
}

#[derive(Serialize, Deserialize, Debug, Clone, PartialEq)]
#[serde(rename_all = "snake_case")]
pub enum DriverKind {
    Flasher,
    General,
    Gi,
    Lamp,
    Solenoid,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(rename_all = "snake_case")]
pub enum ColorName {
    Blue,
    Cyan,
    Green,
    Orange,
    Pink,
    Red,
    White,
    Yellow
}

impl ColorName {
    pub fn to_color(&self) -> render::Color {
        match self {
            Self::Blue => render::Color::new(0, 0, 255, 255),
            Self::Cyan => render::Color::new(0, 255, 255, 255),
            Self::Green => render::Color::new(0, 255, 0, 255),
            Self::Orange => render::Color::new(255, 165, 0, 255),
            Self::Pink => render::Color::new(255, 105, 180, 255 ),
            Self::Red => render::Color::new(255, 0, 0, 255),
            Self::White => render::Color::new(255, 255, 255, 255),
            Self::Yellow => render::Color::new(255, 255, 0, 255),
        }
    }
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(rename_all = "snake_case")]
pub enum Shape {
    Rect,
    Circle,
    Diamond,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct Layout {
    pub shape: Shape,
    pub x: i32,
    pub y: i32,
    pub w: u32,
    pub h: u32,
    #[serde(default)]
    pub color_name: Option<ColorName>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct DriverDef {
    pub address: String,
    pub name: String,
    pub kind: DriverKind,
    pub sort_name: Option<String>,
    pub manual_name: Option<String>,
    #[serde(default)]
    pub unused: bool,
    #[serde(default)]
    pub components: Vec<Component>,
    #[serde(default)]
    pub layout: Vec<Layout>
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct FontDef {
    pub path: String,
    pub point_size: Option<u16>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct FlipperDef {
    pub address_power: String,
    pub address_hold: String,
    pub name: String,
    pub sort_name: Option<String>,
    #[serde(default)]
    pub components: Vec<Component>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct KeyDef {
    pub key: String,
    #[serde(default)]
    pub left_shift: bool,
    pub down: Option<Message>,
    pub up: Option<Message>,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct MatrixDef {
    #[serde(default)]
    pub rows: Vec<Vec<Component>>,
    #[serde(default)]
    pub columns: Vec<Vec<Component>>,
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
pub struct RunGroup {
    pub parent: Option<String>
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct ScriptDef {
    pub module: String,
    #[serde(default)]
    pub group: Option<String>,
    #[serde(default)]
    pub replace: bool,
    #[serde(default)]
    pub test: bool,
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

#[derive(Serialize, Deserialize, Debug, Default, Clone)]
#[serde(rename_all = "snake_case")]
pub enum SwitchNormally {
    #[default]
    Open,
    Closed
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(rename_all = "snake_case")]
pub enum Component {
    Jumper(String),
    Transistor(String),
    Wire(String),
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct SwitchDef {
    pub name: String,
    pub address: String,
    #[serde(default)]
    pub normally: SwitchNormally,
    pub sort_name: Option<String>,
    pub manual_name: Option<String>,
    #[serde(default)]
    pub unused: bool,
    #[serde(default)]
    pub layout: Vec<Layout>,
    #[serde(default)]
    pub components: Vec<Component>,
}

fn default_step() -> i64 { -1 }
fn default_tick() -> f64 { 1.0 }

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct TimerDef {
    pub var: String,
    pub start: i64,
    #[serde(default)]
    pub end: i64,
    #[serde(default = "default_step")]
    pub step: i64,
    #[serde(default = "default_tick")]
    pub tick: f64,
    pub group: Option<String>,
    pub expire_delay: Option<f64>,
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
    pub module_name: Option<String>,

    #[serde(default)] pub displays: HashMap<String, VideoDef>,
    #[serde(default)] pub drivers: HashMap<String, DriverDef>,
    #[serde(default)] pub flippers: HashMap<String, FlipperDef>,
    #[serde(default)] pub fonts: HashMap<String, FontDef>,
    #[serde(default)] pub keyboard: Vec<KeyDef>,
    #[serde(default)] pub matrices: HashMap<String, MatrixDef>,
    #[serde(default)] pub mock: Option<mock::Config>,
    #[serde(default)] pub music: HashMap<String, MusicDef>,
    #[serde(default)] pub namespaces: HashMap<String, HashMap<String, VarDef>>,
    #[serde(default)] pub run_groups: HashMap<String, RunGroup>,
    #[serde(default)] pub scripts: HashMap<String, ScriptDef>,
    #[serde(default)] pub sounds: HashMap<String, SoundDef>,
    #[serde(default)] pub std: Vec<String>,
    #[serde(default)] pub switches: HashMap<String, SwitchDef>,
    #[serde(default)] pub timers: HashMap<String, TimerDef>,
    #[serde(default)] pub vocals: HashMap<String, VocalDef>,
    #[serde(default)] pub vars: HashMap<String, VarDef>,
    #[serde(default)] pub video: HashMap<String, VideoDef>,

    #[cfg(feature = "sdl")]
    pub sdl: Option<crate::sdl::Config>,

}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct Runtime {
    pub debug_config: bool,
    pub prog_name: String,
    pub prog_description: String,
    pub prog_version: String,
    pub prog_date: String,
    pub mode: RunMode,
    pub dirs: Dirs,
    pub error: Option<String>
}

impl Runtime {
    pub fn new(dirs: Dirs) -> Runtime {
        Runtime {
            debug_config: false,
            prog_name: "PROG".to_string(),
            prog_description: "NAME".to_string(),
            prog_version: "ERSION".to_string(),
            prog_date: "DATE".to_string(),
            mode: RunMode::Develop,
            dirs,
            error: None,
        }
    }

    pub fn is_auto_test(&self) -> bool {
        self.mode == RunMode::AutoTest
    }

    pub fn is_develop(&self) -> bool {
        self.mode == RunMode::Develop
    }

    pub fn is_release(&self) -> bool {
        self.mode == RunMode::Release
    }

    pub fn is_shutdown_on_fault(&self) -> bool {
        return self.mode == RunMode::Release || self.mode == RunMode::AutoTest
    }
}

// ----------------------------------------------------------------------------
#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct Dirs {
    pub app: PathBuf,
    pub conf: PathBuf,
    pub data: PathBuf,
    pub scripts: PathBuf
}

impl Dirs {
    pub fn new(app_dir: &Path) -> Dirs {
        Dirs {
            app: app_dir.to_path_buf(),
            conf: app_dir.join("config"),
            data: app_dir.join("data"),
            scripts: app_dir.join("scripts"),
        }
    }
}

impl Default for Dirs {
    fn default() -> Dirs {
        let app_dir = PathBuf::from(env::var_os("SPIN_DIR").unwrap_or(".".into()));
        Dirs::new(&app_dir)
    }
}

pub fn load_config(runtime: &Runtime) -> Result<AppConfig> {
    let files = match find_files(&runtime.dirs.conf) {
        Ok(f) => f,
        Err(e) => return raise!(Error::Config, "{}: {}", runtime.dirs.conf.to_string_lossy(), e)
    };

    if files.is_empty() {
        return raise!(Error::Config, "no configuration files found in '{}'", runtime.dirs.conf.to_string_lossy());
    }

    let mut builder = Figment::new();
    for file in files {
        if runtime.debug_config {
            println!("loading: {}", file.to_string_lossy());
        }
        builder = builder.admerge(Yaml::file(&file));
        if runtime.debug_config {
            if let Err(e) = builder.extract::<AppConfig>() {
                return raise!(Error::Config, "{}", e);
            }
        }
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
        if runtime.debug_config {
            println!("loading: {}", name);
        }
        let Some(conf) = std_config.get(name.as_str()) else {
            return raise!(Error::Config, "no such standard library config: {}", name);
        };
        builder = builder.adjoin(Yaml::string(conf));
        if runtime.debug_config {
            if let Err(e) = builder.extract::<AppConfig>() {
                return raise!(Error::Config, "{}", e);
            }
        }
    }

    let config: AppConfig = match builder.extract() {
        Ok(c) => c,
        Err(e) => return raise!(Error::Config, "{}", e),
    };

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



#[derive(PartialEq)]
pub enum RunMode {
    Dev,
    Prod
}

pub struct Sound {
    id: String,
    path: String,
}

pub struct Config {
    pub mode: RunMode,
    pub sounds: Vec<Sound>
}

impl Config {
    pub fn new(mode: RunMode) -> Self {
        Config {
            mode, sounds:
            Vec::new(),
        }
    }

    pub fn is_dev(&self) -> bool {
        return self.mode == RunMode::Dev
    }
}

impl Default for Config {
    fn default() -> Self {
        Self::new(RunMode::Dev)
    }
}

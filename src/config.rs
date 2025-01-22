
#[derive(Clone, Copy, PartialEq)]
pub enum RunMode {
    Develop,
    Test,
    Release
}

#[derive(Clone)]
pub struct Sound {
    pub id: String,
    pub path: String,
}

impl Sound {
    pub fn new(id: &str, path: &str) -> Sound {
        Self {
            id: id.to_string(),
            path: path.to_string(),
        }
    }
}

#[derive(Clone)]
pub struct Config {
    pub mode: RunMode,
    pub sounds: Vec<Sound>
}

impl Config {
    pub fn new(mode: RunMode) -> Self {
        Config {
            mode,
            sounds: Vec::new(),
        }
    }

    pub fn add_sound(&mut self, s: &Sound) -> &mut Self {
        self.sounds.push(s.clone());
        self
    }

    pub fn is_develop(&self) -> bool {
        return self.mode == RunMode::Develop
    }
}

impl Default for Config {
    fn default() -> Self {
        Self::new(RunMode::Develop)
    }
}

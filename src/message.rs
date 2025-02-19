use crate::prelude::*;

use serde::{Serialize, Deserialize};
use std::sync::mpsc::Sender;
use std::fmt;
use std::collections::HashMap;

#[derive(Serialize, Deserialize, Debug, PartialEq, Copy, Clone)]
#[serde(rename_all = "snake_case")]
pub enum NoteKind {
    Alert,
    Diag,
    Fault,
    Info,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlayMusic {
    pub name: String,
    #[serde(default)]
    pub loops: i32,
    #[serde(default)]
    pub no_restart: bool,
    #[serde(default)]
    pub notify: bool,
}

impl fmt::Display for PlayMusic {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}", self.name)?;
        if self.loops != 0 {
            write!(f, ", loops={}", self.loops)?;
        }
        if self.no_restart {
            write!(f, ", no_restart={}", self.no_restart)?;
        }
        if self.notify {
            write!(f, ", notify={}", self.notify)?;
        }
        Ok(())
    }
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlaySound {
    pub name: String,
    #[serde(default)]
    pub loops: i32,
    #[serde(default)]
    pub notify: bool,
}

impl fmt::Display for PlaySound {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}", self.name)?;
        if self.loops != 0 {
            write!(f, ", loops={}", self.loops)?;
        }
        if self.notify {
            write!(f, ", notify={}", self.notify)?;
        }
        Ok(())
    }
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlayVocal {
    pub name: String,
    #[serde(default)]
    pub notify: bool,
}

impl fmt::Display for PlayVocal {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}", self.name)?;
        if self.notify {
            write!(f, ", notify={}", self.notify)?;
        }
        Ok(())
    }
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Name {
    pub name: String
}

impl fmt::Display for Name {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        if !self.name.is_empty() {
            write!(f, "{}", self.name)
        } else {
            write!(f, "any")
        }
    }
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Note {
    pub kind: NoteKind,
    pub message: String,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PulseDriver {
    pub name: String,
    /// Milliseconds
    pub time: Option<i64>,
}

impl fmt::Display for PulseDriver {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}", self.name)?;
        if let Some(t) = self.time {
            write!(f, "time={}", t)?;
        }
        Ok(())
    }
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PwmDriver {
    pub name: String,
    /// Milliseconds
    pub time_on: i64,
    /// Milliseconds
    pub time_off: i64,
}

impl fmt::Display for PwmDriver {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}, time_on={}, time_off={}", self.name, self.time_on, self.time_off)
    }
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Rejected {
    pub reason: String,
}

impl fmt::Display for Rejected {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}", self.reason)
    }
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct ScheduleDriver {
    pub name: String,
    pub schedule: Vec<(bool, i64)>,
}

impl fmt::Display for ScheduleDriver {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{} schedule={{{}}}", self.name,
            self.schedule
                .iter()
                .map(|s| format!("({}, {})", s.0, s.1))
                .collect::<Vec<String>>()
                .join(", "))
    }
}


#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct SwitchUpdated {
    pub name: String,
    pub active: bool,
}

impl fmt::Display for SwitchUpdated {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{} active={}", self.name, self.active)
    }
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Vars {
    pub ns: Option<String>,
    pub vars: HashMap<String, vars::Value>
}

impl fmt::Display for Vars {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        let mut nvs: Vec<String> = Vec::new();
        if let Some(ns) = &self.ns {
            nvs.push(format!("ns={}", ns));
        }
        for (name, val) in &self.vars {
            nvs.push(format!("{}={}", name, val));
        }
        write!(f, "{}", nvs.join(", "))
    }
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Updated {
    pub name: String,
    pub was: vars::Value,
    pub value: vars::Value,
}

impl fmt::Display for Updated {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}={} (was={})", self.name, self.value, self.was)
    }
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(rename_all = "snake_case")]
pub enum Message {
    Halt,
    Kill(Name),
    KillGroup(Name),
    Note(Note),
    MusicEnded(Name),
    Nop,
    PlayMusic(PlayMusic),
    PlaySound(PlaySound),
    PlayVocal(PlayVocal),
    PulseDriver(PulseDriver),
    PwmDriver(PwmDriver),
    Rejected(Rejected),
    ResetTimer(Name),
    ScriptEnded(Name),
    ScriptKilled(Name),
    Set(Vars),
    ScheduleDriver(ScheduleDriver),
    Shutdown,
    StartDriver(Name),
    StartTimer(Name),
    Silence,
    SoundEnded(Name),
    StopDriver(Name),
    StopMusic(Name),
    StopTimer(Name),
    StopVocal(Name),
    SwitchUpdated(SwitchUpdated),
    Run(Name),
    Tick,
    TimerExpired(Name),
    Updated(Updated),
    VocalEnded(Name),
    Wake,
}

impl fmt::Display for Message {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match &self {
            Message::Halt => write!(f, "halt"),
            Message::Kill(m) => write!(f, "kill: {}", m),
            Message::KillGroup(m) => write!(f, "kill_group: {}", m),
            Message::Note(m) => {
                match m.kind {
                    NoteKind::Alert => write!(f, "(!) {}", m.message),
                    NoteKind::Diag => write!(f, "(?) {}", m.message),
                    NoteKind::Fault => write!(f, "(*) {}", m.message),
                    NoteKind::Info => write!(f, "{}", m.message),
                }
            }
            Message::Nop => Ok(()),
            Message::MusicEnded(m) => write!(f, "music_ended: {}", m),
            Message::PlayMusic(m) => write!(f, "play_music: {}", m),
            Message::PlaySound(m) => write!(f, "play_sound: {}", m),
            Message::PlayVocal(m) => write!(f, "play_vocal: {}", m),
            Message::PulseDriver(m) => write!(f, "pulse_driver: {}", m),
            Message::PwmDriver(m) => write!(f, "pwm_driver: {}", m),
            Message::Rejected(m) => write!(f, "rejected: {}", m),
            Message::ResetTimer(m) => write!(f, "reset_timer: {}", m),
            Message::Run(m) => write!(f, "run: {}", m),
            Message::ScheduleDriver(m) => write!(f, "schedule_driver: {}", m),
            Message::ScriptEnded(m) => write!(f, "script_ended: {}", m),
            Message::ScriptKilled(m) => write!(f, "script_killed: {}", m),
            Message::Set(m) => write!(f, "set: {}", m),
            Message::Shutdown => write!(f, "shutdown"),
            Message::StartDriver(m) => write!(f, "start_driver: {}", m),
            Message::StartTimer(m) => write!(f, "start_timer: {}", m),
            Message::Silence => write!(f, "silence"),
            Message::SoundEnded(m) => write!(f, "sound_ended: {}", m),
            Message::StopDriver(m) => write!(f, "stop_driver: {}", m),
            Message::StopMusic(m) => write!(f, "stop_music: {}", m),
            Message::StopTimer(m) => write!(f, "stop_timer: {}", m),
            Message::StopVocal(m) => write!(f, "stop_vocal: {}", m),
            Message::SwitchUpdated(m) => write!(f, "switch_updated: {}", m),
            Message::Tick => Ok(()),
            Message::TimerExpired(m) => write!(f, "timer_expired: {}", m),
            Message::Updated(m) => write!(f, "updated: {}", m),
            Message::VocalEnded(m) => write!(f, "vocal_ended: {}", m),
            Message::Wake => Ok(()),
        }
    }
}

#[macro_export]
macro_rules! alert {
    ($q:expr, $($args:expr),+) => {
        $q.post(Message::Note(Note{kind: NoteKind::Alert, message: format!($($args),+)}))
    };
}

#[macro_export]
macro_rules! diag {
    ($q:expr, $($args:expr),+) => {
        $q.post(Message::Note(Note{kind: NoteKind::Diag, message: format!($($args),+)}))
    };
}

#[macro_export]
macro_rules! fault {
    ($q:expr, $($args:expr),+) => {
        $q.post(Message::Note(Note{kind: NoteKind::Fault, message: format!($($args),+)}))
    };
}

#[macro_export]
macro_rules! info {
    ($q:expr, $($args:expr),+) => {
        $q.post(Message::Note(Note{kind: NoteKind::Info, message: format!($($args),+)}))
    };
}

pub struct Queue {
    tx: Sender<Message>
}

impl Queue {
    pub fn new(tx: Sender<Message>) -> Queue {
        Queue{tx}
    }

    pub fn post(&self, m: Message) {
        self.tx.send(m).expect("system is shutting down")
    }
}

impl Clone for Queue {
    fn clone(&self) -> Queue {
        Queue{tx: self.tx.clone()}
    }
}

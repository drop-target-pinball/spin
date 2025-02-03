use crate::prelude::*;

use serde::{Serialize, Deserialize};
use std::sync::mpsc::Sender;
use std::fmt;

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
    fn fmt(&self, f: &mut fmt::Formatter) -> FmtResult {
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
    fn fmt(&self, f: &mut fmt::Formatter) -> FmtResult {
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
    fn fmt(&self, f: &mut fmt::Formatter) -> FmtResult {
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
    fn fmt(&self, f: &mut fmt::Formatter) -> FmtResult {
        if self.name != "" {
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
pub struct Var {
    pub name: String,
    pub value: Value,
}

impl fmt::Display for Var {
    fn fmt(&self, f: &mut fmt::Formatter) -> FmtResult {
        write!(f, "{}={}", self.name, self.value)
    }
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct VarChanged {
    pub name: String,
    pub prev: Value,
    pub this: Value,
}

impl fmt::Display for VarChanged {
    fn fmt(&self, f: &mut fmt::Formatter) -> FmtResult {
        write!(f, "{} prev={}, this={}", self.name, self.prev, self.this)
    }
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(rename_all = "snake_case")]
pub enum Message {
    Halt,
    Init,
    Kill(Name),
    KillGroup(Name),
    Note(Note),
    MusicEnded(Name),
    Nop,
    PlayMusic(PlayMusic),
    PlaySound(PlaySound),
    PlayVocal(PlayVocal),
    ScriptEnded(Name),
    SetVar(Var),
    Shutdown,
    Silence,
    SoundEnded(Name),
    StopMusic(Name),
    StopVocal(Name),
    Run(Name),
    Tick,
    VarChanged(VarChanged),
    VocalEnded(Name),
}

impl fmt::Display for Message {
    fn fmt(&self, f: &mut fmt::Formatter) -> FmtResult {
        match &self {
            Message::Halt => write!(f, "halt"),
            Message::Init => write!(f, "init"),
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
            Message::ScriptEnded(m) => write!(f, "script_ended: {}", m),
            Message::SetVar(m) => write!(f, "set_var: {}", m),
            Message::Run(m) => write!(f, "run: {}", m),
            Message::Shutdown => write!(f, "shutdown"),
            Message::Silence => write!(f, "silence"),
            Message::SoundEnded(m) => write!(f, "sound_ended: {}", m),
            Message::StopMusic(m) => write!(f, "stop_music: {}", m),
            Message::StopVocal(m) => write!(f, "stop_vocal: {}", m),
            Message::Tick => Ok(()),
            Message::VarChanged(m) => write!(f, "var_changed: {}", m),
            Message::VocalEnded(m) => write!(f, "vocal_ended: {}", m),
        }
    }
}

#[macro_export]
macro_rules! alert {
    ($tx:expr, $($args:expr),+) => {
        $tx.post(Message::Note(Note{kind: NoteKind::Alert, message: format!($($args),+)}))
    };
}

#[macro_export]
macro_rules! diag {
    ($tx:expr, $($args:expr),+) => {
        $tx.post(Message::Note(Note{kind: NoteKind::Diag, message: format!($($args),+)}))
    };
}

#[macro_export]
macro_rules! fault {
    ($tx:expr, $($args:expr),+) => {
        $tx.post(Message::Note(Note{kind: NoteKind::Fault, message: format!($($args),+)}))
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

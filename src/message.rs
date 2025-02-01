use serde::{Serialize, Deserialize};
use std::sync::mpsc::Sender;
use std::fmt;

#[derive(Serialize, Deserialize, Debug, PartialEq, Copy, Clone)]
#[serde(rename_all = "snake_case")]
pub enum NoteKind {
    Info,
    Alert,
    Fault,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct PlayAudio {
    pub name: String,
    #[serde(default)]
    pub volume: u8,
    #[serde(default)]
    pub loops: i32,
    #[serde(default)]
    pub notify: bool,
}

impl fmt::Display for PlayAudio {
    fn fmt(&self, f: &mut fmt::Formatter) -> Result<(), fmt::Error> {
        write!(f, "{}", self.name)?;
        if self.volume != 0 {
            write!(f, " volume={}", self.loops)?;
        }
        if self.loops != 0 {
            write!(f, " loops={}", self.loops)?;
        }
        if self.notify {
            write!(f, " notify")?;
        }
        Ok(())
    }
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Name {
    pub name: String
}

impl fmt::Display for Name {
    fn fmt(&self, f: &mut fmt::Formatter) -> Result<(), fmt::Error> {
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
#[serde(rename_all = "snake_case")]
pub enum Message {
    Init,
    Note(Note),
    MusicEnded(Name),
    Nop,
    PlayMusic(PlayAudio),
    PlaySound(PlayAudio),
    PlayVocal(PlayAudio),
    ScriptEnded(Name),
    Shutdown,
    StopMusic(Name),
    Run(Name),
    Tick,
    VocalEnded(Name),
}

impl fmt::Display for Message {
    fn fmt(&self, f: &mut fmt::Formatter) -> Result<(), fmt::Error> {
        match &self {
            Message::Init => write!(f, "init"),
            Message::Note(m) => {
                match m.kind {
                    NoteKind::Alert => write!(f, "(!) {}", m.message),
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
            Message::Run(m) => write!(f, "run: {}", m),
            Message::Shutdown => write!(f, "shutdown"),
            Message::StopMusic(m) => write!(f, "stop_music: {}", m),
            Message::Tick => Ok(()),
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

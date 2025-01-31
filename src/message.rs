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

#[derive(Serialize, Deserialize, Debug)]
pub struct PlayAudio {
    pub name: String,
}

impl fmt::Display for PlayAudio {
    fn fmt(&self, f: &mut fmt::Formatter) -> Result<(), fmt::Error> {
        write!(f, "{}", self.name)
    }
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Name {
    pub name: String
}

impl fmt::Display for Name {
    fn fmt(&self, f: &mut fmt::Formatter) -> Result<(), fmt::Error> {
        write!(f, "{}", self.name)
    }
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Note {
    pub kind: NoteKind,
    pub message: String,
}

#[derive(Serialize, Deserialize, Debug)]
#[serde(rename_all = "snake_case")]
pub enum Message {
    Init,
    Note(Note),
    Nop,
    PlaySound(PlayAudio),
    ProcEnded(Name),
    Run(Name),
    Shutdown,
    Tick,
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
            Message::PlaySound(m) => write!(f, "play_sound: {}", m),
            Message::ProcEnded(m) => write!(f, "proc_ended: {}", m),
            Message::Run(m) => write!(f, "run: {}", m),
            Message::Shutdown => write!(f, "shutdown"),
            Message::Tick => Ok(()),
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

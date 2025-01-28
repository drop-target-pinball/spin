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
pub struct Note {
    pub kind: NoteKind,
    pub message: String,
}

#[derive(Serialize, Deserialize, Debug)]
#[serde(rename_all = "snake_case")]
pub enum Message {
    Init,
    Note(Note),
    PlaySound(PlayAudio),
}

impl fmt::Display for Message {
    fn fmt(&self, f: &mut fmt::Formatter) -> Result<(), fmt::Error> {
        match &self {
            Message::Init => write!(f, "init"),
            Message::Note(n) => {
                match n.kind {
                    NoteKind::Alert => write!(f, "(!) {}", n.message),
                    NoteKind::Fault => write!(f, "(*) {}", n.message),
                    NoteKind::Info => write!(f, "{}", n.message),
                }
            }
            Message::PlaySound(a) => write!(f, "play_sound: {}", a),
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

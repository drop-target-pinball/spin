use serde::{Serialize, Deserialize};
use std::collections::VecDeque;

#[derive(Serialize, Deserialize, Debug, PartialEq)]
#[serde(rename_all = "snake_case")]
pub enum NoteKind {
    Info,
    Alert,
    Fault,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct PlayAudio {
    pub id: String,
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

pub struct Queue {
    messages: VecDeque<Message>
}

impl Queue {
    pub fn new() -> Queue {
        Queue {
            messages: VecDeque::new(),
        }
    }

    pub fn push(&mut self, message: Message) {
        self.messages.push_back(message);
    }

    pub fn pop(&mut self) -> Option<Message> {
        self.messages.pop_front()
    }

    pub fn is_empty(&self) -> bool {
        self.messages.is_empty()
    }

    pub fn alert(&mut self, message: &str) {
        let n = Note{kind: NoteKind::Alert, message: message.to_string() };
        self.push(Message::Note(n));
    }

    pub fn fault(&mut self, message: &str) {
        let n = Note{kind: NoteKind::Fault, message: message.to_string() };
        self.push(Message::Note(n));
    }

    pub fn info(&mut self, message: &str) {
        let n = Note{kind: NoteKind::Info, message: message.to_string() };
        self.push(Message::Note(n));
    }
}

#[macro_export]
macro_rules! alert {
    ($queue:expr, $($args:expr),+) => {
        $queue.alert(&format!($($args),+));
    };
}

#[macro_export]
macro_rules! fault {
    ($queue:expr, $($args:expr),+) => {
        $queue.fault(&format!($($args),+))
    };
}

#[macro_export]
macro_rules! info {
    ($queue:expr, $($args:expr),+) => {
        $queue.info(&format!($($args),+));
    };
}
use serde::{Serialize, Deserialize};
use std::collections::VecDeque;

#[derive(Copy, Clone, PartialEq)]
pub enum Topic {
    All,
    Audio,
    Driver,
    Event,
    Render,
    Display,
    Admin,
}

#[derive(Serialize, Deserialize, Debug)]
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

pub struct Packet {
    pub topic: Topic,
    pub device_id: u8,
    pub message: Message,
}

pub struct Queue {
    messages: VecDeque<Packet>
}

impl Queue {
    pub fn new() -> Queue {
        Queue {
            messages: VecDeque::new(),
        }
    }

    pub fn post(&mut self, topic: Topic, device_id: u8, message: Message) {
        self.messages.push_back(Packet{
            topic,
            device_id,
            message
        });
    }

    pub fn process(&mut self) -> Option<Packet> {
        self.messages.pop_front()
    }

    pub fn is_empty(&self) -> bool {
        self.messages.is_empty()
    }

    pub fn alert(&mut self, message: &str) {
        let n = Note{kind: NoteKind::Alert, message: message.to_string() };
        self.post(Topic::Admin, 0, Message::Note(n));
    }

    pub fn fault(&mut self, message: &str) {
        let n = Note{kind: NoteKind::Fault, message: message.to_string() };
        self.post(Topic::Admin, 0, Message::Note(n));
    }

    pub fn info(&mut self, message: &str) {
        let n = Note{kind: NoteKind::Info, message: message.to_string() };
        self.post(Topic::Admin, 0, Message::Note(n));
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
        $queue.fault(&format!($($args),+));
    };
}

#[macro_export]
macro_rules! info {
    ($queue:expr, $($args:expr),+) => {
        $queue.info(&format!($($args),+));
    };
}
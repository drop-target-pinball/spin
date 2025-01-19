use serde::{Serialize, Deserialize};
use std::collections::VecDeque;

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
pub enum Event {
    Note(Note),
    PlaySound(PlayAudio),
}

pub struct Queue {
    events: VecDeque<Event>
}

impl Queue {
    pub fn new() -> Queue {
        Queue {
            events: VecDeque::new(),
        }
    }

    pub fn post(&mut self, e: Event) {
        self.events.push_back(e);
    }

    pub fn process(&mut self) -> Option<Event> {
        self.events.pop_front()
    }

    pub fn is_empty(&self) -> bool {
        self.events.is_empty()
    }

    pub fn alert(&mut self, message: &str) {
        let n = Note{kind: NoteKind::Alert, message: message.to_string() };
        self.post(Event::Note(n));
    }

    pub fn fault(&mut self, message: &str) {
        let n = Note{kind: NoteKind::Fault, message: message.to_string() };
        self.post(Event::Note(n));
    }

    pub fn info(&mut self, message: &str) {
        let n = Note{kind: NoteKind::Info, message: message.to_string() };
        self.post(Event::Note(n));
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
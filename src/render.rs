use crate::prelude::*;
use std::collections::HashMap;
use serde::{Serialize, Deserialize};

pub struct State {
    pub ops: Vec<Instruction>,
    pub videos: HashMap<String, Video>
}

impl Default for State {
    fn default() -> State {
        State {
            ops: Vec::new(),
            videos: HashMap::new(),
        }
    }
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct Instruction {
    pub device: String,
    #[serde(default)]
    pub layer: usize,
    #[serde(default)]
    pub priority: i32,
    pub op: Op,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct Color {
    pub r: u8,
    pub g: u8,
    pub b: u8,
    pub a: u8,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct Rect {
    pub x: i32,
    pub y: i32,
    pub w: u32,
    pub h: u32,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(rename_all = "snake_case")]
pub enum Op {
    Color(Color),
    FillRect(Rect)
}
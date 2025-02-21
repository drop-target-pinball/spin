use crate::prelude::*;
use std::collections::HashMap;
use serde::{Serialize, Deserialize};

pub struct State {
    pub videos: HashMap<String, Video>
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

#[cfg(feature = "sdl")]
impl Color {
    pub fn new(r: u8, g: u8, b: u8, a: u8) -> Self {
        Self { r: r, g: g, b: b, a: a }
    }

    pub fn to_sdl(&self) -> sdl2::pixels::Color {
        sdl2::pixels::Color {
            r: self.r,
            g: self.g,
            b: self.b,
            a: self.a,
        }
    }
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct Rect {
    pub x: i32,
    pub y: i32,
    pub w: u32,
    pub h: u32,
}

#[cfg(feature = "sdl")]
impl Rect {
    pub fn to_sdl(&self) -> sdl2::rect::Rect {
        sdl2::rect::Rect::new(
            self.x,
            self.y,
            self.w,
            self.h,
        )
    }
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct FillRect {
    pub rect: Rect,
    pub color: Color,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct DrawText {
    pub text: String,
    pub font: String,
    pub color: Color,
    #[serde(default)]
    pub x: i32,
    #[serde(default)]
    pub y: i32,
    #[serde(default)]
    pub center_x: bool,
    #[serde(default)]
    pub center_y: bool,
    #[serde(default)]
    pub right: bool,
    #[serde(default)]
    pub bottom: bool,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(rename_all = "snake_case")]
pub enum Op {
    DrawText(DrawText),
    FillRect(FillRect),
    New(Color),
}
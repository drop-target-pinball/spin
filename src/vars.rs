use crate::prelude::*;

use serde::{Serialize, Deserialize};
use std::collections::HashMap;
use std::fmt;

#[derive(Serialize, Deserialize, Debug)]
pub struct VarsBox {
    pub vars: Vars
}

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(rename_all = "snake_case")]
pub enum Value {
    Int(i64),
    Float(f64),
    String(String),
    Bool(bool),
}

impl fmt::Display for Value {
    fn fmt(&self, f: &mut fmt::Formatter) -> FmtResult {
        match self {
            Value::Int(i) => write!(f, "{}", i),
            Value::Float(fl) => write!(f, "{}", fl),
            Value::String(s) => write!(f, "'{}'", s),
            Value::Bool(b) => write!(f, "{}", b),
        }
    }
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Vars {
    pub elapsed: u64,
    store: HashMap<String,Value>,
}

impl Default for Vars {
    fn default() -> Vars {
        Vars {
            elapsed: 0,
            store: HashMap::new(),
        }
    }
}

impl Vars {
    fn update(&mut self, queue: &mut Queue, name: &str, prev: Value, this: &Value) {
        let msg = VarChanged{
            name: name.to_string(),
            prev,
            this: this.clone()
        };

        self.store.insert(name.to_string(), this.clone());
        queue.post(Message::VarChanged(msg));
    }

    pub fn define(&mut self, queue: &mut Queue, name: &str, value: &Value) {
        if let Some(_) = self.store.get(name) {
            fault!(queue, "variable already defined: {}", name);
            return;
        }
        self.store.insert(name.to_string(), value.clone());
    }

    pub fn set(&mut self, queue: &mut Queue, name: &str, this: &Value) {
        let prev = match self.store.get(name) {
            Some(v) => v,
            None => {
                fault!(queue, "variable not defined: {}", name);
                return;
            }
        };

        match (prev, this) {
            (Value::Int(_), Value::Int(_)) => self.update(queue, name, prev.clone(), this),
            (Value::Float(_), Value::Float(_)) => self.update(queue, name, prev.clone(), this),
            (Value::String(_), Value::String(_)) => self.update(queue, name, prev.clone(), this),
            (Value::Bool(_), Value::Bool(_)) => self.update(queue, name, prev.clone(), this),
            (got, _) => {
                fault!(queue, "invalid type, expected int, got {}", got);
            }
        }
    }

    pub fn set_int(&mut self, queue: &mut Queue, name: &str, i: i64) {
        self.set(queue, name, &Value::Int(i));
    }

}

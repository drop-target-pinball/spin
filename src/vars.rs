use crate::prelude::*;

use serde::{Serialize, Deserialize};
use std::collections::HashMap;
use std::fmt;

#[derive(Debug)]
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
    Vars(Vars),
}

impl fmt::Display for Value {
    fn fmt(&self, f: &mut fmt::Formatter) -> FmtResult {
        match self {
            Value::Int(i) => write!(f, "{}", i),
            Value::Float(fl) => write!(f, "{}", fl),
            Value::String(s) => write!(f, "'{}'", s),
            Value::Bool(b) => write!(f, "{}", b),
            Value::Vars(v) => write!(f, "{{ {} }}", v),
        }
    }
}

pub type Namespaces = HashMap<String, Vec<config::Var>>;

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct Vars {
    pub elapsed: u64,
    store: HashMap<String, Value>,
}

impl Vars {
    pub fn new() -> Self {
        Self {
            elapsed: 0,
            store: HashMap::new(),
        }
    }

    fn update(&mut self, queue: &mut Queue, name: &str, prev: Value, this: &Value) {
        let msg = VarChanged{
            name: name.to_string(),
            prev,
            this: this.clone()
        };

        self.store.insert(name.to_string(), this.clone());
        queue.post(Message::VarChanged(msg));
    }

    pub fn define(&mut self, queue: &mut Queue, spaces: &Namespaces, name: &str, kind: &config::VarKind) {
        if self.store.contains_key(name) {
            fault!(queue, "variable already defined: {}", name);
            return;
        }
        let value = match kind {
            config::VarKind::Int{default} => Value::Int(*default),
            config::VarKind::Float{default} => Value::Float(*default),
            config::VarKind::String{default} => Value::String(default.clone()),
            config::VarKind::Bool{default} => Value::Bool(*default),
            config::VarKind::Namespace{name} => {
                let defs = match spaces.get(name) {
                    Some(v) => v,
                    None => {
                        fault!(queue, "unknown namespace: {}", name);
                        return;
                    }
                };
                let mut vars = Vars::new();
                for def in defs {
                    vars.define(queue, spaces, &def.name, &def.kind);
                }
                Value::Vars(vars)
            }
        };
        self.store.insert(name.to_string(), value);
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
            (Value::Vars(_), Value::Vars(_)) => {
                fault!(queue, "cannot set vars '{}'", name);
            },
            (p, t) => {
                fault!(queue, "invalid type, expected {}, got {}", p, t);
            }
        }
    }

    pub fn set_int(&mut self, queue: &mut Queue, name: &str, i: i64) {
        self.set(queue, name, &Value::Int(i));
    }
}

impl Default for Vars {
    fn default() -> Self {
        Self::new()
    }
}

impl fmt::Display for Vars {
    fn fmt(&self, f: &mut fmt::Formatter) -> FmtResult {
        let mut kvs: Vec<String> = Vec::new();
        for (k, v) in &self.store {
            kvs.push(format!("{}={}", k, v));
        }
        write!(f, "{}", kvs.join(" "))
    }
}

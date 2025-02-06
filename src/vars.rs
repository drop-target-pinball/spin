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

impl Value {
    pub fn as_int(&self) -> i64 {
        match self {
            Value::Int(i) => *i,
            _ =>  panic!("not an integer: {}", self)
        }
    }
}

impl fmt::Display for Value {
    fn fmt(&self, f: &mut fmt::Formatter) -> FmtResult {
        match self {
            Value::Int(i) => write!(f, "{}", i),
            Value::Float(fl) => write!(f, "{}", fl),
            Value::String(s) => write!(f, "'{}'", s),
            Value::Bool(b) => write!(f, "{}", b),
            Value::Vars(v) => write!(f, "{{ {:?} }}", v),
        }
    }
}

pub type Namespaces = HashMap<String, Vec<config::Var>>;
pub type Vars = HashMap<String, Value>;

fn update(env: &mut Env, name: &str, prev: Value, this: &Value) {
    let msg = Updated{
        name: name.to_string(),
        was: prev,
        value: this.clone()
    };

    env.vars.insert(name.to_string(), this.clone());
    env.queue.post(Message::Updated(msg));
}

pub fn define(queue: &mut Queue, vars: &mut Vars, spaces: &HashMap<String, HashMap<String, config::Var>>, name: &str, kind: &config::VarKind) {
    if vars.contains_key(name) {
        fault!(queue, "variable already defined: {}", name);
        return;
    }
    let value = match kind {
        config::VarKind::Int(i) => Value::Int(*i),
        config::VarKind::Float(f) => Value::Float(*f),
        config::VarKind::String(s) => Value::String(s.clone()),
        config::VarKind::Bool(b) => Value::Bool(*b),
        config::VarKind::Namespace{name} => {
            let defs = match spaces.get(name) {
                Some(v) => v,
                None => {
                    fault!(queue, "unknown namespace: {}", name);
                    return;
                }
            };
            let mut sub_vars = Vars::new();
            for (name, def) in defs {
                define(queue, &mut sub_vars, spaces, &name, &def.kind);
            }
            Value::Vars(sub_vars)
        }
    };
    vars.insert(name.to_string(), value);
}

pub fn set(env: &mut Env, name: &str, this: &Value) {
    let prev = match env.vars.get(name) {
        Some(v) => v,
        None => {
            fault!(env.queue, "variable not defined: {}", name);
            return;
        }
    };

    match (prev, this) {
        (Value::Int(_), Value::Int(_)) => update(env, name, prev.clone(), this),
        (Value::Float(_), Value::Float(_)) => update(env, name, prev.clone(), this),
        (Value::String(_), Value::String(_)) => update(env, name, prev.clone(), this),
        (Value::Bool(_), Value::Bool(_)) => update(env, name, prev.clone(), this),
        (Value::Vars(_), Value::Vars(_)) => {
            fault!(env.queue, "cannot set vars '{}'", name);
        },
        (p, t) => {
            fault!(env.queue, "invalid type, expected {}, got {}", p, t);
        }
    }
}
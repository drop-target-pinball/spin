use crate::prelude::*;
use crate::{Result, Error};
use serde::{Serialize, Deserialize};
use std::collections::HashMap;
use std::fmt;

#[derive(Serialize, Deserialize, Debug, Clone, PartialEq)]
#[serde(rename_all = "snake_case")]
pub enum Value {
    Int(i64),
    Float(f64),
    String(String),
    Bool(bool),
}

impl Value {
    pub fn as_f64(&self) -> f64 {
        match self {
            Value::Float(f) => *f,
            _ => panic!("not a float: {}", self),
        }
    }
    pub fn as_i64(&self) -> i64 {
        match self {
            Value::Int(i) => *i,
            _ => panic!("not an integer: {}", self)
        }
    }

    pub fn kind(&self) -> &str {
        match self {
            Self::Int(_) => "int",
            Self::Float(_) => "float",
            Self::String(_) => "string",
            Self::Bool(_) => "bool",
        }
    }
}

impl fmt::Display for Value {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match self {
            Value::Int(i) => write!(f, "{}", i),
            Value::Float(fl) => write!(f, "{}", fl),
            Value::String(s) => write!(f, "'{}'", s),
            Value::Bool(b) => write!(f, "{}", b),
        }
    }
}

pub type Vars = HashMap<String, Value>;

fn update(vars: &mut HashMap<String, Value>, namespace: Namespace, name: &str, prev: Value, this: &Value) -> Result<Option<Updated>> {
    if prev == *this {
        return Ok(None);
    }
    let msg = Updated{
        namespace,
        name: name.to_string(),
        was: prev,
        value: this.clone()
    };

    vars.insert(name.to_string(), this.clone());
    Ok(Some(msg))
}

pub fn define(queue: &mut Queue, vars: &mut Vars, name: &str, kind: &VarKind) {
    if vars.contains_key(name) {
        fault!(queue, "variable already defined: {}", name);
        return;
    }
    let value = match kind {
        VarKind::Int(i) => Value::Int(*i),
        VarKind::Float(f) => Value::Float(*f),
        VarKind::String(s) => Value::String(s.clone()),
        VarKind::Bool(b) => Value::Bool(*b),
    };
    vars.insert(name.to_string(), value);
}

pub fn set(vars: &mut HashMap<String, Value>, namespace: Namespace, name: &str, this: &Value) -> Result<Option<Updated>> {
    let prev = match vars.get(name) {
        Some(v) => v,
        None => return Err(Error::NotDefined(name.to_string())),
    };

    match (prev, this) {
        (Value::Int(_), Value::Int(_)) => update(vars, namespace, name, prev.clone(), this),
        (Value::Float(_), Value::Float(_)) => update(vars, namespace, name, prev.clone(), this),
        (Value::String(_), Value::String(_)) => update(vars, namespace, name, prev.clone(), this),
        (Value::Bool(_), Value::Bool(_)) => update(vars, namespace, name, prev.clone(), this),
        (p, t) => Err(Error::InvalidType(p.kind().to_string(), t.kind().to_string())),
    }
}



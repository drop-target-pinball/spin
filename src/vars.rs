use crate::prelude::*;
use crate::{Result, Error};
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

    pub fn kind(&self) -> &str {
        match self {
            Self::Int(_) => "int",
            Self::Float(_) => "float",
            Self::String(_) => "string",
            Self::Bool(_) => "bool",
            Self::Vars(_) => "vars",
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
            Value::Vars(v) => write!(f, "{{ {:?} }}", v),
        }
    }
}

pub type Namespaces = HashMap<String, Vec<VarDef>>;
pub type Vars = HashMap<String, Value>;

fn update(vars: &mut HashMap<String, Value>, name: &str, prev: Value, this: &Value) -> Result<Updated> {
    let msg = Updated{
        name: name.to_string(),
        was: prev,
        value: this.clone()
    };

    vars.insert(name.to_string(), this.clone());
    Ok(msg)
}

pub fn define(queue: &mut Queue, vars: &mut Vars, spaces: &HashMap<String, HashMap<String, VarDef>>, name: &str, kind: &VarKind) {
    if vars.contains_key(name) {
        fault!(queue, "variable already defined: {}", name);
        return;
    }
    let value = match kind {
        VarKind::Int(i) => Value::Int(*i),
        VarKind::Float(f) => Value::Float(*f),
        VarKind::String(s) => Value::String(s.clone()),
        VarKind::Bool(b) => Value::Bool(*b),
        VarKind::Namespace{name} => {
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

fn set_var(vars: &mut HashMap<String, Value>, name: &str, this: &Value) -> Result<Updated> {
    let prev = match vars.get(name) {
        Some(v) => v,
        None => return Err(Error::NotDefined(name.to_string())),
    };

    match (prev, this) {
        (Value::Int(_), Value::Int(_)) => update(vars, name, prev.clone(), this),
        (Value::Float(_), Value::Float(_)) => update(vars, name, prev.clone(), this),
        (Value::String(_), Value::String(_)) => update(vars, name, prev.clone(), this),
        (Value::Bool(_), Value::Bool(_)) => update(vars, name, prev.clone(), this),
        (Value::Vars(_), Value::Vars(_)) => {
            raise!(Error::InvalidArgument, "cannot set '{}', use set_ns instead", name)
        },
        (p, t) => Err(Error::InvalidType(p.kind().to_string(), t.kind().to_string())),
    }
}

pub fn set(vars: &mut HashMap<String, Value>, maybe_ns: &Option<String>, name: &str, this: &Value) -> Result<Updated> {
    let ns= match maybe_ns {
        Some(ns) => ns,
        None => return set_var(vars, name, this)
    };
    let Some(ns_val) = vars.get_mut(ns) else {
        return raise!(Error::InvalidArgument, "no such namespace: {}", ns);
    };

    let ns_kind = ns_val.kind().to_string();
    match { ns_val } {
        &mut Value::Vars(ref mut sub_vars) => {
            set_var(sub_vars, name, this)
        }
        _ => {
            Err(Error::InvalidType("vars".to_string(), ns_kind))
        }
    }
}


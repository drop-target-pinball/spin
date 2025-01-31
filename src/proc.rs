use crate::prelude::*;

use std::env;
use mlua::prelude::*;
use mlua::{Function, Table, Value};

static SCRIPTS: [(&str, &[u8]); 2] = [
    ("spin.lua", include_bytes!("spin.lua")),
    ("engine.lua", include_bytes!("engine.lua")),
];

pub struct Env {
    lua: Lua,
    process: Function,
}

impl Env {
    pub fn new(conf: &config::App) -> Result<Env> {
        // Setup path for use when loading project-specific files
        let root = conf.app_dir.to_string_lossy();
        env::set_var("LUA_PATH",
        format!("{}/proc/?.lua;{}/proc/?/?.lua", root, root));

        let lua = Lua::new();
        for (name, data) in SCRIPTS {
            let chunk = lua.load(data).set_name(name);
            if let Err(e) = chunk.exec() {
                return raise!(Error::ProcExec, "{}", e);
            }
        }

        let globals = lua.globals();
        let lua_engine: Table = match globals.get("engine") {
            Ok(p) => p,
            Err(_) => return raise!(Error::ProcEnv, "'engine' not found in globals")
        };

        let process: Function = match lua_engine.get("process") {
            Ok(p) => p,
            Err(_) => return raise!(Error::ProcEnv, "'process' not found in 'engine'")
        };

        let lua_spin: Table = match globals.get("spin") {
            Ok(p) => p,
            Err(_) => return raise!(Error::ProcEnv, "'spin' not found in globals")
        };

        let lua_conf = match lua.to_value(&conf) {
            Ok(v) => v,
            Err(e) => return raise!(Error::ProcEnv, "unable to convert config: {}", e)
        };

        if let Err(e) =  lua_spin.set("conf", lua_conf) {
            return raise!(Error::ProcEnv, "unable to set config: {}", e);
        }

        Ok(Env{lua, process})
    }

    pub fn load(&self, name: &str, data: &[u8]) -> Result<()> {
        let chunk = self.lua.load(data).set_name(name);
        match chunk.exec() {
            Ok(_) => Ok(()),
            Err(e) => raise!(Error::ProcExec, "{}", e)
        }
    }

    pub fn process(&self, msg: &Message) -> Result<Vec<Message>> {
        let lua_msg = match self.lua.to_value(&msg) {
            Ok(m) => m,
            Err(e) => return raise!(Error::ProcExec, "cannot convert message to lua table: {}", e)
        };

        let result = match self.process.call::<Value>(&lua_msg) {
            Ok(r) => r,
            Err(e) => return raise!(Error::ProcExec, "{}", e)
        };

        let rets = match result {
            Value::Table(t) => t,
            Value::Nil => return Ok(Vec::new()),
            _ => return raise!(Error::ProcExec, "invalid lua return type: {:?}", result)
        };

        let mut msgs: Vec<Message> = Vec::new();

        for ret in rets.sequence_values::<Value>() {
            match ret {
                Err(e) => return raise!(Error::ProcExec, "expected table in returns: {}", e),
                Ok(tbl) => {
                    match self.lua.from_value(tbl) {
                        Ok(m) => msgs.push(m),
                        Err(e) => return raise!(Error::ProcExec, "invalid return value: {}", e),
                    }
                }
            }
        }
        Ok(msgs)
    }

    pub fn lua(&mut self) -> &mut Lua {
        &mut self.lua
    }
}

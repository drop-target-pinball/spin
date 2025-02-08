use crate::prelude::*;

use std::env;
use std::sync::{Arc, Mutex};
use mlua::prelude::*;

static SCRIPTS: [(&str, &[u8]); 4] = [
    ("spin.lua", include_bytes!("spin.lua")),
    ("std.lua", include_bytes!("std.lua")),
    ("message.lua", include_bytes!("message.lua")),
    ("game.lua", include_bytes!("std/scripts/game.lua")),
];

pub struct Env {
    lua: Lua,
    state: Arc<Mutex<State>>,
    spin: LuaTable,
    post: LuaFunction,
}

impl Env {
    pub fn new(state: Arc<Mutex<State>>) -> SpinResult<Env> {
        let s = unwrap!(state.lock());
        // Setup path for use when loading project-specific files
        let root = s.conf.app_dir.to_string_lossy();
        env::set_var("LUA_PATH",
        format!("{}/scripts/?.lua;{}/scripts/?/?.lua", root, root));

        let lua = Lua::new();
        for (name, data) in SCRIPTS {
            let chunk = lua.load(data).set_name(name);
            if let Err(e) = chunk.exec() {
                return raise!(Error::ScriptExec, "{}", e);
            }
        }

        let globals = lua.globals();
        let spin: LuaTable = match globals.get("spin") {
            Ok(p) => p,
            Err(_) => return raise!(Error::ScriptEnv, "'spin' not found in globals")
        };

        let post: LuaFunction = match spin.get("post") {
            Ok(p) => p,
            Err(_) => return raise!(Error::ScriptEnv, "'post' not found in 'spin'")
        };

        let lua_conf = match lua.to_value(&s.conf) {
            Ok(v) => v,
            Err(e) => return raise!(Error::ScriptEnv, "unable to convert config: {}", e)
        };

        if let Err(e) = spin.set("conf", lua_conf) {
            return raise!(Error::ScriptEnv, "unable to set config: {}", e);
        }

        drop(s);
        Ok(Env{lua, state, spin, post})
    }

    pub fn send_vars(&self) -> SpinResult<()> {
        let vars = &mut unwrap!(self.state.lock()).vars;
        let lua_vars= match self.lua.to_value(vars) {
            Ok(v) => v,
            Err(e) => return raise!(Error::ScriptEnv, "unable to convert vars: {}", e),
        };

        match self.spin.set("vars", &lua_vars) {
            Ok(()) => Ok(()),
            Err(e) => raise!(Error::ScriptEnv, "unable to send vars: {}", e)
        }
    }

    pub fn recv_vars(&self) -> SpinResult<()> {
        // let vars_box = &mut unwrap!(self.vars.lock());

        // let lua_vars: LuaValue = match self.spin.get("vars") {
        //     Ok(v) => v,
        //     Err(e) => return raise!(Error::ScriptEnv, "unable to receive vars: {}", e)
        // };

        // let vars = match self.lua.from_value(lua_vars) {
        //     Ok(v) => v,
        //     Err(e) => return raise!(Error::ScriptEnv, "unable to convert vars: {}", e)
        // };
        // vars_box.vars = vars;
        Ok(())
    }

    pub fn load_string(&self, name: &str, data: &str) -> LuaChunk {
        self.lua.load(data.to_string()).set_name(name)
    }

    pub fn exec(&self, name: &str, data: &[u8]) -> SpinResult<()> {
        let chunk = self.lua.load(data).set_name(name);
        match chunk.exec() {
            Ok(_) => Ok(()),
            Err(e) => raise!(Error::ScriptExec, "{}", e)
        }
    }

    pub fn process(&self, msg: &Message) -> SpinResult<Vec<Message>> {
        let lua_msg = match self.lua.to_value(&msg) {
            Ok(m) => m,
            Err(e) => return raise!(Error::ScriptExec, "cannot convert message to lua table: {}", e)
        };

        let results = match self.post.call::<LuaMultiValue>(&lua_msg) {
            Ok(r) => r,
            Err(e) => return raise!(Error::ScriptExec, "{}", e)
        };

        let result = &results[0];
        let rets = match result {
            LuaValue::Table(t) => t,
            LuaValue::Nil => return Ok(Vec::new()),
            _ => return raise!(Error::ScriptExec, "invalid lua return type: {:?}", result)
        };

        let mut msgs: Vec<Message> = Vec::new();

        for ret in rets.sequence_values::<LuaValue>() {
            match ret {
                Err(e) => return raise!(Error::ScriptExec, "expected table in returns: {}", e),
                Ok(tbl) => {
                    match self.lua.from_value(tbl) {
                        Ok(m) => msgs.push(m),
                        Err(e) => return raise!(Error::ScriptExec, "invalid return value: {}", e),
                    }
                }
            }
        }
        Ok(msgs)
    }
}

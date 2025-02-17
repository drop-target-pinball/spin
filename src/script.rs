use crate::prelude::*;

use std::env;
use std::sync::{Arc, Mutex};
use mlua::prelude::*;
use crate::{Error, Result};

static SCRIPTS: [(&str, &[u8]); 9] = [
    ("std.lua", include_bytes!("std.lua")),
    ("check.lua", include_bytes!("check.lua")),
    ("render.lua", include_bytes!("render.lua")),
    ("spin.lua", include_bytes!("spin.lua")),
    ("message.lua", include_bytes!("message.lua")),
    ("test.lua", include_bytes!("test.lua")),

    ("dmd.lua", include_bytes!("std/scripts/dmd.lua")),
    ("game.lua", include_bytes!("std/scripts/game.lua")),
    ("service.lua", include_bytes!("std/scripts/service.lua")),
];

static TEST_SCRIPTS: [(&str, &[u8]); 1] = [
    ("game_test.lua", include_bytes!("../tests/game_test.lua")),
];

pub struct Env {
    lua: Lua,
    state: Arc<Mutex<State>>,
    spin: LuaTable,
    render: LuaTable,
    post: LuaFunction,
}

impl Env {
    pub fn new(state: Arc<Mutex<State>>) -> Result<Env> {
        let s = unwrap!(state.lock());
        // Setup path for use when loading project-specific files
        let root = s.runtime.dirs.app.to_string_lossy();
        env::set_var("LUA_PATH",
        format!("{}/scripts/?.lua;{}/scripts/?/?.lua", root, root));

        let lua = unsafe { Lua::unsafe_new() };
        for (name, data) in SCRIPTS {
            let chunk = lua.load(data).set_name(name);
            if let Err(e) = chunk.exec() {
                return raise!(Error::ScriptExec, "{}", e);
            }
        }
        if !s.runtime.is_release() {
            for (name, data) in TEST_SCRIPTS {
                let chunk = lua.load(data).set_name(name);
                if let Err(e) = chunk.exec() {
                    return raise!(Error::ScriptExec, "{}", e);
                }
            }
        }

        let globals = lua.globals();
        let spin: LuaTable = match globals.get("spin") {
            Ok(p) => p,
            Err(_) => return raise!(Error::ScriptEnv, "'spin' not found in globals")
        };

        let render: LuaTable = match globals.get("_render") {
            Ok(p) => p,
            Err(_) => return raise!(Error::ScriptEnv, "'_render' not found in globals")
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

        let lua_runtime = match lua.to_value(&s.runtime) {
            Ok(v) => v,
            Err(e) => return raise!(Error::ScriptEnv, "unable to convert runtime: {}", e)
        };

        if let Err(e) = spin.set("runtime", lua_runtime) {
            return raise!(Error::ScriptEnv, "unable to set runtime: {}", e);
        }

        let init: LuaFunction = match spin.get("_init") {
            Ok(p) => p,
            Err(_) => return raise!(Error::ScriptEnv, "'_init' not found in 'spin'")
        };

        match init.call::<bool>(()) {
            Ok(r) => r,
            Err(e) => return raise!(Error::ScriptExec, "_init failed: {}", e)
        };

        drop(s);
        Ok(Env{lua, state, spin, render, post})
    }

    pub fn send_vars(&self) -> Result<()> {
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

    pub fn recv_vars(&self) -> Result<()> {
        let state = &mut unwrap!(self.state.lock());

        let lua_ops: LuaTable = match self.render.get("ops") {
            Ok(v) => v,
            Err(e) => return raise!(Error::ScriptEnv, "unable to receive vars: {}", e)
        };

        let mut ops: Vec<render::Instruction> = Vec::new();
        for v in lua_ops.sequence_values::<LuaValue>() {
            match v {
                Err(e) => return raise!(Error::ScriptExec, "expected table in ops: {}", e),
                Ok(tbl) => {
                    let tbl_msg = tbl.clone();
                    match self.lua.from_value(tbl) {
                        Ok(o) => ops.push(o),
                        Err(e) => return raise!(Error::ScriptExec, "invalid return value: {}\n{}", e, value_to_string(&tbl_msg)),
                    }
                }
            }
        }
        state.render_list = ops.clone();
        match lua_ops.clear() {
            Ok(()) => Ok(()),
            Err(e) => raise!(Error::ScriptEnv, "unable to clear ops table: {}", e),
        }
    }

    pub fn load_string(&self, name: &str, data: &str) -> LuaChunk {
        self.lua.load(data.to_string()).set_name(name)
    }

    pub fn exec(&self, name: &str, data: &[u8]) -> Result<()> {
        let chunk = self.lua.load(data).set_name(name);
        match chunk.exec() {
            Ok(_) => Ok(()),
            Err(e) => raise!(Error::ScriptExec, "{}", e)
        }
    }

    pub fn process(&self, msg: &Message) -> Result<Vec<Message>> {
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

fn value_to_string(val: &LuaValue) -> String {
    let mut s = String::new();
    match val {
        LuaValue::Table(tbl) => {
            let mut kvs = Vec::new();
            for pair in tbl.pairs::<LuaValue, LuaValue>() {
                let (k, v) = pair.unwrap();
                kvs.push(format!("{} = {}", value_to_string(&k), value_to_string(&v)));
            }
            s += &format!("{{ {} }}", kvs.join(", "));
        },
        LuaValue::Boolean(b) => s += &format!("{}", b),
        LuaValue::Error(e) => s += &format!("error({})", e.to_string()),
        LuaValue::Function(_) => s += &format!("function()"),
        LuaValue::Integer(i) => s += &format!("{}", i),
        LuaValue::Number(n) => s += &format!("{}", n),
        LuaValue::String(st) => s += &format!("'{}'", st.to_string_lossy()),
        _ => s += "other",
    }
    s
}
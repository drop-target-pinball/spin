use crate::prelude::*;

use std::env;
use std::sync::{Arc, Mutex};
use std::collections::HashMap;
use mlua::prelude::*;
use crate::{Error, Result};

static SCRIPTS: [(&str, &[u8]); 8] = [
    ("std.lua", include_bytes!("std.lua")),
    ("check.lua", include_bytes!("check.lua")),
    ("render.lua", include_bytes!("render.lua")),
    ("spin.lua", include_bytes!("spin.lua")),
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
            try_script!(chunk.exec());
        }
        if !s.runtime.is_release() {
            for (name, data) in TEST_SCRIPTS {
                let chunk = lua.load(data).set_name(name);
                try_script!(chunk.exec());
            }
        }

        let globals = lua.globals();
        let spin: LuaTable = try_script!(globals.get("spin"));
        let render: LuaTable = try_script!(globals.get("_render"));
        let post: LuaFunction = try_script!(spin.get("post"));

        let lua_conf: LuaTable = try_script!(spin.get("conf"));
        let conf_value: LuaValue = try_script!(lua.to_value(&s.conf));
        let conf_table = match conf_value {
            LuaValue::Table(t) => t,
            _ => return raise!(Error::Script, "not a table"),
        };
        for pair in  conf_table.pairs::<String, LuaValue>() {
            let (k, v) = try_script!(pair);
            try_script!(lua_conf.set(k, v));
        }

        let lua_runtime = try_script!(lua.to_value(&s.runtime));
        try_script!(spin.set("runtime", lua_runtime));

        let lua_switches = try_script!(lua.to_value(&s.switches));
        try_script!(spin.set("switches", lua_switches));

        let init: LuaFunction = try_script!(spin.get("_init"));
        try_script!(init.call::<bool>(()));

        drop(s);
        Ok(Env{lua, state, spin, render, post})
    }

    pub fn send_vars(&self) -> Result<()> {
        let s  = &mut self.state.lock().unwrap();

        let mut vars: LuaTable = try_script!(self.spin.get("raw_vars"));
        let mut settings: LuaTable = try_script!(self.spin.get("raw_settings"));
        let players: LuaTable = try_script!(self.spin.get("raw_players"));

        try_script!(send_var_group(&s.vars, &mut vars));
        try_script!(send_var_group(&s.settings, &mut settings));
        for (i, player) in s.players.iter().enumerate() {
            let mut lua_player: LuaTable = try_script!(players.get(i + 1));
            try_script!(send_var_group(player, &mut lua_player));
        }

        let lua_switches = try_script!(self.lua.to_value(&s.switches));
        try_script!(self.spin.set("switches", lua_switches));
        Ok(())
    }

    pub fn recv_vars(&self) -> Result<()> {
        let state = &mut unwrap!(self.state.lock());

        let lua_ops: LuaTable = try_script!(self.render.get("ops"));
        let mut ops: Vec<render::Instruction> = Vec::new();
        for v in lua_ops.sequence_values::<LuaValue>() {
            match v {
                Err(e) => return raise!(Error::Script, "expected table in ops: {}", e),
                Ok(tbl) => {
                    let tbl_msg = tbl.clone();
                    match self.lua.from_value(tbl) {
                        Ok(o) => ops.push(o),
                        Err(e) => return raise!(Error::Script, "invalid return value: {}\n{}", e, value_to_string(&tbl_msg)),
                    }
                }
            }
        }
        state.render_ops = ops.clone();
        try_script!(lua_ops.clear());

        Ok(())
    }

    pub fn load_string(&self, name: &str, data: &str) -> LuaChunk {
        self.lua.load(data.to_string()).set_name(name)
    }

    pub fn exec(&self, name: &str, data: &[u8]) -> Result<()> {
        let chunk = self.lua.load(data).set_name(name);
        try_script!(chunk.exec());
        Ok(())
    }

    pub fn process(&self, msg: &Message) -> Result<Vec<Message>> {
        let elapsed = self.state.lock().unwrap().elapsed;

        let lua_elapsed = try_script!(self.lua.to_value(&elapsed));
        let lua_msg = try_script!(self.lua.to_value(&msg));

        let results = try_script!(self.post.call::<LuaMultiValue>((&lua_elapsed, &lua_msg)));
        let result = &results[0];
        let rets = match result {
            LuaValue::Table(t) => t,
            LuaValue::Nil => return Ok(Vec::new()),
            _ => return raise!(Error::Script, "invalid lua return type: {:?}", result)
        };

        let mut msgs: Vec<Message> = Vec::new();
        for ret in rets.sequence_values::<LuaValue>() {
            match ret {
                Err(e) => return raise!(Error::Script, "expected table in returns: {}", e),
                Ok(tbl) => {
                    match self.lua.from_value(tbl) {
                        Ok(m) => msgs.push(m),
                        Err(e) => return raise!(Error::Script, "invalid return value: {}", e),
                    }
                }
            }
        }
        Ok(msgs)
    }
}

fn send_var_group(vars: &HashMap<String,vars::Value>, lua_table: &mut LuaTable) -> Result<()> {
    for (name, value) in vars {
        let name = name.to_string();
        match value {
            vars::Value::Int(i) => try_script!(lua_table.set(name, *i)),
            vars::Value::Float(f) => try_script!(lua_table.set(name, *f)),
            vars::Value::String(s) => try_script!(lua_table.set(name, s.clone())),
            vars::Value::Bool(b) => try_script!(lua_table.set(name, *b)),
        }
    }
    Ok(())
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
        LuaValue::Error(e) => s += &format!("error({})", e),
        LuaValue::Function(_) => s += "function()",
        LuaValue::Integer(i) => s += &format!("{}", i),
        LuaValue::Number(n) => s += &format!("{}", n),
        LuaValue::String(st) => s += &format!("'{}'", st.to_string_lossy()),
        _ => s += "other",
    }
    s
}
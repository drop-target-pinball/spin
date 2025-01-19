use std::cmp::max;
use std::time::{Duration, Instant};
use std::thread;
use std::path::Path;

use mlua::prelude::*;

fn main() -> LuaResult<()> {
    let lua = Lua::new();
    let chunk = lua.load(Path::new("./src/main.lua"));

    if let Err(e) = chunk.exec() {
        println!("error: {}", e);
        return Err(e);
    }

    loop {
        let begin = Instant::now();
        let next = begin + Duration::from_millis(1000);

        lua.load("tick()").exec()?;

        let end = Instant::now();
        let sleep = max(Duration::ZERO, next - end);
        thread::sleep(sleep);
    }
}
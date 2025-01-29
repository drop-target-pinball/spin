use crate::prelude::*;
use mlua::Lua;

pub fn init_lua(e: &mut Engine, lua: &mut Lua) {
    let mut dev_sdl= Box::new(sdl::SdlDevice::default()
        .with_audio(0,sdl::AudioOptions::default()));
    e.add_device(&mut dev_sdl);
}
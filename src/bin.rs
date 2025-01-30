use std::time::Duration;
use clap::{Parser, crate_name, crate_description, crate_version};
use std::thread;

use spin::prelude::*;

#[derive(Parser)]
struct Cli {
    #[arg(short, long)]
    test: bool,

    #[arg(short, long)]
    release: bool
}

pub fn main()  {
    let cli = Cli::parse();

    let mode = if cli.release {
        config::RunMode::Release
    } else if cli.test {
        config::RunMode::Test
    } else {
        config::RunMode::Develop
    };

    let mut conf = config::new(mode);
    conf.server.enabled = true;
    conf.sounds.push(config::Sound::new("foo", "example/swing.ogg"));

    let mut e = Engine::new(&conf);

    let dev_sdl = sdl::SdlDevice::default()
        .with_audio(0,sdl::AudioOptions::default());

    e.add_device(Box::new(dev_sdl));

    if mode == config::RunMode::Release {
        let logger = Logger::default();
        e.add_device(Box::new(logger));
    } else {
        let console = builtin::Console::new(e.state());
        e.add_device(Box::new(console));
    }

    e.run();
    println!("\n");
}


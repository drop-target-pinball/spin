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

    let dev_sdl = sdl::SdlDevice::default()
        .with_audio(0,sdl::AudioOptions::default());
    // let mut logger = Logger::default();

    let mut e = Engine::new(&conf);
    e.add_device(Box::new(dev_sdl));
    // e.add_device(&mut logger);

    let cons = builtin::Console::new(e.state());
    e.add_device(Box::new(cons));

    let q = e.queue();
    info!(q, "{}: {}, version {}", crate_name!(), crate_description!(), crate_version!());

    e.run();
    println!("\n");
}


use std::time::Duration;
use clap::{Parser, crate_name, crate_description, crate_version};

use spin::prelude::*;

#[derive(Parser)]
struct Cli {
    #[arg(short, long)]
    test: bool,

    #[arg(short, long)]
    release: bool
}

pub fn main() {
    let cli = Cli::parse();

    let mode = if cli.release {
        RunMode::Release
    } else if cli.test {
        RunMode::Test
    } else {
        RunMode::Develop
    };

    let mut conf = Config::new(mode);
    conf.sounds.push(Sound::new("foo", "sample/swing.ogg"));

    let dev_sdl = sdl::SdlDevice::default()
        .with_audio(0,sdl::AudioOptions::default());
    let logger = Logger::default();

    let mut e = Engine::new(conf);
    e.add_device(Box::new(dev_sdl));
    e.add_device(Box::new(logger));

    info!(e.queue, "{}: {}, version {}", crate_name!(), crate_description!(), crate_version!());
    e.init();

    info!(e.queue, "ready");
    e.tick();

    e.queue.push(Message::PlaySound(PlayAudio{name: "foo".to_string() }));
    e.tick();
    std::thread::sleep(Duration::from_secs(5));
}
use std::{process::ExitCode, time::Duration};
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

pub fn main() -> ExitCode {
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

    let mut dev_sdl = sdl::SdlDevice::default()
        .with_audio(0,sdl::AudioOptions::default());
    let mut logger = Logger::default();

    let mut e = Engine::new(conf);
    e.add_device(&mut dev_sdl);
    e.add_device(&mut logger);

    info!(e.queue(), "{}: {}, version {}", crate_name!(), crate_description!(), crate_version!());

    let q = e.queue();
    thread::spawn(move || {
        std::thread::sleep(Duration::from_secs(1));
        q.post(Message::PlaySound(PlayAudio{name: "foo".to_string() }));
    });

    match e.run() {
        Ok(()) => ExitCode::SUCCESS,
        Err(e) => {
            eprintln!("error: {}", e);
            ExitCode::FAILURE
        }
    }

}
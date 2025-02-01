use clap::{Parser, crate_name, crate_description, crate_version};
use std::process::ExitCode;

use spin::prelude::*;

#[derive(Parser)]
struct Cli {
    #[arg(short, long)]
    test: bool,

    #[arg(short, long)]
    release: bool
}

pub fn main() -> ExitCode  {
    let cli = Cli::parse();

    let mode = if cli.release {
        config::RunMode::Release
    } else if cli.test {
        config::RunMode::Test
    } else {
        config::RunMode::Develop
    };

    let conf_dir = config::app_dir().join("config");
    let mut conf = match config::load(&conf_dir) {
        Ok(c) => c,
        Err(e) => {
            eprintln!("{}", e);
            return ExitCode::FAILURE;
        }
    };
    conf.app_dir = config::app_dir();

    let mut e = Engine::new(&conf);

    let dev_sdl = sdl::SdlDevice::default()
        .with_audio(0,sdl::AudioOptions::default());
    e.add_device(Box::new(dev_sdl));

    if mode == config::RunMode::Release {
        let logger = builtin::Logger::default();
        e.add_device(Box::new(logger));
    } else {
        let console = builtin::Console::new(e.state());
        e.add_device(Box::new(console));
    }

    info!(e.queue(), "{}: {}, version {}", crate_name!(), crate_description!(), crate_version!());
    e.run();
    println!("");

    ExitCode::SUCCESS
}


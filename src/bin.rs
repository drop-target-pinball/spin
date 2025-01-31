use clap::{Parser, crate_name, crate_description, crate_version};
use std::fs;
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

    let conf_file = config::app_dir().join("config.yaml");
    let conf_text = match fs::read_to_string(&conf_file) {
        Ok(f) => f,
        Err(e) => {
            eprintln!("error: unable to read config file '{}': {}", &conf_file.to_string_lossy(), e);
            return ExitCode::FAILURE;
        }
    };

    let mut conf: config::App = match serde_yml::from_str(&conf_text) {
        Ok(c) => c,
        Err(e) => {
            eprintln!("configuration error: {}", e);
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


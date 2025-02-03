use clap::{Parser, crate_name, crate_description, crate_version};
use std::{os::fd::AsRawFd, process::ExitCode};

use spin::prelude::*;

#[derive(Parser)]
struct Cli {
    #[arg(short, long)]
    test: bool,

    #[arg(short, long)]
    release: bool
}

pub fn main() -> ExitCode  {
    // let v = config::Var{
    //     name: "test".to_string(),
    //     value: Value::Int(0),
    //     kind: config::VarKind::Int{default: 0},
    // };
    // let x = serde_yml::to_string(&v).unwrap();
    // println!("{}", x);
    // return ExitCode::SUCCESS;

    let cli = Cli::parse();

    let mode = if cli.release {
        config::RunMode::Release
    } else if cli.test {
        config::RunMode::Test
    } else {
        config::RunMode::Develop
    };

    let conf = match config::load(&config::app_dir()) {
        Ok(c) => c,
        Err(e) => {
            eprintln!("{}", e);
            return ExitCode::FAILURE;
        }
    };

    let mut e = Engine::new(&conf);

    let dev_sdl = sdl::SdlDevice::default()
        .with_audio(0,sdl::AudioOptions::default());
    e.add_device(Box::new(dev_sdl));

    let store = builtin::Store::new();
    e.add_device(Box::new(store));
    let validator = builtin::Validator::new();
    e.add_device(Box::new(validator));

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

    let tos = termios::Termios::from_fd(0).unwrap();
    termios::tcsetattr(std::io::stdin().as_raw_fd(), termios::TCSADRAIN, &tos).unwrap();

    ExitCode::SUCCESS
}


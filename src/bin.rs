use clap::{Parser, crate_name, crate_description, crate_version};
use std::{env, os::fd::AsRawFd, process::ExitCode};

use spin::prelude::*;

#[derive(Parser)]
struct Cli {
    #[arg(short, long)]
    /// testing mode - enable hardware devices
    test: bool,

    #[arg(short, long)]
    /// release mode - headless and panic on fault
    release: bool,

    /// run this script at startup
    run_script: Option<String>
}

pub fn main() -> ExitCode  {
    let cli = Cli::parse();

    let mode = if cli.release {
        RunMode::Release
    } else if cli.test {
        RunMode::Test
    } else {
        RunMode::Develop
    };

    let dirs = Dirs::default();
    let conf = match load_config(&dirs) {
        Ok(c) => c,
        Err(e) => {
            eprintln!("{}", e);
            return ExitCode::FAILURE;
        }
    };

    let runtime = Runtime {
        prog_name: crate_name!().to_string(),
        prog_description: crate_description!().to_string(),
        prog_version: crate_version!().to_string(),
        prog_date: env::var("SPIN_BUILD_DATE").unwrap_or_default(),
        mode,
        dirs,
    };

    let mut e = Engine::new(conf.clone(), runtime);

    #[cfg(feature = "sdl")] {
        if let Some(sdl_conf) = &conf.sdl {
            let device = crate::sdl::Device::new(&conf, sdl_conf);
            e.add_device(Box::new(device));
        }
    }

    let store = builtin::Store::new();
    e.add_device(Box::new(store));
    let validator = builtin::Validator::default();
    e.add_device(Box::new(validator));

    if mode == RunMode::Release {
        let logger = builtin::Logger::default();
        e.add_device(Box::new(logger));
    } else {
        let console = builtin::Console::new(e.state());
        e.add_device(Box::new(console));
    }

    info!(e.queue(), "{}: {}, version {}", crate_name!(), crate_description!(), crate_version!());
    e.run(cli.run_script);
    println!();

    if mode != RunMode::Release {
        let tos = termios::Termios::from_fd(0).unwrap();
        termios::tcsetattr(std::io::stdin().as_raw_fd(), termios::TCSADRAIN, &tos).unwrap();
    }

    ExitCode::SUCCESS
}


use clap::{Parser, crate_name, crate_description, crate_version};
use std::{env, os::fd::AsRawFd, process::ExitCode};

use spin::prelude::*;

#[derive(Parser)]
struct Cli {
    #[arg(long)]
    /// print filenames of config files being loaded
    debug_config: bool,

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
        RunMode::PlayTest
    } else {
        RunMode::Develop
    };

    let dirs = Dirs::default();
    let mut runtime = Runtime::new(dirs);

    runtime.debug_config = cli.debug_config;
    runtime.prog_name = crate_name!().to_string();
    runtime.prog_description = crate_description!().to_string();
    runtime.prog_version = crate_version!().to_string();
    runtime.prog_date = env::var("SPIN_BUILD_DATE").unwrap_or_default();
    runtime.mode = mode;

    let conf = match load_config(&runtime) {
        Ok(c) => c,
        Err(e) => {
            eprintln!("{}", e);
            return ExitCode::FAILURE;
        }
    };

    let mut e = Engine::new(conf.clone(), runtime.clone());

    #[cfg(feature = "sdl")] {
        if conf.sdl.is_some() {
            let device = crate::sdl::Device::new(&conf, &runtime);
            e.add_device(Box::new(device));
        }
    }

    let store = Store::new();
    e.add_device(Box::new(store));
    let validator = Validator::default();
    e.add_device(Box::new(validator));

    if mode == RunMode::Release {
        let logger = Logger::default();
        e.add_device(Box::new(logger));
    } else {
        let console = Console::new(e.state());
        e.add_device(Box::new(console));
    }

    if mode == RunMode::Develop || mode == RunMode::AutoTest {
        if let Some(mock_conf) = conf.mock {
            let mock_device = mock::Device::new(&mock_conf);
            e.add_device(Box::new(mock_device));
        }
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


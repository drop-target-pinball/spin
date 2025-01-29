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
        RunMode::Release
    } else if cli.test {
        RunMode::Test
    } else {
        RunMode::Develop
    };

    let mut conf = Config::new(mode);

    conf.sounds.push(Sound::new("foo", "example/swing.ogg"));

    let mut dev_sdl = sdl::SdlDevice::default()
        .with_audio(0,sdl::AudioOptions::default());
    let mut logger = Logger::default();

    let mut e = Engine::new(conf);
    e.add_device(&mut dev_sdl);
    e.add_device(&mut logger);

    let q = e.queue();
    info!(q, "{}: {}, version {}", crate_name!(), crate_description!(), crate_version!());

    #[cfg(feature = "server")] {
        if mode != RunMode::Release {
            //let svr = server::new(e.queue(), "0.0.0.0:7746");
            use rocket::tokio::runtime::Runtime;
            thread::spawn(move || {
                let rt = Runtime::new().unwrap();
                rt.block_on( async move {
                    server::run().await
                });
            });
        }
    }

    thread::spawn(move || {
        std::thread::sleep(Duration::from_secs(1));
        q.post(Message::Run(Run{name: "hello".to_string()}));
    });

    e.run();

}
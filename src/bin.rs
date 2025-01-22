use std::time::Duration;

use spin::prelude::*;

pub fn main() {
    println!("hello world");

    let mut conf = Config::new(RunMode::Develop);
    conf.sounds.push(Sound::new("foo", "sample/swing.ogg"));

    let dev_sdl = sdl::SdlDevice::new().unwrap()
        .with_audio(0,sdl::AudioOptions::default()).unwrap();
    let logger = Logger::default();

    let mut e = Engine::new(conf);
    e.add_device(Box::new(dev_sdl));
    e.add_device(Box::new(logger));

    e.init();

    e.queue.push(Message::PlaySound(PlayAudio{name: "foo".to_string() }));
    e.tick();
    std::thread::sleep(Duration::from_secs(1));
}
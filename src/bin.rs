use spin::prelude::*;

pub fn main() {
    println!("hello world");

    let mut conf = Config::new(RunMode::Develop);
    conf.sounds.push(Sound::new("foo", "tmp/foo.wav"));

    let dev_sdl = sdl::SdlDevice::new().unwrap()
        .with_audio(0,sdl::AudioOptions::default()).unwrap();
    let logger = Logger::default();

    let mut e = Engine::new(conf);
    e.add_device(Box::new(dev_sdl));
    e.add_device(Box::new(logger));

    e.init();

}
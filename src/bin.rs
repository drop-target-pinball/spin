use spin::prelude::*;

pub fn main() {
    println!("hello world");

    let mut dev_sdl = sdl::Device::new(0).unwrap();
    let mut audio = sdl::Audio::new();
    dev_sdl.add_system(&mut audio);


    let mut conf = Config::new(RunMode::Devel);
    conf.add_sound(&Sound::new("test", "test.wav"));

    let mut logger = Logger::default();
    let mut e = Engine::new(conf);

    e.add_device(&mut dev_sdl);
    e.add_device(&mut logger);

    e.init();

}
use spin::prelude::*;
use std::time::{Duration, Instant};
use std::thread;

fn run_lua_test(eng: &mut Engine, name: &str) -> Option<String> {
    let run_start = Instant::now();
    let rate = Duration::from_micros(16670);
    eng.main = name.to_string();
    eng.shutdown = false;

    let queue = eng.queue();
    queue.post(Message::Halt);
    eng.tick(run_start.elapsed());
    queue.post(Message::Run(Name{name: name.to_string()}));

    let mut i = 0;
    while !eng.shutdown {
        i = i + 1;
        if i > 60 * 60 * 5 {
            panic!("test timeout exceeded")
        }
        let frame_start = Instant::now();
        eng.tick(run_start.elapsed());

        let frame_time = frame_start.elapsed();
        if let Some(remaining) = rate.checked_sub(frame_time) {
            thread::sleep(remaining);
        }
    }
    eng.error()
}

pub fn main() {
    let mut runtime = Runtime::new(Dirs::default());
    runtime.mode = RunMode::AutoTest;
    let conf = match load_config(&runtime) {
        Ok(c) => c,
        Err(e) => panic!("{}", e),
    };

    let mut eng = Engine::new(conf.clone(), runtime.clone());

    #[cfg(feature = "sdl")] {
        if conf.sdl.is_some() {
            let device = crate::sdl::Device::new(&conf, &runtime);
            eng.add_device(Box::new(device));
        }
    }

    let store = builtin::Store::new();
    eng.add_device(Box::new(store));
    let validator = builtin::Validator::default();
    eng.add_device(Box::new(validator));
    let logger = builtin::Logger::default();
    eng.add_device(Box::new(logger));
    eng.init();

    let mut test_count = 0;
    let mut test_passed = 0;
    let mut test_failed = 0;

    for (name, def) in conf.scripts {
        if !def.test {
            continue
        }

        println!("running test: {}", name);
        match run_lua_test(&mut eng, &name) {
            None => {
                println!("pass: {}", name);
                test_count += 1;
                test_passed += 1;
            },
            Some(e) => {
                println!("fail: {}, reason: {}", name, e);
                test_count += 1;
                test_failed += 1;
            }
        }
    }

    println!("{} test(s), {} passed, {} failed", test_count, test_passed, test_failed);
    if test_failed > 0 {
        panic!("test failed");
    }
}




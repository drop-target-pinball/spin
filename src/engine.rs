use crate::prelude::*;
use std::sync::mpsc::Receiver;
use std::sync::Arc;
use std::{
    sync::mpsc::{self, TryRecvError},
    thread,
    time::{self, Duration},
};
use std::sync::Mutex;
use std::collections::HashMap;


pub struct State {
    pub conf: AppConfig,
    pub queue: Queue,
    pub vars: vars::Vars,
    pub render_list: Vec<render::Instruction>,
}


pub struct Engine<'a> {
    queue: Queue,
    state: Arc<Mutex<State>>,
    r_state: render::State,
    script_env: script::Env,

    pub rx: Receiver<Message>,
    devices: Vec<Box<dyn Device + 'a>>,
    shutdown: bool,
}

impl<'a> Engine<'a> {
    pub fn new(conf: AppConfig) -> Self {
        let mut videos = HashMap::new();
        for (name, c) in &conf.video {
            videos.insert(name.to_string(), Video::new(&c));
        }

        let (tx, rx) = mpsc::channel();
        let queue = Queue::new(tx);
        let state = Arc::new(Mutex::new(State {
            conf,
            queue: queue.clone(),
            vars: vars::Vars::new(),
            render_list: Vec::new(),
        }));
        let script_env = unwrap!(script::Env::new(state.clone()));

        Self {
            queue,
            state,
            r_state: render::State::default(),
            rx,
            devices: Vec::new(),
            script_env,
            shutdown: false,
        }
    }

    pub fn add_device(&mut self, d: Box<dyn Device>) {
        self.devices.push(d)
    }

    pub fn queue(&self) -> Queue {
        self.queue.clone()
    }

    pub fn state(&self) -> Arc<Mutex<State>> {
        self.state.clone()
    }

    pub fn init(&mut self) {
        for d in &mut self.devices {
            let mut s = unwrap!(self.state.lock());
            d.init(&mut Globals { s: &mut s, r: &mut self.r_state });
        }
        info!(self.queue, "ready");
        self.process_queue(time::Duration::ZERO);
    }

    pub fn tick(&mut self, elapsed: time::Duration) {
        self.queue.post(Message::Poll);
        self.process_queue(elapsed);
        self.queue.post(Message::Tick);
        self.queue.post(Message::Present);
        self.process_queue(elapsed);
    }

    pub fn run(&mut self, init_script: Option<String>) {
        let run_start = time::Instant::now();
        let rate = Duration::from_micros(16670);

        self.init();

        if let Some(init) = init_script {
            let msg = Name{name: init};
            self.queue.post(Message::Run(msg));
        }
        // FIXME: Add in control-c handler for release mode
        // let running = Arc::new(AtomicBool::new(true));
        // let running_2 = running.clone();

        // let result = ctrlc::set_handler(move || {
        //     running_2.store(false, Ordering::SeqCst);
        // });
        // if let Err(e) = result {
        //     panic!("unable to set signal handler: {}", e);
        // }

        while /*running.load(Ordering::SeqCst) &&*/ !self.shutdown {
            let frame_start = time::Instant::now();
            self.tick(run_start.elapsed());

            let frame_time = frame_start.elapsed();
            if let Some(remaining) = rate.checked_sub(frame_time) {
                thread::sleep(remaining);
            }
        }

        if !self.shutdown {
            self.queue.post(Message::Shutdown);
            self.tick(run_start.elapsed());
        }
    }

    fn process_queue(&mut self, elapsed: time::Duration) {
        let messages = self.process_queue_rust(elapsed);
        self.process_queue_lua(messages);
    }

    fn process_queue_rust(&mut self, elapsed: time::Duration) -> Vec<Message> {
        let mut state = &mut unwrap!(self.state.lock());
        state.vars.insert("elapsed".to_string(), vars::Value::Int(elapsed.as_millis() as i64));

        let mut messages: Vec<Message> = Vec::new();
        loop {
            if self.shutdown {
                break
            }
            match self.rx.try_recv() {
                Err(TryRecvError::Empty) => break,
                Err(TryRecvError::Disconnected) => panic!("channel closed"),
                Ok(msg) => {
                    for dev in &mut self.devices {
                        dev.process(&mut state, &msg);
                    }
                    match &msg {
                        Message::Note(n) => {
                            if state.conf.is_release() && n.kind == NoteKind::Fault {
                                self.shutdown = true
                            }
                        }
                        Message::Shutdown => self.shutdown = true,
                        _ => (),
                    }
                    messages.push(msg);
                }
            }
        }
        messages
    }

    fn process_queue_lua(&mut self, messages: Vec<Message>) {
        if let Err(e) = self.script_env.send_vars() {
            fault!(self.queue, "{}", e);
            return
        }
        for msg in messages {
            match self.script_env.process(&msg) {
                Ok(returns) => {
                    for ret in returns {
                        self.queue.post(ret);
                    }
                }
                Err(e) => fault!(self.queue, "{}", e),
            }
        }
        if let Err(e) = self.script_env.recv_vars() {
            fault!(self.queue, "{}", e);
        }
    }

}


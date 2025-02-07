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
use std::rc::Rc;
use std::cell::RefCell;

pub struct Env<'a> {
    pub conf: &'a AppConfig,
    pub vars: &'a mut vars::Vars,
    pub videos: &'a mut HashMap<String, Rc<RefCell<Video>>>,
    pub queue: Queue,
}

impl<'a> Env<'a> {
    pub fn new(
        conf: &'a AppConfig,
        vars: &'a mut vars::Vars,
        videos: &'a mut HashMap<String, Rc<RefCell<Video>>>,
        queue: Queue) -> Self
    {
        Self { conf, vars, videos, queue }
    }
}

pub struct State {
    pub conf: AppConfig,
    pub vars_box: Arc<Mutex<vars::VarsBox>>,
    pub queue: Queue,
}

pub trait Device {
    fn process(&mut self, e: &mut Env, msg: &Message);
}

pub struct Engine<'a> {
    conf: AppConfig,
    vars_box: Arc<Mutex<vars::VarsBox>>,
    queue: Queue,
    script_env: script::Env,

    pub rx: Receiver<Message>,
    devices: Vec<Box<dyn Device + 'a>>,
    shutdown: bool,

    videos: HashMap<String, Rc<RefCell<Video>>>,
}

impl<'a> Engine<'a> {
    pub fn new(conf: &AppConfig) -> Self {
        let vars_box = vars::VarsBox{ vars: vars::Vars::new() };
        let arc_vars_box = Arc::new(Mutex::new(vars_box));

        let script_env = unwrap!(script::Env::new(conf, arc_vars_box.clone()));
        let (tx, rx) = mpsc::channel();

        let mut videos = HashMap::new();
        for (name, c) in &conf.video {
            videos.insert(
                name.to_string(),
                Rc::new(RefCell::new(Video::new(&c)))
            );
        }

        Self {
            conf: conf.clone(),
            vars_box: arc_vars_box,
            queue: Queue::new(tx),
            rx,
            devices: Vec::new(),
            videos,
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

    pub fn state(&self) -> State {
        State {
            conf: self.conf.clone(),
            vars_box: self.vars_box.clone(),
            queue: self.queue.clone(),
        }
    }

    pub fn init(&mut self) {
        self.queue.post(Message::Init);
        self.process_queue(time::Duration::ZERO);
        info!(self.queue, "ready");
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
        let vars = &mut unwrap!(self.vars_box.lock()).vars;
        let mut env: Env = Env::new(
            &self.conf, vars,
            &mut self.videos,
            self.queue.clone()
        );
        env.vars.insert("elapsed".to_string(), vars::Value::Int(elapsed.as_millis() as i64));

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
                        dev.process(&mut env, &msg);
                    }
                    match &msg {
                        Message::Note(n) => {
                            if env.conf.is_release() && n.kind == NoteKind::Fault {
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


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
    pub elapsed: i64,
    pub conf: AppConfig,
    pub runtime: Runtime,
    pub queue: Queue,
    pub vars: vars::Vars,
    pub render_ops: Vec<render::Instruction>,
}

pub struct Engine<'a> {
    queue: Queue,
    state: Arc<Mutex<State>>,
    r_state: render::State,
    script_env: script::Env,
    devices: Vec<Box<dyn Device + 'a>>,

    pub rx: Receiver<Message>,
    pub main: String,
    pub shutdown: bool,
}

impl<'a> Engine<'a> {
    pub fn new(conf: AppConfig, runtime: Runtime) -> Self {
        let (tx, rx) = mpsc::channel();
        let queue = Queue::new(tx);

        let mut videos = HashMap::new();
        for (name, c) in &conf.video {
            videos.insert(name.to_string(), Video::new(&c));
        }
        let r_state = render::State{videos};

        let state = Arc::new(Mutex::new(State {
            elapsed: 0,
            conf,
            runtime,
            queue: queue.clone(),
            vars: vars::Vars::new(),
            render_ops: Vec::new(),
        }));

        let script_env = unwrap!(script::Env::new(state.clone()));
        Self {
            queue,
            state,
            r_state,
            rx,
            devices: Vec::new(),
            script_env,
            main: "".to_string(),
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

    pub fn error(&self) -> Option<String> {
        let s = self.state.lock().unwrap();
        s.runtime.error.clone()
    }

    pub fn tick(&mut self, elapsed: time::Duration) {
        self.poll(elapsed);
        self.queue.post(Message::Tick);
        self.process_queue();
        self.render();
        self.present();
    }

    pub fn init(&mut self) {
        let mut s: std::sync::MutexGuard<'_, State> = self.state.lock().unwrap();
        for d in &mut self.devices {
            d.init(&mut s, &mut self.r_state);
        }
    }

    fn poll(&mut self, elapsed: time::Duration) {
        let mut s = self.state.lock().unwrap();
        s.elapsed = elapsed.as_millis() as i64;
        for d in &mut self.devices {
            if let Err(e) = d.poll(&mut s) {
                fault!(s.queue, "{}", e);
            }
        }
    }

    fn render(&mut self) {
        if let Err(e) = self.script_env.recv_vars() {
            fault!(self.queue, "{}", e);
        }
        let mut s = unwrap!(self.state.lock());
        s.render_ops.sort_by_key(|e| e.priority);
        for d in &mut self.devices {
            d.render(&mut s,&mut self.r_state);
        }
        s.render_ops.clear();
    }

    fn present(&mut self) {
        let mut s = unwrap!(self.state.lock());
        for d in &mut self.devices {
            d.present(&mut s, &mut self.r_state);
        }
    }

    pub fn run(&mut self, init_script: Option<String>) {
        let run_start = time::Instant::now();
        let rate = Duration::from_micros(16670);

        self.process_queue();
        self.init();
        info!(self.queue, "ready");

        #[cfg(feature = "debug_fps")]
        diag!(self.queue, "debug_fps enabled");

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

        self.tick(run_start.elapsed());
        while /*running.load(Ordering::SeqCst) &&*/ !self.shutdown {
            let frame_start = time::Instant::now();
            self.tick(run_start.elapsed());

            let frame_time = frame_start.elapsed();
            if let Some(remaining) = rate.checked_sub(frame_time) {
                thread::sleep(remaining);
            } else {
                #[cfg(feature = "debug_fps")]
                diag!(self.queue, "late frame: {} ms", frame_time.as_millis());
            }
        }

        if !self.shutdown {
            self.queue.post(Message::Shutdown);
            self.tick(run_start.elapsed());
        }
    }

    fn process_queue(&mut self) {
        let messages = self.process_queue_rust();
        self.process_queue_lua(messages);
    }

    fn process_queue_rust(&mut self) -> Vec<Message> {
        let mut state = &mut unwrap!(self.state.lock());
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
                            if state.runtime.is_shutdown_on_fault() && n.kind == NoteKind::Fault {
                                state.runtime.error = Some(n.message.clone());
                                self.shutdown = true
                            }
                        }
                        Message::Shutdown => self.shutdown = true,
                        Message::ScriptEnded(m) => {
                            if m.name == self.main {
                                self.shutdown = true
                            }
                        },
                        Message::Halt => {
                            // clear queue
                            loop {
                                match self.rx.try_recv() {
                                    Err(TryRecvError::Empty) => break,
                                    Err(TryRecvError::Disconnected) => panic!("channel closed"),
                                    _ => ()
                                }
                            }
                        }
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
        // if let Err(e) = self.script_env.recv_vars() {
        //     fault!(self.queue, "{}", e);
        // }
    }

}


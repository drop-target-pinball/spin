use crate::prelude::*;
use std::sync::atomic::{AtomicBool, Ordering};
use std::sync::mpsc::Receiver;
use std::sync::Arc;
use std::{
    sync::mpsc::{self, TryRecvError},
    thread,
    time::{self, Duration},
};
use mlua::prelude::*;

static SCRIPTS: [(&str, &[u8]); 2] = [
    ("spin.lua", include_bytes!("spin.lua")),
    ("engine.lua", include_bytes!("engine.lua")),
];

pub struct Env {
    pub conf: config::App,
    pub vars: Vars,
}

impl Env {
    pub fn new(conf: config::App, vars: Vars) -> Self {
        Self { conf, vars }
    }
}

impl Default for Env {
    fn default() -> Self {
        Self {
            conf: config::App::default(),
            vars: Vars::default(),
        }
    }
}

pub trait Device {
    fn process(&mut self, e: &mut Env, q: &mut Queue, msg: &Message);
}

pub struct Engine<'e> {
    pub env: Env,
    pub queue: Queue,
    pub rx: Receiver<Message>,
    devices: Vec<&'e mut dyn Device>,
    lua: Lua,
    lua_process: mlua::Function,
}

impl<'e> Engine<'e> {
    pub fn new(conf: &config::App) -> Self {
        let (lua, lua_process) = match lua_init() {
            Ok((l, p)) => (l, p),
            Err(e) => panic!("lua setup failure: {}", e)
        };
        let env = Env::new(conf.clone(), Vars::default());
        let (tx, rx) = mpsc::channel();
        Engine {
            env,
            queue: Queue::new(tx),
            rx,
            devices: Vec::new(),
            lua,
            lua_process,
        }
    }

    pub fn add_device(&mut self, d: &'e mut dyn Device) {
        self.devices.push(d);
    }

    pub fn queue(&self) -> Queue {
        self.queue.clone()
    }

    pub fn init(&mut self) {
        self.queue.post(Message::Init);
        self.process_queue();
        info!(self.queue, "ready");
    }

    pub fn tick(&mut self) {
        self.env.vars.now = time::Instant::now();
        self.process_queue();
    }

    pub fn run(&mut self) {
        let rate = Duration::from_micros(16670);

        self.init();

        let running = Arc::new(AtomicBool::new(true));
        let running_2 = running.clone();

        let result = ctrlc::set_handler(move || {
            running_2.store(false, Ordering::SeqCst);
        });
        if let Err(e) = result {
            panic!("unable to set signal handler: {}", e);
        }

        while running.load(Ordering::SeqCst) {
            let t0 = time::Instant::now();
            self.tick();

            let elapsed = t0.elapsed();
            let remaining = rate - elapsed;
            if remaining > Duration::ZERO {
                thread::sleep(remaining);
            }
        }

        self.queue.post(Message::Shutdown);
        self.tick();
    }

    fn process_queue(&mut self) {
        // Send each message in the queue to every system for processing.
        loop {
            // As messages are being processed, the systems may want to
            // generate additional messages. Create a new sender for those.
            let mut q = self.queue.clone();

            match self.rx.try_recv() {
                Err(TryRecvError::Empty) => break,
                Err(TryRecvError::Disconnected) => panic!("channel closed"),
                Ok(msg) => {
                    for dev in &mut self.devices {
                        dev.process(&mut self.env, &mut q, &msg);
                    }
                    self.process_lua_message(&msg);
                    Self::process(&mut self.env, &mut q, &msg);
                }
            }
        }
    }

    fn process_lua_message(&mut self, msg: &Message) {
        let lua_msg = match self.lua.to_value(&msg) {
            Ok(m) => m,
            Err(e) => {
                fault!(self.queue, "cannot convert message for lua: {}", e);
                return;
            }
        };

        let result = match self.lua_process.call::<mlua::Value>(&lua_msg) {
            Ok(r) => r,
            Err(e) => {
                fault!(self.queue, "lua execution failed: {}", e);
                return;
            }
        };

        match result {
            mlua::Value::Nil => (),
            mlua::Value::Table(t) => self.process_lua_returns(&t),
            _ => fault!(self.queue, "invalid return type from lua: {:?}", result)
        }
    }

    fn process_lua_returns(&mut self, ret: &mlua::Table) {
        for tbls in ret.sequence_values::<mlua::Value>() {
            match tbls {
                Err(e) => fault!(self.queue, "expected lua table: {}", e),
                Ok(tbl) => {
                    match self.lua.from_value(tbl) {
                        Ok(m) => self.queue.post(m),
                        Err(e) => fault!(self.queue, "invalid return: {}", e)
                    };
                }
            }
        }
    }

    fn process(e: &mut Env, _: &mut Queue, msg: &Message)  {
        match msg {
            Message::Note(n) => {
                if e.conf.is_develop() && n.kind == NoteKind::Fault {
                    panic!("fault: {}", n.message);
                }
            }
            _ => ()
        }
    }

}

impl Default for Engine<'_> {
    fn default() -> Self {
        Engine::new(&config::App::default())
    }
}

fn lua_init() -> mlua::Result<(Lua, mlua::Function)> {
    let lua = Lua::new();
    for (name, data) in SCRIPTS {
        let chunk = lua.load(data).set_name(name);
        if let Err(e) = chunk.exec() {
            panic!("{}", e);
        }
    }

    let globals = lua.globals();
    let lua_engine: mlua::Table = match globals.get("engine") {
        Ok(p) => p,
        Err(e) => panic!("{}", e),
    };

    let lua_process: mlua::Function = match lua_engine.get("process") {
        Ok(p) => p,
        Err(e) => panic!("{}", e),
    };
    Ok((lua, lua_process))
}
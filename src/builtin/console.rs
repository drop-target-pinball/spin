use crate::prelude::*;
use rustyline::{config, DefaultEditor, ExternalPrinter};
use std::{os::fd::AsRawFd, thread};
use mlua::{Error, MultiValue};
use ansi_term::Color;
use termios::Termios;
use std::path::PathBuf;
use std::sync::{Arc, Mutex};

static GRAY: Color = Color::Fixed(8);
static BRIGHT_RED: Color = Color::Fixed(9);
static BRIGHT_YELLOW: Color = Color::Fixed(11);

static GLOBALS: &[u8] = include_bytes!("console.lua");

pub struct Console<'a> {
    out: Box<dyn ExternalPrinter + 'a>,
    original_mode: Termios,
}

impl Console<'_> {
    pub fn new(state: Arc<Mutex<State>>) -> Self {
        let mut editor = unwrap!(DefaultEditor::with_config(
            config::Builder::new()
                .edit_mode(config::EditMode::Emacs)
                .build()
        ));
        let out = unwrap!(editor.create_external_printer());

        thread::spawn(move || {
            run(editor, state);
        });

        // rustyline is going to switch the input to raw mode. Save the
        // current mode so we can restore on exit
        let original_mode = unwrap!(Termios::from_fd(0));
        Self {
            out: Box::new(out),
            original_mode,
        }
    }

    fn log(&mut self, s: &mut State, text: &str) {
        if let Err(e) = self.checked_log(s, text) {
            panic!("fault: unable to write to console: {}", e)
        }
    }

    fn checked_log(&mut self, s: &mut State, text: &str) -> rustyline::Result<()> {
        let elapsed = s.vars["elapsed"].as_int();
        let fmt_uptime = format!("[{:10.3}]", elapsed as f64 / 1000.0);
        self.out.print(format!("{} {}\n", Color::Blue.bold().paint(fmt_uptime), text))?;
        Ok(())
    }

}

impl Drop for Console<'_> {
    fn drop(&mut self) {
        // Disable raw mode. This is necessary for when the cli exits without
        // shutting down the thread that reads from stdin
        unwrap!(termios::tcsetattr(std::io::stdin().as_raw_fd(), termios::TCSADRAIN, &self.original_mode));
    }
}

impl Device for Console<'_> {
    fn init(&mut self, _: &mut Globals) {}

    fn process(&mut self, s: &mut State, msg: &Message) {
        match msg {
            Message::Note(n) => {
                match n.kind {
                    NoteKind::Alert => self.log(s, &format!("{}", BRIGHT_YELLOW.bold().paint(msg.to_string()))),
                    NoteKind::Diag => self.log(s, &format!("{}", Color::Cyan.bold().paint(msg.to_string()))),
                    NoteKind::Info => self.log(s, &format!("{}", Color::Cyan.bold().paint(msg.to_string()))),
                    NoteKind::Fault => self.log(s, &format!("{}", BRIGHT_RED.bold().paint(msg.to_string()))),
                }
            }
            _ => {
                let text: String = msg.to_string();
                if !text.is_empty() {
                    self.log(s, &format!("{}", GRAY.bold().paint(text)));
                }
            }
        }
    }

    fn render(&mut self, _: &mut render::State) {}
    fn present(&mut self, _: &render::State) {}
}

fn run(mut editor: DefaultEditor, state: Arc<Mutex<State>>) {
    let queue = unwrap!(state.lock()).queue.clone();

    let script_env: script::Env = unwrap!(script::Env::new(state.clone()));
    unwrap!(script_env.exec("console.lua", GLOBALS));

    let history_file: Option<PathBuf> = home::home_dir()
        .map(|h| h.join(".local/share/spin.history"));

    if let Some(h) = &history_file {
        if editor.load_history(&h).is_err() {
            // ok
        }
    }

    loop {
        let mut prompt = "spin> ";
        let mut line = String::new();

        loop {
            if let Some(h) = &history_file {
                if editor.save_history(&h).is_err() {
                    // ok
                }
            }

            match editor.readline(prompt) {
                Ok(input) => line.push_str(&input),
                Err(_) => {
                    queue.post(Message::Shutdown);
                    return
                }
            }

            if line.trim() == "quit" || line.trim() == "exit" {
                queue.post(Message::Shutdown);
                return
            }

            // Get a fresh copy of the variable state
            post(&script_env, &queue, Message::Nop);

            match script_env.load_string("cli", &line).eval::<MultiValue>() {
                Ok(values) => {
                    unwrap!(editor.add_history_entry(line));
                    println!(
                        "{}",
                        values
                            .iter()
                            .map(|value| format!("{:#?}", value))
                            .collect::<Vec<_>>()
                            .join("\t")
                    );
                    // Process any generated messages from lua
                    post(&script_env, &queue, Message::Nop);
                    break;
                }
                Err(Error::SyntaxError {
                    incomplete_input: true,
                    ..
                }) => {
                    // continue reading input and append it to `line`
                    line.push('\n'); // separate input lines
                    prompt = ">> ";
                }
                Err(e) => {
                    unwrap!(editor.add_history_entry(line));
                    println!("{}", BRIGHT_RED.bold().paint(e.to_string()));
                    break;
                }
            }
        }
    }
}

fn post(proc_env: &script::Env, queue: &Queue, msg: Message) {
    unwrap!(proc_env.send_vars());
    let messages = match proc_env.process(&msg) {
        Ok(m) => m,
        Err(e) => {
            println!("{}", BRIGHT_RED.bold().paint(e.to_string()));
            Vec::new()
        }
    };
    unwrap!(proc_env.recv_vars());

    for msg in messages {
        queue.post(msg);
    }
}
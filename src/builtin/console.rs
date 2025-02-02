use crate::prelude::*;
use rustyline::{config, DefaultEditor, ExternalPrinter};
use std::{os::fd::AsRawFd, thread};
use mlua::{Error, MultiValue};
use ansi_term::Color;
use termios::Termios;
use std::path::PathBuf;

static GRAY: Color = Color::Fixed(8);
static BRIGHT_RED: Color = Color::Fixed(9);
static BRIGHT_YELLOW: Color = Color::Fixed(11);

static GLOBALS: &[u8] = include_bytes!("console.lua");

pub struct Console<'c> {
    out: Box<dyn ExternalPrinter + 'c>,
    original_mode: Termios,
}

impl<'c> Console<'c> {
    pub fn new(state: State) -> Self {
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

    fn log(&mut self, env: &mut Env, text: &str) {
        if let Err(e) = self.checked_log(env, text) {
            panic!("fault: unable to write to console: {}", e)
        }
    }

    fn checked_log(&mut self, env: &mut Env, text: &str) -> rustyline::Result<()> {
        let elapsed = env.vars.elapsed;
        let fmt_uptime = format!("[{:10.3}]", elapsed as f64 / 1000.0);
        self.out.print(format!("{} {}\n", Color::Blue.bold().paint(fmt_uptime), text))?;
        Ok(())
    }

}

impl<'c> Drop for Console<'c> {
    fn drop(&mut self) {
        // Disable raw mode. This is necessary for when the cli exits without
        // shutting down the thread that reads from stdin
        unwrap!(termios::tcsetattr(std::io::stdin().as_raw_fd(), termios::TCSADRAIN, &self.original_mode));
    }
}

impl<'c> Device for Console<'c> {
    fn process(&mut self, env: &mut Env, msg: &Message) {
        match msg {
            Message::Note(n) => {
                match n.kind {
                    NoteKind::Info => self.log(env, &format!("{}", Color::Cyan.bold().paint(msg.to_string()))),
                    NoteKind::Alert => self.log(env, &format!("{}", BRIGHT_YELLOW.bold().paint(msg.to_string()))),
                    NoteKind::Fault => self.log(env, &format!("{}", BRIGHT_RED.bold().paint(msg.to_string()))),
                }
            }
            Message::Nop => (),
            Message::Tick => (),
            _ => self.log(env, &format!("{}", GRAY.bold().paint(msg.to_string()))),
        }
    }
}

fn run(mut editor: DefaultEditor, mut state: State) {
    let script_env: script::Env = unwrap!(script::Env::new(&state.conf, state.vars_box.clone()));
    unwrap!(script_env.exec("console.lua", GLOBALS));

    let history_file: Option<PathBuf> = match home::home_dir() {
        Some(h) => Some(h.join(".local/share/spin.history")),
        None => None
    };

    if let Some(h) = &history_file {
        if let Err(_) = editor.load_history(&h) {
            // ok
        }
    }

    loop {
        let mut prompt = "spin> ";
        let mut line = String::new();

        loop {
            if let Some(h) = &history_file {
                if let Err(_) = editor.save_history(&h) {
                    // ok
                }
            }

            match editor.readline(prompt) {
                Ok(input) => line.push_str(&input),
                Err(_) => {
                    state.queue.post(Message::Shutdown);
                    return
                }
            }

            if line.trim() == "quit" || line.trim() == "exit" {
                state.queue.post(Message::Shutdown);
                return
            }

            match script_env.load_string("cli", &line).eval::<MultiValue>() {
                Ok(values) => {
                    editor.add_history_entry(line).unwrap();
                    println!(
                        "{}",
                        values
                            .iter()
                            .map(|value| format!("{:#?}", value))
                            .collect::<Vec<_>>()
                            .join("\t")
                    );
                    post(&script_env, &mut state,Message::Nop);
                    break;
                }
                Err(Error::SyntaxError {
                    incomplete_input: true,
                    ..
                }) => {
                    // continue reading input and append it to `line`
                    line.push_str("\n"); // separate input lines
                    prompt = ">> ";
                }
                Err(e) => {
                    editor.add_history_entry(line).unwrap();
                    println!("{}", BRIGHT_RED.bold().paint(e.to_string()));
                    break;
                }
            }
        }
    }
}

fn post(proc_env: &script::Env, state: &mut State, msg: Message) {
    unwrap!(proc_env.send_vars());
    let messages = match proc_env.process(&msg) {
        Ok(m) => m,
        Err(e) => {
            println!("{}", BRIGHT_RED.bold().paint(e.to_string()));
            Vec::new()
        }
    };
    unwrap!(proc_env.recv_vars());

    let vars = &mut unwrap!(state.vars_box.lock()).vars;
    let eng_env = Env::new(&state.conf, vars, state.queue.clone());

    for msg in messages {
        eng_env.queue.post(msg);
    }
}
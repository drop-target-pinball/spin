use crate::prelude::*;
use rustyline::{DefaultEditor, ExternalPrinter};
use std::thread;
use mlua::{Error, MultiValue};
use ansi_term::Color;

pub struct Console<'c> {
    out: Box<dyn ExternalPrinter + 'c>,
}

impl<'c> Console<'c> {
    pub fn new(state: State) -> Self {
        let mut editor = DefaultEditor::new().expect("failed to create editor");
        let out = editor.create_external_printer().expect("failed to create printer");

        thread::spawn(move || {
            run(editor, state);
        });

        Self {
            out: Box::new(out),
        }
    }

    fn log(&mut self, env: &mut Env, text: &str) {
        if let Err(e) = self.checked_log(env, text) {
            panic!("fault: unable to write to console: {}", e)
        }
    }

    fn checked_log(&mut self, env: &mut Env, text: &str) -> rustyline::Result<()> {
        let elapsed = env.vars.elapsed;
        let fmt_uptime = format!("[{:10.3}]", elapsed.as_secs_f32());
        self.out.print(format!("{} {}", Color::Yellow.paint(fmt_uptime), text))?;
        Ok(())
    }

}

impl<'c> Device for Console<'c> {
    fn process(&mut self, env: &mut Env, msg: &Message) {
        match msg {
            Message::Note(_) => self.log(env, &msg.to_string()),
            _ => self.log(env, &format!("> {}", msg))
        }
    }
}

fn run(mut editor: DefaultEditor, state: State) {
    let mut proc_env: proc::Env = proc::Env::new().unwrap();

    loop {
        let mut prompt = "spin> ";
        let mut line = String::new();


        loop {
            match editor.readline(prompt) {
                Ok(input) => line.push_str(&input),
                Err(_) => return,
            }

            match proc_env.lua().load(&line).eval::<MultiValue>() {
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

                    let mut vars = state.vars.lock().unwrap();
                    let eng_env = Env::new(&state.conf, &mut vars, state.queue.clone());

                    let messages = proc_env.process(&Message::Nop).unwrap();
                    for msg in messages {
                        eng_env.queue.post(msg);
                    }

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
                    eprintln!("error: {}", e);
                    break;
                }
            }
        }
    }
}

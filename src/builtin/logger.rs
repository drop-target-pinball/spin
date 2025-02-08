use crate::prelude::*;
use std::io;

pub struct Logger<W> {
    out: W,
}

impl<W> Logger<W>
    where W: io::Write
{
    pub fn new(out: W) -> Self {
        Logger{out}
    }

    fn log(&mut self, s: &mut State, text: &str) {
        // In the event that the note could not be logged, panic when in
        // debug mode and simply write to standard error when in
        // production mode
        if let Err(e) = self.checked_log(s, text) {
            if s.conf.is_develop() {
                panic!("fault: unable to log: {}", e)
            } else {
                eprintln!("fault: unable to log: {}", e)
            }
        }
    }

    fn checked_log(&mut self, s: &mut State, text: &str) -> io::Result<()> {
        let elapsed = s.vars["elapsed"].as_int();
        let fmt_uptime = format!("[{:10.3}]", elapsed as f64 / 1000.0);
        writeln!(self.out, "{} {}", fmt_uptime, text)
    }
}

impl Default for Logger<io::Stdout> {
    fn default() -> Self {
        let out = io::stdout();
        Logger::new(out)
    }
}

impl<W> Device for Logger<W>
where W: io::Write {
    fn process(&mut self, s: &mut State, msg: &Message) {
        match msg {
            Message::Note(_) => self.log(s, &msg.to_string()),
            _ => {
                if s.conf.is_develop() {
                    let text: String = msg.to_string();
                    if !text.is_empty() {
                        self.log(s, &format!("> {}", text));
                    }
                }
            }
        }
    }

    fn render(&mut self, _: &mut render::State) {}
}

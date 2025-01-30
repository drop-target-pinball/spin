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

    fn log(&mut self, env: &mut Env, text: &str) {
        // In the event that the note could not be logged, panic when in
        // debug mode and simply write to standard error when in
        // production mode
        if let Err(e) = self.checked_log(env, text) {
            if env.conf.is_develop() {
                panic!("fault: unable to log: {}", e)
            } else {
                eprintln!("fault: unable to log: {}", e)
            }
        }
    }

    fn checked_log(&mut self, env: &mut Env, text: &str) -> io::Result<()> {
        let elapsed = env.vars.now.duration_since(env.vars.uptime);
        let fmt_uptime = format!("[{:10.3}]", elapsed.as_secs_f32());
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
    fn process(&mut self, env: &mut Env, _: &mut Queue, msg: &Message) {
        match msg {
            Message::Note(_) => self.log(env, &msg.to_string()),
            _ => {
                if env.conf.is_develop() {
                    self.log(env, &format!("> {}", msg));
                }
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_note_info() {
        let mut e = Engine::default();
        let mut buf = Vec::new();
        let mut logger= Logger::new(&mut buf);
        e.add_device(&mut logger);

        let q = e.queue();
        info!(q, "this is a test");
        e.tick();

        let want = "[     0.000] this is a test\n";
        let have = String::from_utf8(buf).unwrap();
        assert_eq!(have, want);
    }

    #[test]
    fn test_note_alert() {
        let mut e = Engine::default();
        let mut buf = Vec::new();
        let mut logger= Logger::new(&mut buf);
        e.add_device(&mut logger);

        let q = e.queue();
        alert!(q, "this is a test");
        e.tick();

        let want = "[     0.000] (!) this is a test\n";
        let have = String::from_utf8(buf).unwrap();
        assert_eq!(have, want);
    }

    #[test]
    fn test_note_fault() {
        let conf = config::new(config::RunMode::Release);
        let mut e = Engine::new(&conf);
        let mut buf = Vec::new();
        let mut logger= Logger::new(&mut buf);
        e.add_device(&mut logger);

        let q = e.queue();
        fault!(q, "this is a test");
        e.tick();

        let want = "[     0.000] (*) this is a test\n";
        let have = String::from_utf8(buf).unwrap();
        assert_eq!(have, want);
    }
}
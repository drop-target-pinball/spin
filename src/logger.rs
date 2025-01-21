use crate::prelude::*;
use std::io;

pub struct Logger<W> {
    out: W,
}

impl<W> Logger<W>
    where W: io::Write
{
    pub fn new(out: W) -> Self {
        Logger{
            out
        }
    }

    fn log(&mut self, ctx: &mut Context, n: &Note) {
        // In the event that the note could not be logged, panic when in
        // debug mode and simply write to standard error when in
        // production mode
        if let Err(e) = self.checked_log(ctx, n) {
            if ctx.conf.is_debug() {
                panic!("fault: unable to log: {}", e)
            } else {
                eprintln!("fault: unable to log: {}", e)
            }
        }
    }

    fn checked_log(&mut self, ctx: &mut Context, n: &Note) -> io::Result<()> {
        let elapsed = ctx.state.now.duration_since(ctx.state.uptime);
        let fmt_uptime = format!("[{:10.3}]", elapsed.as_secs_f32());
        match n.kind {
            NoteKind::Info => writeln!(self.out, "{} {}", fmt_uptime, n.message),
            NoteKind::Alert => writeln!(self.out, "{} (!) {}", fmt_uptime, n.message),
            NoteKind::Fault => {
                if ctx.conf.is_debug() {
                    panic!("fault: {}", n.message);
                }
                writeln!(self.out, "{} (*) {}", fmt_uptime, n.message)
            }
        }
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
    fn id(&self) -> u8 {
        0
    }

    fn topic(&self) -> Topic {
        Topic::Admin
    }

    fn process(&mut self, ctx: &mut Context, _: Topic, msg: &Message) {
        match msg {
            Message::Note(n) => self.log(ctx, n),
            _ => (),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_note() {
        let mut e = Engine::default();
        let mut buf = Vec::new();
        let mut logger = Logger::new(&mut buf);
        e.add_device(&mut logger);

        info!(e.queue, "this is a test");
        e.tick();

        let want = "[     0.000] this is a test\n";
        let have = String::from_utf8(buf).unwrap();
        assert_eq!(have, want);
    }

    #[test]
    fn test_alert() {
        let mut e = Engine::default();
        let mut buf = Vec::new();
        let mut logger = Logger::new(&mut buf);
        e.add_device(&mut logger);

        alert!(e.queue, "this is a test");
        e.tick();

        let want = "[     0.000] (!) this is a test\n";
        let have = String::from_utf8(buf).unwrap();
        assert_eq!(have, want);
    }

    #[test]
    #[should_panic(expected = "fault: this is a test")]
    fn test_fault() {
        let mut e = Engine::default();
        let mut buf = Vec::new();
        let mut logger = Logger::new(&mut buf);
        e.add_device(&mut logger);

        fault!(e.queue, "this is a test");
        e.tick();
    }

    #[test]
    fn test_fault_prod() {
        let mut e = Engine::new(Config::new(RunMode::Prod));
        let mut buf = Vec::new();
        let mut logger = Logger::new(&mut buf);
        e.add_device(&mut logger);

        fault!(e.queue, "this is a test");
        e.tick();
    }
}
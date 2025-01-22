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

    fn log(&mut self, env: &mut Env, n: &Note) {
        // In the event that the note could not be logged, panic when in
        // debug mode and simply write to standard error when in
        // production mode
        if let Err(e) = self.checked_log(env, n) {
            if env.conf.is_develop() {
                panic!("fault: unable to log: {}", e)
            } else {
                eprintln!("fault: unable to log: {}", e)
            }
        }
    }

    fn checked_log(&mut self, env: &mut Env, n: &Note) -> io::Result<()> {
        let elapsed = env.vars.now.duration_since(env.vars.uptime);
        let fmt_uptime = format!("[{:10.3}]", elapsed.as_secs_f32());
        match n.kind {
            NoteKind::Info => writeln!(self.out, "{} {}", fmt_uptime, n.message),
            NoteKind::Alert => writeln!(self.out, "{} (!) {}", fmt_uptime, n.message),
            NoteKind::Fault => writeln!(self.out, "{} (*) {}", fmt_uptime, n.message),
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
    fn process(&mut self, env: &mut Env, _: &mut Queue, msg: &Message) -> bool {
        match msg {
            Message::Note(n) => self.log(env, n),
            _ => return false,
        }
        true
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_note() {
        let mut e = Env::default();
        let mut buf = Vec::new();
        let mut logger: Logger<&mut Vec<u8>> = Logger::new(&mut buf);
        logger.process(&mut e, &mut Queue::new(), &Message::Note(Note {
            kind: NoteKind::Info,
            message: "this is a test".to_string(),
        }));

        let want = "[     0.000] this is a test\n";
        let have = String::from_utf8(buf).unwrap();
        assert_eq!(have, want);
    }

    #[test]
    fn test_alert() {
        let mut e = Env::default();
        let mut buf = Vec::new();
        let mut logger = Logger::new(&mut buf);
        logger.process(&mut e, &mut Queue::new(), &Message::Note(Note {
            kind: NoteKind::Alert,
            message: "this is a test".to_string(),
        }));

        let want = "[     0.000] (!) this is a test\n";
        let have = String::from_utf8(buf).unwrap();
        assert_eq!(have, want);
    }

    #[test]
    fn test_fault() {
        let mut e = Env::default();
        let mut buf = Vec::new();
        let mut logger = Logger::new(&mut buf);
        logger.process(&mut e, &mut Queue::new(), &Message::Note(Note {
            kind: NoteKind::Fault,
            message: "this is a test".to_string(),
        }));

        let want = "[     0.000] (*) this is a test\n";
        let have = String::from_utf8(buf).unwrap();
        assert_eq!(have, want);
    }
}
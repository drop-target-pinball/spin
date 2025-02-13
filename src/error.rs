#[derive(thiserror::Error, Debug)]
pub enum Error {
    #[error("configuration error: {0}")]
    Config(String),

    #[error("{0}")]
    Load(String),

    #[error("invalid format: {0}")]
    InvalidFormat(String),

    #[error("rendering error: {0}")]
    Render(String),

    #[error("presentation error: {0}")]
    Present(String),

    #[error("invalid script environment: {0}")]
    ScriptEnv(String),

    #[error("{0}")]
    ScriptExec(String),
}

pub type Result<T> = std::result::Result<T, Error>;

#[macro_export]
macro_rules! raise {
    ($id:expr, $($args:expr),+) => {
        Err($id(format!($($args),+)))
    };
}

#[macro_export]
macro_rules! chain {
    ($expr:expr, $err:expr) => {
        match $expr {
            Ok(v) => v,
            Err(e) => return Err($err(e.to_string())),
        }
    };
}

#[macro_export]
macro_rules! try_present {
    ($expr:expr) => {
        chain!($expr, Error::Present)
    }
}
#[derive(thiserror::Error, Debug)]
pub enum Error {
    #[error("configuration error: {0}")]
    Config(String),

    #[error("initialization error: {0}")]
    Init(String),

    #[error("invalid argument: {0}")]
    InvalidArgument(String),

    #[error("invalid format: {0}")]
    InvalidFormat(String),

    #[error("invalid type: expected {0}, got {1}")]
    InvalidType(String, String),

    #[error("not defined: {0}")]
    NotDefined(String),

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

#[macro_export]
macro_rules! try_render {
    ($expr:expr) => {
        chain!($expr, Error::Render)
    }
}
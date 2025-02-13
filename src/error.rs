#[derive(thiserror::Error, Debug)]
pub enum Error {
    #[error("configuration error: {0}")]
    Config(String),

    #[error("{0}")]
    Load(String),

    #[error("invalid format: {0}")]
    InvalidFormat(String),

    #[error("rendering error: {0}")]
    RenderError(String),

    #[error("invalid script environment: {0}")]
    ScriptEnv(String),

    #[error("{0}")]
    ScriptExec(String),
}

pub type SpinResult<T> = std::result::Result<T, Error>;

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
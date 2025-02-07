#[derive(thiserror::Error, Debug)]
pub enum Error {
    #[error("configuration error: {0}")]
    Config(String),

    #[error("invalid script environment: {0}")]
    ScriptEnv(String),

    #[error("{0}")]
    ScriptExec(String),
}

pub type SpinResult<T> = std::result::Result<T, Error>;
pub type FmtResult = std::result::Result<(), std::fmt::Error>;

#[macro_export]
macro_rules! raise {
    ($id:expr, $($args:expr),+) => {
        Err($id(format!($($args),+)))
    };
}

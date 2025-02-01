#[derive(thiserror::Error, Debug)]
pub enum Error {
    #[error("configuration error: {0}")]
    Config(String),

    #[error("invalid script environment: {0}")]
    ScriptEnv(String),

    #[error("script execution error: {0}")]
    ScriptExec(String),
}

pub type Result<T> = std::result::Result<T, Error>;

#[macro_export]
macro_rules! raise {
    ($id:expr, $($args:expr),+) => {
        Err($id(format!($($args),+)))
    };
}

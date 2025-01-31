#[derive(thiserror::Error, Debug)]
pub enum Error {
    #[error("configuration error: {0}")]
    Config(String),

    #[error("invalid procedure environment: {0}")]
    ProcEnv(String),

    #[error("procedure execution error: {0}")]
    ProcExec(String),
}

pub type Result<T> = std::result::Result<T, Error>;

#[macro_export]
macro_rules! raise {
    ($id:expr, $($args:expr),+) => {
        Err($id(format!($($args),+)))
    };
}

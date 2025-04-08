#[derive(thiserror::Error, Debug)]
pub enum Error {
    #[error("configuration error: {0}")]
    Config(String),

    #[error("device error: {0}")]
    Device(String),

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

    #[error("{0}")]
    Script(String),

    #[error("unexpected type: {0}")]
    UnexpectedType(String)
}

pub type Result<T> = std::result::Result<T, Error>;

#[macro_export]
macro_rules! raise {
    ($id:expr, $($args:expr),+) => {
        Err($id(format!($($args),+)))
    };
}

#[cfg(not(feature = "debug_faults"))]
#[macro_export]
macro_rules! chain {
    ($expr:expr, $err:expr) => {
        match $expr {
            Ok(v) => v,
            Err(e) => return Err($err(e.to_string())),
        }
    };
}

#[cfg(feature = "debug_faults")]
#[macro_export]
macro_rules! chain {
    ($expr:expr, $err:expr) => {
        match $expr {
            Ok(v) => v,
            Err(e) => panic!("{}", e.to_string()),
        }
    };
}

#[macro_export]
macro_rules! try_device {
    ($expr:expr) => {
        chain!($expr, Error::Init)
    }
}

#[macro_export]
macro_rules! try_init {
    ($expr:expr) => {
        chain!($expr, Error::Init)
    }
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

#[macro_export]
macro_rules! try_script {
    ($expr:expr) => {
        chain!($expr, Error::Script)
    }
}
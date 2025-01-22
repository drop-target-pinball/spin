#[derive(thiserror::Error, Debug)]
pub enum Error {
    #[error("{0}: {1}")]
    Device(String, String),
}

pub type Result<T> = std::result::Result<T, Error>;


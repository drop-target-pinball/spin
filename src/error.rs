#[derive(thiserror::Error, Debug)]
pub enum Error {
    #[error("{0}: {1}")]
    Device(String, String),
}

pub type Result<T> = std::result::Result<T, Error>;

pub fn device_error<T>(what: &str, reason: String) -> Result<T> {
    Err(Error::Device(what.to_string(), reason))
}

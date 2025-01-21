#[derive(thiserror::Error, Debug)]
pub enum Error {
    #[error("device error: {0}")]
    Device(String),
}

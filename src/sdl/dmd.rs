use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize, Debug, Clone)]
#[serde(deny_unknown_fields)]
pub struct DmdConfig {
    pub display: String,
}
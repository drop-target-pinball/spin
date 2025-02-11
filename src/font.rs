use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize, Debug, Default, Clone)]
#[serde(rename_all = "PascalCase")]
pub struct FontConfig {
    #[serde(default)]
    pub offset_y: i32
}
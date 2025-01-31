use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize, Debug)]
pub struct VarsBox {
    pub vars: Vars
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Vars {
    pub elapsed: u64
}

impl Default for Vars {
    fn default() -> Vars {
        Vars {
            elapsed: 0
        }
    }
}
use super::Lane;
use serde::{Serialize, Deserialize};

#[derive(Debug, Serialize, Deserialize)]
pub struct Champion {
    name: String,
    lanes: Vec<Lane>,
}

impl Champion {
    pub fn new(name: String, lanes: Vec<Lane>) -> Self {
        Self { name, lanes }
    }
}

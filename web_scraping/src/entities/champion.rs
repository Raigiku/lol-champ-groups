use super::Lane;
use serde::Serialize;

#[derive(Debug, Serialize)]
pub struct Champion {
    name: String,
    lanes: Vec<Lane>,
}

impl Champion {
    pub fn new(name: String, lanes: Vec<Lane>) -> Self {
        Self { name, lanes }
    }
}

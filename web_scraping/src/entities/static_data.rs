use serde::Deserialize;
use std::{collections::HashMap, error::Error, fs};

#[derive(Deserialize)]
pub struct StaticData {
    lanes: Vec<StaticDataLane>,
    champion_names: Vec<String>,
    stats_xpaths: HashMap<String, String>,
}

impl StaticData {
    pub fn from_json_file(path: &str) -> Result<Self, Box<dyn Error>> {
        let json_data = fs::read_to_string(path)?;
        let result: Self = serde_json::from_str(&json_data)?;
        Ok(result)
    }

    pub fn lanes(&self) -> &Vec<StaticDataLane> {
        &self.lanes
    }

    pub fn champion_names(&self) -> &Vec<String> {
        &self.champion_names
    }

    pub fn stats_xpaths(&self) -> &HashMap<String, String> {
        &self.stats_xpaths
    }
}

#[derive(Deserialize)]
pub struct StaticDataLane {
    name: String,
    pick_percentage_xpath: String,
}

impl StaticDataLane {
    pub fn name(&self) -> &String {
        &self.name
    }

    pub fn pick_percentage_xpath(&self) -> &String {
        &self.pick_percentage_xpath
    }
}

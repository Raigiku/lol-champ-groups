use crate::constants::*;
use serde::Serialize;
use std::{collections::HashMap, error::Error};

#[derive(Debug, Serialize)]
pub struct Lane {
    name: String,
    pick_percentage: f32,
    physical_damage: f32,
    magic_damage: f32,
    true_damage: f32,
    total_damage: f32,
    damage_taken: f32,
    healing: f32,
    kills: f32,
    deaths: f32,
    assists: f32,
    max_kill_spree: f32,
    gold: f32,
    minions_killed: f32,
    jungle_cs: f32,
    enemy_jungle_cs: f32,
    team_jungle_cs: f32,
}

impl Lane {
    pub fn from_website(
        name: String,
        pick_percentage: String,
        stats: HashMap<String, String>,
    ) -> Result<Self, Box<dyn Error>> {
        let physical_damage = stats
            .get(&key_physical_damage())
            .unwrap()
            .replace(",", "")
            .parse::<f32>()?;
        let magic_damage = stats
            .get(&key_magic_damage())
            .unwrap()
            .replace(",", "")
            .parse::<f32>()?;
        let true_damage = stats
            .get(&key_true_damage())
            .unwrap()
            .replace(",", "")
            .parse::<f32>()?;
        let total_damage = stats
            .get(&key_total_damage())
            .unwrap()
            .replace(",", "")
            .parse::<f32>()?;
        let damage_taken = stats
            .get(&key_damage_taken())
            .unwrap()
            .replace(",", "")
            .parse::<f32>()?;
        let healing = stats
            .get(&key_healing())
            .unwrap()
            .replace(",", "")
            .parse::<f32>()?;
        let kills = stats
            .get(&key_kills())
            .unwrap()
            .replace(",", "")
            .parse::<f32>()?;
        let deaths = stats
            .get(&key_deaths())
            .unwrap()
            .replace(",", "")
            .parse::<f32>()?;
        let assists = stats
            .get(&key_assists())
            .unwrap()
            .replace(",", "")
            .parse::<f32>()?;
        let max_kill_spree = stats
            .get(&key_max_kill_spree())
            .unwrap()
            .replace(",", "")
            .parse::<f32>()?;
        let gold = stats
            .get(&key_gold())
            .unwrap()
            .replace(",", "")
            .parse::<f32>()?;
        let minions_killed = stats
            .get(&key_minions_killed())
            .unwrap()
            .replace(",", "")
            .parse::<f32>()?;
        let jungle_cs = stats
            .get(&key_jungle_cs())
            .unwrap()
            .replace(",", "")
            .parse::<f32>()?;
        let enemy_jungle_cs = stats
            .get(&key_enemy_jungle_cs())
            .unwrap()
            .replace(",", "")
            .parse::<f32>()?;
        let team_jungle_cs = stats
            .get(&key_team_jungle_cs())
            .unwrap()
            .replace(",", "")
            .parse::<f32>()?;
        let pick_percentage = pick_percentage.replace("%", "").parse::<f32>()?;

        Ok(Self {
            name,
            pick_percentage,
            assists,
            damage_taken,
            deaths,
            enemy_jungle_cs,
            gold,
            healing,
            jungle_cs,
            kills,
            magic_damage,
            max_kill_spree,
            minions_killed,
            physical_damage,
            team_jungle_cs,
            total_damage,
            true_damage,
        })
    }
}

mod constants;
mod entities;

use std::{collections::HashMap, error::Error, fs, thread, time::Duration};
use thirtyfour::prelude::*;
use tokio;

use entities::{Champion, Lane, StaticData};

async fn text_from_element_by_xpath(
    driver: &WebDriver,
    xpath: &String,
) -> Result<String, Box<dyn Error>> {
    Ok(driver
        .find_element(By::XPath(xpath))
        .await?
        .text()
        .await?
        .clone())
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    let static_data = StaticData::from_json_file("src/static_data.json")?;

    let caps = DesiredCapabilities::chrome();
    let driver: WebDriver = WebDriver::new("http://localhost:4444", &caps).await?;

    let lolalytics_url = "https://lolalytics.com/lol";
    let mut champions = Vec::<Champion>::new();

    // let json = fs::read_to_string("champions.json")?;
    // champions = serde_json::from_str(&json)?;

    for champion_name in static_data.champion_names().iter() {
        let mut lanes = Vec::<Lane>::new();
        for lane in static_data.lanes().iter() {
            let url = format!(
                "{}/{}/?tier=all&patch=10.10&lane={}",
                lolalytics_url,
                champion_name,
                lane.name()
            );
            driver.get(url).await?;

            thread::sleep(Duration::from_secs(10));

            let mut stats = HashMap::new();
            for (stat, xpath) in static_data.stats_xpaths().iter() {
                let stat_value: String = text_from_element_by_xpath(&driver, xpath).await?;
                stats.insert(stat.clone(), stat_value);
            }

            let lane_pick_percentage_value: String =
                text_from_element_by_xpath(&driver, lane.pick_percentage_xpath()).await?;
            lanes.push(Lane::from_website(
                lane.name().clone(),
                lane_pick_percentage_value,
                stats,
            )?);
        }
        champions.push(Champion::new(champion_name.clone(), lanes));
        serde_json::to_writer_pretty(&fs::File::create("champions.json")?, &champions)?;
    }

    Ok(())
}

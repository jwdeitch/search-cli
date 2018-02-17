use reqwest;
use std::io::{Error, ErrorKind, Read, Result};

pub fn get_http(path: &str) -> Result<Vec<u8>> {
    let mut contents = Vec::new();
    match reqwest::Url::parse(path) {
        Ok(url) => {
            reqwest::get(url)
                .map_err(|_e| Error::new(ErrorKind::NotConnected, "reqwest"))
                .and_then(|mut r| r.read_to_end(&mut contents))?;
        }
        Err(e) => {
            println!("ERR: {:?}", e)
        }
    }
    Ok(contents)
}
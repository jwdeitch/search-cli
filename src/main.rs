extern crate reqwest;
extern crate serde_json;
extern crate base64;
extern crate url;

use std::io::{Stdout, stdout};
use std::env;
use url::form_urlencoded::byte_serialize;

mod lambda;
mod get;

fn main() {
    let mut args: Vec<String> = env::args().collect();

    if args.len() == 0 {
        println!("No arguments passed");
        return;
    }

    args.remove(0);

    let mut stdout: Stdout = stdout();
    let urlencoded: String = byte_serialize(&args.join("").as_bytes()).collect();

    lambda::inline_image(&mut stdout, &wra_endpoint(&urlencoded));
}

fn wra_endpoint(query: &str) -> String {
    return match env::var_os("WRA_API_ID") {
        Some(val) => format!("http://api.wolframalpha.com/v1/simple?appid={0}&i={1}", val.into_string().expect("cannot convert WRA_API_ID to string"), query),
        None => format!("http://api.wolframalpha.com/v1/simple?appid={0}&i={1}", "35TP3H-VAE68AAT2Y", query) // pls dont steal
    };
}
extern crate reqwest;
extern crate serde_json;
extern crate base64;

use std::io::{Stdout, stdout};

mod lambda;
mod get;

fn main() {
    let mut test: Stdout = stdout();

    let test3 = String::from("http://s.rsa.pub/hl32xzhs6aeb98g.jpg");
    match lambda::inline_image(&mut test, &test3) {
        Ok(()) => print!("ok"),
        Err(_) => print!("err"),
    }
}

fn wra_endpoint() -> String {
    return match env::var_os("WRA_API_ID") {
        Some(val) => format!("http://api.wolframalpha.com/v1/simple?appid=", val),
        None => format!("http://api.wolframalpha.com/v1/simple?appid=", "ewfwef")
    }
}
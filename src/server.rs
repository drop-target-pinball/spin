use crate::prelude::*;
use rocket::{get, routes};

struct State {
    conf: config::App,
    queue: Queue
}

#[get("/")]
fn index() -> &'static str {
    "Hello, world!"
}

pub async fn run(conf: config::Server) {
    let log_level = conf.log_level.unwrap_or(String::from("off"));
    let figment = rocket::Config::figment()
        .merge(("host", conf.host))
        .merge(("port", conf.port))
        .merge(("log_level", log_level));

    match rocket::custom(figment).mount("/", routes![index]).launch().await {
        Ok(_) => println!("*** GOT IT"),
        Err(e) => println!("*** ERR: {}", e)
    }
}


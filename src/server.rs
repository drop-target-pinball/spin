use crate::prelude::*;
use std::io::Read;
use std::net::{TcpListener, TcpStream};

use rocket::{get, routes};


#[get("/")]
fn index() -> &'static str {
    "Hello, world!"
}

pub async fn run() {
    match rocket::build().mount("/", routes![index]).launch().await {
        Ok(_) => (),
        Err(e) => println!("*** ERR: {}", e)
    }
}

// pub struct Server {
//     listener: TcpListener,
//     queue: Queue,
// }

// pub fn new(queue: Queue, addr: &str) -> Server {
//     let listener = TcpListener::bind(addr).unwrap_or_else(|e| {
//         panic!("unable to bind to {}: {}", addr, e);
//     });
//     info!(queue, "http server listening on {}", addr);
//     Server{listener, queue}
// }

// impl Server {
//     pub fn run(&self) {
//         for stream in self.listener.incoming() {
//             match stream {
//                 Ok(mut s) => self.handle_connection(&mut s),
//                 Err(e) => {
//                     fault!(self.queue, "http server error: {}", e);
//                     return;
//                 }
//             }
//         }
//     }

//     fn handle_connection(&self, stream: &mut TcpStream) {
//         let mut headers = [httparse::EMPTY_HEADER; 64];
//         let mut buf = [0u8; 4096];

//         info!(self.queue, "got a connection");
//         let mut req = httparse::Request::new(&mut headers);

//         loop {
//             if let Err(e) = stream.read(&mut buf) {
//                 alert!(self.queue, "http: error reading into buffer: {}", e);
//                 return
//             }
//             match req.parse(&mut buf) {
//                 Err(e) => {
//                     alert!(self.queue, "http: error parsing headers: {}", e);
//                     return
//                 }
//                 Ok(status) => {
//                     if status.is_complete() {
//                         break
//                     }
//                 }
//             }
//         }
//         println!("***** HEADERS: {:?}", headers);
//     }
// }
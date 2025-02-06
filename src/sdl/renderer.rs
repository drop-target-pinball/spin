// use crate::prelude::*;

// use sdl2::surface::Surface;
// use sdl2::pixels::PixelFormatEnum;
// use sdl2::VideoSubsystem;

// pub struct Renderer<'a> {
//     _video: VideoSubsystem,
//     id: u8,
//     surface: Surface<'a>,
// }

// impl<'a> Renderer<'a> {
//     pub fn new(ctx: &sdl2::Sdl, id: u8, conf: &config::Renderer) -> Self {
//         let _video = match ctx.video( ){
//             Ok(v) => v,
//             Err(reason) => {
//                 panic!("unable to open SDL video: {}", reason);
//             }
//         };

//         let surface = match Surface::new(
//             conf.width,
//             conf.height,
//             PixelFormatEnum::RGBA8888,
//         ) {
//             Ok(s) => s,
//             Err(e) => panic!("unable to create rendering surface: {}", e),
//         };

//         Self { _video, id, surface }
//     }
// }

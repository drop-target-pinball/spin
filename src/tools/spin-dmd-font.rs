use std::process::ExitCode;
use spin::prelude::*;
use clap::Parser;
use std::fs;
use sdl2::surface::Surface;
use sdl2::event::Event;
use sdl2::keyboard::Keycode;

#[derive(Parser)]
#[command(version, about, long_about = None)]
struct Cli {
    #[arg(short, long)]
    #[clap(default_value_t = 1)]
    scale: u8,

    dmd_file: String
}


pub fn display(name: &str, frame: &Surface<'static>, scale: u8) -> Result<(), String> {
    let sdl = sdl2::init()?;
    let video = sdl.video()?;

    let w = frame.width() * scale as u32;
    let h = frame.height() * scale as u32;

    let win = video.window(name, w, h)
        .position(0, 0)
        .hidden()
        .build().map_err(|e| e.to_string())?;

    let mut canvas = win.into_canvas()
        .accelerated()
        .present_vsync()
        .build().map_err(|e| e.to_string())?;

    let tc = canvas.texture_creator();
    let tex = tc
        .create_texture_from_surface(frame)
        .map_err(|e| e.to_string())?;

    canvas.copy(&tex, None, None).map_err(|e| e.to_string())?;
    canvas.window_mut().show();

    let mut event_pump = sdl.event_pump().map_err(|e| e.to_string())?;
    'running: loop {
        for event in event_pump.poll_iter() {
            match event {
                Event::Quit { .. }
                | Event::KeyDown {
                    keycode: Some(Keycode::Escape),
                    ..
                } => break 'running,
                _ => {}
            }
        }
        canvas.present();
    }
    Ok(())
}

pub fn main() -> ExitCode {
    let cli = Cli::parse();

    let data = match fs::read(&cli.dmd_file) {
        Ok(d) => d,
        Err(e) => {
            eprintln!("error: unable to open '{}': {}", cli.dmd_file, e);
            return ExitCode::FAILURE
        }
    };

    let frames = match sdl::decode_dmd(&data) {
        Ok(f) => f,
        Err(e) => {
            eprintln!("error: unable to decode DMD '{}': {}", cli.dmd_file, e);
            return ExitCode::FAILURE
        }
    };

    match display(&cli.dmd_file, &frames[0], cli.scale) {
        Ok(_) => ExitCode::SUCCESS,
        Err(e) => {
            eprintln!("error: unable to display font: {}", e);
            ExitCode::FAILURE
        }
    }
}
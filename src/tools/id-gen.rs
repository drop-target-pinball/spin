use std::process::ExitCode;
use std::fmt::{Result, Write};
use std::fs;
use spin::prelude::*;

pub fn main() -> ExitCode {
    let conf = match load_config(&app_dir()) {
        Ok(c) => c,
        Err(e) => {
            eprintln!("{}", e);
            return ExitCode::FAILURE;
        }
    };

    let Some(main_module) = conf.module_name else {
        eprintln!("no module name defined in configuration");
        return ExitCode::FAILURE;
    };

    let mut ids: Vec<String> = Vec::new();

    let add = |ids: &mut Vec<String>, name: String| {
        let name = name.trim().to_owned();
        assert!(!name.is_empty());
        ids.push(name);
    };

    for v in conf.music.keys()   { add(&mut ids, v.into()) }
    for v in conf.sounds.keys()  { add(&mut ids, v.into()) }
    for v in conf.vocals.keys()  { add(&mut ids, v.into()) }

    for (name, s) in conf.scripts {
        add(&mut ids, name);
        if !s.group.is_empty() {
            add(&mut ids, s.group);
        }
    }

    ids.sort();
    ids.dedup();

    let mut lua = String::new();
    if let Err(e) = format_source(&ids, &mut lua) {
        eprintln!("error formatting source: {}", e);
        return ExitCode::FAILURE;
    };

    let id_file = conf.scripts_dir.join(format!("{}.lua", main_module));
    match fs::write(id_file, &lua) {
        Ok(_) => ExitCode::SUCCESS,
        Err(e) => {
            eprintln!("unable to write: {}", e);
            ExitCode::FAILURE
        }
    }
}

pub fn format_source(ids: &Vec<String>, f: &mut String) -> Result {
    writeln!(f, "-- Code generated by 'id-gen'; DO NOT EDIT\n")?;
    writeln!(f, "local pub = {{}}\n")?;

    for id in ids {
        writeln!(f, "pub.{} = \"{}\"", id.to_uppercase(), id)?;
    }
    writeln!(f, "\nreturn pub")?;
    Ok(())
}
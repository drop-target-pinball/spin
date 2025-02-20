use std::process::Command;

fn main() {
    let build_date_cmd = Command::new("date")
        .arg("+%d %b %Y")
        .output()
        .expect("unable to get current date");

    let build_date = String::from_utf8_lossy(&build_date_cmd.stdout);
    println!("cargo:rustc-env=SPIN_BUILD_DATE={}", build_date);

}
use std::process::Command;

fn main() {
    let build_date_cmd = Command::new("date")
        .arg("+%d %b %Y")
        .output()
        .expect("unable to get current date");

    let build_date = String::from_utf8_lossy(&build_date_cmd.stdout);
    println!("cargo:rustc-env=SPIN_BUILD_DATE={}", build_date);

    #[cfg(feature = "proc")] {
        use std::path::PathBuf;

        // Tell cargo to look for shared libraries in the specified directory
        println!("cargo:rustc-link-search=/path/to/lib");

        // Tell cargo to tell rustc to link the system bzip2
        // shared library.
        println!("cargo:rustc-link-lib=bz2");

        // The bindgen::Builder is the main entry point
        // to bindgen, and lets you build up options for
        // the resulting bindings.
        let bindings = bindgen::Builder::default()
            // The input header we would like to generate
            // bindings for.
            .header("src/proc/wrapper.h")
            // Tell cargo to invalidate the built crate whenever any of the
            // included header files changed.
            .parse_callbacks(Box::new(bindgen::CargoCallbacks::new()))
            // Finish the builder and generate the bindings.
            .generate()
            // Unwrap the Result and panic on failure.
            .expect("Unable to generate bindings");

        // Write the bindings to the $OUT_DIR/bindings.rs file.
        let out_path = PathBuf::from("src/proc");
        bindings
            .write_to_file(out_path.join("bindings.rs"))
            .expect("Couldn't write bindings!");
    }
}
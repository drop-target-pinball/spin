package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/drop-target-pinball/spin/v2"
	"github.com/hashicorp/hcl/v2"
)

var (
	Json    bool
	Module  bool
	Version bool
)

func main() {
	log.SetFlags(0)

	flag.Usage = func() {
		fmt.Printf("Usage: %v [options] <dir>\n", os.Args[0])
		fmt.Print(`
Parses an HCL configuration directory.

If no options are specified, the program produces no output and exits with a
return code of zero upon success. Otherwise, diagnostics are printed to the
console and the program exits with a return code of one.

Use -json to print the configuration as a JSON document.

`)
		flag.PrintDefaults()
	}

	flag.BoolVar(&Json, "j", false, "emit JSON document")
	flag.BoolVar(&Module, "m", false, "load as a module")
	flag.BoolVar(&Version, "v", false, "version")
	flag.Parse()

	if Version {
		fmt.Println(spin.Banner())
		return
	}

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	dir := flag.Arg(0)

	config := spin.NewConfig()
	var diags hcl.Diagnostics

	if Module {
		diags = spin.LoadModule(config, dir)
	} else {
		diags = spin.LoadConfig(config, dir)
	}

	for _, d := range diags {
		fmt.Println(d)
	}
	if diags.HasErrors() {
		os.Exit(1)
	}
	if Json {
		text, err := json.MarshalIndent(config, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(text))
	}
}

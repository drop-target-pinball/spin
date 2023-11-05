package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/drop-target-pinball/spin/v2"
	"github.com/hashicorp/hcl/v2"
)

var (
	flagHelp bool
	flagJSON bool
)

func main() {
	prog := path.Base(os.Args[0])

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <config_file...>\n", prog)
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, `
Parses an HCL configuration file.

If no options are specified, the program produces no output and exits with a
return code of zero upon success. Otherwise, diagnostics are printed to the
console and the program exits with a return code of one.

Use -json to print the configuration as a JSON document.
`)
	}

	flag.BoolVar(&flagJSON, "json", false, "print config as JSON")
	flag.BoolVar(&flagHelp, "help", false, "display help message")

	flag.Parse()
	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	conf := spin.NewConfig()
	for _, arg := range flag.Args() {
		if err := conf.AddFile(arg); err != nil {
			diag, ok := err.(hcl.Diagnostics)
			if !ok {
				fmt.Fprintln(os.Stderr, err)
			} else {
				for _, err := range diag.Errs() {
					fmt.Fprintln(os.Stderr, err)
				}
			}
			os.Exit(1)
		}
	}

	if flagJSON {
		out, err := json.MarshalIndent(conf, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(out))
	}
}

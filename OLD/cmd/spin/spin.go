package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/drop-target-pinball/spin/v2"
	"github.com/drop-target-pinball/spin/v2/pkg/sdl"
)

var settings spin.Settings

func main() {
	flag.StringVar(&settings.ConfigFile, "c", "project.hcl", "configuration file")
	flag.BoolVar(&settings.DevMode, "d", false, "development mode")
	flag.StringVar(&settings.Dir, "p", "./project", "project directory")
	flag.Parse()

	spin.AddDeviceFactory(sdl.AudioFactory{})

	eng := spin.NewEngine(&settings)
	if err := eng.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	eng.Run()
}

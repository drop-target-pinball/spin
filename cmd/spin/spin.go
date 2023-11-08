package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/drop-target-pinball/spin/v2"
	"github.com/drop-target-pinball/spin/v2/pkg/sdl"
)

var settings spin.Settings

func main() {
	flag.StringVar(&settings.ConfigFile, "c", "project.hcl", "configuration file")
	flag.StringVar(&settings.Dir, "d", "./project", "project directory")
	flag.Parse()

	spin.AddNewDeviceFunc(sdl.AudioHandlerName, sdl.NewAudioDevice)

	eng := spin.NewEngine(&settings)
	if err := eng.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	msg := spin.Play{
		ID:     "service_enter",
		Repeat: true,
		Notify: true,
	}
	eng.Send(msg)
	time.Sleep(1 * time.Second)
}

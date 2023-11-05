package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/drop-target-pinball/spin/v2"
)

var settings spin.Settings

func main() {
	flag.StringVar(&settings.Dir, "d", "./project", "project directory")
	flag.Parse()

	eng, err := spin.NewEngine(&settings)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(eng.Config)

	msg := spin.Play{
		ID:     "music",
		Repeat: true,
		Notify: true,
	}
	eng.Send(msg)
}

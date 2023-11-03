package main

import (
	"fmt"
	"os"

	"github.com/drop-target-pinball/spin/v2"
)

var settings spin.Settings

func main() {
	eng, err := spin.NewEngine(settings)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(eng.Config)
}

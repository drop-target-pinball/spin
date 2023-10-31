package main

import (
	"fmt"

	"github.com/drop-target-pinball/spin/v2"
)

var (
	runPort int
	varPort int
)

func main() {
	config, err := spin.LoadConfig()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("%+v\n", config)
}

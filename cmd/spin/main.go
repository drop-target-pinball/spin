package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/drop-target-pinball/spin/v2"
)

var (
	RedisAddr string
	Version   bool
)

func main() {
	log.SetFlags(0)

	flag.Usage = func() {
		fmt.Printf("Usage: %v [options]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.StringVar(&RedisAddr, "r", spin.DefaultRedisAddr, "redis address")
	flag.BoolVar(&Version, "v", false, "version")
	flag.Parse()

	fmt.Println(spin.Banner())
	if Version {
		return
	}

	settings := spin.DefaultSettings()
	e := spin.NewEngine(settings)
	err := e.Run()

	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

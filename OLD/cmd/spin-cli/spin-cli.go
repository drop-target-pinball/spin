package main

import (
	"flag"
	"log"

	"github.com/drop-target-pinball/spin/v2"
	"github.com/drop-target-pinball/spin/v2/app/cli"
)

var (
	addr string
)

func main() {
	log.SetFlags(0)
	flag.StringVar(&addr, "addr", "localhost:1080", "redis address to runtime database")
	flag.Parse()

	pin, err := spin.NewClient(addr)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	app := cli.NewApp(pin)
	app.Run()
}

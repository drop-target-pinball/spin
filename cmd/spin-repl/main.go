package main

import (
	"log"

	"github.com/drop-target-pinball/spin/app"
	"github.com/drop-target-pinball/spin/prog/boot"
	"github.com/drop-target-pinball/spin/prog/jdx"
)

func main() {
	log.SetFlags(0)
	opts := app.DefaultOptions()
	eng := app.NewEngine(opts)

	boot.Load(eng)
	jdx.Load(eng)

	repl := app.NewREPL(eng)
	go repl.Run()

	eng.Run()
}

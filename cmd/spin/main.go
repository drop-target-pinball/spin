package main

import (
	"log"

	"github.com/drop-target-pinball/spin/app"
	"github.com/drop-target-pinball/spin/prog"
)

func main() {
	log.SetFlags(0)
	opts := app.DefaultOptions()
	eng := app.NewEngine(opts)

	prog.Load(eng)

	repl := app.NewREPL(eng)
	go repl.Run()

	eng.Run()
}

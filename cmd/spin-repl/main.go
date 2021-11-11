package main

import (
	"log"

	"github.com/drop-target-pinball/spin/app"
	"github.com/drop-target-pinball/spin/prog/jdx"
)

func main() {
	log.SetFlags(0)
	eng := app.NewEngine(app.DefaultOptions())
	eng.Start()

	jdx.Load(eng)
	repl := app.NewREPL(eng)
	repl.Run()
}

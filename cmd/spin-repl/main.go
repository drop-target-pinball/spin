package main

import (
	"log"

	"github.com/drop-target-pinball/spin/app"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/prog/boot"
	"github.com/drop-target-pinball/spin/prog/jdx"
	"github.com/drop-target-pinball/spin/prog/sandbox"
)

func main() {
	log.SetFlags(0)
	opts := app.DefaultOptions()
	eng := app.NewEngine(opts)

	jd.Load(eng)
	boot.Load(eng)
	jdx.Load(eng)
	sandbox.Load(eng)

	repl := app.NewREPL(eng)
	go repl.Run()

	eng.Run()
}

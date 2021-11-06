package main

import (
	"log"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/app"
	"github.com/drop-target-pinball/spin/prog/jdx"
	"github.com/drop-target-pinball/spin/system/console"
	"github.com/drop-target-pinball/spin/system/sdl"
)

func main() {
	log.SetFlags(0)
	eng := spin.NewEngine()
	sdl.NewAudioSystem(eng)
	console.NewLoggingSystem(eng)
	eng.Start()

	jdx.Load(eng)
	repl := app.NewREPL(eng)
	repl.Run()
}

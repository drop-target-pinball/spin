package main

import (
	"flag"
	"log"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/app"
	"github.com/drop-target-pinball/spin/prog"
)

var (
	appOptions  = app.DefaultOptions()
	spinOptions = spin.DefaultOptions()

	noAudio bool
	script  string
)

func main() {
	log.SetFlags(0)

	flag.BoolVar(&spinOptions.RegisterEOS, "eos", false, "register end-of-stroke switches")
	flag.BoolVar(&noAudio, "no-audio", false, "disable audio")
	flag.BoolVar(&appOptions.WithPROC, "proc", false, "use P-ROC")
	flag.StringVar(&script, "script", "", "play script when starting")
	flag.Parse()

	if noAudio {
		appOptions.WithAudio = false
	}
	eng := app.NewEngine(appOptions, spinOptions)
	prog.Load(eng)

	repl := app.NewREPL(eng)
	go repl.Run()

	if script != "" {
		eng.Do(spin.PlayScript{ID: script})
	}
	eng.Run()
}

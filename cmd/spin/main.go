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
)

func main() {
	log.SetFlags(0)

	flag.BoolVar(&appOptions.WithPROC, "proc", false, "use P-ROC")
	flag.BoolVar(&spinOptions.RegisterEOS, "eos", false, "register end-of-stroke switches")
	flag.Parse()

	eng := app.NewEngine(appOptions, spinOptions)
	prog.Load(eng)

	repl := app.NewREPL(eng)
	go repl.Run()

	eng.Run()
}

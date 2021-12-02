package main

import (
	"flag"
	"log"

	"github.com/drop-target-pinball/spin/app"
	"github.com/drop-target-pinball/spin/prog"
)

var (
	options = app.DefaultOptions()
)

func main() {
	log.SetFlags(0)

	flag.BoolVar(&options.WithPROC, "proc", false, "use P-ROC")
	flag.Parse()

	eng := app.NewEngine(options)
	prog.Load(eng)

	repl := app.NewREPL(eng)
	go repl.Run()

	eng.Run()
}

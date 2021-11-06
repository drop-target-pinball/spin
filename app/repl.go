package app

import (
	"fmt"
	"io"
	"log"

	"github.com/chzyer/readline"
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/prog/jdx"
)

type REPL struct {
	eng *spin.Engine
	rl  *readline.Instance
}

func NewREPL(eng *spin.Engine) *REPL {
	rl, err := readline.NewEx(&readline.Config{
		Prompt: "spin> ",
	})
	if err != nil {
		log.Fatalf("unable to initialize readline: %v", err)
	}

	return &REPL{
		eng: eng,
		rl:  rl,
	}
}

func (r *REPL) Run() error {
	for {
		line, err := r.rl.Readline()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		result := r.Eval(line)
		if result != "" {
			fmt.Println(result)
		}
	}
}

func (r *REPL) Eval(line string) string {
	if line == "1" {
		r.eng.Do(spin.PlayMusic{ID: jdx.MainTheme})
	}
	if line == "2" {
		r.eng.Do(spin.PlayMusic{ID: jdx.MultiballTheme})
	}
	if line == "3" {
		r.eng.Do(spin.PlayScript{ID: jdx.SniperMode})
	}
	if line == "4" {
		r.eng.Do(spin.PlayScript{ID: jdx.SniperHunt})
	}
	if line == "5" {
		r.eng.Do(spin.PlayScript{ID: jdx.SniperFall})
	}
	if line == "9" {
		r.eng.Post(spin.SwitchEvent{ID: "popper"})
	}
	if line == "0" {
		r.eng.Do(spin.StopMusic{})
		r.eng.Do(spin.StopScript{ID: "*"})
	}
	return ""
}

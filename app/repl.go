package app

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/chzyer/readline"
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/prog/jdx"
	"github.com/drop-target-pinball/spin/terminal"
	"github.com/drop-target-pinball/spin/terminal/ansi"
)

const (
	progName = "spin"
)

type REPL struct {
	eng *spin.Engine
	rl  *readline.Instance
	out *terminal.Writer
}

func NewREPL(eng *spin.Engine) *REPL {
	rl, err := readline.NewEx(&readline.Config{
		Prompt: fmt.Sprintf("%v%v%v> ",
			ansi.LightGreen,
			progName,
			ansi.Reset),
	})
	if err != nil {
		log.Fatalf("unable to initialize readline: %v", err)
	}

	out := terminal.NewWriter(os.Stdout)
	out.RefreshFunc = func() { rl.Refresh() }
	log.SetOutput(out)

	return &REPL{
		eng: eng,
		rl:  rl,
		out: out,
	}
}

func (r *REPL) Run() error {
	for {
		line, err := r.rl.Readline()
		fmt.Printf("%v%v%v> %v%v\n",
			ansi.PreviousLine,
			ansi.BrightBlack,
			progName,
			line,
			ansi.Reset,
		)
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

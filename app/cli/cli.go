package cli

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"reflect"
	"strings"

	"github.com/adrg/xdg"
	"github.com/alexflint/go-arg"
	"github.com/chzyer/readline"
	"github.com/drop-target-pinball/spin/v2"
)

type messages struct {
	Load *spin.Load `arg:"subcommand:load"`
	Play *spin.Play `arg:"subcommand:play"`
}

type App struct {
	stream *spin.StreamClient
	rl     *readline.Instance
	out    *Writer
}

func NewApp(stream *spin.StreamClient) *App {
	stateDir := path.Join(xdg.StateHome, "spin")
	os.MkdirAll(stateDir, 0o750)
	historyFile := path.Join(stateDir, "history")

	rl, err := readline.NewEx(&readline.Config{
		Prompt: fmt.Sprintf("%v%v%v> ",
			AnsiLightGreen,
			spin.ShortName,
			AnsiReset),
		HistoryFile: historyFile,
	})
	if err != nil {
		log.Fatalf("unable to initialize readline: %v", err)
	}

	out := NewWriter(os.Stdout)
	out.RefreshFunc = func() { rl.Refresh() }
	log.SetOutput(out)

	return &App{
		stream: stream,
		rl:     rl,
		out:    out,
	}
}

func (a *App) Run() error {
	for {
		line, err := a.rl.Readline()
		fmt.Printf("%v%v%v> %v%v\n",
			AnsiPreviousLine,
			AnsiBrightBlack,
			spin.ShortName,
			line,
			AnsiReset,
		)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if err := a.Eval(line); err != nil {
			fmt.Printf("ERROR: %v\n", err)
		}
	}
}

func (a *App) Eval(line string) error {
	var msgs messages
	os.Args = append([]string{spin.ShortName}, strings.Split(line, " ")...)

	if err := arg.Parse(&msgs); err != nil {
		return err
	}
	val := reflect.ValueOf(msgs)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if !field.IsNil() {
			return a.stream.Send(field.Interface())
		}
	}
	return fmt.Errorf("FIXME")
}

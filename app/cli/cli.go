package cli

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"reflect"
	"strings"
	"time"
	"unicode"

	"github.com/adrg/xdg"
	"github.com/alexflint/go-arg"
	"github.com/chzyer/readline"
	"github.com/drop-target-pinball/spin/v2"
)

type Quit struct{}

type messages struct {
	Load *spin.Load `arg:"subcommand:load"`
	Play *spin.Play `arg:"subcommand:play"`
	Quit *struct{}  `arg:"subcommand:quit"`
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
	go a.HandleMessages()

	fmt.Fprintf(a.out, "%v%v\n", AnsiClearScreen, AnsiMoveToBottom)
	more := true
	for more {
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
		more, err = a.Eval(line)
		if err != nil {
			a.Error("%v", err)
		}
	}
	return nil
}

func (a *App) Eval(line string) (bool, error) {
	var msgs messages
	os.Args = append([]string{spin.ShortName}, strings.Split(line, " ")...)

	if err := arg.Parse(&msgs); err != nil {
		return true, err
	}
	if msgs.Quit != nil {
		return false, nil
	}
	val := reflect.ValueOf(msgs)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if !field.IsNil() {
			return true, a.stream.Send(field.Interface())
		}
	}
	panic("should be unreachable")
}

func (a *App) HandleMessages() {
	connected := true

	for {
		msg, err := a.stream.Read()
		if err != nil {
			if connected {
				a.Error("read error: %v", err)
				connected = false
			}
			time.Sleep(1 * time.Second)
			continue
		}
		fmt.Fprintln(a.out, formatMessage(msg))
	}
}

func (a *App) Error(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(a.out, "%v(*) %v%v\n", AnsiBrightRed, msg, AnsiReset)
}

func formatMessage(msg any) string {
	var out []string
	typ := reflect.TypeOf(msg)
	out = append(out, camelToSnake(typ.Name()))
	val := reflect.ValueOf(msg)
	for i := 0; i < typ.NumField(); i++ {
		fieldType := typ.Field(i)
		fieldValue := val.FieldByName(fieldType.Name)
		if fieldValue.IsZero() {
			continue
		}
		name := "--" + camelToSnake(fieldType.Name) + "="
		tag := fieldType.Tag.Get("arg")
		if strings.Contains(tag, "positional") {
			name = ""
		}
		val := fieldValue.Interface()
		out = append(out, fmt.Sprintf("%v%v", name, val))
	}
	return strings.Join(out, " ")
}

func camelToSnake(str string) string {
	var out strings.Builder
	for i, ch := range str {
		if unicode.IsUpper(ch) {
			if i != 0 {
				out.WriteRune('_')
			}
			out.WriteRune(unicode.ToLower(ch))
		} else {
			out.WriteRune(ch)
		}
	}
	return out.String()
}

package app

import (
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/chzyer/readline"
	"github.com/drop-target-pinball/spin"
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
	historyFile := ""
	homedir, err := os.UserHomeDir()
	if err != nil {
		spin.Warn("unable to determine home directory: %v", err)
	} else {
		historyFile = homedir + "/.spin-history"
	}

	rl, err := readline.NewEx(&readline.Config{
		Prompt: fmt.Sprintf("%v%v%v> ",
			ansi.LightGreen,
			progName,
			ansi.Reset),
		HistoryFile: historyFile,
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
	defer os.Exit(0)

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
		if err := r.Eval(line); err != nil {
			spin.Error(err.Error())
		}
	}
}

func (r *REPL) Eval(line string) error {
	fields := strings.Fields(line)
	if len(fields) == 0 {
		return nil
	}
	name := fields[0]
	props := make(map[string]string)
	args := fields[1:]
	for _, arg := range args {
		kv := strings.Split(arg, "=")
		if len(kv) != 2 {
			return fmt.Errorf("expecting property: %v", kv)
		}
		props[kv[0]] = kv[1]
	}
	proto, ok := r.getPrototype(name)
	if !ok {
		return fmt.Errorf("unknown action or event: %v", name)
	}

	t := reflect.TypeOf(proto)
	// First get the interface{} value
	intf := reflect.ValueOf(&proto).Elem()
	// Create a new value based on the underlying concrete value
	v := reflect.New(intf.Elem().Type()).Elem()
	for key, val := range props {
		f, ok := t.FieldByName(key)
		if !ok {
			return fmt.Errorf("unknown property: %v", key)
		}
		switch f.Type.Kind() {
		case reflect.String:
			v.FieldByName(key).SetString(val)
		case reflect.Int:
			iVal, err := strconv.Atoi(val)
			if err != nil {
				return fmt.Errorf("not an integer: %v", val)
			}
			v.FieldByName(key).SetInt(int64(iVal))
		case reflect.Float64:
			fVal, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return fmt.Errorf("not a float: %v", val)
			}
			v.FieldByName(key).SetFloat(fVal)
		default:
			return fmt.Errorf("cannot handle type %v: %v", f.Type.Kind(), val)
		}
	}

	switch p := v.Interface().(type) {
	case spin.Action:
		r.eng.Do(p)
	case spin.Event:
		r.eng.Post(p)
	default:
		return fmt.Errorf("unexpected payload")
	}
	return nil
}

func (r *REPL) getPrototype(name string) (interface{}, bool) {
	act, ok := r.eng.Actions[name]
	if ok {
		return act, true
	}
	evt, ok := r.eng.Events[name]
	if ok {
		return evt, true
	}
	return nil, false
}

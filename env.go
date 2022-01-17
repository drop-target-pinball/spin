package spin

import (
	"time"

	"github.com/drop-target-pinball/coroutine"
)

type ScriptEnv struct {
	Config Config
	eng    *Engine
	co     *coroutine.C
}

func newScriptEnv(eng *Engine, co *coroutine.C) *ScriptEnv {
	return &ScriptEnv{
		Config: eng.Config,
		eng:    eng,
		co:     co,
	}
}

func (e *ScriptEnv) Do(act Action) {
	e.eng.Do(act)
}

func (e *ScriptEnv) Post(evt Event) {
	e.eng.Post(evt)
}

func (e *ScriptEnv) Sleep(ms int) bool {
	return e.co.Sleep(time.Duration(ms) * time.Millisecond)
}

func (e *ScriptEnv) WaitFor(events ...coroutine.Event) (coroutine.Event, bool) {
	return e.co.WaitFor(events...)
}

func (e *ScriptEnv) WaitForUntil(ms int, s ...coroutine.Event) (coroutine.Event, bool) {
	return e.co.WaitForUntil(time.Duration(ms)*time.Millisecond, s...)
}

func (e *ScriptEnv) Display(id string) Display {
	return e.eng.Display(id)
}

func (e *ScriptEnv) NewCoroutine(fn ScriptFn) {
	e.co.New(func(co *coroutine.C) {
		fn(newScriptEnv(e.eng, co))
	})
}

func (e *ScriptEnv) RegisterVars(name string, vars interface{}) {
	e.eng.RegisterVars(name, vars)
}

func (e *ScriptEnv) GetVars(name string) (interface{}, bool) {
	return e.eng.GetVars(name)
}

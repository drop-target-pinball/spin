package system

import (
	"context"
	"log"
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/coroutine"
)

type env struct {
	eng      *spin.Engine
	displays map[string]spin.Display
	ctx      *coroutine.Context
	//eventQueue chan spin.Event
}

func (e *env) Do(act spin.Action) {
	e.eng.Do(act)
}

func (e *env) Post(evt spin.Event) {
	e.eng.Post(evt)
}

func (e *env) WaitForUntil(d time.Duration, s ...coroutine.Selector) coroutine.Selector {
	return e.ctx.WaitForUntil(d, s...)
}

func (e *env) WaitFor(d time.Duration) bool {
	return e.ctx.WaitFor(d)
}

func (e *env) WaitUntil(s ...coroutine.Selector) coroutine.Selector {
	return e.ctx.WaitUntil(s...)
}

// func (e *env) EventQueue() chan spin.Event {
// 	return e.eventQueue
// }

func (e *env) Display(id string) spin.Display {
	r, ok := e.displays[id]
	if !ok {
		log.Panicf("no such display: %v", id)
	}
	return r
}

func (e *env) Int(ns string, id string) int {
	return e.eng.Namespaces.Get(ns).Int(id)
}

func (e *env) SetInt(ns string, id string, val int) {
	e.eng.Namespaces.Get(ns).SetInt(id, val)
}

func (e *env) AddInt(ns string, id string, val int) {
	e.eng.Namespaces.Get(ns).AddInt(id, val)
}

func (e *env) String(ns string, id string) string {
	return e.eng.Namespaces.Get(ns).String(id)
}

func (e *env) SetString(ns string, id string, val string) {
	e.eng.Namespaces.Get(ns).SetString(id, val)
}

type ScriptRunner struct {
	eng      *spin.Engine
	scripts  map[string]spin.Script
	running  map[string]context.CancelFunc
	displays map[string]spin.Display
	env      map[string]spin.Env
	runner   *coroutine.Runner

	//mutex    sync.Mutex
}

func RegisterScriptRunner(eng *spin.Engine) {
	sys := &ScriptRunner{
		eng:      eng,
		scripts:  make(map[string]spin.Script),
		running:  make(map[string]context.CancelFunc),
		displays: make(map[string]spin.Display),
		env:      make(map[string]spin.Env),
		runner:   coroutine.NewRunner(),
	}
	eng.RegisterActionHandler(sys)
	eng.RegisterEventHandler(sys)
	eng.RegisterServer(sys)
}

func (s *ScriptRunner) HandleAction(action spin.Action) {
	switch act := action.(type) {
	case spin.RegisterDisplaySDL:
		s.registerDisplaySDL(act)
	case spin.RegisterScript:
		s.registerScript(act)
	case spin.PlayScript:
		s.playScript(act)
	case spin.StopScript:
		s.stopScripts(act)
	}
}

func (s *ScriptRunner) HandleEvent(evt spin.Event) {
	s.runner.Post(evt)
}

func (s *ScriptRunner) Service() {
	s.runner.Service()
}

func (s *ScriptRunner) registerDisplaySDL(act spin.RegisterDisplaySDL) {
	s.displays[act.ID] = act.Display
}

func (s *ScriptRunner) registerScript(a spin.RegisterScript) {
	s.scripts[a.ID] = a.Script
	s.running[a.ID] = nil
}

func (s *ScriptRunner) playScript(a spin.PlayScript) {
	// s.mutex.Lock()
	// defer s.mutex.Unlock()

	scr, ok := s.scripts[a.ID]
	if !ok {
		spin.Warn("%v: no such script", a.ID)
		return
	}
	if cancel := s.running[a.ID]; cancel != nil {
		cancel()
	}

	// env := &env{
	// 	eng: s.eng,
	// 	// eventQueue: make(chan spin.Event, 10),
	// 	displays: s.displays,
	// }
	//s.env[a.ID] = env

	ctx, cancel := context.WithCancel(context.Background())
	s.running[a.ID] = cancel

	s.runner.Create(ctx, func(ctx *coroutine.Context) {
		e := &env{
			eng: s.eng,
			// eventQueue: make(chan spin.Event, 10),
			displays: s.displays,
			ctx:      ctx,
		}
		scr(e)
	})

	// go func() {
	// 	scr(ctx, env)
	// 	s.stopScript(a.ID)
	// }()
}

func (s *ScriptRunner) stopScript(id string) {
	// s.mutex.Lock()
	// defer s.mutex.Unlock()

	cancel, ok := s.running[id]
	if !ok {
		spin.Warn("%v: no such script", id)
		return
	}
	if cancel == nil {
		return
	}
	cancel()
	s.running[id] = nil
	s.env[id] = nil
}

func (s *ScriptRunner) stopScripts(a spin.StopScript) {
	if a.ID == "*" {
		for id := range s.running {
			s.stopScript(id)
		}
	} else {
		s.stopScript(a.ID)
	}
}

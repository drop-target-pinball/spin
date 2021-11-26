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
}

func (e *env) Do(act spin.Action) {
	e.eng.Do(act)
}

func (e *env) Post(evt spin.Event) {
	e.eng.Post(evt)
}

func (e *env) Sleep(d time.Duration) bool {
	return e.ctx.Sleep(d)
}

func (e *env) WaitFor(s ...coroutine.Selector) (coroutine.Selector, bool) {
	return e.ctx.WaitFor(s...)
}

func (e *env) Display(id string) spin.Display {
	r, ok := e.displays[id]
	if !ok {
		log.Panicf("no such display: %v", id)
	}
	return r
}

func (e *env) Derive() (context.Context, context.CancelFunc) {
	return e.ctx.Derive()
}

func (e *env) NewCoroutine(ctx context.Context, scr spin.Script) {
	coroutine.Create(ctx, func(ctx *coroutine.Context) {
		e := &env{
			eng:      e.eng,
			displays: e.displays,
			ctx:      ctx,
		}
		scr(e)
	})

}

type ScriptRunner struct {
	eng      *spin.Engine
	scripts  map[string]spin.Script
	running  map[string]context.CancelFunc
	displays map[string]spin.Display
}

func RegisterScriptRunner(eng *spin.Engine) {
	sys := &ScriptRunner{
		eng:      eng,
		scripts:  make(map[string]spin.Script),
		running:  make(map[string]context.CancelFunc),
		displays: make(map[string]spin.Display),
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
	coroutine.Post(evt)
}

func (s *ScriptRunner) Service() {
	coroutine.Service()
}

func (s *ScriptRunner) registerDisplaySDL(act spin.RegisterDisplaySDL) {
	s.displays[act.ID] = act.Display
}

func (s *ScriptRunner) registerScript(a spin.RegisterScript) {
	s.scripts[a.ID] = a.Script
	s.running[a.ID] = nil
}

func (s *ScriptRunner) playScript(a spin.PlayScript) {
	scr, ok := s.scripts[a.ID]
	if !ok {
		spin.Warn("%v: no such script", a.ID)
		return
	}
	if cancel := s.running[a.ID]; cancel != nil {
		cancel()
	}

	ctx, cancel := context.WithCancel(context.Background())
	s.running[a.ID] = cancel

	coroutine.Create(ctx, func(ctx *coroutine.Context) {
		e := &env{
			eng:      s.eng,
			displays: s.displays,
			ctx:      ctx,
		}
		scr(e)
	})
}

func (s *ScriptRunner) stopScript(id string) {
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

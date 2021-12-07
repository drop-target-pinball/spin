package spin

import (
	"context"
	"log"
	"time"

	"github.com/drop-target-pinball/spin/coroutine"
)

const (
	FrameDuration = 16670 * time.Microsecond
)

type Script func(Env)

type Env struct {
	Config   Config
	eng      *Engine
	displays map[string]Display
	ctx      *coroutine.Context
}

func (e Env) Do(act Action) {
	e.eng.Do(act)
}

func (e Env) Post(evt Event) {
	e.eng.Post(evt)
}

func (e Env) Sleep(d time.Duration) bool {
	return e.ctx.Sleep(d)
}

func (e Env) WaitFor(s ...coroutine.Selector) (coroutine.Selector, bool) {
	return e.ctx.WaitFor(s...)
}

func (e Env) WaitForUntil(d time.Duration, s ...coroutine.Selector) (coroutine.Selector, bool) {
	return e.ctx.WaitForUntil(d, s...)
}

func (e Env) Display(id string) Display {
	r, ok := e.displays[id]
	if !ok {
		log.Panicf("no such display: %v", id)
	}
	return r
}

func (e Env) Derive() (context.Context, context.CancelFunc) {
	return e.ctx.Derive()
}

var anonCounter = 1

func (e Env) NewCoroutine(ctx context.Context, scr Script) {
	coroutine.New(ctx, func(ctx *coroutine.Context) {
		e := Env{
			Config:   e.eng.Config,
			eng:      e.eng,
			displays: e.displays,
			ctx:      ctx,
		}
		scr(e)
	})
}

func (e Env) Vars(name string) (interface{}, bool) {
	return e.eng.Vars(name)
}

func (e Env) RegisterVars(name string, vars interface{}) {
	e.eng.RegisterVars(name, vars)
}

type scriptSystem struct {
	eng      *Engine
	scripts  map[string]Script
	running  map[string]context.CancelFunc
	displays map[string]Display
}

func registerScriptSystem(eng *Engine) {
	sys := &scriptSystem{
		eng:      eng,
		scripts:  make(map[string]Script),
		running:  make(map[string]context.CancelFunc),
		displays: make(map[string]Display),
	}
	eng.RegisterActionHandler(sys)
	eng.RegisterEventHandler(sys)
	eng.RegisterServer(sys)
}

func (s *scriptSystem) HandleAction(action Action) {
	switch act := action.(type) {
	case RegisterDisplay:
		s.registerDisplay(act)
	case RegisterScript:
		s.registerScript(act)
	case PlayScript:
		s.playScript(act)
	case StopScript:
		s.stopScripts(act)
	}
}

func (s *scriptSystem) HandleEvent(evt Event) {
	coroutine.Post(evt)
}

func (s *scriptSystem) Service() {
	coroutine.Service()
}

func (s *scriptSystem) registerDisplay(act RegisterDisplay) {
	s.displays[act.ID] = act.Display
}

func (s *scriptSystem) registerScript(a RegisterScript) {
	s.scripts[a.ID] = a.Script
	s.running[a.ID] = nil
}

func (s *scriptSystem) playScript(a PlayScript) {
	scr, ok := s.scripts[a.ID]
	if !ok {
		Warn("%v: no such script", a.ID)
		return
	}
	if cancel := s.running[a.ID]; cancel != nil {
		cancel()
	}

	ctx, cancel := context.WithCancel(context.Background())
	s.running[a.ID] = cancel

	coroutine.New(ctx, func(ctx *coroutine.Context) {
		e := Env{
			Config:   s.eng.Config,
			eng:      s.eng,
			displays: s.displays,
			ctx:      ctx,
		}
		scr(e)
	})
}

func (s *scriptSystem) stopScript(id string) {
	cancel, ok := s.running[id]
	if !ok {
		Warn("%v: no such script", id)
		return
	}
	if cancel == nil {
		return
	}
	cancel()
	s.running[id] = nil
}

func (s *scriptSystem) stopScripts(a StopScript) {
	if a.ID == "*" {
		for id := range s.running {
			s.stopScript(id)
		}
	} else {
		s.stopScript(a.ID)
	}
}

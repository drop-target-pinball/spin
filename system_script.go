package spin

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/drop-target-pinball/spin/coroutine"
)

const (
	ScopeInit     = "init"
	ScopeRoot     = "spin"
	ScopeProgram  = "spin.program"
	ScopeGame     = "spin.program.game"
	ScopeBall     = "spin.program.game.ball"
	ScopeMode     = "spin.program.game.ball.mode"
	ScopePriority = "spin.program.game.ball.priority"
)

const (
	DebugScripts = "Scripts"
)

const (
	FrameDuration = 16670 * time.Microsecond
)

type ScriptFn func(Env)

type script struct {
	id        string
	scope     string
	fn        ScriptFn
	coroutine coroutine.ID
}

type Env struct {
	Config   Config
	eng      *Engine
	displays map[string]Display
	ctx      *coroutine.Context
	scripts  map[string]*script
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

func (e Env) Context() context.Context {
	return e.ctx.Context()
}

func (e Env) IsActive(id string) bool {
	script, ok := e.scripts[id]
	if !ok {
		return false
	}
	return coroutine.IsActive(script.coroutine)
}

func (e Env) NewCoroutine(ctx context.Context, scr ScriptFn) {
	coroutine.New(ctx, func(ctx *coroutine.Context) {
		e := Env{
			Config:   e.eng.Config,
			eng:      e.eng,
			displays: e.displays,
			ctx:      ctx,
			scripts:  e.scripts,
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
	scripts  map[string]*script
	displays map[string]Display
}

func registerScriptSystem(eng *Engine) {
	sys := &scriptSystem{
		eng:      eng,
		scripts:  make(map[string]*script),
		displays: make(map[string]Display),
	}
	eng.RegisterActionHandler(sys)
	eng.RegisterEventHandler(sys)
	eng.RegisterServer(sys)
}

func (s *scriptSystem) HandleAction(action Action) {
	switch act := action.(type) {
	case Debug:
		s.debug(act)
	case RegisterDisplay:
		s.registerDisplay(act)
	case RegisterScript:
		s.registerScript(act)
	case PlayScript:
		s.playScript(act)
	case StopScope:
		s.stopScope(act)
	case StopScript:
		s.stopScript(act)
	}
}

func (s *scriptSystem) HandleEvent(event Event) {
	coroutine.Post(event)
}

func (s *scriptSystem) Service() {
	coroutine.Service()
}

func (s *scriptSystem) debug(evt Debug) {
	switch evt.ID {
	case DebugScripts:
		s.debugScripts()
	}
}

func (s *scriptSystem) debugScripts() {
	running := make([]string, 0)
	for _, script := range s.scripts {
		if coroutine.IsActive(script.coroutine) {
			running = append(running, fmt.Sprintf("%v: %v", script.scope, script.id))
		}
	}
	sort.Strings(running)
	for _, name := range running {
		Log(name)
	}
}

func (s *scriptSystem) registerDisplay(act RegisterDisplay) {
	s.displays[act.ID] = act.Display
}

func (s *scriptSystem) registerScript(a RegisterScript) {
	scope := a.Scope
	if scope == "" {
		scope = "spin"
	}
	s.scripts[a.ID] = &script{
		id:    a.ID,
		scope: scope,
		fn:    a.Script,
	}
}

func (s *scriptSystem) playScript(a PlayScript) {
	scr, ok := s.scripts[a.ID]
	if !ok {
		Warn("no such script: %v", a.ID)
		return
	}
	coroutine.Cancel(scr.coroutine)
	id := coroutine.New(context.Background(), func(ctx *coroutine.Context) {
		e := Env{
			Config:   s.eng.Config,
			eng:      s.eng,
			displays: s.displays,
			ctx:      ctx,
			scripts:  s.scripts,
		}
		scr.fn(e)
	})
	scr.coroutine = id
}

func (s *scriptSystem) stopScript(act StopScript) {
	scr, ok := s.scripts[act.ID]
	if !ok {
		Warn("no such script: %v", act.ID)
		return
	}
	coroutine.Cancel(scr.coroutine)
}

func (s *scriptSystem) stopScope(a StopScope) {
	stopPrefix := a.ID + "."
	for _, scr := range s.scripts {
		scope := scr.scope + "."
		if strings.HasPrefix(scope, stopPrefix) {
			coroutine.Cancel(scr.coroutine)
		}
	}
}

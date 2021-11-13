package system

import (
	"context"
	"log"

	"github.com/drop-target-pinball/spin"
)

type env struct {
	eng        *spin.Engine
	eventQueue chan spin.Event
}

func (e *env) Do(act spin.Action) {
	e.eng.Do(act)
}

func (e *env) Post(evt spin.Event) {
	e.eng.Post(evt)
}

func (e *env) EventQueue() chan spin.Event {
	return e.eventQueue
}

func (e *env) RenderTargetSDL(name string) *spin.RenderTargetSDL {
	r, ok := e.eng.RenderTargetSDL[name]
	if !ok {
		log.Panicf("no such SDL renderer: %v", name)
	}
	return r
}

type ScriptRunner struct {
	eng     *spin.Engine
	scripts map[string]spin.Script
	running map[string]context.CancelFunc
	env     map[string]spin.Env
}

func NewScriptRunner(eng *spin.Engine) *ScriptRunner {
	sys := &ScriptRunner{
		eng:     eng,
		scripts: make(map[string]spin.Script),
		running: make(map[string]context.CancelFunc),
		env:     make(map[string]spin.Env),
	}
	eng.RegisterActionHandler(sys)
	eng.RegisterEventHandler(sys)
	return sys
}

func (s *ScriptRunner) HandleAction(a spin.Action) {
	switch action := a.(type) {
	case spin.RegisterScript:
		s.registerScript(action)
	case spin.PlayScript:
		s.playScript(action)
	case spin.StopScript:
		s.stopScripts(action)
	}
}

func (s *ScriptRunner) HandleEvent(evt spin.Event) {
	for _, env := range s.env {
		if env != nil {
			env.EventQueue() <- evt
		}
	}
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

	env := &env{eng: s.eng, eventQueue: make(chan spin.Event, 1)}
	s.env[a.ID] = env

	ctx, cancel := context.WithCancel(context.Background())
	s.running[a.ID] = cancel

	go func() {
		scr(ctx, env)
		s.stopScript(a.ID)
	}()
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
	close(s.env[id].EventQueue())
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

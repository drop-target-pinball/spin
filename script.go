package spin

import (
	"context"
	"time"
)

type Env struct {
	eng    *Engine
	Events chan Event
}

func (e *Env) Do(act Action) {
	e.eng.Do(act)
}

func (e *Env) Post(evt Event) {
	e.eng.Post(evt)
}

type Script func(context.Context, *Env)

type ScriptSystem struct {
	eng     *Engine
	scripts map[string]Script
	running map[string]context.CancelFunc
	env     map[string]*Env
}

func NewScriptSystem(eng *Engine) *ScriptSystem {
	sys := &ScriptSystem{
		eng:     eng,
		scripts: make(map[string]Script),
		running: make(map[string]context.CancelFunc),
		env:     make(map[string]*Env),
	}
	eng.RegisterActionHandler(sys)
	eng.RegisterEventHandler(sys)
	return sys
}

func (s *ScriptSystem) HandleAction(a Action) {
	switch action := a.(type) {
	case RegisterScript:
		s.registerScript(action)
	case PlayScript:
		s.playScript(action)
	case StopScript:
		s.stopScripts(action)
	}
}

func (s *ScriptSystem) HandleEvent(evt Event) {
	for _, env := range s.env {
		if env != nil {
			env.Events <- evt
		}
	}
}

func (s *ScriptSystem) Service() {}

func (s *ScriptSystem) registerScript(a RegisterScript) {
	s.scripts[a.ID] = a.Script
	s.running[a.ID] = nil
}

func (s *ScriptSystem) playScript(a PlayScript) {
	scr, ok := s.scripts[a.ID]
	if !ok {
		Warn("%v: no such script", a.ID)
		return
	}
	if cancel := s.running[a.ID]; cancel != nil {
		cancel()
	}

	env := &Env{eng: s.eng, Events: make(chan Event, 1)}
	s.env[a.ID] = env

	ctx, cancel := context.WithCancel(context.Background())
	s.running[a.ID] = cancel

	go func() {
		scr(ctx, env)
		s.stopScript(a.ID)
	}()
}

func (s *ScriptSystem) stopScript(id string) {
	cancel, ok := s.running[id]
	if !ok {
		Warn("%v: no such script", id)
		return
	}
	if cancel == nil {
		return
	}
	cancel()
	close(s.env[id].Events)
	s.running[id] = nil
	s.env[id] = nil
}

func (s *ScriptSystem) stopScripts(a StopScript) {
	if a.ID == "*" {
		for id := range s.running {
			s.stopScript(id)
		}
	} else {
		s.stopScript(a.ID)
	}
}

func Wait(ctx context.Context, d time.Duration) bool {
	select {
	case <-time.After(d):
		return false
	case <-ctx.Done():
		return true
	}
}

func WaitForSwitch(ctx context.Context, e *Env, id string) (bool, SwitchEvent) {
	for {
		select {
		case event := <-e.Events:
			sw, ok := event.(SwitchEvent)
			if ok && sw.ID == id {
				return false, sw
			}
		case <-ctx.Done():
			return true, SwitchEvent{}
		}
	}
}

func WaitForMessage(ctx context.Context, e *Env, id string) (bool, Message) {
	for {
		select {
		case event := <-e.Events:
			msg, ok := event.(Message)
			if ok && msg.ID == id {
				return false, msg
			}
		case <-ctx.Done():
			return true, Message{}
		}
	}
}

func WaitForSwitchUntil(ctx context.Context, e *Env, id string, d time.Duration) (bool, SwitchEvent) {
	timer := time.After(d)
	for {
		select {
		case event := <-e.Events:
			sw, ok := event.(SwitchEvent)
			if ok && sw.ID == id {
				return false, sw
			}
		case <-timer:
			return false, SwitchEvent{}
		case <-ctx.Done():
			return true, SwitchEvent{}
		}
	}
}

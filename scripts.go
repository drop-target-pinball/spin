package spin

import (
	"context"
	"time"

	"github.com/drop-target-pinball/spin/coroutine"
)

func RenderFrameScript(e Env, fn func(Env)) context.CancelFunc {
	ctx, cancel := e.Derive()
	e.NewCoroutine(ctx, func(e Env) {
		for {
			fn(e)
			if done := e.Sleep(FrameDuration); done {
				return
			}
		}
	})
	return cancel
}

func CountdownScript(e Env, timer *int, tickMs int, end Event) context.CancelFunc {
	ctx, cancel := e.Derive()
	e.NewCoroutine(ctx, func(e Env) {
		for *timer > 0 {
			if done := e.Sleep(time.Duration(tickMs) * time.Millisecond); done {
				return
			}
			*timer -= 1
		}
		e.Post(end)
	})
	return cancel
}

type opSleep struct {
	t time.Duration
}

type opDo struct {
	act Action
}

type opPost struct {
	evt Event
}

type opStart struct {
	seq *Sequencer
}

type opWaitFor struct {
	selectors []coroutine.Selector
}

type opLoop struct {
	n int
}

type opFunc struct {
	fn func()
}

type Sequencer struct {
	ops    []interface{}
	defers []Action
}

func NewSequencer() *Sequencer {
	return &Sequencer{
		ops:    make([]interface{}, 0),
		defers: make([]Action, 0),
	}
}

func (s *Sequencer) Sleep(ms int) *Sequencer {
	s.ops = append(s.ops, opSleep{time.Duration(ms) * time.Millisecond})
	return s
}

func (s *Sequencer) Do(act Action) *Sequencer {
	s.ops = append(s.ops, opDo{act})
	return s
}

func (s *Sequencer) Post(evt Event) *Sequencer {
	s.ops = append(s.ops, opPost{evt})
	return s
}

func (s *Sequencer) WaitFor(selectors ...coroutine.Selector) *Sequencer {
	s.ops = append(s.ops, opWaitFor{selectors})
	return s
}

func (s *Sequencer) Loop() *Sequencer {
	s.ops = append(s.ops, opLoop{})
	return s
}

func (s *Sequencer) Start(seq *Sequencer) *Sequencer {
	s.ops = append(s.ops, opStart{seq})
	return s
}

func (s *Sequencer) Func(fn func()) *Sequencer {
	s.ops = append(s.ops, opFunc{fn})
	return s
}

func (s *Sequencer) Defer(act Action) *Sequencer {
	s.defers = append(s.defers, act)
	return s
}

func (s *Sequencer) Run0(ctx context.Context, env Env) {
	exit := func() {
		for _, act := range s.defers {
			env.Do(act)
		}
	}

	env.NewCoroutine(ctx, func(e Env) {
		pc := 0
		for {
			if pc >= len(s.ops) {
				break
			}
			operation := s.ops[pc]
			switch op := operation.(type) {
			case opSleep:
				if done := e.Sleep(op.t); done {
					exit()
					return
				}
			case opDo:
				e.Do(op.act)
			case opPost:
				e.Post(op.evt)
			case opWaitFor:
				if _, done := e.WaitFor(op.selectors...); done {
					exit()
					return
				}
			case opStart:
				op.seq.Run0(ctx, e)
			case opLoop:
				pc = 0
				continue
			}
			pc += 1
		}
		exit()
	})
}

func (s *Sequencer) Run(e Env) {
	cancels := make([]Action, 0)

	defer func() {
		for _, act := range s.defers {
			e.Do(act)
		}
		for _, act := range cancels {
			e.Do(act)
		}
	}()

	pc := 0
	for {
		if pc >= len(s.ops) {
			break
		}
		operation := s.ops[pc]
		switch op := operation.(type) {
		case opSleep:
			if done := e.Sleep(op.t); done {
				return
			}
			cancels = nil
		case opDo:
			e.Do(op.act)
			switch act := op.act.(type) {
			case PlaySpeech:
				cancels = append(cancels, StopSpeech{ID: act.ID})
			case PlaySound:
				cancels = append(cancels, StopSound{ID: act.ID})
			}
		case opPost:
			e.Post(op.evt)
		case opWaitFor:
			if _, done := e.WaitFor(op.selectors...); done {
				return
			}
			cancels = nil
		case opFunc:
			op.fn()
		case opLoop:
			pc = 0
			continue
		}
		pc += 1
	}
}

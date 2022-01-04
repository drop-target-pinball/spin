package spin

import (
	"time"

	"github.com/drop-target-pinball/spin/coroutine"
)

type opSetIntVar struct {
	ptr *int
	val int
}

type opDo struct {
	act Action
}

type opFunc struct {
	fn func()
}

type opLoop struct {
}

type opPost struct {
	evt Event
}

type opRepeat struct {
	n int
}

type opSequence struct {
	seq *Sequencer
}

type opSleep struct {
	t time.Duration
}

type opWaitFor struct {
	selectors []coroutine.Selector
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

func (s *Sequencer) Defer(act Action) *Sequencer {
	s.defers = append(s.defers, act)
	return s
}

func (s *Sequencer) Do(act Action) *Sequencer {
	s.ops = append(s.ops, opDo{act})
	return s
}

func (s *Sequencer) Func(fn func()) *Sequencer {
	s.ops = append(s.ops, opFunc{fn})
	return s
}

func (s *Sequencer) Loop() *Sequencer {
	s.ops = append(s.ops, opLoop{})
	return s
}

func (s *Sequencer) Post(evt Event) *Sequencer {
	s.ops = append(s.ops, opPost{evt})
	return s
}

func (s *Sequencer) Repeat(n int) *Sequencer {
	s.ops = append(s.ops, opRepeat{n})
	return s
}

func (s *Sequencer) SetIntVar(ptr *int, val int) *Sequencer {
	s.ops = append(s.ops, opSetIntVar{ptr, val})
	return s
}

func (s *Sequencer) Sequence(seq *Sequencer) *Sequencer {
	s.ops = append(s.ops, opSequence{seq})
	return s
}

func (s *Sequencer) Sleep(ms int) *Sequencer {
	s.ops = append(s.ops, opSleep{time.Duration(ms) * time.Millisecond})
	return s
}

func (s *Sequencer) WaitFor(selectors ...coroutine.Selector) *Sequencer {
	s.ops = append(s.ops, opWaitFor{selectors})
	return s
}

func (s *Sequencer) Run(e Env) bool {
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
	repeat := false
	n := 0

	for {
		if pc >= len(s.ops) {
			break
		}
		operation := s.ops[pc]
		switch op := operation.(type) {
		case opSetIntVar:
			*op.ptr = op.val
		case opSleep:
			if done := e.Sleep(op.t); done {
				return true
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
		case opRepeat:
			if !repeat {
				repeat = true
				n = op.n
			} else {
				n -= 1
			}
			if n > 0 {
				pc = 0
				continue
			}
			repeat = false
		case opSequence:
			if done := op.seq.Run(e); done {
				return true
			}
		case opWaitFor:
			if _, done := e.WaitFor(op.selectors...); done {
				return true
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
	return false
}

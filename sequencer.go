package spin

import (
	"time"

	"github.com/drop-target-pinball/coroutine"
)

type Sequencer struct {
	seq *coroutine.Sequencer
	env *ScriptEnv
}

func NewSequencer(env *ScriptEnv) *Sequencer {
	return &Sequencer{
		seq: coroutine.NewSequencer(),
		env: env,
	}
}

func (s *Sequencer) Defer(act Action) {
	s.seq.Defer(func() { s.env.Do(act) })
}

func (s *Sequencer) Do(act Action) {
	s.seq.Do(func() { s.env.Do(act) })
	switch a := act.(type) {
	case PlaySpeech:
		s.seq.Cancel(func() { s.env.Do(StopSpeech{ID: a.ID}) })
	case PlaySound:
		s.seq.Cancel(func() { s.env.Do(StopSound{ID: a.ID}) })
	}
}

func (s *Sequencer) DoFunc(fn func()) {
	s.seq.Do(fn)
}

// func (s *Sequencer) DoRun(fn func() bool) {
// 	s.seq.DoRun(fn)
// }

func (s *Sequencer) Loop() {
	s.seq.Loop()
}

func (s *Sequencer) LoopN(n int) {
	s.seq.LoopN(n)
}

func (s *Sequencer) Post(evt Event) {
	s.seq.Do(func() { s.env.Post(evt) })
}

func (s *Sequencer) Sleep(ms int) {
	s.seq.Sleep(time.Duration(ms) * time.Millisecond)
}

func (s *Sequencer) WaitFor(events ...coroutine.Event) {
	s.seq.WaitFor(events...)
}

func (s *Sequencer) WaitForUntil(ms int, events ...coroutine.Event) {
	s.seq.WaitForUntil(time.Duration(ms)*time.Millisecond, events...)
}

func (s *Sequencer) Run() bool {
	return s.seq.Run(s.env.co)
}

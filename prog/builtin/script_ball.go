package builtin

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/coroutine"
)

func ballDrainScript(e spin.Env) {
	for {
		if _, done := e.WaitFor(spin.SwitchEvent{ID: e.Config.SwitchDrain}); done {
			return
		}
		e.Post(spin.BallDrainEvent{})
	}
}

func ballWillDrainScript(e spin.Env) {
	events := make([]coroutine.Selector, len(e.Config.SwitchWillDrain))
	for i, sw := range e.Config.SwitchWillDrain {
		events[i] = spin.SwitchEvent{ID: sw}
	}

	for {
		if _, done := e.WaitFor(events...); done {
			return
		}
		e.Post(spin.BallWillDrainEvent{})
	}
}

func ballTrackerScript(e spin.Env) {
	ctx, _ := e.Derive()
	e.NewCoroutine(ctx, ballDrainScript)
	e.NewCoroutine(ctx, ballWillDrainScript)

	for {
		if _, done := e.WaitFor(spin.Done{}); done {
			return
		}
	}
}

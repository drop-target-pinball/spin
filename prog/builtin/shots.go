package builtin

import (
	"time"

	"github.com/drop-target-pinball/spin"
)

func ShotTrapScript(e spin.Env, sw string, shot string, t time.Duration) {
	for {
		if _, done := e.WaitFor(spin.SwitchEvent{ID: sw}); done {
			return
		}
		evt, done := e.WaitForUntil(t, spin.SwitchEvent{ID: sw, Released: true})
		if done {
			return
		}
		if evt != (spin.SwitchEvent{ID: sw, Released: true}) {
			e.Post(spin.ShotEvent{ID: shot})
		}
	}
}

func ShotSequenceScript(e spin.Env, switches []string, shot string, t time.Duration) {
	for {
		if _, done := e.WaitFor(spin.SwitchEvent{ID: switches[0]}); done {
			return
		}
		shotMade := true
		for _, sw := range switches[1:] {
			evt, done := e.WaitForUntil(t, spin.SwitchEvent{ID: sw})
			if done {
				return
			}
			if evt == nil {
				shotMade = false
				break
			}
		}
		if shotMade {
			e.Post(spin.ShotEvent{ID: shot})
		}
	}
}

func ShotSwitchScript(e spin.Env, sw string, shot string) {
	for {
		if _, done := e.WaitFor(spin.SwitchEvent{ID: sw}); done {
			return
		}
		e.Post(spin.ShotEvent{ID: shot})
	}
}

package builtin

import (
	"time"

	"github.com/drop-target-pinball/spin"
)

func ShotTrapScript(e spin.Env, sw string, shot string, t time.Duration) {
	spin.Log("** STARTED: %v", shot)
	for {
		if _, done := e.WaitFor(spin.SwitchEvent{ID: sw}); done {
			return
		}
		spin.Log("** ON")
		evt, done := e.WaitForUntil(t, spin.SwitchEvent{ID: sw, Released: true})
		if done {
			return
		}
		spin.Log("** OFF")
		if evt != (spin.SwitchEvent{ID: sw, Released: true}) {
			e.Post(spin.ShotEvent{ID: shot})
		}
	}
}

package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

func programScript(e *spin.ScriptEnv) {
	for {
		e.Do(spin.PlayScript{ID: ScriptGame})
		if _, done := e.WaitFor(spin.GameOverEvent{}); done {
			return
		}
		e.Do(spin.PlayScript{ID: ScriptAttractMode})
		evt, done := e.WaitFor(
			spin.SwitchEvent{ID: jd.SwitchStartButton},
			spin.SwitchEvent{ID: jd.SwitchSuperGameButton},
		)
		if done {
			return
		}
		if evt == (spin.SwitchEvent{ID: jd.SwitchSuperGameButton}) {
			break
		}
	}
}

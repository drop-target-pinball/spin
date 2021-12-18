package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

func programScript(e spin.Env) {
	for {
		e.Do(spin.PlayScript{ID: ScriptAttractMode})
		_, done := e.WaitFor(
			spin.SwitchEvent{ID: jd.SwitchStartButton},
			spin.SwitchEvent{ID: jd.SwitchSuperGameButton},
		)
		if done {
			return
		}
		e.Do(spin.StopScope{ID: spin.ScopeMode})
		e.Do(spin.PlayScript{ID: ScriptGame})
		if _, done := e.WaitFor(spin.GameOverEvent{}); done {
			return
		}
	}
}

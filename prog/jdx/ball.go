package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

func ballScript(e spin.Env) {
	e.Do(spin.FlippersOn{})
	e.Do(spin.AutoPulseOn{ID: jd.AutoSlingLeft})
	e.Do(spin.AutoPulseOn{ID: jd.AutoSlingRight})

	e.Do(spin.PlayScript{ID: ScriptSling})
	e.Do(spin.PlayScript{ID: ScriptLeftPopperShot})
	e.Do(spin.PlayScript{ID: ScriptLeftShooterLaneShot})
	e.Do(spin.PlayScript{ID: ScriptRightPopperShot})
	e.Do(spin.PlayScript{ID: jd.ScriptInactiveGlobe})
}

func slingScript(e spin.Env) {
	for {
		if _, done := e.WaitFor(
			spin.SwitchEvent{ID: jd.SwitchLeftSling},
			spin.SwitchEvent{ID: jd.SwitchRightSling},
		); done {
			return
		}
		e.Do(spin.AwardScore{Val: ScoreSling})
	}
}

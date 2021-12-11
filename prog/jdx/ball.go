package jdx

import (
	"time"

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
	e.Do(spin.PlayScript{ID: ScriptOutlane})
	e.Do(spin.PlayScript{ID: ScriptReturnLane})
	e.Do(spin.PlayScript{ID: ScriptRightPopperShot})
	e.Do(spin.PlayScript{ID: jd.ScriptInactiveGlobe})

	e.Do(spin.PlayScript{ID: jd.ScriptRaiseDropTargetsWhenAllDown})

	if done := e.Sleep(1000 * time.Millisecond); done {
		return
	}
	e.Do(spin.PlayScript{ID: jd.ScriptRaiseDropTargets})
}

func outlaneScript(e spin.Env) {
	for {
		if _, done := e.WaitFor(
			spin.SwitchEvent{ID: jd.SwitchLeftOutlane},
			spin.SwitchEvent{ID: jd.SwitchRightOutlane},
		); done {
			return
		}
		e.Do(spin.PlaySound{ID: SoundBallLost})
		e.Do(spin.AwardScore{Val: ScoreOutlane})
	}
}

func returnLaneScript(e spin.Env) {
	for {
		if _, done := e.WaitFor(
			spin.SwitchEvent{ID: jd.SwitchLeftReturnLane},
			spin.SwitchEvent{ID: jd.SwitchInnerRightReturnLane},
			spin.SwitchEvent{ID: jd.SwitchOuterRightReturnLane},
		); done {
			return
		}
		e.Do(spin.PlaySound{ID: SoundReturnLane})
		e.Do(spin.AwardScore{Val: ScoreReturnLane})
	}
}

func slingScript(e spin.Env) {
	for {
		if _, done := e.WaitFor(
			spin.SwitchEvent{ID: jd.SwitchLeftSling},
			spin.SwitchEvent{ID: jd.SwitchRightSling},
		); done {
			return
		}
		e.Do(spin.PlaySound{ID: SoundSling})
		e.Do(spin.AwardScore{Val: ScoreSling})
	}
}

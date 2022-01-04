package jdx

import (
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/prog/builtin"
)

func ballScript(e spin.Env) {
	startOfBallReset(e)

	e.Do(spin.FlippersOn{})
	e.Do(spin.AutoPulseOn{ID: jd.AutoSlingLeft})
	e.Do(spin.AutoPulseOn{ID: jd.AutoSlingRight})

	e.Do(spin.PlayScript{ID: builtin.ScriptBallTracker})
	e.Do(spin.PlayScript{ID: ScriptSling})
	e.Do(spin.PlayScript{ID: ScriptOutlane})
	e.Do(spin.PlayScript{ID: ScriptReturnLane})
	e.Do(spin.PlayScript{ID: jd.ScriptInactiveGlobe})
	e.Do(spin.PlayScript{ID: jd.ScriptRaiseDropTargetsWhenAllDown})
	e.Do(spin.PlayScript{ID: ScriptDebugExtraBall})
	e.Do(spin.PlayScript{ID: jd.ScriptRaiseDropTargets})
	e.Do(spin.PlayScript{ID: ScriptChain})

	ctx, cancel := e.Derive()
	defer cancel()
	e.NewCoroutine(ctx, modeScript)

	for {
		evt, done := e.WaitFor(spin.BallDrainEvent{})
		if done {
			return
		}
		e := evt.(spin.BallDrainEvent)
		if e.BallsInPlay == 0 {
			break
		}
	}
	e.Do(spin.AdvanceGame{})
}

func modeScript(e spin.Env) {
	e.Do(spin.PlayScript{ID: ScriptPlungeMode})
	if _, done := e.WaitFor(spin.ScriptFinishedEvent{ID: ScriptPlungeMode}); done {
		return
	}
	e.Do(spin.StopScope{ID: spin.ScopeMode})
	e.Do(spin.PlayScript{ID: ScriptBasicMode})

	for {
		if _, done := e.WaitFor(spin.Done{}); done {
			return
		}
	}
}

func debugExtraBallScript(e spin.Env) {
	e.Do(spin.DriverOn{ID: jd.LampBuyInButton})
	for {
		if _, done := e.WaitFor(spin.SwitchEvent{ID: jd.SwitchBuyInButton}); done {
			return
		}
		e.Do(spin.DriverPulse{ID: jd.CoilTrough})
		if done := e.Sleep(1 * time.Second); done {
			return
		}
		e.Do(spin.DriverPulse{ID: jd.CoilRightShooterLane})
	}
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
		e.Do(spin.AwardScore{Val: ScoreOutlane * Multiplier(e)})
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
		e.Do(spin.AwardScore{Val: ScoreReturnLane * Multiplier(e)})
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
		e.Do(spin.AwardScore{Val: ScoreSling * Multiplier(e)})
	}
}

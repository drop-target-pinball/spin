package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

func ballScript(e *spin.ScriptEnv) {
	vars := GetVars(e)

	vars.BadImpersonatorBonus = 0
	vars.ManhuntBonus = 0
	vars.MeltdownBonus = 0
	vars.PursuitBonus = 0
	vars.SafecrackerBonus = 0
	vars.SniperBonus = 0
	vars.SniperScore = 0
	vars.StakeoutBonus = 0
	vars.TankBonus = 0

	if vars.SelectedMode == 0 {
		vars.SelectedMode = ModePursuit // FIXME
	}

	startBase(e)
	e.Do(spin.PlayScript{ID: ScriptChain})
	e.Do(spin.PlayScript{ID: ScriptPlungeMode})

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

	stopBase(e)
	e.Do(spin.StopScriptGroup{ID: spin.ScriptGroupMode})
	e.Do(spin.StopScriptGroup{ID: spin.ScriptGroupBall})
	e.Do(spin.AdvanceGame{})
}

func startBase(e *spin.ScriptEnv) {
	e.Do(spin.PlayScript{ID: jd.ScriptGIOn})
	e.Do(spin.FlippersOn{})
	e.Do(spin.AutoPulseOn{ID: jd.AutoSlingLeft})
	e.Do(spin.AutoPulseOn{ID: jd.AutoSlingRight})
	e.Do(spin.PlayScript{ID: jd.ScriptRaiseDropTargets})

	e.Do(spin.PlayScript{ID: jd.ScriptInactiveGlobe})
	e.Do(spin.PlayScript{ID: jd.ScriptRaiseDropTargetsWhenAllDown})
	e.NewCoroutine(defaultSlingLoop)
	e.NewCoroutine(defaultOutlaneLoop)
	e.NewCoroutine(defaultReturnLaneLoop)
	e.NewCoroutine(defaultScoreLoop)
	e.NewCoroutine(defaultLeftShooterLaneLoop)
	e.NewCoroutine(defaultRightShooterLaneLoop)
	e.NewCoroutine(defaultLeftPopperLoop)
	e.NewCoroutine(defaultRightPopperLoop)
}

func stopBase(e *spin.ScriptEnv) {
	e.Do(spin.FlippersOff{})
	e.Do(spin.AutoPulseOff{ID: jd.AutoSlingLeft})
	e.Do(spin.AutoPulseOff{ID: jd.AutoSlingRight})

	e.Do(spin.StopScript{ID: jd.ScriptInactiveGlobe})
	e.Do(spin.StopScript{ID: jd.ScriptRaiseDropTargetsWhenAllDown})
}

func baseScript(e *spin.ScriptEnv) {
	startBase(e)
	e.Do(spin.AddBall{})
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
	stopBase(e)
	e.Do(spin.PlayScript{ID: jd.ScriptGIOff})
}

func defaultOutlaneLoop(e *spin.ScriptEnv) {
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

func defaultReturnLaneLoop(e *spin.ScriptEnv) {
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

func defaultScoreLoop(e *spin.ScriptEnv) {
	vars := GetVars(e)
	spin.RenderFrameLoop(e, func(e *spin.ScriptEnv) {
		render := vars.Mode == ModeNone || vars.Mode == ModePlunge
		if render {
			spin.ScorePanel(e)
		}
	})
}

func defaultSlingLoop(e *spin.ScriptEnv) {
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

func defaultLeftShooterLaneLoop(e *spin.ScriptEnv) {
	vars := GetVars(e)
	for {
		if done := spin.WaitForBallArrivalLoop(e, jd.SwitchLeftShooterLane, 1000); done {
			return
		}
		if vars.Mode == ModeAirRaid {
			continue
		}

		e.Do(spin.PlayScript{ID: jd.ScriptRaiseDropTargets})
		if done := e.Sleep(500); done {
			return
		}
		e.Do(spin.DriverPulse{ID: jd.CoilLeftShooterLane})
	}
}

func defaultRightShooterLaneLoop(e *spin.ScriptEnv) {
	vars := GetVars(e)
	for {
		if done := spin.WaitForBallArrivalLoop(e, jd.SwitchRightShooterLane, 1000); done {
			return
		}
		if vars.Mode == ModePlunge {
			continue
		}
		e.Do(spin.DriverPulse{ID: jd.CoilRightShooterLane})
	}
}

func defaultLeftPopperLoop(e *spin.ScriptEnv) {
	for {
		if done := spin.WaitForBallArrivalLoop(e, jd.SwitchLeftPopper, 500); done {
			return
		}
		for i := 0; i < 3; i++ {
			e.Do(spin.DriverPulse{ID: jd.FlasherSubwayExit})
			if done := e.Sleep(200); done {
				return
			}
		}
		e.Do(spin.DriverPulse{ID: jd.CoilLeftPopper})
	}
}

func defaultRightPopperLoop(e *spin.ScriptEnv) {
	vars := GetVars(e)
	for {
		if done := spin.WaitForBallArrivalLoop(e, jd.SwitchRightPopper, 500); done {
			return
		}
		if vars.Mode == ModeSniper {
			continue
		}
		e.Do(spin.DriverPulse{ID: jd.CoilRightPopper})
	}
}

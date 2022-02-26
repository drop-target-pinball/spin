package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/proc"
)

func ballScript(e *spin.ScriptEnv) {
	vars := GetVars(e)

	vars.BadImpersonatorBonus = 0
	vars.ManhuntBonus = 0
	vars.MeltdownBonus = 0
	vars.Multiplier = 1
	vars.PursuitBonus = 0
	vars.SafecrackerBonus = 0
	vars.SniperBonus = 0
	vars.SniperScore = 0
	vars.StakeoutBonus = 0
	vars.TankBonus = 0
	vars.LeftRampsMade = 0
	vars.RightRampsMade = 0
	vars.TopLeftRampsMade = 0

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
	e.Do(spin.PlayScript{ID: ScriptLightBallLock})
	e.Do(spin.PlayScript{ID: ScriptBallLock})
	e.Do(spin.PlayScript{ID: ScriptBallSaver})
	e.Do(spin.PlayScript{ID: ScriptCrimeScenes})
	e.NewCoroutine(defaultSlingLoop)
	e.NewCoroutine(defaultOutlaneLoop)
	e.NewCoroutine(defaultReturnLaneLoop)
	e.NewCoroutine(defaultScoreLoop)
	e.NewCoroutine(defaultLeftShooterLaneLoop)
	e.NewCoroutine(defaultRightShooterLaneLoop)
	e.NewCoroutine(defaultLeftPopperLoop)
	e.NewCoroutine(defaultRightPopperLoop)
	e.NewCoroutine(defaultPostLoop)
	e.NewCoroutine(defaultMysteryLoop)
	e.NewCoroutine(defaultLeftRampLoop)
	e.NewCoroutine(defaultRightRampLoop)
	e.NewCoroutine(defaultTopLeftRampLoop)
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
	r := e.Display("").Open(0)
	defer r.Close()

	vars := GetVars(e)
	spin.RenderFrameLoop(e, func(e *spin.ScriptEnv) {
		render := vars.Mode == ModeNone || vars.Mode == ModePlunge
		if render {
			spin.ScorePanel(e, r)
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
		if done := spin.WaitForBallArrivalLoop(e, jd.SwitchLeftShooterLane, 500); done {
			return
		}
		if vars.Mode == ModeAirRaid {
			continue
		}

		e.Do(spin.PlayScript{ID: jd.ScriptRaiseDropTargets})
		if done := e.Sleep(250); done {
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

func defaultPostLoop(e *spin.ScriptEnv) {
	sounds := map[spin.SwitchEvent]string{
		{ID: jd.SwitchLeftPost}:  SoundLeftPost,
		{ID: jd.SwitchRightPost}: SoundRightPost,
	}

	for {
		evt, done := e.WaitFor(
			spin.SwitchEvent{ID: jd.SwitchLeftPost},
			spin.SwitchEvent{ID: jd.SwitchRightPost})
		if done {
			return
		}
		sound := sounds[evt.(spin.SwitchEvent)]
		e.Do(spin.AwardScore{Val: ScorePost})
		e.Do(spin.PlaySound{ID: sound})
	}
}

func defaultMysteryLoop(e *spin.ScriptEnv) {
	for {
		if _, done := e.WaitFor(spin.SwitchEvent{ID: jd.SwitchMysteryTarget}); done {
			return
		}
		e.Do(spin.AwardScore{Val: ScoreMystery})
		e.Do(spin.PlaySound{ID: SoundMystery})
	}
}

func defaultLeftRampLoop(e *spin.ScriptEnv) {
	vars := GetVars(e)
	for {
		if _, done := e.WaitFor(spin.SwitchEvent{ID: jd.SwitchLeftRampExit}); done {
			return
		}
		vars.LeftRampsMade += 1
		score := vars.LeftRampsMade * ScoreLeftRampN
		if score > MaxRampScore {
			score = MaxRampScore
		}
		e.Do(spin.AwardScore{Val: score})

		if vars.Mode == ModeNone && vars.StartModeLeft {
			continue
		}
		if vars.LocksReady > vars.BallsLocked {
			continue
		}
		e.Do(spin.PlayScript{ID: ScriptLeftRampAward})
	}
}

func defaultRightRampLoop(e *spin.ScriptEnv) {
	vars := GetVars(e)
	for {
		if _, done := e.WaitFor(spin.SwitchEvent{ID: jd.SwitchRightRampExit}); done {
			return
		}
		vars.RightRampsMade += 1
		score := vars.RightRampsMade * ScoreRightRampN
		if score > MaxRampScore {
			score = MaxRampScore
		}
		e.Do(spin.AwardScore{Val: score})
		e.Do(spin.PlayScript{ID: ScriptRightRampAward})
	}
}

func defaultTopLeftRampLoop(e *spin.ScriptEnv) {
	vars := GetVars(e)
	for {
		if _, done := e.WaitFor(spin.SwitchEvent{ID: jd.SwitchTopLeftRampExit}); done {
			return
		}
		vars.TopLeftRampsMade += 1
		score := vars.TopLeftRampsMade * ScoreTopLeftRampN
		if score > MaxRampScore {
			score = MaxRampScore
		}
		e.Do(spin.AwardScore{Val: score})
		e.Do(spin.PlayScript{ID: ScriptTopLeftRampAward})
	}
}

func ballSaverScript(e *spin.ScriptEnv) {
	vars := spin.GetGameVars(e)
	vars.BallSave = true
	defer func() { vars.BallSave = false }()

	e.Do(spin.DriverOn{ID: jd.LampDrainShield})
	if _, done := e.WaitFor(jd.PlayfieldSwitches...); done {
		return
	}

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.Do(proc.DriverSchedule{ID: jd.LampDrainShield, Schedule: proc.BlinkSchedule, Now: true})
		s.Sleep(5_000)
		s.Do(proc.DriverSchedule{ID: jd.LampDrainShield, Schedule: proc.HurryUpBlinkSchedule, Now: true})
		s.Sleep(2_000)
		s.Do(spin.DriverOff{ID: jd.LampDrainShield})
		s.Do(spin.PlaySpeech{ID: SpeechDrainShieldDeactivated})
		s.WaitFor(spin.SpeechFinishedEvent{})

		s.Run()
	})

	evt, done := e.WaitForUntil(10_000, spin.BallDrainEvent{}, spin.BallWillDrainEvent{})
	if done {
		return
	}
	e.Do(spin.DriverOff{ID: jd.LampDrainShield})
	if evt == nil {
		return
	}
	e.Do(spin.PlaySpeech{ID: SpeechDontMove})
}

func leftRampAwardScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(spin.PriorityAnnounce)
	defer r.Close()

	vars := GetVars(e)
	score := vars.LeftRampsMade * ScoreLeftRampN
	if score > MaxRampScore {
		score = MaxRampScore
	}

	ScoreAndLabelPanel(e, r, score, "RAMP AWARD")

	s := spin.NewSequencer(e)

	s.Do(spin.PlaySound{ID: SoundMotorcycleRamp})
	s.Sleep(2_000)

	s.Run()
}

func rightRampAwardScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(spin.PriorityAnnounce)
	defer r.Close()

	vars := GetVars(e)
	score := vars.RightRampsMade * ScoreRightRampN
	if score > MaxRampScore {
		score = MaxRampScore
	}

	ScoreAndLabelPanel(e, r, score, "RAMP AWARD")

	s := spin.NewSequencer(e)

	s.Do(spin.PlaySound{ID: SoundRightRamp})
	s.Sleep(2_000)

	s.Run()
}

func topLeftRampAwardScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(spin.PriorityAnnounce)
	defer r.Close()

	vars := GetVars(e)
	score := vars.TopLeftRampsMade * ScoreTopLeftRampN
	if score > MaxRampScore {
		score = MaxRampScore
	}

	ScoreAndLabelPanel(e, r, score, "RAMP AWARD")

	s := spin.NewSequencer(e)

	s.Do(spin.PlaySound{ID: SoundTopLeftRamp})
	s.Sleep(2_000)

	s.Run()
}

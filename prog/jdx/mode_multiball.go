package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/proc"
)

func multiballScript(e *spin.ScriptEnv) {
	switches := spin.GetResourceVars(e).Switches

	e.Do(spin.StopScriptGroup{ID: spin.ScriptGroupMode})
	e.Do(spin.StopScript{ID: ScriptLightBallLock})
	e.Do(spin.StopScript{ID: ScriptBallLock})
	e.Do(spin.StopScript{ID: ScriptChain})

	defer func() {
		e.Do(spin.StopScriptGroup{ID: spin.ScriptGroupMode})
		e.Do(spin.DriverOff{ID: jd.FlasherJudgeDeath})
		e.Do(spin.DriverOff{ID: jd.FlasherJudgeFear})
		e.Do(spin.DriverOff{ID: jd.FlasherJudgeFire})
		e.Do(spin.DriverOff{ID: jd.FlasherJudgeMortis})
	}()

	e.Do(spin.AddBall{N: 3})

	if switches[jd.SwitchRightShooterLane].Active {
		e.Do(spin.DriverPulse{ID: jd.CoilRightShooterLane})
	}

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.Do(spin.PlayMusic{ID: MusicMultiballIntro, Notify: true, Loops: 1})
		s.WaitFor(spin.MusicFinishedEvent{})
		s.Do(spin.PlayMusic{ID: MusicMultiballLoop})
		s.Run()
	})

	e.Do(spin.PlayScript{ID: ScriptMultiballLightJackpot})

	for {
		evt, done := e.WaitFor(spin.BallDrainEvent{})
		if done {
			return
		}
		if evt.(spin.BallDrainEvent).BallsInPlay == 1 {
			break
		}
	}

	e.Do(spin.PlayScript{ID: ScriptChain})
	e.Do(spin.PlayScript{ID: ScriptLightBallLock})
	e.Do(spin.PlayScript{ID: ScriptBallLock})
	e.Do(spin.PlayMusic{ID: MusicMain})
}

func multiballLightJackpotScript(e *spin.ScriptEnv) {
	vars := GetVars(e)
	vars.MultiballShotsMade = 0
	shotsNeeded := MultiballShotsToLightJackpot[vars.DarkJudgeSelected]

	e.Do(spin.PlayScript{ID: ScriptLeftRampRunway})
	e.Do(spin.PlayScript{ID: ScriptRightRampRunway})
	e.Do(spin.PlayScript{ID: ScriptRightPopperRunway})

	defer func() {
		e.Do(spin.StopScript{ID: ScriptLeftRampRunway})
		e.Do(spin.StopScript{ID: ScriptRightRampRunway})
		e.Do(spin.StopScript{ID: ScriptRightPopperRunway})
	}()

	e.Do(proc.DriverSchedule{ID: jd.DarkJudgeFlashers[vars.DarkJudgeSelected], Schedule: 0x80})

	for vars.MultiballShotsMade < shotsNeeded {
		_, done := e.WaitFor(
			spin.SwitchEvent{ID: jd.SwitchLeftRampExit},
			spin.SwitchEvent{ID: jd.SwitchRightRampExit},
			spin.SwitchEvent{ID: jd.SwitchRightPopper},
		)
		if done {
			return
		}
		vars.MultiballShotsMade++
		if vars.MultiballShotsMade == shotsNeeded {
			break
		}
		e.Do(spin.PlaySound{ID: SoundAnnounce})
	}
	e.Do(spin.PlayScript{ID: ScriptMultiballJackpot})
}

func multiballJackpotScript(e *spin.ScriptEnv) {
	vars := GetVars(e)

	e.Do(spin.PlaySpeech{ID: SpeechJackpotIsLit})
	e.Do(spin.PlayScript{ID: ScriptJackpotRunway})
	defer e.Do(spin.StopScript{ID: ScriptJackpotRunway})

	e.NewCoroutine(watchCenterDropTargetLoop) // FIXME: This should be shared somewhere

	if _, done := e.WaitFor(
		spin.SwitchEvent{ID: jd.SwitchSubwayEnter1},
		spin.SwitchEvent{ID: jd.SwitchSubwayEnter2}); done {
		return
	}
	vars.DarkJudgeSelected += 1
	vars.MultiballShotsMade = 0
	e.Do(spin.PlayScript{ID: ScriptDarkJudgeContained})
	e.Do(spin.PlayScript{ID: ScriptMultiballLightJackpot})
}

func darkJudgeContainedScript(e *spin.ScriptEnv) {
	vars := GetVars(e)

	speechJudges := []string{
		SpeechJudgeMortis,
		SpeechJudgeFire,
		SpeechJudgeFear,
	}

	contained := vars.DarkJudgeSelected - 1
	if contained < 0 {
		contained = 0
	}
	if contained > 2 {
		contained = 2
	}
	e.Do(spin.DriverOff{ID: jd.DarkJudgeFlashers[contained]})

	s := spin.NewSequencer(e)

	s.Do(spin.PlaySpeech{ID: speechJudges[contained], Notify: true})
	s.WaitFor(spin.SpeechFinishedEvent{})
	s.Do(spin.PlaySpeech{ID: SpeechContained})
	s.Run()
}

func leftRampRunwayScript(e *spin.ScriptEnv) {
	e.Do(proc.DriverSchedule{ID: jd.LampLeftModeStart, Schedule: 0xff})
	e.Do(proc.DriverSchedule{ID: jd.LampLock1, Schedule: (0xff << 2)})
	e.Do(proc.DriverSchedule{ID: jd.LampLock2, Schedule: (0xff << 4)})
	e.Do(proc.DriverSchedule{ID: jd.LampLock3, Schedule: (0xff << 6)})
	e.Do(proc.DriverSchedule{ID: jd.FlasherLeftPursuit, Schedule: 0x8000})

	e.WaitFor(spin.Done{})

	e.Do(spin.DriverOff{ID: jd.LampLeftModeStart})
	e.Do(spin.DriverOff{ID: jd.LampLock1})
	e.Do(spin.DriverOff{ID: jd.LampLock2})
	e.Do(spin.DriverOff{ID: jd.LampLock3})
	e.Do(spin.DriverOff{ID: jd.FlasherLeftPursuit})
}

func rightRampRunwayScript(e *spin.ScriptEnv) {
	e.Do(proc.DriverSchedule{ID: jd.LampRightRampCrimeSceneGreen, Schedule: 0xff})
	e.Do(proc.DriverSchedule{ID: jd.LampRightRampCrimeSceneYellow, Schedule: 0xff << 2})
	e.Do(proc.DriverSchedule{ID: jd.LampRightRampCrimeSceneRed, Schedule: 0xff << 4})
	e.Do(proc.DriverSchedule{ID: jd.LampRightRampCrimeSceneWhite, Schedule: 0xff << 6})
	e.Do(proc.DriverSchedule{ID: jd.FlasherRightPursuit, Schedule: 0x8000})

	e.WaitFor(spin.Done{})

	e.Do(spin.DriverOff{ID: jd.LampRightRampCrimeSceneGreen})
	e.Do(spin.DriverOff{ID: jd.LampRightRampCrimeSceneYellow})
	e.Do(spin.DriverOff{ID: jd.LampRightRampCrimeSceneRed})
	e.Do(spin.DriverOff{ID: jd.LampRightRampCrimeSceneWhite})
	e.Do(spin.DriverOff{ID: jd.FlasherRightPursuit})
}

func rightPopperRunwayScript(e *spin.ScriptEnv) {
	e.Do(proc.DriverSchedule{ID: jd.LampRightModeStart, Schedule: 0xff})
	e.Do(proc.DriverSchedule{ID: jd.LampRightPopperCrimeSceneGreen, Schedule: 0xff << 2})
	e.Do(proc.DriverSchedule{ID: jd.LampRightPopperCrimeSceneYellow, Schedule: 0xff << 4})
	e.Do(proc.DriverSchedule{ID: jd.LampRightPopperCrimeSceneRed, Schedule: 0xff << 6})
	e.Do(proc.DriverSchedule{ID: jd.LampRightPopperCrimeSceneWhite, Schedule: 0x8000})

	e.WaitFor(spin.Done{})

	e.Do(spin.DriverOff{ID: jd.LampRightModeStart})
	e.Do(spin.DriverOff{ID: jd.LampRightPopperCrimeSceneGreen})
	e.Do(spin.DriverOff{ID: jd.LampRightPopperCrimeSceneYellow})
	e.Do(spin.DriverOff{ID: jd.LampRightPopperCrimeSceneRed})
	e.Do(spin.DriverOff{ID: jd.LampRightPopperCrimeSceneWhite})
}

func jackpotRunwayScript(e *spin.ScriptEnv) {
	e.Do(proc.DriverSchedule{ID: jd.LampAwardBadImpersonator, Schedule: 0xff})
	e.Do(proc.DriverSchedule{ID: jd.LampAwardSafeCracker, Schedule: 0xff})
	e.Do(proc.DriverSchedule{ID: jd.LampDropTargetJ, Schedule: 0xff << 2})
	e.Do(proc.DriverSchedule{ID: jd.LampDropTargetE, Schedule: 0xff << 2})
	e.Do(proc.DriverSchedule{ID: jd.LampDropTargetU, Schedule: 0xff << 4})
	e.Do(proc.DriverSchedule{ID: jd.LampDropTargetG, Schedule: 0xff << 4})
	e.Do(proc.DriverSchedule{ID: jd.LampDropTargetD, Schedule: 0xff << 6})
	e.Do(proc.DriverSchedule{ID: jd.LampMultiballJackpot, Schedule: 0xff << 8})

	e.WaitFor(spin.Done{})

	e.Do(spin.DriverOff{ID: jd.LampDropTargetJ})
	e.Do(spin.DriverOff{ID: jd.LampDropTargetE})
	e.Do(spin.DriverOff{ID: jd.LampDropTargetU})
	e.Do(spin.DriverOff{ID: jd.LampDropTargetG})
	e.Do(spin.DriverOff{ID: jd.LampDropTargetD})
	e.Do(spin.DriverOff{ID: jd.LampAwardBadImpersonator})
	e.Do(spin.DriverOff{ID: jd.LampAwardSafeCracker})
	e.Do(spin.DriverOff{ID: jd.LampMultiballJackpot})
}

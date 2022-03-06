package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/proc"
)

var speechJudges = []string{
	"",
	SpeechJudgeMortis,
	SpeechJudgeFire,
	SpeechJudgeFear,
}

const darkJudgeSchedule = 0x80

func multiballScript(e *spin.ScriptEnv) {
	vars := GetVars(e)
	switches := spin.GetResourceVars(e).Switches

	e.Do(spin.StopScriptGroup{ID: spin.ScriptGroupMode})
	e.Do(spin.StopScript{ID: ScriptLightBallLock})
	e.Do(spin.StopScript{ID: ScriptBallLock})
	e.Do(spin.StopScript{ID: ScriptChain})
	e.Do(spin.StopScript{ID: ScriptCrimeScenes})

	defer func() {
		e.Do(spin.StopScriptGroup{ID: spin.ScriptGroupMode})
		e.Do(spin.DriverOff{ID: jd.FlasherJudgeDeath})
		e.Do(spin.DriverOff{ID: jd.FlasherJudgeFear})
		e.Do(spin.DriverOff{ID: jd.FlasherJudgeFire})
		e.Do(spin.DriverOff{ID: jd.FlasherJudgeMortis})
	}()

	e.Do(spin.AddBall{N: 3})

	if vars.DarkJudgeSelected == jd.DarkJudgeNone {
		vars.DarkJudgeSelected = jd.DarkJudgeMortis
	}
	vars.MultiballShotsLeft = MultiballShotsToLightJackpot[vars.DarkJudgeSelected]

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

	if vars.DarkJudgeSelected == jd.DarkJudgeDeath {
		e.Do(spin.PlayScript{ID: ScriptJudgeDeath})
	} else {
		e.Do(spin.PlayScript{ID: ScriptMultiballAnnounce})
	}
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
	e.Do(spin.PlayScript{ID: ScriptCrimeScenes})

	e.Do(spin.PlayMusic{ID: MusicMain})
}

func multiballFrame(e *spin.ScriptEnv, r spin.Renderer) {
	g := r.Graphics()

	g.Y = 5
	g.Font = spin.FontPfRondaSevenBold16
	r.Print(g, "MULTIBALL")
}

func shotsToGoFrame(e *spin.ScriptEnv, r spin.Renderer) {
	g := r.Graphics()
	vars := GetVars(e)

	r.Fill(spin.ColorOff)

	g.Font = spin.FontPfRondaSevenBold16
	g.X = 16
	g.Y = 5
	g.AnchorX = spin.AnchorLeft
	r.Print(g, "%v", vars.MultiballShotsLeft)

	g.Font = spin.FontPfRondaSeven8
	g.X = (r.Width() / 2) + 8
	g.Y = 8
	g.AnchorX = spin.AnchorCenter
	shots := "SHOTS"
	if vars.MultiballShotsLeft == 1 {
		shots = "SHOT"
	}
	r.Print(g, "%v TO LIGHT", shots)

	g.Y = 18
	g.Font = spin.FontPfRondaSevenBold8
	r.Print(g, "JACKPOT")
}

func jackpotIsLitFrame(e *spin.ScriptEnv, r spin.Renderer) {
	g := r.Graphics()

	g.Font = spin.FontPfRondaSevenBold8
	g.Y = 8
	r.Print(g, "JACKPOT")

	g.Y = 18
	r.Print(g, "IS LIT")
}

func multiballAnnounceScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(0)
	defer r.Close()
	vars := GetVars(e)

	judgeCallout := speechJudges[vars.DarkJudgeSelected]

	s := spin.NewSequencer(e)

	s.DoFunc(func() { multiballFrame(e, r) })
	s.Sleep(500)
	s.Do(spin.PlaySpeech{ID: SpeechPrepareToFaceTheImmortal, Notify: true, Duck: 0.5})
	s.Sleep(1_500)
	s.DoFunc(func() { shotsToGoFrame(e, r) })
	s.Sleep(2_000)
	s.Do(spin.PlaySpeech{ID: judgeCallout, Notify: true, Duck: 0.5})
	s.Run()
}

func multiballShotsToGoScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(spin.PriorityAnnounce)
	defer r.Close()

	s := spin.NewSequencer(e)

	s.DoFunc(func() { shotsToGoFrame(e, r) })
	s.Sleep(2_000)
	s.Run()
}

func multiballJackpotIsLitScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(spin.PriorityAnnounce)
	defer r.Close()

	s := spin.NewSequencer(e)

	s.Do(spin.PlaySpeech{ID: SpeechJackpotIsLit, Notify: true, Duck: 0.5, Priority: spin.PriorityAudioModeCallout})
	s.DoFunc(func() { jackpotIsLitFrame(e, r) })
	s.Sleep(2_000)
	s.Run()
}

func multiballLightJackpotScript(e *spin.ScriptEnv) {
	vars := GetVars(e)
	vars.MultiballShotsLeft = MultiballShotsToLightJackpot[vars.DarkJudgeSelected]

	e.Do(spin.PlayScript{ID: ScriptLeftRampRunway})
	e.Do(spin.PlayScript{ID: ScriptRightRampRunway})
	e.Do(spin.PlayScript{ID: ScriptRightPopperRunway})

	defer func() {
		e.Do(spin.StopScript{ID: ScriptLeftRampRunway})
		e.Do(spin.StopScript{ID: ScriptRightRampRunway})
		e.Do(spin.StopScript{ID: ScriptRightPopperRunway})
	}()

	e.Do(proc.DriverSchedule{ID: jd.DarkJudgeFlashers[vars.DarkJudgeSelected], Schedule: darkJudgeSchedule})

	for vars.MultiballShotsLeft > 0 {
		_, done := e.WaitFor(
			spin.SwitchEvent{ID: jd.SwitchLeftRampExit},
			spin.SwitchEvent{ID: jd.SwitchRightRampExit},
			spin.SwitchEvent{ID: jd.SwitchRightPopper},
		)
		if done {
			return
		}
		vars.MultiballShotsLeft--
		if vars.MultiballShotsLeft > 0 {
			e.Do(spin.PlaySound{ID: SoundAnnounce})
			e.Do(spin.PlayScript{ID: ScriptMultiballShotsToGo})
		}
	}
	e.Do(spin.PlayScript{ID: ScriptMultiballJackpot})
}

func multiballJackpotScript(e *spin.ScriptEnv) {
	e.Do(spin.PlayScript{ID: ScriptMultiballJackpotIsLit})
	e.Do(spin.PlayScript{ID: ScriptJackpotRunway})
	defer e.Do(spin.StopScript{ID: ScriptJackpotRunway})

	e.NewCoroutine(watchCenterDropTargetRoutine) // FIXME: This should be shared somewhere

	if _, done := e.WaitFor(
		spin.SwitchEvent{ID: jd.SwitchSubwayEnter1},
		spin.SwitchEvent{ID: jd.SwitchSubwayEnter2}); done {
		return
	}
	e.Do(spin.PlayScript{ID: ScriptDarkJudgeContained})
}

func darkJudgeContainedScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(spin.PriorityAnnounce)
	defer r.Close()

	vars := GetVars(e)

	if vars.DarkJudgeSelected == jd.DarkJudgeNone {
		vars.DarkJudgeSelected = jd.DarkJudgeMortis
	}
	contained := vars.DarkJudgeSelected

	if vars.DarkJudgeSelected == jd.DarkJudgeDeath {
		vars.DarkJudgeSelected = jd.DarkJudgeNone
		vars.LeftPopperManual = true
		e.Do(spin.PlayScript{ID: ScriptMultiballTransition})
		return
	} else {
		vars.DarkJudgeSelected += 1
	}
	vars.MultiballShotsLeft = MultiballShotsToLightJackpot[vars.DarkJudgeSelected]

	jackpot := ScoreMultiballJackpot0 + (ScoreMultiballJackpotN * contained)
	e.Do(spin.AwardScore{Val: jackpot})
	ScoreAndLabelPanel(e, r, jackpot, "JACKPOT")

	e.Do(spin.DriverOff{ID: jd.DarkJudgeFlashers[contained]})
	e.Do(spin.PlayScript{ID: ScriptMultiballLightJackpot})

	s := spin.NewSequencer(e)

	s.Do(spin.PlaySound{ID: SoundMultiballJackpot, Notify: true, Duck: 0.25})
	s.Do(spin.PlaySpeech{ID: speechJudges[contained], Notify: true})
	s.WaitFor(spin.SpeechFinishedEvent{})
	s.Do(spin.PlaySpeech{ID: SpeechContained})
	s.Sleep(3_000)
	s.Do(spin.PlayScript{ID: ScriptMultiballShotsToGo})
	done := s.Run()
	if done {
		return
	}
	r.Close()

	if vars.DarkJudgeSelected == jd.DarkJudgeDeath {
		s := spin.NewSequencer(e)
		s.Sleep(3_000)
		e.Do(spin.PlayScript{ID: ScriptJudgeDeath})
		s.Run()
	} else if vars.DarkJudgeSelected != jd.DarkJudgeNone {
		e.Do(spin.PlaySpeech{ID: speechJudges[vars.DarkJudgeSelected]})
	}
}

func judgeDeathScript(e *spin.ScriptEnv) {
	s := spin.NewSequencer(e)

	s.Do(spin.PlayScript{ID: ScriptMultiballShotsToGo})

	s.Do(spin.PlaySpeech{ID: SpeechMyNameIsDeath, Notify: true})
	s.WaitFor(spin.SpeechFinishedEvent{})
	s.Sleep(500)
	s.Do(spin.PlaySpeech{ID: SpeechDeath})
	s.Sleep(500)
	s.Do(spin.PlaySpeech{ID: SpeechGreetingsJudgeAnderson, Notify: true})
	s.WaitFor(spin.SpeechFinishedEvent{})
	s.Sleep(500)
	s.Do(spin.PlaySpeech{ID: SpeechYouAreDoomed})

	s.Run()
}

func multiballTransitionScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(0)
	defer r.Close()

	jackpot := ScoreMultiballJackpot0 + (ScoreMultiballJackpotN * jd.DarkJudgeDeath)
	e.Do(spin.AwardScore{Val: jackpot})
	e.Do(spin.StopScript{ID: ScriptPlungeMode})
	ScoreAndLabelPanel(e, r, jackpot, "JACKPOT")

	s := spin.NewSequencer(e)

	s.Do(spin.StopScriptGroup{ID: ScriptGroupNoMultiball})
	s.Do(spin.StopAudio{})
	s.Do(spin.FlippersOff{})
	s.Do(spin.PlayScript{ID: jd.ScriptGIOff})
	s.Do(proc.DriverSchedule{ID: jd.FlasherJudgeDeath, Schedule: darkJudgeSchedule})

	s.Do(spin.PlaySound{ID: SoundMultiballJackpot, Notify: true, Duck: 0.25})
	s.WaitFor(spin.SoundFinishedEvent{ID: SoundMultiballJackpot})
	s.Do(spin.PlaySpeech{ID: SpeechYouCannotContainMe, Notify: true})
	s.WaitFor(spin.SpeechFinishedEvent{})
	s.Do(spin.PlaySpeech{ID: SpeechIThoughtIStoppedYou, Notify: true})
	s.WaitFor(spin.SpeechFinishedEvent{})
	s.Do(spin.PlaySpeech{ID: SpeechFoolDeath, Notify: true})
	s.WaitFor(spin.SpeechFinishedEvent{})
	s.Do(spin.PlaySpeech{ID: SpeechYouCannotKillWhatDoesNotLive, Notify: true})
	s.WaitFor(spin.SpeechFinishedEvent{})
	s.DoFunc(func() { r.Fill(spin.ColorOff) })
	s.Sleep(1_000)
	s.Do(spin.PlayScript{ID: ScriptDeadworldMode})

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
	e.Do(proc.DriverSchedule{ID: jd.LampAwardSniper, Schedule: 0xff})
	e.Do(proc.DriverSchedule{ID: jd.LampRightPopperCrimeSceneGreen, Schedule: 0xff << 2})
	e.Do(proc.DriverSchedule{ID: jd.LampRightPopperCrimeSceneYellow, Schedule: 0xff << 4})
	e.Do(proc.DriverSchedule{ID: jd.LampRightPopperCrimeSceneRed, Schedule: 0xff << 6})
	e.Do(proc.DriverSchedule{ID: jd.LampRightPopperCrimeSceneWhite, Schedule: 0x8000})

	e.WaitFor(spin.Done{})

	e.Do(spin.DriverOff{ID: jd.LampAwardSniper})
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

func multiballJudgeDeathScript(e *spin.ScriptEnv) {
	vars := GetVars(e)
	vars.DarkJudgeSelected = jd.DarkJudgeDeath
	e.Do(spin.PlayScript{ID: ScriptMultiball})
}

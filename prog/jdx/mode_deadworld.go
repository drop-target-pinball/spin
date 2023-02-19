package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/proc"
)

func deadworldModeScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(0)
	defer r.Close()

	game := spin.GetGameVars(e)
	switches := spin.GetResourceVars(e).Switches
	vars := GetVars(e)
	player := spin.GetPlayerVars(e)

	vars.StartScore = player.Score
	vars.ShotsToLowerBarriers = 40

	e.Do(spin.StopScriptGroup{ID: ScriptGroupNoMultiball})
	e.Do(spin.PlayScript{ID: ScriptDeadworldIntro})
	if _, done := e.WaitFor(spin.ScriptFinishedEvent{ID: ScriptDeadworldIntro}); done {
		return
	}

	defer func() {
		e.Do(spin.StopScript{ID: ScriptJackpotRunway})
		e.Do(spin.DriverOff{ID: jd.LampAdvanceCrimeLevel})
		e.Do(spin.DriverOff{ID: jd.LampMystery})
	}()

	e.Do(spin.FlippersOn{})
	e.Do(spin.AddBall{N: e.Config.NumBalls - game.BallsInPlay})
	e.NewCoroutine(deadwordBallSaverRoutine)
	e.NewCoroutine(func(e *spin.ScriptEnv) {
		spin.RenderFrameLoop(e, func(e *spin.ScriptEnv) {
			TimerAndScorePanel(e, r, "DEADWORLD", vars.ShotsToLowerBarriers, player.Score, "SHOOT ANY TARGET")
		})
	})

	if switches[jd.SwitchLeftPopper].Active {
		e.Do(spin.PlayScript{ID: ScriptLeftPopperEject})
	}
	if switches[jd.SwitchRightShooterLane].Active {
		e.Do(spin.DriverPulse{ID: jd.CoilRightShooterLane})
	}
	vars.LeftPopperManual = false
	//e.NewCoroutine(watchCenterDropTargetRoutine) // FIXME: move to central location

	var lost, next bool

	for !lost && !next {
		evt, done := e.WaitFor(
			spin.BallDrainEvent{},
			spin.SwitchEvent{ID: jd.SwitchSubwayEnter1},
			spin.SwitchEvent{ID: jd.SwitchSubwayEnter2},
			spin.SwitchEvent{ID: jd.SwitchBankTargets},
			spin.SwitchEvent{ID: jd.SwitchMysteryTarget},
			spin.SwitchEvent{ID: jd.SwitchLeftPost},
			spin.SwitchEvent{ID: jd.SwitchRightPost},
			spin.SwitchEvent{ID: jd.SwitchDropTargetJ},
			spin.SwitchEvent{ID: jd.SwitchDropTargetU},
			spin.SwitchEvent{ID: jd.SwitchDropTargetD},
			spin.SwitchEvent{ID: jd.SwitchDropTargetG},
			spin.SwitchEvent{ID: jd.SwitchDropTargetE},
		)
		if done {
			return
		}
		switch evt.(type) {
		case spin.BallDrainEvent:
			if game.BallsInPlay <= 1 && !game.BallSave {
				lost = true
			}
		case spin.SwitchEvent:
			vars.ShotsToLowerBarriers -= 1
			if vars.ShotsToLowerBarriers == 0 {
				next = true
				break
			}
			//e.Do(spin.PlayScript{ID: ScriptDeadworldBarrierShotsToGo})
		}
	}

	if lost {
		panic("we are done")
	}
	e.Do(spin.PlayScript{ID: ScriptDeadworldBarriersDown})
}

func deadworldIntroScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(spin.PriorityAnnounce)
	defer r.Close()

	e.NewCoroutine(deadworldMusicRoutine)

	s := spin.NewSequencer(e)

	s.DoFunc(func() { deadworldHasMaterializedFrame(e, r) })
	s.Do(proc.DriverSchedule{ID: jd.FlasherJudgeDeath, Schedule: darkJudgeSchedule})
	s.Do(spin.PlaySpeech{ID: SpeechImAnitaMann, Notify: true})
	s.WaitFor(spin.SpeechFinishedEvent{})
	s.Do(spin.PlaySpeech{ID: SpeechAndWelcomeToSuperGame, Notify: true})
	s.WaitFor(spin.SpeechFinishedEvent{})
	s.Sleep(1_000)
	s.DoFunc(func() { deadworldBarrierShotsToGoFrame(e, r) })
	s.Do(spin.PlayScript{ID: ScriptJackpotRunway})
	s.Do(spin.PlayScript{ID: jd.ScriptGIOn})
	s.Do(proc.DriverSchedule{ID: jd.LampAdvanceCrimeLevel, Schedule: proc.BlinkSchedule})
	s.Do(proc.DriverSchedule{ID: jd.LampMystery, Schedule: proc.BlinkSchedule})
	s.Do(spin.PlaySpeech{ID: SpeechDeadworldHasMaterialized, Notify: true})
	s.WaitFor(spin.SpeechFinishedEvent{})
	s.Sleep(500)
	s.Do(spin.PlaySpeech{ID: SpeechDeadworldDefensiveBarriersAtFullPower, Notify: true})
	s.WaitFor(spin.SpeechFinishedEvent{})

	s.Run()
}

func deadworldHasMaterializedFrame(e *spin.ScriptEnv, r spin.Renderer) {
	g := r.Graphics()

	g.Font = spin.FontPfRondaSevenBold8
	g.Y = 8
	r.Print(g, "DEADWORLD")

	g.Font = spin.FontPfRondaSeven8
	g.Y = 18
	r.Print(g, "HAS MATERIALIZED")
}

func deadworldBarrierShotsToGoFrame(e *spin.ScriptEnv, r spin.Renderer) {
	g := r.Graphics()
	vars := GetVars(e)

	r.Fill(spin.ColorOff)

	g.Font = spin.FontPfRondaSevenBold16
	g.X = 30
	g.Y = 6
	g.AnchorX = spin.AnchorRight
	r.Print(g, "%v", vars.ShotsToLowerBarriers)

	g.Font = spin.FontPfRondaSeven8
	g.X = (r.Width() / 2) + 14
	g.Y = 8
	g.AnchorX = spin.AnchorCenter
	targets := "TARGETS"
	if vars.ShotsToLowerBarriers == 1 {
		targets = "TARGET"
	}
	r.Print(g, "%v TO LOWER", targets)

	g.Y = 18
	g.Font = spin.FontPfRondaSevenBold8
	r.Print(g, "BARRIERS")
}

func deadworldBarrierShotsToGoScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(spin.PriorityAnnounce)
	defer r.Close()

	s := spin.NewSequencer(e)

	s.DoFunc(func() { deadworldBarrierShotsToGoFrame(e, r) })
	e.Do(spin.PlaySound{ID: SoundAnnounce})
	s.Sleep(2_000)
	s.Run()
}

func deadworldMusicRoutine(e *spin.ScriptEnv) {
	s := spin.NewSequencer(e)

	s.Do(spin.PlayMusic{ID: MusicSuperGame, Loops: 1, Notify: true})
	s.WaitFor(spin.MusicFinishedEvent{})
	s.Do(spin.PlayMusic{ID: MusicSuperGame2})

	s.Run()
}

func deadwordBallSaverRoutine(e *spin.ScriptEnv) {
	game := spin.GetGameVars(e)
	game.BallSave = true

	s := spin.NewSequencer(e)
	s.Do(spin.DriverOn{ID: jd.LampDrainShield})
	s.Defer(spin.DriverOff{ID: jd.LampDrainShield})
	s.Sleep(23_000)
	s.Do(proc.DriverSchedule{ID: jd.LampDrainShield, Schedule: proc.BlinkSchedule, Now: true})
	s.Sleep(5_000)
	s.Do(proc.DriverSchedule{ID: jd.LampDrainShield, Schedule: proc.HurryUpBlinkSchedule, Now: true})
	s.Sleep(2_000)
	s.DoFunc(func() { game.BallSave = false })
	s.Do(spin.PlaySpeech{ID: SpeechDrainShieldDeactivated})
	s.Run()
}

func deadworldBarriersDownFrame(e *spin.ScriptEnv, r spin.Renderer) {
	g := r.Graphics()

	g.Font = spin.FontPfRondaSeven8
	g.Y = 8
	r.Print(g, "BARRIERS DOWN")

	g.Font = spin.FontPfRondaSevenBold8
	g.Y = 18
	r.Print(g, "SHOOT LEFT RAMP")
}

func deadworldShotsToGoFrame(e *spin.ScriptEnv, r spin.Renderer) {
	g := r.Graphics()
	vars := GetVars(e)

	r.Fill(spin.ColorOff)

	g.Font = spin.FontPfRondaSevenBold16
	g.X = 30
	g.Y = 6
	g.AnchorX = spin.AnchorRight
	r.Print(g, "%v", vars.ShotsToDestroyDeadworld)

	g.Font = spin.FontPfRondaSeven8
	g.X = (r.Width() / 2) + 14
	g.Y = 8
	g.AnchorX = spin.AnchorCenter
	shots := "SHOTS"
	if vars.ShotsToDestroyDeadworld == 1 {
		shots = "SHOT"
	}
	r.Print(g, "%v FOR", shots)

	g.Y = 18
	g.Font = spin.FontPfRondaSevenBold8
	r.Print(g, "100 MILLION")
}

func deadworldShotsToGoScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(spin.PriorityAnnounce)
	defer r.Close()

	deadworldShotsToGoFrame(e, r)
	e.Sleep(2_000)
}

func deadworldBarriersDownIntroScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(spin.PriorityAnnounce)
	defer r.Close()

	s := spin.NewSequencer(e)

	s.DoFunc(func() { deadworldBarriersDownFrame(e, r) })
	s.Do(spin.PlaySpeech{ID: SpeechDeadworldDefensiveBarriersAreDownFireAtWill, Notify: true, Duck: 0.5})
	s.WaitFor(spin.SpeechFinishedEvent{})
	s.Sleep(500)
	s.DoFunc(func() { deadworldShotsToGoFrame(e, r) })
	s.Do(spin.PlaySpeech{ID: SpeechLoadPlanetForSuperJackpot, Notify: true, Duck: 0.5})
	s.WaitFor(spin.SpeechFinishedEvent{})
	s.DoFunc(func() { r.Close() })
	s.Sleep(5_000)
	s.Do(spin.PlaySpeech{ID: SpeechGoForTheHundredMillion, Notify: true, Duck: 0.5})
	s.WaitFor(spin.SpeechFinishedEvent{})

	s.Run()
}

func deadworldBarriersDownScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(0)
	defer r.Close()

	vars := GetVars(e)
	player := spin.GetPlayerVars(e)
	game := spin.GetGameVars(e)

	vars.ShotsToDestroyDeadworld = 1

	e.Do(spin.PlayScript{ID: ScriptDeadworldBarriersDownIntro})
	e.Do(spin.AddBall{N: e.Config.NumBalls - game.BallsInPlay})
	e.Do(spin.PlayScript{ID: ScriptLeftRampRunway})

	defer func() {
		e.Do(spin.DriverOff{ID: jd.FlasherJudgeDeath})
		e.Do(spin.StopScript{ID: ScriptLeftRampRunway})
	}()

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		spin.RenderFrameLoop(e, func(e *spin.ScriptEnv) {
			TimerAndScorePanel(e, r, "DEADWORLD", vars.ShotsToDestroyDeadworld, player.Score, "SHOOT LEFT RAMP")
		})
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		for {
			if _, done := e.WaitFor(spin.SwitchEvent{ID: jd.SwitchLeftRampEnter}); done {
				return
			}
			// FIXME: make a common function with multiball?
			// From: https://github.com/preble/JD-pyprocgame/blob/master/multiball.py#L56
			e.Do(proc.DriverSchedule{ID: jd.CoilDiverter, Schedule: 0xfff, CycleSeconds: 1, Now: true})
		}
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		for {
			if _, done := e.WaitFor(spin.SwitchEvent{ID: jd.SwitchLeftRampToLock}); done {
				return
			}
			vars.ShotsToDestroyDeadworld--
			if vars.ShotsToDestroyDeadworld == 0 {
				break
			}
			e.Do(spin.PlaySound{ID: SoundDeadworldExplosion, Duck: 0.5, Notify: true})
			e.Do(proc.DriverSchedule{ID: jd.FlasherGlobe, Schedule: proc.FlasherBlinkSchedule, Now: true, CycleSeconds: 4})
			e.Do(spin.PlayScript{ID: ScriptDeadworldShotsToGo})
		}
		e.Post(spin.AdvanceEvent{})
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		for {
			evt, done := e.WaitFor(spin.BallDrainEvent{})
			if done {
				return
			}
			if evt.(spin.BallDrainEvent).BallsInPlay == 1 {
				break
			}
		}
		e.Post(spin.TimeoutEvent{})
	})

	event, done := e.WaitFor(spin.AdvanceEvent{}, spin.TimeoutEvent{})
	if done {
		return
	}
	switch event.(type) {
	case spin.AdvanceEvent:
		e.Do(spin.PlayScript{ID: ScriptDeadworldComplete})
	case spin.TimeoutEvent:
		panic("we are done")
	}

}

func deadworldCompleteScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(spin.PriorityAnnounce)
	defer r.Close()

	vars := GetVars(e)
	player := spin.GetPlayerVars(e)
	player.Score += 100_000_000
	total := player.Score - vars.StartScore

	e.Do(spin.StopAudio{})
	e.Do(spin.FlippersOff{})

	s := spin.NewSequencer(e)
	s.Do(spin.PlaySound{ID: SoundDeadworldExplosion})
	s.Do(spin.DriverPulse{ID: jd.FlasherGlobe})
	s.DoFunc(func() { r.Fill(spin.ColorOn) })
	s.Sleep(200)
	s.DoFunc(func() { r.Fill(spin.ColorOff) })
	s.Sleep(200)
	s.LoopN(4)
	if done := s.Run(); done {
		return
	}

	s = spin.NewSequencer(e)
	s.Sleep(2_000)
	s.Do(spin.PlaySound{ID: SoundApplause})
	s.Sleep(1_000)
	s.Do(spin.PlayMusic{ID: MusicDeadworldComplete, Notify: true})
	s.Sleep(2_000)
	s.DoFunc(func() { OneLinePanel(e, r, "CONGRATULATIONS") })
	s.Do(spin.PlaySpeech{ID: SpeechCongratulations, Notify: true, Duck: 0.5})
	s.WaitFor(spin.SpeechFinishedEvent{})
	s.Sleep(1_000)
	s.Do(spin.PlaySpeech{ID: SpeechOneHundredMillion})
	s.DoFunc(func() { OneLineBigPanel(e, r, "ONE") })
	s.Sleep(500)
	s.DoFunc(func() { OneLineBigPanel(e, r, "HUNDRED") })
	s.Sleep(750)
	s.DoFunc(func() { OneLineBigPanel(e, r, "MILLION") })
	s.Sleep(3_000)
	s.DoFunc(func() { ModeAndScorePanel(e, r, "DEADWORLD TOTAL", total) })
	s.Sleep(1_000)
	s.Do(spin.PlaySpeech{ID: SpeechCantWeJustBeFriends})
	s.Sleep(2_000)
	s.Do(spin.FadeOutMusic{Time: 5_000})
	s.WaitFor(spin.MusicFinishedEvent{})
	s.DoFunc(func() { r.Fill(spin.ColorOff) })
	s.Sleep(1_000)
	s.Do(spin.PlaySpeech{ID: SpeechSuperGameHasEnded, Notify: true})
	s.WaitFor(spin.SpeechFinishedEvent{})
	s.Sleep(1_000)

	if done := s.Run(); done {
		return
	}

	e.Do(spin.PlayMusic{ID: MusicAttackFromMars})
	multiballEnd(e)
}

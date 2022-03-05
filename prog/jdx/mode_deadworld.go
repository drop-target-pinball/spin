package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/proc"
)

func deadworldModeScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(0)

	defer func() { panic("we are done") }()
	game := spin.GetGameVars(e)
	switches := spin.GetResourceVars(e).Switches
	vars := GetVars(e)
	player := spin.GetPlayerVars(e)

	vars.ShotsToLowerBarriers = 50

	e.Do(spin.StopScript{ID: ScriptLightBallLock})
	e.Do(spin.StopScript{ID: ScriptBallLock})
	e.Do(spin.StopScript{ID: ScriptChain})
	e.Do(spin.StopScript{ID: ScriptCrimeScenes})
	e.Do(spin.StopScript{ID: ScriptPlungeMode})
	e.Do(spin.StopScript{ID: ScriptBallSaver})

	e.Do(spin.PlayScript{ID: ScriptDeadworldIntro})
	if _, done := e.WaitFor(spin.ScriptFinishedEvent{ID: ScriptDeadworldIntro}); done {
		return
	}

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

}

func deadworldIntroScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(spin.PriorityAnnounce)
	defer r.Close()

	e.NewCoroutine(deadworldMusicRoutine)

	s := spin.NewSequencer(e)

	s.DoFunc(func() { deadworldHasMaterializedFrame(e, r) })
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

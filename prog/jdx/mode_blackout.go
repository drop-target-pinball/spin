package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/proc"
)

func blackoutModeScript(e *spin.ScriptEnv) {
	r, _ := e.Display("").Renderer("")

	e.Do(spin.PlayMusic{ID: MusicMode1})

	vars := GetVars(e)
	player := spin.GetPlayerVars(e)

	vars.Mode = ModeBlackout
	defer func() { vars.Mode = ModeNone }()
	vars.Multiplier = 2
	defer func() { vars.Multiplier = 1 }()

	e.Do(proc.DriverSchedule{ID: jd.FlasherBlackout, Schedule: proc.FlasherBlinkSchedule})
	defer e.Do(spin.DriverOff{ID: jd.FlasherBlackout})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.Do(spin.PlaySpeech{ID: SpeechMegaCityOneIsBlackedOutBeOnTheAlertForLooters, Notify: true, Duck: 0.5})
		s.WaitFor(spin.SpeechFinishedEvent{})
		s.Do(spin.PlaySpeech{ID: SpeechSendBackupUnits, Notify: true})
		s.WaitFor(spin.SpeechFinishedEvent{})

		s.Run()
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.Do(spin.DriverOff{ID: jd.GI1})
		s.Sleep(200)
		s.Do(spin.DriverOff{ID: jd.GI2})
		s.Sleep(300)
		s.Do(spin.DriverOff{ID: jd.GI3})
		s.Sleep(100)
		s.Do(spin.DriverOff{ID: jd.GI4})
		s.Sleep(300)
		s.Do(spin.DriverOff{ID: jd.GI5})
		s.Sleep(200)

		s.Run()
	})
	defer e.Do(spin.PlayScript{ID: jd.ScriptGIOn})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		ModeIntroScript(e, "BLACKOUT", "EVERYTHING", "2X")
		spin.RenderFrameLoop(e, func(e *spin.ScriptEnv) {
			ModeAndScorePanel(e, r, "BLACKOUT", player.Score)
		})
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.WaitFor(spin.SwitchEvent{ID: jd.SwitchTopLeftRampExit})
		s.Do(spin.PlayScript{ID: ScriptBlackoutJackpot})
		s.Loop()

		s.Run()
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.WaitFor(spin.BallDrainEvent{})
		s.Post(spin.AdvanceEvent{})

		s.Run()
	})

	e.Do(spin.AddBall{})
	if _, done := e.WaitFor(spin.AdvanceEvent{}); done {
		return
	}
	e.Do(spin.PlayMusic{ID: MusicMain})
}

func blackoutJackpotScript(e *spin.ScriptEnv) {
	r, _ := e.Display("").Renderer(spin.LayerPriority)
	defer r.Clear()

	ScoreAndLabelPanel(e, r, ScoreBlackoutJackpot, "JACKPOT")

	s := spin.NewSequencer(e)

	s.Do(spin.AwardScore{Val: ScoreBlackoutJackpot})
	s.Do(spin.PlaySound{ID: SoundBlackoutJackpot, Notify: true, Duck: 0.25})
	s.WaitFor(spin.SoundFinishedEvent{ID: SoundBlackoutJackpot})

	s.Run()
}

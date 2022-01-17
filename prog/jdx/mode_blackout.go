package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
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

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.Do(spin.PlaySpeech{ID: SpeechMegaCityOneIsBlackedOutBeOnTheAlertForLooters, Notify: true, Duck: 0.5})
		s.WaitFor(spin.SpeechFinishedEvent{})
		s.Do(spin.PlaySpeech{ID: SpeechSendBackupUnits, Notify: true})
		s.WaitFor(spin.SpeechFinishedEvent{})

		s.Run()
	})

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
	e.Post(spin.ScriptFinishedEvent{ID: ScriptBlackoutMode})
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

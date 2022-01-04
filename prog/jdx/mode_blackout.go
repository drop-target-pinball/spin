package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

func blackoutJackpotScript(e spin.Env) {
	r, _ := e.Display("").Renderer(spin.LayerPriority)
	defer r.Clear()

	scoreAndLabelPanel(e, r, ScoreBlackoutJackpot, "JACKPOT")

	spin.NewSequencer().
		Do(spin.AwardScore{Val: ScoreBlackoutJackpot}).
		Do(spin.PlaySound{ID: SoundBlackoutJackpot, Notify: true, Duck: 0.25}).
		WaitFor(spin.SoundFinishedEvent{ID: SoundBlackoutJackpot}).
		Run(e)
}

func blackoutModeScript(e spin.Env) {
	r, _ := e.Display("").Renderer("")
	defer r.Clear()

	e.Do(spin.PlayMusic{ID: MusicMode1})

	vars := GetVars(e)
	player := spin.GetPlayerVars(e)
	vars.Multiplier = 2
	defer func() { vars.Multiplier = 1 }()

	e.NewCoroutine(e.Context(), func(e spin.Env) {
		spin.NewSequencer().
			Do(spin.PlaySpeech{ID: SpeechMegaCityOneIsBlackedOutBeOnTheAlertForLooters, Notify: true, Duck: 0.5}).
			WaitFor(spin.SpeechFinishedEvent{}).
			Do(spin.PlaySpeech{ID: SpeechSendBackupUnits}).
			Run(e)
	})

	e.NewCoroutine(e.Context(), func(e spin.Env) {
		spin.NewSequencer().
			WaitFor(spin.ShotEvent{ID: jd.ShotTopLeftRamp}).
			Do(spin.PlayScript{ID: ScriptBlackoutJackpot}).
			Loop().
			Run(e)
	})

	e.NewCoroutine(e.Context(), func(e spin.Env) {
		spin.NewSequencer().
			WaitFor(spin.BallDrainEvent{}).
			Post(spin.AdvanceEvent{}).
			Run(e)
	})

	e.Do(spin.AddBall{})

	modeText := [3]string{"BLACKOUT", "EVERYTHING", "2X"}
	if done := modeIntroVideo(e, modeText); done {
		return
	}

	spin.RenderFrameScript(e, func(e spin.Env) {
		modeAndScorePanel(e, r, "BLACKOUT", player.Score)
	})

	if _, done := e.WaitFor(spin.AdvanceEvent{}); done {
		return
	}
	e.Do(spin.PlayMusic{ID: MusicMain})
	e.Post(spin.ScriptFinishedEvent{ID: ScriptBlackoutMode})
}

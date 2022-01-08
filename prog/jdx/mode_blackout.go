package jdx

// func blackoutModeScript(e spin.Env) {
// 	r, _ := e.Display("").Renderer("")
// 	defer r.Clear()

// 	e.Do(spin.PlayMusic{ID: MusicMode1})

// 	vars := GetVars(e)
// 	player := spin.GetPlayerVars(e)
// 	vars.Multiplier = 2
// 	defer func() { vars.Multiplier = 1 }()

// 	e.NewCoroutine(e.Context(), func(e spin.Env) {
// 		spin.NewSequencer().
// 			Do(spin.PlaySpeech{ID: SpeechMegaCityOneIsBlackedOutBeOnTheAlertForLooters, Notify: true, Duck: 0.5}).
// 			WaitFor(spin.SpeechFinishedEvent{}).
// 			Do(spin.PlaySpeech{ID: SpeechSendBackupUnits, Notify: true}).
// 			WaitFor(spin.SpeechFinishedEvent{}).
// 			Run(e)
// 	})

// 	e.NewCoroutine(e.Context(), func(e spin.Env) {
// 		if done := ModeIntroSequence(e, "BLACKOUT", "EVERYTHING", "2X").Run(e); done {
// 			return
// 		}
// 		spin.RenderFrameScript(e, func(e spin.Env) {
// 			ModeAndScorePanel(e, r, "BLACKOUT", player.Score)
// 		})
// 		e.WaitFor(spin.Done{})
// 	})

// 	e.NewCoroutine(e.Context(), func(e spin.Env) {
// 		spin.NewSequencer().
// 			WaitFor(spin.ShotEvent{ID: jd.ShotTopLeftRamp}).
// 			Do(spin.PlayScript{ID: ScriptBlackoutJackpot}).
// 			Loop().
// 			Run(e)
// 	})

// 	e.NewCoroutine(e.Context(), func(e spin.Env) {
// 		spin.NewSequencer().
// 			WaitFor(spin.BallDrainEvent{}).
// 			Post(spin.AdvanceEvent{}).
// 			Run(e)
// 	})

// 	e.Do(spin.AddBall{})
// 	if _, done := e.WaitFor(spin.AdvanceEvent{}); done {
// 		return
// 	}
// 	e.Do(spin.PlayMusic{ID: MusicMain})
// 	e.Post(spin.ScriptFinishedEvent{ID: ScriptBlackoutMode})
// }

// func blackoutJackpotScript(e spin.Env) {
// 	r, _ := e.Display("").Renderer(spin.LayerPriority)
// 	defer r.Clear()

// 	scoreAndLabelPanel(e, r, ScoreBlackoutJackpot, "JACKPOT")

// 	spin.NewSequencer().
// 		Do(spin.AwardScore{Val: ScoreBlackoutJackpot}).
// 		Do(spin.PlaySound{ID: SoundBlackoutJackpot, Notify: true, Duck: 0.25}).
// 		WaitFor(spin.SoundFinishedEvent{ID: SoundBlackoutJackpot}).
// 		Run(e)
// }

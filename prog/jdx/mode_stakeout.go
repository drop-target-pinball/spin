package jdx

// func stakeoutScoreFrame(e spin.Env) {
// 	vars := GetVars(e)
// 	player := spin.GetPlayerVars(e)
// 	r, g := e.Display("").Renderer("")

// 	r.Fill(spin.ColorBlack)
// 	g.Y = 2
// 	g.Font = builtin.FontPfArmaFive8
// 	r.Print(g, "STAKEOUT")

// 	g.AnchorY = spin.AnchorBottom
// 	g.Y = r.Height()
// 	r.Print(g, "SHOOT RIGHT RAMP")

// 	g.AnchorX = spin.AnchorLeft
// 	g.X = 5
// 	g.AnchorY = spin.AnchorMiddle
// 	g.Y = r.Height() / 2
// 	g.Font = builtin.Font14x10
// 	r.Print(g, "%v", vars.Timer)

// 	g.X = r.Width() - 2
// 	g.AnchorX = spin.AnchorRight
// 	g.Font = builtin.Font09x7
// 	r.Print(g, spin.FormatScore("%v", player.Score))
// }

// func stakeoutInterestingScript(e spin.Env) {
// 	r, _ := e.Display("").Renderer(spin.LayerPriority)
// 	defer r.Clear()

// 	vars := GetVars(e)
// 	callouts := []string{
// 		SpeechIWonderWhatsOverThere,
// 		SpeechIWonderWhatsDownThere,
// 	}

// 	callout := callouts[vars.StakeoutCallout]
// 	vars.StakeoutCallout += 1
// 	if vars.StakeoutCallout >= len(callouts) {
// 		vars.StakeoutCallout = 0
// 	}

// 	vars.StakeoutBonus += ScoreStakeoutN
// 	scoreAndLabelPanel(e, r, ScoreStakeoutN, "AWARDED")

// 	spin.NewSequencer().
// 		Do(spin.PlaySpeech{ID: callout, Priority: spin.PriorityAudioModeCallout}).
// 		Sleep(2_500).
// 		Do(spin.PlaySpeech{ID: SpeechInteresting, Priority: spin.PriorityAudioModeCallout}).
// 		Run(e)
// }

// func stakeoutWatchRampScript(e spin.Env) {
// 	for {
// 		if _, done := e.WaitFor(spin.ShotEvent{ID: jd.ShotRightRamp}); done {
// 			return
// 		}
// 		e.Do(spin.PlayScript{ID: ScriptStakeoutInteresting})
// 	}
// }

// func stakeoutCompleteScript(e spin.Env) {
// 	r, _ := e.Display("").Renderer(spin.LayerPriority)
// 	defer r.Clear()

// 	vars := GetVars(e)
// 	ModeAndScorePanel(e, r, "STAKEOUT TOTAL", vars.StakeoutBonus)
// 	e.Do(spin.PlayMusic{ID: MusicMain})

// 	if !e.IsActive(ScriptStakeoutInteresting) {
// 		e.Do(spin.PlaySound{ID: SoundSuccess, Duck: 0.5})
// 	}
// 	if done := e.Sleep(3_000 * time.Millisecond); done {
// 		e.Do(spin.StopSound{ID: SoundSuccess})
// 		return
// 	}
// }

// func stakeoutModeScript(e spin.Env) {
// 	vars := GetVars(e)

// 	e.Do(spin.PlayMusic{ID: MusicMode2})
// 	vars.StakeoutBonus = ScoreStakeout0

// 	e.NewCoroutine(e.Context(), stakeoutWatchRampScript)
// 	vars.Timer = 30
// 	spin.CountdownScript(e, &vars.Timer, 1000, spin.TimeoutEvent{})

// 	e.NewCoroutine(e.Context(), func(e spin.Env) {
// 		spin.NewSequencer().
// 			Do(spin.PlaySpeech{ID: SpeechImStakingOutACrackHouseInSectorTwentyThree}).
// 			Sleep(15_000).
// 			Do(spin.PlaySpeech{ID: SpeechShootRightRamp}).
// 			Sleep(10_000).
// 			Do(spin.PlaySpeech{ID: SpeechShootRightRamp}).
// 			Run(e)
// 	})

// 	modeText := [3]string{"STAKEOUT", "SHOOT", "RIGHT RAMP"}
// 	if done := modeIntroVideo(e, modeText); done {
// 		return
// 	}
// 	spin.RenderFrameScript(e, stakeoutScoreFrame)

// 	e.WaitFor(spin.TimeoutEvent{})
// 	e.Do(spin.PlayScript{ID: ScriptStakeoutComplete})
// 	e.Post(spin.ScriptFinishedEvent{ID: ScriptStakeoutMode})
// }

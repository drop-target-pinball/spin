package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

func manhuntModeScript(e *spin.ScriptEnv) {
	r, _ := e.Display("").Renderer("")

	e.Do(spin.PlayMusic{ID: MusicMode1})

	vars := GetVars(e)
	player := spin.GetPlayerVars(e)
	vars.Timer = 30
	vars.ManhuntBonus = ScoreManhunt0

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.Do(spin.PlaySpeech{ID: SpeechSuspiciousCharacterReportedInEugeneBlock, Notify: true, Vol: 64})
		s.WaitFor(spin.SpeechFinishedEvent{})
		s.Do(spin.PlaySound{ID: SoundWalking, Loop: true})
		s.Defer(spin.StopSound{ID: SoundWalking})

		s.DoFunc(func() {
			e.NewCoroutine(func(e *spin.ScriptEnv) {
				s := spin.NewSequencer(e)
				s.Sleep(2000)
				s.Do(spin.PlaySound{ID: SoundManhuntSingleFire})
				s.Loop()
				s.Run()
			})
		})

		s.Sleep(10_000)
		s.Do(spin.PlaySpeech{ID: SpeechShootLeftRamp})
		s.Sleep(3_000)
		s.Do(spin.PlaySpeech{ID: SpeechStop})
		s.Sleep(1_000)
		s.Do(spin.PlaySpeech{ID: SpeechOrIWillShoot})
		s.Sleep(3_000)
		s.Do(spin.PlaySpeech{ID: SpeechFreeze})
		s.Sleep(4_000)
		s.Do(spin.PlaySpeech{ID: SpeechShootLeftRamp})
		s.Sleep(5_000)
		s.Run()
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		ModeIntroScript(e, "MANHUNT", "SHOOT", "LEFT RAMP")
		spin.RenderFrameScript(e, func(e *spin.ScriptEnv) {
			TimerAndScorePanel(e, r, "MANHUNT", vars.Timer, player.Score, "SHOOT LEFT RAMP")
		})
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		spin.CountdownScript(e, &vars.Timer, 1000, spin.TimeoutEvent{})
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.WaitFor(spin.ShotEvent{ID: jd.ShotLeftRamp})
		s.DoFunc(func() { vars.ManhuntBonus += ScoreManhuntN })
		s.Do(spin.PlaySound{ID: SoundManhuntAutoFire})
		s.Loop()

		s.Run()
	})

	if _, done := e.WaitFor(spin.TimeoutEvent{}); done {
		return
	}
	e.Do(spin.PlayScript{ID: ScriptManhuntComplete})
	e.Post(spin.ScriptFinishedEvent{ID: ScriptManhuntMode})
}

func manhuntCompleteScript(e *spin.ScriptEnv) {
	r, _ := e.Display("").Renderer(spin.LayerPriority)
	defer r.Clear()

	e.Do(spin.PlayMusic{ID: MusicMain})

	vars := GetVars(e)

	s := spin.NewSequencer(e)
	s.Sleep(1_000)
	s.DoFunc(func() { ModeAndScorePanel(e, r, "MANHUNT TOTAL", vars.ManhuntBonus) })
	s.Do(spin.PlaySound{ID: SoundSuccess, Duck: 0.5})
	s.Sleep(3_000)
	s.Run()
}

// func manhuntScoreFrame(e spin.Env) {
// 	vars := GetVars(e)
// 	player := spin.GetPlayerVars(e)

// 	r, g := e.Display("").Renderer("")

// 	r.Fill(spin.ColorBlack)
// 	g.Y = 2
// 	g.Font = builtin.FontPfArmaFive8
// 	r.Print(g, "MANHUNT")

// 	g.AnchorX = spin.AnchorLeft
// 	g.X = 5
// 	g.AnchorY = spin.AnchorMiddle
// 	g.Y = r.Height()/2 + 6
// 	g.Font = builtin.Font14x10
// 	r.Print(g, "%v", vars.Timer)

// 	g.X = r.Width() - 2
// 	g.AnchorX = spin.AnchorRight
// 	g.Font = builtin.Font09x7
// 	r.Print(g, spin.FormatScore("%v", player.Score))
// }

// func manhuntTotalFrame(e spin.Env) {
// 	r, g := e.Display("").Renderer(spin.LayerPriority)
// 	vars := GetVars(e)

// 	r.Fill(spin.ColorBlack)
// 	g.Y = 2
// 	g.Font = builtin.FontPfArmaFive8
// 	r.Print(g, "MANHUNT TOTAL")
// 	g.Y = 12

// 	g.Font = builtin.Font14x10
// 	r.Print(g, spin.FormatScore("%v", vars.ManhuntBonus))
// }

// func manhuntWatchRamp(e spin.Env) {
// 	vars := GetVars(e)
// 	for {
// 		if _, done := e.WaitFor(spin.ShotEvent{ID: jd.ShotLeftRamp}); done {
// 			return
// 		}
// 		vars.ManhuntBonus += ScoreManhuntN
// 		e.Do(spin.PlaySound{ID: SoundManhuntAutoFire})
// 	}
// }

// func manhuntCompleteScript(e spin.Env) {
// 	defer e.Display("").Clear(spin.LayerPriority)
// 	manhuntTotalFrame(e)
// 	e.Do(spin.PlayMusic{ID: MusicMain})
// 	e.Do(spin.PlaySound{ID: SoundSuccess, Duck: 0.5})
// 	if done := e.Sleep(3_000 * time.Millisecond); done {
// 		e.Do(spin.StopSound{ID: SoundSuccess})
// 		return
// 	}
// }

// func manhuntModeScript(e spin.Env) {
// 	vars := GetVars(e)
// 	ctx, cancel := e.Derive()
// 	defer cancel()

// 	e.Do(spin.PlayMusic{ID: MusicMode1})
// 	vars.ManhuntBonus = ScoreManhunt0

// 	gunfire := spin.NewSequencer().
// 		Sleep(2000).
// 		Do(spin.PlaySound{ID: SoundManhuntSingleFire}).
// 		Loop()

// 	e.NewCoroutine(e.Context(), func(e spin.Env) {
// 		spin.NewSequencer().
// 			Do(spin.PlaySpeech{ID: SpeechSuspiciousCharacterReportedInEugeneBlock, Notify: true}).
// 			WaitFor(spin.SpeechFinishedEvent{}).
// 			Do(spin.PlaySound{ID: SoundWalking, Loop: true}).
// 			Func(func() { gunfire.Run(e) }).
// 			Sleep(10_000).
// 			Do(spin.PlaySpeech{ID: SpeechShootLeftRamp}).
// 			Sleep(3_000).
// 			Do(spin.PlaySpeech{ID: SpeechStop}).
// 			Sleep(1_000).
// 			Do(spin.PlaySpeech{ID: SpeechOrIWillShoot}).
// 			Sleep(3_000).
// 			Do(spin.PlaySpeech{ID: SpeechFreeze}).
// 			Sleep(4_000).
// 			Do(spin.PlaySpeech{ID: SpeechShootLeftRamp}).
// 			Run(e)
// 	})

// 	defer e.Do(spin.StopSound{ID: SoundWalking})

// 	e.NewCoroutine(ctx, manhuntWatchRamp)
// 	vars.Timer = 30
// 	spin.CountdownScript(e, &vars.Timer, 1000, spin.TimeoutEvent{})

// 	modeText := [3]string{"MANHUNT", "SHOOT", "LEFT RAMP"}
// 	if done := modeIntroVideo(e, modeText); done {
// 		return
// 	}
// 	spin.RenderFrameScript(e, manhuntScoreFrame)

// 	if _, done := e.WaitFor(spin.TimeoutEvent{}); done {
// 		return
// 	}
// 	e.Do(spin.PlayScript{ID: ScriptManhuntComplete})
// 	e.Post(spin.ScriptFinishedEvent{ID: ScriptManhuntMode})
// }

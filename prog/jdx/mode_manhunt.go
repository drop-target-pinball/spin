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

	vars.Mode = ModeManhunt
	defer func() { vars.Mode = ModeNone }()
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
		spin.RenderFrameLoop(e, func(e *spin.ScriptEnv) {
			TimerAndScorePanel(e, r, "MANHUNT", vars.Timer, player.Score, "SHOOT LEFT RAMP")
		})
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		spin.CountdownLoop(e, &vars.Timer, 1000, spin.TimeoutEvent{})
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.WaitFor(spin.SwitchEvent{ID: jd.SwitchLeftRampExit})
		s.DoFunc(func() { vars.ManhuntBonus += ScoreManhuntN })
		s.Do(spin.PlaySound{ID: SoundManhuntAutoFire})
		s.Loop()

		s.Run()
	})

	if _, done := e.WaitFor(spin.TimeoutEvent{}); done {
		return
	}
	e.Do(spin.PlayScript{ID: ScriptManhuntComplete})
}

func manhuntCompleteScript(e *spin.ScriptEnv) {
	r, _ := e.Display("").Renderer(spin.LayerPriority)
	defer r.Clear()

	e.Do(spin.PlayMusic{ID: MusicMain})

	vars := GetVars(e)
	player := spin.GetPlayerVars(e)

	s := spin.NewSequencer(e)
	s.DoFunc(func() { TimerAndScorePanel(e, r, "MANHUNT", vars.Timer, player.Score, "SHOOT LEFT RAMP") })
	s.Sleep(1_000)
	s.DoFunc(func() { ModeAndScorePanel(e, r, "MANHUNT TOTAL", vars.ManhuntBonus) })
	s.Do(spin.PlaySound{ID: SoundSuccess, Duck: 0.5})
	s.Sleep(3_000)
	s.Run()
}

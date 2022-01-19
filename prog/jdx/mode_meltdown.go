package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

func meltdownModeScript(e *spin.ScriptEnv) {
	r, _ := e.Display("").Renderer("")

	e.Do(spin.PlayMusic{ID: MusicMode1})

	vars := GetVars(e)
	player := spin.GetPlayerVars(e)

	vars.Mode = ModeMeltdown
	defer func() { vars.Mode = ModeNone }()
	vars.Timer = 30
	vars.MeltdownBonus = ScoreMeltdown0

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.Do(spin.PlaySpeech{ID: SpeechContainmentFailureAtThreeMeterIsland, Notify: true, Duck: 0.5})
		s.WaitFor(spin.SpeechFinishedEvent{})
		if done := s.Run(); done {
			return
		}

		s = spin.NewSequencer(e)
		s.Defer(spin.StopSound{ID: SoundMeltdownKlaxon})
		s.Sleep(4_000)
		s.Do(spin.PlaySound{ID: SoundMeltdownCracking})
		s.Loop()
		s.Run()
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.WaitFor(spin.SwitchEvent{ID: jd.SwitchCaptiveBall1})
		s.Do(spin.PlaySpeech{ID: SpeechReactorOneStabilized, Duck: 0.5, Priority: spin.PriorityAudioModeCallout})
		s.DoFunc(func() { vars.MeltdownBonus = ScoreMeltdown1 })

		s.WaitFor(spin.SwitchEvent{ID: jd.SwitchCaptiveBall2})
		s.Do(spin.PlaySpeech{ID: SpeechReactorTwoStabilized, Duck: 0.5, Priority: spin.PriorityAudioModeCallout})
		s.DoFunc(func() { vars.MeltdownBonus = ScoreMeltdown2 })

		s.WaitFor(spin.SwitchEvent{ID: jd.SwitchCaptiveBall3})
		s.DoFunc(func() { vars.MeltdownBonus = ScoreMeltdown3 })
		s.Post(spin.AdvanceEvent{})

		s.Run()
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		spin.CountdownLoop(e, &vars.Timer, 1000, spin.TimeoutEvent{})
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		if done := ModeIntroScript(e, "MELTDOWN", "SHOOT", "CAPTIVE BALLS"); done {
			return
		}
		spin.RenderFrameLoop(e, func(e *spin.ScriptEnv) {
			TimerAndScorePanel(e, r, "MELTDOWN", vars.Timer, player.Score, "SHOOT CAPTIVE BALLS")
		})
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		spin.WatcherTimerLoop(e, &vars.Timer, func(timer int) {
			if timer < 27 && timer > 0 {
				vol := 0
				if timer == 20 || timer == 10 {
					vol = 64
				}
				e.Do(spin.PlaySound{ID: SoundMeltdownKlaxon, Vol: vol})
			}
			switch timer {
			// case 20:
			// 	e.Do(spin.PlaySpeech{ID: SpeechAllReactorsApprochingCriticalMass})
			case 10:
				e.Do(spin.PlaySpeech{ID: SpeechMeltdownIsImminent})
			case 4:
				e.Do(spin.PlaySpeech{ID: SpeechFour})
			case 3:
				e.Do(spin.PlaySpeech{ID: SpeechThree})
			case 2:
				e.Do(spin.PlaySpeech{ID: SpeechTwo})
			case 1:
				e.Do(spin.PlaySpeech{ID: SpeechOne})

			}
		})
	})

	evt, done := e.WaitFor(
		spin.AdvanceEvent{},
		spin.TimeoutEvent{},
	)
	if done {
		return
	}
	if evt == (spin.TimeoutEvent{}) {
		e.Do(spin.PlayScript{ID: ScriptMeltdownIncomplete})
	} else {
		e.Do(spin.PlayScript{ID: ScriptMeltdownComplete})
	}
}

func meltdownIncompleteScript(e *spin.ScriptEnv) {
	r, _ := e.Display("").Renderer(spin.LayerPriority)
	defer r.Clear()

	e.Do(spin.PlayMusic{ID: MusicMain})

	vars := GetVars(e)

	s := spin.NewSequencer(e)

	s.Do(spin.PlaySound{ID: SoundMeltdownExplosion})
	s.Sleep(2_000)

	s.DoFunc(func() { ModeAndScorePanel(e, r, "MELTDOWN TOTAL", vars.MeltdownBonus) })
	s.Do(spin.PlaySound{ID: SoundSuccess})
	s.Sleep(2_500)

	s.Run()
}

func meltdownCompleteScript(e *spin.ScriptEnv) {
	r, _ := e.Display("").Renderer(spin.LayerPriority)
	defer r.Clear()

	e.Do(spin.PlayMusic{ID: MusicMain})

	vars := GetVars(e)
	ModeAndScorePanel(e, r, "MELTDOWN TOTAL", vars.MeltdownBonus)

	s := spin.NewSequencer(e)

	s.Do(spin.PlaySpeech{ID: SpeechAllReactorsStabilized})
	s.Sleep(3_000)
	s.Do(spin.PlaySpeech{ID: SpeechThreeMeterIslandIsSecured, Notify: true})
	s.WaitFor(spin.SpeechFinishedEvent{})

	s.Run()
}

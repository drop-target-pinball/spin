package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

func sniperModeScript(e *spin.ScriptEnv) {
	r, _ := e.Display("").Renderer("")

	e.Do(spin.PlayMusic{ID: MusicMode1})

	vars := GetVars(e)
	vars.SniperScore = ScoreSniperStart

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.Do(spin.PlaySpeech{ID: SpeechSniperIsShootingIntoCrowdFromJohnsonTower, Duck: 0.5})
		s.Sleep(4_000)
		s.Do(spin.PlaySpeech{ID: SpeechShootSniperTower})
		s.Sleep(1_000)

		s.DoFunc(func() {
			e.NewCoroutine(func(e *spin.ScriptEnv) {
				spin.ScoreHurryUpScript(e,
					&vars.SniperScore,
					160, // tick ms
					ScoreSniperDec,
					ScoreSniperEnd,
					spin.TimeoutEvent{},
				)
			})
		})

		s.DoFunc(func() {
			e.NewCoroutine(func(e *spin.ScriptEnv) {
				s := spin.NewSequencer(e)

				s.Do(spin.PlaySound{ID: SoundGunLoadSniper})
				s.Sleep(1_500)
				s.Do(spin.PlaySound{ID: SoundGunFire})
				s.Sleep(1_500)
				s.Loop()

				s.Run()
			})
		})

		s.Run()
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		if done := ModeIntroScript(e, "SNIPER", "SHOOT", "SNIPER TOWER"); done {
			return
		}
		spin.RenderFrameScript(e, func(e *spin.ScriptEnv) {
			ModeAndScorePanel(e, r, "SNIPER", vars.SniperScore)
		})
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.WaitFor(spin.ShotEvent{ID: jd.ShotRightPopper})
		s.Post(spin.AdvanceEvent{})

		s.Run()
	})

	evt, done := e.WaitFor(
		spin.AdvanceEvent{},
		spin.TimeoutEvent{},
	)
	if done {
		return
	}
	if evt == (spin.TimeoutEvent{}) {
		e.Do(spin.PlayMusic{ID: MusicMain})
		e.Post(spin.ScriptFinishedEvent{ID: ScriptSniperMode})
		return
	}
	e.Do(spin.PlayScript{ID: ScriptSniperMode2})
}

func sniperMode2Script(e *spin.ScriptEnv) {
	r, _ := e.Display("").Renderer("")

	vars := GetVars(e)
	vars.Timer = 10
	vars.SniperBonus = vars.SniperScore

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.Do(spin.PlaySound{ID: SoundSuccess, Duck: 0.5})
		s.Sleep(1_000)
		s.Do(spin.DriverPulse{ID: jd.CoilRightPopper})
		s.Sleep(1_500)

		s.Do(spin.PlaySpeech{ID: SpeechShootSniperTower, Notify: true})
		s.WaitFor(spin.SpeechFinishedEvent{})
		s.Do(spin.PlaySpeech{ID: SpeechAaaaah, Notify: true})
		s.Sleep(3_000)
		s.Do(spin.PlaySpeech{ID: SpeechItsALongWayDown, Notify: true})
		s.Sleep(2_500)
		s.Do(spin.PlaySpeech{ID: SpeechAaaaah})
		s.Sleep(3_500)
		s.Do(spin.PlaySpeech{ID: SpeechICanSeeMyHouseFromHere})
		s.Sleep(2_000)
		s.Do(spin.PlaySpeech{ID: SpeechAaaaah})
		s.WaitFor(spin.SpeechFinishedEvent{})

		s.Run()
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.WaitFor(spin.ShotEvent{ID: jd.ShotRightPopper})
		s.Post(spin.AdvanceEvent{})

		s.Run()
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		if done := ModeAndBlinkingScoreScript(e, r, "SNIPER", vars.SniperScore); done {
			return
		}
		e.NewCoroutine(func(e *spin.ScriptEnv) {
			spin.CountdownScript(e, &vars.Timer, 1500, spin.TimeoutEvent{})
		})
		spin.RenderFrameScript(e, func(e *spin.ScriptEnv) {
			TimerAndScorePanel(e, r, "SNIPER", vars.Timer, vars.SniperScore, "")
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
		e.Do(spin.PlayScript{ID: ScriptSniperIncomplete})
	} else {
		e.Do(spin.PlayScript{ID: ScriptSniperComplete})
	}
	e.Post(spin.ScriptFinishedEvent{ID: ScriptSniperMode})
}

func sniperIncompleteScript(e *spin.ScriptEnv) {
	r, _ := e.Display("").Renderer(spin.LayerPriority)
	defer r.Clear()

	vars := GetVars(e)
	e.Do(spin.PlayMusic{ID: MusicMain})
	ModeAndScorePanel(e, r, "SNIPER TOTAL", vars.SniperBonus)

	s := spin.NewSequencer(e)

	s.Do(spin.PlaySound{ID: SoundSniperSplat})
	s.Sleep(1_000)
	s.Do(spin.PlaySpeech{ID: SpeechSniperEliminated, Notify: true})
	s.Sleep(2_000)

	s.Run()
}

func sniperCompleteScript(e *spin.ScriptEnv) {
	r, _ := e.Display("").Renderer(spin.LayerPriority)
	defer r.Clear()

	vars := GetVars(e)
	vars.SniperBonus += vars.SniperScore
	e.Do(spin.PlayMusic{ID: MusicMain})
	TimerAndScorePanel(e, r, "SNIPER", vars.Timer, vars.SniperScore, "")

	s := spin.NewSequencer(e)

	s.Do(spin.PlaySound{ID: SoundSniperSplat})
	s.Sleep(1_000)
	s.Do(spin.PlaySpeech{ID: SpeechSniperEliminated, Notify: true})
	s.Sleep(2_000)
	if done := s.Run(); done {
		return
	}

	e.Do(spin.PlaySound{ID: SoundSuccess})
	ModeAndBlinkingScoreScript(e, r, "SNIPER", vars.SniperBonus)
}

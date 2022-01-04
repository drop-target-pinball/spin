package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

func sniperModeScript(e spin.Env) {
	r, _ := e.Display("").Renderer("")
	defer r.Clear()

	e.Do(spin.PlayMusic{ID: MusicMode1})

	vars := GetVars(e)
	vars.SniperScore = ScoreSniperStart

	gunFire := spin.NewSequencer().
		Do(spin.PlaySound{ID: SoundGunLoadSniper}).
		Sleep(1_500).
		Do(spin.PlaySound{ID: SoundGunFire}).
		Sleep(1_500).
		Loop()

	e.NewCoroutine(e.Context(), func(e spin.Env) {
		spin.NewSequencer().
			Do(spin.PlaySpeech{ID: SpeechSniperIsShootingIntoCrowdFromJohnsonTower, Duck: 0.5}).
			Sleep(4_000).
			Do(spin.PlaySpeech{ID: SpeechShootSniperTower}).
			Sleep(1_000).
			Func(func() {
				spin.ScoreHurryUpScript(e,
					&vars.SniperScore,
					160, // tick ms
					ScoreSniperDec,
					ScoreSniperEnd,
					spin.TimeoutEvent{},
				)
			}).
			Sequence(gunFire).
			Run(e)
	})

	e.NewCoroutine(e.Context(), func(e spin.Env) {
		if done := ModeIntroSequence(e, "SNIPER", "SHOOT", "SNIPER TOWER").Run(e); done {
			return
		}
		spin.RenderFrameScript(e, func(e spin.Env) {
			ModeAndScorePanel(e, r, "SNIPER", vars.SniperScore)
		})
		e.WaitFor(spin.Done{})
	})

	e.NewCoroutine(e.Context(), func(e spin.Env) {
		spin.NewSequencer().
			WaitFor(spin.ShotEvent{ID: jd.ShotRightPopper}).
			Post(spin.AdvanceEvent{}).
			Run(e)
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

func sniperMode2Script(e spin.Env) {
	r, _ := e.Display("").Renderer("")
	defer r.Clear()

	vars := GetVars(e)
	vars.Timer = 10
	vars.SniperBonus = vars.SniperScore

	e.NewCoroutine(e.Context(), func(e spin.Env) {
		spin.NewSequencer().
			Do(spin.PlaySound{ID: SoundSuccess, Duck: 0.5}).
			Sleep(1_000).
			Do(spin.DriverPulse{ID: jd.CoilRightPopper}).
			Sleep(1_500).
			Do(spin.PlaySpeech{ID: SpeechShootSniperTower, Notify: true}).
			WaitFor(spin.SpeechFinishedEvent{}).
			Do(spin.PlaySpeech{ID: SpeechAaaaah}).
			Sleep(3_000).
			Do(spin.PlaySpeech{ID: SpeechItsALongWayDown}).
			Sleep(2_500).
			Do(spin.PlaySpeech{ID: SpeechAaaaah}).
			Sleep(3_500).
			Do(spin.PlaySpeech{ID: SpeechICanSeeMyHouseFromHere}).
			Sleep(2_000).
			Do(spin.PlaySpeech{ID: SpeechAaaaah}).
			WaitFor(spin.SpeechFinishedEvent{}).
			Run(e)
	})

	e.NewCoroutine(e.Context(), func(e spin.Env) {
		spin.NewSequencer().
			WaitFor(spin.ShotEvent{ID: jd.ShotRightPopper}).
			Post(spin.AdvanceEvent{}).
			Run(e)
	})

	if done := ModeAndBlinkingScoreSequence(e, r, "SNIPER", vars.SniperScore); done {
		return
	}
	spin.RenderFrameScript(e, func(e spin.Env) {
		TimerAndScorePanel(e, r, "SNIPER", vars.Timer, vars.SniperScore, "")
	})
	spin.CountdownScript(e, &vars.Timer, 1500, spin.TimeoutEvent{})

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
	e.Post(spin.ScriptFinishedEvent{ID: ScriptPursuitMode})
}

func sniperIncompleteScript(e spin.Env) {
	r, _ := e.Display("").Renderer(spin.LayerPriority)
	defer r.Clear()

	vars := GetVars(e)
	e.Do(spin.PlayMusic{ID: MusicMain})
	ModeAndScorePanel(e, r, "SNIPER TOTAL", vars.SniperBonus)

	spin.NewSequencer().
		Do(spin.PlaySound{ID: SoundSniperSplat}).
		Sleep(1_000).
		Do(spin.PlaySpeech{ID: SpeechSniperEliminated, Notify: true}).
		Sleep(2_000).
		Run(e)
}

func sniperCompleteScript(e spin.Env) {
	r, _ := e.Display("").Renderer(spin.LayerPriority)
	defer r.Clear()

	vars := GetVars(e)
	vars.SniperBonus += vars.SniperScore
	e.Do(spin.PlayMusic{ID: MusicMain})
	TimerAndScorePanel(e, r, "SNIPER", vars.Timer, vars.SniperScore, "")

	if done := spin.NewSequencer().
		Do(spin.PlaySound{ID: SoundSniperSplat}).
		Sleep(1_000).
		Do(spin.PlaySpeech{ID: SpeechSniperEliminated, Notify: true}).
		Sleep(2_000).
		Run(e); done {
		return
	}

	e.Do(spin.PlaySound{ID: SoundSuccess})
	ModeAndBlinkingScoreSequence(e, r, "SNIPER", vars.SniperBonus)
}

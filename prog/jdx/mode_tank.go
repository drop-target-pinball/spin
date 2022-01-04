package jdx

import (
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

func tankModeScript(e spin.Env) {
	r, _ := e.Display("").Renderer("")
	defer r.Clear()

	e.Do(spin.PlayMusic{ID: MusicMode2})

	vars := GetVars(e)
	player := spin.GetPlayerVars(e)
	vars.Timer = 30
	vars.TankBonus = ScoreTank0

	tankFire := spin.NewSequencer().
		Do(spin.PlaySound{ID: SoundTankFire, Vol: 100}).
		Sleep(1_500).
		Loop()

	e.NewCoroutine(e.Context(), func(e spin.Env) {
		spin.NewSequencer().
			Do(spin.PlaySpeech{ID: SpeechBattleTankSightedInSectorSix, Notify: true, Duck: 0.5}).
			WaitFor(spin.SpeechFinishedEvent{}).
			Sleep(1_000).
			Sequence(tankFire).
			Run(e)
	})

	e.NewCoroutine(e.Context(), func(e spin.Env) {
		if done := ModeIntroSequence(e, "BATTLE TANK", "SHOOT", "GREEN ARROWS").Run(e); done {
			return
		}
		spin.RenderFrameScript(e, func(e spin.Env) {
			TimerAndScorePanel(e, r, "BATTLE TANK", vars.Timer, player.Score, "SHOOT GREEN ARROWS")
		})
		e.WaitFor(spin.Done{})
	})

	e.NewCoroutine(e.Context(), tankSequenceScript)
	spin.CountdownScript(e, &vars.Timer, 1000, spin.TimeoutEvent{})

	evt, done := e.WaitFor(
		spin.AdvanceEvent{},
		spin.TimeoutEvent{},
	)
	if done {
		return
	}
	if evt == (spin.TimeoutEvent{}) {
		e.Do(spin.PlayScript{ID: ScriptTankIncomplete})
	} else {
		e.Do(spin.PlayScript{ID: ScriptTankComplete})
	}
	e.Post(spin.ScriptFinishedEvent{ID: ScriptTankMode})

}

func tankSequenceScript(e spin.Env) {
	vars := GetVars(e)

	shots := map[interface{}]bool{
		spin.ShotEvent{ID: jd.ShotLeftRamp}:        false,
		spin.ShotEvent{ID: jd.ShotTopLeftRamp}:     false,
		spin.SwitchEvent{ID: jd.SwitchBankTargets}: false,
	}

	vars.TankBonus = ScoreTank0
	hits := 0
	for hits < 3 {
		evt, done := e.WaitFor(
			spin.ShotEvent{ID: jd.ShotLeftRamp},
			spin.ShotEvent{ID: jd.ShotTopLeftRamp},
			spin.SwitchEvent{ID: jd.SwitchBankTargets},
		)
		if done {
			return
		}
		if shots[evt] {
			continue
		}
		hits += 1
		shots[evt] = true
		e.Do(spin.PlayScript{ID: ScriptTankHit})
	}
	vars.TankBonus = ScoreTank3
	e.Post(spin.AdvanceEvent{})
}

func tankHitScript(e spin.Env) {
	vars := GetVars(e)
	vars.TankHits += 1

	var atPercent string
	switch vars.TankHits {
	case 1:
		atPercent = SpeechTwentyFivePercent
		vars.TankBonus = ScoreTank1
	case 2:
		atPercent = SpeechSixtyPercent
		vars.TankBonus = ScoreTank2
	}
	if atPercent == "" {
		return
	}

	spin.NewSequencer().
		Do(spin.PlaySpeech{ID: SpeechBattleTankDamageAt, Notify: true, Duck: 0.5}).
		WaitFor(spin.SpeechFinishedEvent{}).
		Do(spin.PlaySpeech{ID: atPercent, Notify: true, Duck: 0.5}).
		WaitFor(spin.SpeechFinishedEvent{}).
		Run(e)
}

func tankIncompleteScript(e spin.Env) {
	r, _ := e.Display("").Renderer(spin.LayerPriority)
	defer r.Clear()

	vars := GetVars(e)
	e.Do(spin.PlayMusic{ID: MusicMain})
	ModeAndScorePanel(e, r, "BATTLE TANK TOTAL", vars.TankBonus)
	e.Sleep(3_000 * time.Millisecond)
}

func tankCompleteScript(e spin.Env) {
	r, _ := e.Display("").Renderer(spin.LayerPriority)
	defer r.Clear()

	vars := GetVars(e)
	vars.TankHits = 0
	vars.TankBonus = ScoreTank0

	e.Do(spin.PlayMusic{ID: MusicMain})
	ModeAndScorePanel(e, r, "BATTLE TANK TOTAL", vars.TankBonus)

	spin.NewSequencer().
		Do(spin.PlaySound{ID: SoundTankDestroyed}).
		Sleep(1_000).
		Do(spin.PlaySpeech{ID: SpeechBattleTankDestroyed, Notify: true}).
		Sleep(3_000).
		Run(e)
}

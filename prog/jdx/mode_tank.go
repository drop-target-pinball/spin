package jdx

import (
	"context"
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/prog/builtin"
)

const (
	MessageTankAdvance = "jdx.MessageTankAdvance"
	MessageTankTimeout = "jdx.MessageTankTimeout"
)

func tankCountdownFrame(e spin.Env) {
	r, g := e.Display("").Renderer("")
	vars := GetVars(e)

	r.Fill(spin.ColorBlack)
	g.Y = 2
	g.Font = builtin.FontPfArmaFive8
	r.Print(g, "BATTLE TANK")
	g.Y = 12

	g.Font = builtin.Font14x10
	r.Print(g, "%v", vars.TankTimer)
}

func tankDestroyedFrame(e spin.Env) {
	r, g := e.Display("").Renderer("")

	r.Fill(spin.ColorBlack)
	g.Y = 2
	g.Font = builtin.FontPfArmaFive8
	r.Print(g, "BATTLE TANK")
	g.Y = 12

	g.Font = builtin.Font14x10
	r.Print(g, spin.FormatScore("%v", ScoreTank3))
}

func tankCountdownVideoScript(e spin.Env) {
	vars := GetVars(e)
	modeText := [3]string{"BATTLE TANK", "SHOOT", "BATTLE TANK"}
	if done := modeIntroVideo(e, modeText); done {
		return
	}

	vars.TankTimer = 30
	tankCountdownFrame(e)
	if done := e.Sleep(200 * time.Millisecond); done {
		return
	}

	for vars.TankTimer > 0 {
		if done := e.Sleep(1000 * time.Millisecond); done {
			return
		}
		vars.TankTimer -= 1
		tankCountdownFrame(e)
	}
	e.Post(spin.Message{ID: MessageTankTimeout})
}

func tankCountdownAudioScript(e spin.Env) {
	e.Do(spin.PlayMusic{ID: MusicMode2})
	e.Do(spin.PlaySpeech{ID: SpeechBattleTankSightedInSectorSix, Notify: true})
	if _, done := e.WaitFor(spin.SpeechFinishedEvent{}); done {
		e.Do(spin.StopSpeech{ID: SpeechBattleTankSightedInSectorSix})
		return
	}

	if done := e.Sleep(1000 * time.Millisecond); done {
		return
	}
	for {
		e.Do(spin.PlaySound{ID: SoundTankFire})
		if done := e.Sleep(1500 * time.Millisecond); done {
			return
		}
	}
}

func tankDamageAudioScript(e spin.Env, nHits int) {
	e.Do(spin.PlaySpeech{ID: SpeechBattleTankDamageAt, Notify: true})
	e.Do(spin.MusicVolume{Mul: 0.5})
	defer e.Do(spin.MusicVolume{Mul: 2.0})
	if _, done := e.WaitFor(spin.SpeechFinishedEvent{}); done {
		e.Do(spin.StopSpeech{ID: SpeechBattleTankDamageAt})
		return
	}
	switch nHits {
	case 1:
		e.Do(spin.PlaySpeech{ID: SpeechTwentyFivePercent, Notify: true})
		if _, done := e.WaitFor(spin.SpeechFinishedEvent{}); done {
			e.Do(spin.StopSpeech{ID: SpeechTwentyFivePercent})
			return
		}
	case 2:
		e.Do(spin.PlaySpeech{ID: SpeechSixtyPercent, Notify: true})
		if _, done := e.WaitFor(spin.SpeechFinishedEvent{}); done {
			e.Do(spin.StopSpeech{ID: SpeechSixtyPercent})
			return
		}
	}
}

func tankSequenceScript(e spin.Env) {
	var ctx context.Context
	var cancel context.CancelFunc
	vars := GetVars(e)

	shots := map[interface{}]bool{
		spin.ShotEvent{ID: jd.ShotLeftRamp}:        false,
		spin.ShotEvent{ID: jd.ShotTopLeftRamp}:     false,
		spin.SwitchEvent{ID: jd.SwitchBankTargets}: false,
	}
	scores := map[int]int{
		1: ScoreTank1,
		2: ScoreTank2,
	}

	var nHits int

	vars.TankBonus = ScoreTank0
	for {
		evt, done := e.WaitFor(
			spin.ShotEvent{ID: jd.ShotLeftRamp},
			spin.ShotEvent{ID: jd.ShotTopLeftRamp},
			spin.SwitchEvent{ID: jd.SwitchBankTargets},
		)

		if cancel != nil {
			cancel()
		}
		if done {
			return
		}
		hit := shots[evt]
		if hit {
			continue
		}
		nHits += 1
		if nHits == 3 {
			break
		}
		shots[evt] = true
		vars.TankBonus = scores[nHits]
		ctx, cancel = e.Derive()
		e.NewCoroutine(ctx, func(e spin.Env) { tankDamageAudioScript(e, nHits) })
	}
	vars.TankBonus = ScoreTank3
	e.Post(spin.Message{ID: MessageTankAdvance})
}

func tankCountdownScript(e spin.Env) {
	ctx, cancel := e.Derive()
	defer cancel()

	e.NewCoroutine(ctx, tankCountdownAudioScript)
	e.NewCoroutine(ctx, tankCountdownVideoScript)
	e.NewCoroutine(ctx, tankSequenceScript)
	e.WaitFor(spin.Done{})
}

func tankDestroyedScript(e spin.Env) {
	tankDestroyedFrame(e)

	e.Do(spin.PlaySound{ID: SoundTankDestroyed})
	if done := e.Sleep(1000 * time.Millisecond); done {
		return
	}

	e.Do(spin.PlaySpeech{ID: SpeechBattleTankDestroyed, Notify: true})
	if _, done := e.WaitFor(spin.SpeechFinishedEvent{}); done {
		e.Do(spin.StopSpeech{ID: SpeechBattleTankDestroyed})
		return
	}
	e.Post(spin.Message{ID: MessageTankAdvance})
}

func tankModeScript(e spin.Env) {
	e.Do(spin.PlayScript{ID: ScriptTankCountdown})
	evt, done := e.WaitFor(
		spin.Message{ID: MessageTankAdvance},
		spin.Message{ID: MessageTankTimeout},
	)
	e.Do(spin.StopScript{ID: ScriptTankCountdown})
	if done {
		return
	}
	e.Do(spin.PlayMusic{ID: MusicMain})
	if evt == (spin.Message{ID: MessageTankAdvance}) {
		e.Do(spin.PlayScript{ID: ScriptTankDestroyed})
		if _, done := e.WaitFor(spin.Message{ID: MessageTankAdvance}); done {
			return
		}
	}
	e.Post(spin.ScriptFinishedEvent{ID: ScriptTankMode})
}

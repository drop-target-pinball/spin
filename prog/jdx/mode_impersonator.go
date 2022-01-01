package jdx

import (
	"math/rand"
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/prog/builtin"
)

const (
	MessageImpersonatorHit = "jdx.MessageImpersonatorHit"
)

var crowdSounds = []string{
	SpeechBoo,
	SpeechYouSuck,
	SpeechBoo,
	SpeechGoHome,
}

var hitSounds = []string{
	SoundBadImpersonatorGunFire,
	SoundBadImpersonatorThrow,
	SoundShock,
}

func impersonatorCountdownFrame(e spin.Env) {
	r, g := e.Display("").Renderer("")
	vars := GetVars(e)

	r.Fill(spin.ColorBlack)
	g.Y = 2
	g.Font = builtin.FontPfArmaFive8
	r.Print(g, "BAD IMPERSONATOR")
	g.Y = 12

	g.Font = builtin.Font14x10
	r.Print(g, "%v", vars.Timer)
}

func impersonatorTotalFrame(e spin.Env) {
	r, g := e.Display("").Renderer("")
	vars := GetVars(e)

	r.Fill(spin.ColorBlack)
	g.Y = 2
	g.Font = builtin.FontPfArmaFive8
	r.Print(g, "IMPERSONATOR TOTAL")
	g.Y = 12

	g.Font = builtin.Font14x10
	r.Print(g, spin.FormatScore("%v", vars.BadImpersonatorBonus))
}

func impersonatorCountdownVideoScript(e spin.Env) {
	vars := GetVars(e)

	modeText := [3]string{"BAD IMPERSONATOR", "SHOOT LIT", "DROP TARGETS"}
	if done := modeIntroVideo(e, modeText); done {
		e.Do(spin.StopSpeech{ID: SpeechCivilDisorderHasEruptedInHeitschMusicHall})
		return
	}

	vars.Timer = 25
	for vars.Timer > 0 {
		impersonatorCountdownFrame(e)
		if done := e.Sleep(1000 * time.Millisecond); done {
			return
		}
		vars.Timer -= 1
	}
	impersonatorCountdownFrame(e)
	e.Post(spin.TimeoutEvent{ID: ScriptBadImpersonatorMode})
}

func impersonatorIntroAudioScript(e spin.Env) {
	e.Do(spin.PlayMusic{ID: MusicBadImpersonator})
	e.Do(spin.MusicVolume{Mul: 0.5})
	e.Do(spin.PlaySpeech{ID: SpeechCivilDisorderHasEruptedInHeitschMusicHall, Notify: true})
	_, done := e.WaitFor(spin.SpeechFinishedEvent{})
	e.Do(spin.MusicVolume{Mul: 2})
	if done {
		e.Do(spin.StopSpeech{ID: SpeechCivilDisorderHasEruptedInHeitschMusicHall})
	} else {
		e.Do(spin.PlayScript{ID: ScriptBadImpersonatorCountdownAudio})
	}
}

func impersonatorCountdownAudioScript(e spin.Env) {
	if done := e.Sleep(2000 * time.Millisecond); done {
		return
	}
	sound := 0
	for {
		t := rand.Intn(3000) + 1500
		e.Do(spin.PlaySpeech{ID: crowdSounds[sound]})
		if done := e.Sleep(time.Duration(t) * time.Millisecond); done {
			e.Do(spin.StopSpeech{ID: crowdSounds[sound]})
			return
		}
		sound += 1
		if sound >= len(crowdSounds) {
			sound = 0
		}
	}
}

func impersonatorLightDropTargets(e spin.Env) {
	vars := GetVars(e)
	vars.BadImpersonatorTargets = jd.DropTargetJ | jd.DropTargetU

	longWait := time.Duration(3000 * time.Millisecond)
	shortWait := time.Duration(1000 * time.Millisecond)

	wait := longWait
	left := true
	for {
		for i, lamp := range jd.DropTargetLamps {
			if vars.BadImpersonatorTargets&(1<<i) != 0 {
				e.Do(spin.DriverOn{ID: lamp})
			} else {
				e.Do(spin.DriverOff{ID: lamp})
			}
		}
		if done := e.Sleep(wait); done {
			return
		}
		if left {
			vars.BadImpersonatorTargets <<= 1
			if vars.BadImpersonatorTargets == jd.DropTargetG|jd.DropTargetE {
				left = false
				wait = longWait
			} else {
				wait = shortWait
			}
		} else {
			vars.BadImpersonatorTargets >>= 1
			if vars.BadImpersonatorTargets == jd.DropTargetJ|jd.DropTargetU {
				left = true
				wait = longWait
			} else {
				wait = shortWait
			}
		}
	}
}

func impersonatorHitScript(e spin.Env) {
	vars := GetVars(e)
	vars.BadImpersonatorBonus += ScoreBadImpersonatorN

	e.Do(spin.StopMusic{ID: MusicBadImpersonator})
	e.Do(spin.StopScript{ID: ScriptBadImpersonatorCountdownAudio})
	sound := rand.Intn(len(hitSounds))
	e.Do(spin.PlaySound{ID: hitSounds[sound]})
	if done := e.Sleep(1000 * time.Millisecond); done {
		return
	}
	e.Do(spin.PlayMusic{ID: MusicBadImpersonator})
	e.Do(spin.PlayScript{ID: ScriptBadImpersonatorCountdownAudio})
}

func impersonatorWatchDropTargets(e spin.Env) {
	vars := GetVars(e)
	for {
		evt, done := e.WaitFor(jd.SwitchAnyDropTarget...)
		if done {
			return
		}
		switchEvt := evt.(spin.SwitchEvent)
		idx := jd.DropTargetIndexes[switchEvt.ID]
		if vars.BadImpersonatorTargets&(1<<idx) != 0 {
			e.Do(spin.PlayScript{ID: ScriptBadImpersonatorHit})
		}
	}
}

func impersonatorCompleteScript(e spin.Env) {
	e.Do(spin.PlayMusic{ID: MusicMain})
	e.Do(spin.MusicVolume{Mul: 0.5})
	defer e.Do(spin.MusicVolume{Mul: 2.0})

	e.Do(spin.PlaySound{ID: SoundSuccess})
	impersonatorTotalFrame(e)
	if done := e.Sleep(2000 * time.Millisecond); done {
		return
	}
	e.Post(spin.AdvanceEvent{ID: ScriptBadImpersonatorMode})
}

func impersonatorCountdownScript(e spin.Env) {
	ctx, cancel := e.Derive()

	e.NewCoroutine(ctx, impersonatorIntroAudioScript)
	e.NewCoroutine(ctx, impersonatorCountdownVideoScript)
	e.NewCoroutine(ctx, impersonatorLightDropTargets)
	e.NewCoroutine(ctx, impersonatorWatchDropTargets)

	e.WaitFor(spin.TimeoutEvent{ID: ScriptBadImpersonatorMode})
	e.Do(spin.StopScript{ID: ScriptBadImpersonatorCountdownAudio})
	cancel()
}

func impersonatorModeScript(e spin.Env) {
	vars := GetVars(e)
	vars.BadImpersonatorBonus = ScoreBadImpersonator0

	e.Do(spin.PlayScript{ID: ScriptBadImpersonatorCountdown})
	if _, done := e.WaitFor(spin.TimeoutEvent{ID: ScriptBadImpersonatorMode}); done {
		return
	}

	e.Do(spin.PlayScript{ID: ScriptBadImpersonatorComplete})
	if _, done := e.WaitFor(spin.AdvanceEvent{ID: ScriptBadImpersonatorMode}); done {
		return
	}
	e.Post(spin.ScriptFinishedEvent{ID: ScriptBadImpersonatorMode})
}

package jdx

import (
	"math/rand"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

var hitSounds = []string{
	SoundBadImpersonatorGunFire,
	SoundBadImpersonatorThrow,
	SoundShock,
}

func impersonatorModeScript(e *spin.ScriptEnv) {
	r, _ := e.Display("").Renderer("")

	vars := GetVars(e)
	player := spin.GetPlayerVars(e)

	vars.Mode = ModeBadImpersonator
	defer func() { vars.Mode = ModeNone }()
	vars.Timer = 30
	vars.BadImpersonatorBonus = ScoreBadImpersonator0

	e.Do(spin.PlayScript{ID: ScriptBadImpersonatorCrowd})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)
		s.Do(spin.PlaySpeech{ID: SpeechCivilDisorderHasEruptedInHeitschMusicHall, Notify: true, Duck: 0.5})
		s.WaitFor(spin.SpeechFinishedEvent{})
		s.Run()
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		if done := ModeIntroScript(e, "BAD IMPERSONATOR", "SHOOT LIT", "DROP TARGETS"); done {
			return
		}
		spin.RenderFrameLoop(e, func(e *spin.ScriptEnv) {
			TimerAndScorePanel(e, r, "BAD IMPERSONATOR", vars.Timer, player.Score, "SHOOT LIT DROP TARGETS")
		})
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		spin.CountdownLoop(e, &vars.Timer, 1000, spin.TimeoutEvent{})
	})

	e.NewCoroutine(impersonatorLightDropTargets)
	e.NewCoroutine(impersonatorWatchDropTargets)

	if _, done := e.WaitFor(spin.TimeoutEvent{}); done {
		return
	}
	e.Do(spin.StopScript{ID: ScriptBadImpersonatorCrowd})
	e.Do(spin.PlayScript{ID: ScriptBadImpersonatorComplete})
	e.Post(spin.ScriptFinishedEvent{ID: ScriptBadImpersonatorMode})
}

func impersonatorCrowdScript(e *spin.ScriptEnv) {
	e.Do(spin.PlayMusic{ID: MusicBadImpersonator})

	s := spin.NewSequencer(e)

	s.Sleep(4_000)
	s.Do(spin.PlaySpeech{ID: SpeechBoo})
	s.Sleep(4_000)
	s.Do(spin.PlaySpeech{ID: SpeechYouSuck})
	s.Sleep(4_000)
	s.Do(spin.PlaySpeech{ID: SpeechBoo})
	s.Sleep(4_000)
	s.Do(spin.PlaySpeech{ID: SpeechGoHome})
	s.Loop()

	s.Run()
}

func impersonatorLightDropTargets(e *spin.ScriptEnv) {
	vars := GetVars(e)
	vars.BadImpersonatorTargets = jd.DropTargetJ | jd.DropTargetU

	longWait := 3000
	shortWait := 1000

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

func impersonatorWatchDropTargets(e *spin.ScriptEnv) {
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

func impersonatorHitScript(e *spin.ScriptEnv) {
	vars := GetVars(e)
	vars.BadImpersonatorBonus += ScoreBadImpersonatorN
	sound := rand.Intn(len(hitSounds))

	s := spin.NewSequencer(e)

	s.Do(spin.StopMusic{ID: MusicBadImpersonator})
	s.Do(spin.StopScript{ID: ScriptBadImpersonatorCrowd})
	s.Do(spin.PlaySound{ID: hitSounds[sound]})
	s.Sleep(1000)
	s.Do(spin.PlayMusic{ID: MusicBadImpersonator})
	s.Do(spin.PlayScript{ID: ScriptBadImpersonatorCrowd})
	s.Run()
}

func impersonatorCompleteScript(e *spin.ScriptEnv) {
	r, _ := e.Display("").Renderer(spin.LayerPriority)
	defer r.Clear()

	e.Do(spin.PlayMusic{ID: MusicMain})

	vars := GetVars(e)
	ModeAndScorePanel(e, r, "BAD IMPERSONATOR TOTAL", vars.BadImpersonatorBonus)

	s := spin.NewSequencer(e)
	s.Do(spin.PlaySound{ID: SoundSuccess, Duck: 0.5})
	s.Sleep(2_000)
	s.Run()
}

package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

func stakeoutModeScript(e *spin.ScriptEnv) {
	r, _ := e.Display("").Renderer("")

	e.Do(spin.PlayMusic{ID: MusicMode2})

	vars := GetVars(e)
	player := spin.GetPlayerVars(e)
	vars.Timer = 30
	vars.ManhuntBonus = ScoreStakeout0

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.Do(spin.PlaySpeech{ID: SpeechImStakingOutACrackHouseInSectorTwentyThree})
		s.Sleep(15_000)
		s.Do(spin.PlaySpeech{ID: SpeechShootRightRamp})
		s.Sleep(10_000)
		s.Do(spin.PlaySpeech{ID: SpeechShootRightRamp})

		s.Run()
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		ModeIntroScript(e, "STAKEOUT", "SHOOT", "RIGHT RAMP")
		spin.RenderFrameLoop(e, func(e *spin.ScriptEnv) {
			TimerAndScorePanel(e, r, "STAKEOUT", vars.Timer, player.Score, "SHOOT RIGHT RAMP")
		})
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		spin.CountdownLoop(e, &vars.Timer, 1000, spin.TimeoutEvent{})
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.WaitFor(spin.ShotEvent{ID: jd.ShotRightRamp})
		s.Do(spin.PlayScript{ID: ScriptStakeoutInteresting})
		s.Loop()

		s.Run()
	})

	if _, done := e.WaitFor(spin.TimeoutEvent{}); done {
		return
	}
	e.Do(spin.PlayScript{ID: ScriptStakeoutComplete})
	e.Post(spin.ScriptFinishedEvent{ID: ScriptStakeoutMode})
}

func stakeoutInterestingScript(e *spin.ScriptEnv) {
	r, _ := e.Display("").Renderer(spin.LayerPriority)
	defer r.Clear()

	vars := GetVars(e)

	callouts := []string{
		SpeechIWonderWhatsOverThere,
		SpeechIWonderWhatsDownThere,
	}
	callout := callouts[vars.StakeoutCallout]

	vars.StakeoutCallout += 1
	if vars.StakeoutCallout >= len(callouts) {
		vars.StakeoutCallout = 0
	}
	vars.StakeoutBonus += ScoreStakeoutN

	ScoreAndLabelPanel(e, r, ScoreStakeoutN, "AWARDED")

	s := spin.NewSequencer(e)
	s.Do(spin.PlaySpeech{ID: callout, Priority: spin.PriorityAudioModeCallout})
	s.Sleep(2_500)
	s.Do(spin.PlaySpeech{ID: SpeechInteresting, Priority: spin.PriorityAudioModeCallout})
	s.Run()
}

func stakeoutCompleteScript(e *spin.ScriptEnv) {
	r, _ := e.Display("").Renderer(spin.LayerPriority)
	defer r.Clear()

	e.Do(spin.PlayMusic{ID: MusicMain})

	vars := GetVars(e)

	s := spin.NewSequencer(e)
	s.Sleep(1_000)
	s.DoFunc(func() { ModeAndScorePanel(e, r, "STAKEOUT TOTAL", vars.StakeoutBonus) })
	s.Do(spin.PlaySound{ID: SoundSuccess, Duck: 0.5})
	s.Sleep(3_000)
	s.Run()
}

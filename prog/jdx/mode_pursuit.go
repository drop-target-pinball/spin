package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

func pursuitCompleteScript(e spin.Env) {
	r, _ := e.Display("").Renderer(spin.LayerPriority)
	defer r.Clear()

	e.Do(spin.PlayMusic{ID: MusicMain})

	vars := GetVars(e)
	modeAndScorePanel(e, r, "PURSUIT TOTAL", vars.PursuitBonus)

	spin.NewSequencer().
		Do(spin.PlaySound{ID: SoundPursuitMissile}).
		Sleep(500).
		Do(spin.PlaySound{ID: SoundPursuitExplosion}).
		Sleep(1_000).
		Do(spin.PlaySpeech{ID: SpeechYourDrivingDaysAreOverPunk, Notify: true}).
		WaitFor(spin.SpeechFinishedEvent{}).
		Sleep(1_000).
		Run(e)
}

func pursuitIncompleteScript(e spin.Env) {
	r, _ := e.Display("").Renderer(spin.LayerPriority)
	defer r.Clear()

	e.Do(spin.PlayMusic{ID: MusicMain})

	vars := GetVars(e)
	modeAndScorePanel(e, r, "PURSUIT TOTAL", vars.PursuitBonus)

	spin.NewSequencer().
		Do(spin.PlaySpeech{ID: SpeechDreddToControl, Notify: true}).
		WaitFor(spin.SpeechFinishedEvent{}).
		Sleep(1_000).
		Do(spin.PlaySpeech{ID: SpeechSuspectGotAway, Notify: true}).
		WaitFor(spin.SpeechFinishedEvent{}).
		Run(e)
}

func pursuitModeScript(e spin.Env) {
	r, _ := e.Display("").Renderer("")
	defer r.Clear()

	e.Do(spin.PlayMusic{ID: MusicMode2})

	vars := GetVars(e)
	vars.Timer = 30
	vars.PursuitBonus = ScorePursuit0

	e.NewCoroutine(e.Context(), func(e spin.Env) {
		spin.NewSequencer().
			Do(spin.PlaySpeech{ID: SpeechDreddToControl, Notify: true}).
			WaitFor(spin.SpeechFinishedEvent{}).
			Do(spin.PlaySpeech{ID: SpeechImInPursuitOfAStolenVehicle, Notify: true}).
			WaitFor(spin.SpeechFinishedEvent{}).
			Sleep(500).
			Do(spin.PlaySound{ID: SoundPoliceSiren, Loop: true}).
			Do(spin.StopSound{ID: SoundPursuitEngine}).
			Defer(spin.StopSound{ID: SoundPoliceSiren}).
			Defer(spin.StopSound{ID: SoundPursuitEngine}).
			Sleep(3_000).
			Do(spin.PlaySound{ID: SoundMotorRev}).
			Sleep(2_000).
			Do(spin.PlaySound{ID: SoundMotorRev}).
			Sleep(4_000).
			Do(spin.PlaySound{ID: SoundTireSqueal1}).
			Sleep(4_000).
			Do(spin.PlaySound{ID: SoundMotorRev}).
			Sleep(4_000).
			Do(spin.PlaySound{ID: SoundTireSqueal2}).
			Sleep(4_000).
			Do(spin.PlaySound{ID: SoundTireSqueal1}).
			Sleep(4_000).
			Do(spin.PlaySound{ID: SoundMotorRev}).
			WaitFor(spin.Done{}).
			Run(e)
	})

	e.NewCoroutine(e.Context(), func(e spin.Env) {
		spin.NewSequencer().
			WaitFor(spin.ShotEvent{ID: jd.ShotRightRamp}).
			SetIntVar(&vars.PursuitBonus, ScorePursuit1).
			Do(spin.PlaySound{ID: SoundPursuitMissile}).
			WaitFor(spin.ShotEvent{ID: jd.ShotLeftRamp}).
			SetIntVar(&vars.PursuitBonus, ScorePursuit2).
			Do(spin.PlaySound{ID: SoundPursuitMissile}).
			WaitFor(spin.ShotEvent{ID: jd.ShotRightRamp}).
			SetIntVar(&vars.PursuitBonus, ScorePursuit3).
			Post(spin.AdvanceEvent{}).
			Run(e)
	})

	spin.CountdownScript(e, &vars.Timer, 1000, spin.TimeoutEvent{})
	modeText := [3]string{"PURSUIT", "SHOOT", "FLASHING RAMP"}
	if done := modeIntroVideo(e, modeText); done {
		return
	}

	spin.RenderFrameScript(e, func(e spin.Env) {
		timerAndScorePanel(e, r, "PURSUIT", "SHOOT FLASHING RAMP")
	})

	evt, done := e.WaitFor(
		spin.AdvanceEvent{},
		spin.TimeoutEvent{})
	if done {
		return
	}
	if evt == (spin.TimeoutEvent{}) {
		e.Do(spin.PlayScript{ID: ScriptPursuitIncomplete})
	} else {
		e.Do(spin.PlayScript{ID: ScriptPursuitComplete})
	}
	e.Post(spin.ScriptFinishedEvent{ID: ScriptPursuitMode})
}

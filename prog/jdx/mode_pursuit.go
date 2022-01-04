package jdx

import (
	"math/rand"
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

const (
	MessagePursuitAdvance = "jdx.MessagePursuitAdvance"
	MessagePursuitTimeout = "jdx.MessagePursuitTimeout"
)

var pursuitSounds = []string{
	SoundMotorRev,
	SoundTireSqueal1,
	SoundTireSqueal2,
}

func pursuitCompleteScript(e spin.Env) {
	defer e.Display("").Clear(spin.LayerPriority)

	vars := GetVars(e)
	modeTotalPanel(e, "PURSUIT TOTAL", vars.PursuitBonus)
	e.Do(spin.PlayMusic{ID: MusicMain})

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
	defer e.Display("").Clear(spin.LayerPriority)

	vars := GetVars(e)
	modeTotalPanel(e, "PURSUIT TOTAL", vars.PursuitBonus)
	e.Do(spin.PlayMusic{ID: MusicMain})

	spin.NewSequencer().
		Do(spin.PlaySpeech{ID: SpeechDreddToControl, Notify: true}).
		WaitFor(spin.SpeechFinishedEvent{}).
		Sleep(1_000).
		Do(spin.PlaySpeech{ID: SpeechSuspectGotAway, Notify: true}).
		WaitFor(spin.SpeechFinishedEvent{}).
		Run(e)
}

func pursuitModeScript(e spin.Env) {
	defer e.Display("").Clear("")
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
			Sleep(1_000).
			Func(func() {
				for {
					t := rand.Intn(3000) + 1500
					sound := rand.Intn(len(pursuitSounds))

					e.Do(spin.PlaySound{ID: pursuitSounds[sound]})
					if done := e.Sleep(time.Duration(t) * time.Millisecond); done {
						e.Do(spin.StopSound{ID: pursuitSounds[sound]})
						return
					}
				}
			}).
			Run(e)
	})

	e.NewCoroutine(e.Context(), func(e spin.Env) {
		spin.NewSequencer().
			WaitFor(spin.ShotEvent{ID: jd.ShotRightRamp}).
			Do(spin.SetIntVar{Var: &vars.PursuitBonus, Val: ScorePursuit1}).
			Do(spin.PlaySound{ID: SoundPursuitMissile}).
			WaitFor(spin.ShotEvent{ID: jd.ShotLeftRamp}).
			Do(spin.SetIntVar{Var: &vars.PursuitBonus, Val: ScorePursuit2}).
			Do(spin.PlaySound{ID: SoundPursuitMissile}).
			WaitFor(spin.ShotEvent{ID: jd.ShotRightRamp}).
			Do(spin.SetIntVar{Var: &vars.PursuitBonus, Val: ScorePursuit3}).
			Post(spin.AdvanceEvent{}).
			Run(e)
	})

	spin.CountdownScript(e, &vars.Timer, 1000, spin.TimeoutEvent{})
	modeText := [3]string{"PURSUIT", "SHOOT", "FLASHING RAMP"}
	if done := modeIntroVideo(e, modeText); done {
		return
	}

	spin.RenderFrameScript(e, func(e spin.Env) {
		timerAndScorePanel(e, "PURSUIT", "SHOOT FLASHING RAMP")
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

package jdx

import (
	"math/rand"
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/prog/builtin"
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

func pursuitCountdownFrame(e spin.Env, seconds int) {
	r, g := e.Display("").Renderer("")

	r.Fill(spin.ColorBlack)
	g.Y = 2
	g.Font = builtin.FontPfArmaFive8
	r.Print(g, "PURSUIT")
	g.Y = 12

	g.Font = builtin.Font14x10
	r.Print(g, "%v", seconds)
}

func pursuitCaughtFrame(e spin.Env) {
	r, g := e.Display("").Renderer("")

	r.Fill(spin.ColorBlack)
	g.Y = 2
	g.Font = builtin.FontPfArmaFive8
	r.Print(g, "PURSUIT")
	g.Y = 12

	g.Font = builtin.Font14x10
	r.Print(g, spin.FormatScore("%v", ScorePursuit3))
}

func pursuitCountdownVideoScript(e spin.Env) {
	modeText := [3]string{"PURSUIT", "SHOOT", "LIT RAMPS"}
	if done := modeIntroVideo(e, modeText); done {
		return
	}

	seconds := 30
	pursuitCountdownFrame(e, seconds)
	if done := e.Sleep(200 * time.Millisecond); done {
		return
	}

	for seconds > 0 {
		if done := e.Sleep(1000 * time.Millisecond); done {
			return
		}
		seconds -= 1
		pursuitCountdownFrame(e, seconds)
	}
	e.Post(spin.Message{ID: MessagePursuitTimeout})
}

func pursuitCountdownAudioScript(e spin.Env) {
	e.Do(spin.PlayMusic{ID: MusicMode2})
	e.Do(spin.PlaySpeech{ID: SpeechDreddToControl, Notify: true})
	if _, done := e.WaitFor(spin.SpeechFinishedEvent{}); done {
		e.Do(spin.StopSpeech{ID: SpeechDreddToControl})
		return
	}

	e.Do(spin.PlaySpeech{ID: SpeechImInPursuitOfAStolenVehicle, Notify: true})
	if _, done := e.WaitFor(spin.SpeechFinishedEvent{}); done {
		e.Do(spin.StopSpeech{ID: SpeechImInPursuitOfAStolenVehicle})
		return
	}

	if done := e.Sleep(300 * time.Millisecond); done {
		return
	}

	e.Do(spin.PlaySound{ID: SoundPoliceSiren, Loops: -1})
	defer e.Do(spin.StopSound{ID: SoundPoliceSiren})
	e.Do(spin.PlaySound{ID: SoundPursuitEngine})
	defer e.Do(spin.StopSound{ID: SoundPursuitEngine})
	if done := e.Sleep(1000 * time.Millisecond); done {
		return
	}

	for {
		t := rand.Intn(3000) + 1500
		sound := rand.Intn(len(pursuitSounds))

		e.Do(spin.PlaySound{ID: pursuitSounds[sound]})
		if done := e.Sleep(time.Duration(t) * time.Millisecond); done {
			e.Do(spin.StopSound{ID: pursuitSounds[sound]})
			return
		}
	}
}

func pursuitSequenceScript(e spin.Env) {
	vars := GetVars(e)

	vars.PursuitBonus = ScorePursuit0
	if _, done := e.WaitFor(spin.ShotEvent{ID: jd.ShotRightRamp}); done {
		return
	}

	e.Do(spin.PlaySound{ID: SoundPursuitMissile})
	vars.PursuitBonus = ScorePursuit1
	if _, done := e.WaitFor(spin.ShotEvent{ID: jd.ShotLeftRamp}); done {
		return
	}

	e.Do(spin.PlaySound{ID: SoundPursuitMissile})
	vars.PursuitBonus = ScorePursuit2
	if _, done := e.WaitFor(spin.ShotEvent{ID: jd.ShotRightRamp}); done {
		return
	}

	e.Do(spin.PlaySound{ID: SoundPursuitMissile})
	vars.PursuitBonus = ScorePursuit3

	if done := e.Sleep(500 * time.Millisecond); done {
		return
	}
	e.Post(spin.Message{ID: MessagePursuitAdvance})
}

func pursuitCountdownScript(e spin.Env) {
	ctx, cancel := e.Derive()
	defer cancel()

	e.NewCoroutine(ctx, pursuitCountdownAudioScript)
	e.NewCoroutine(ctx, pursuitCountdownVideoScript)
	e.NewCoroutine(ctx, pursuitSequenceScript)
	e.WaitFor(spin.Done{})
}

func pursuitEscapeScript(e spin.Env) {
	e.Do(spin.PlayMusic{ID: MusicMain})
	e.Do(spin.PlaySpeech{ID: SpeechDreddToControl, Notify: true})
	if _, done := e.WaitFor(spin.SpeechFinishedEvent{}); done {
		e.Do(spin.StopSpeech{ID: SpeechDreddToControl})
		return
	}

	if done := e.Sleep(500 * time.Millisecond); done {
		return
	}

	e.Do(spin.PlaySpeech{ID: SpeechSuspectGotAway, Notify: true})
	if _, done := e.WaitFor(spin.SpeechFinishedEvent{}); done {
		e.Do(spin.StopSpeech{ID: SpeechSuspectGotAway})
		return
	}
	e.Post(spin.Message{ID: MessagePursuitAdvance})
}

func pursuitCaughtScript(e spin.Env) {
	e.Do(spin.PlayMusic{ID: MusicMain})
	pursuitCaughtFrame(e)

	e.Do(spin.PlaySound{ID: SoundPursuitExplosion})
	if done := e.Sleep(1000 * time.Millisecond); done {
		return
	}

	e.Do(spin.PlaySpeech{ID: SpeechYourDrivingDaysAreOverPunk, Notify: true})
	if _, done := e.WaitFor(spin.SpeechFinishedEvent{}); done {
		e.Do(spin.StopSpeech{ID: SpeechYourDrivingDaysAreOverPunk})
		return
	}
	e.Post(spin.Message{ID: MessagePursuitAdvance})
}

func pursuitModeScript(e spin.Env) {
	e.Do(spin.PlayScript{ID: ScriptPursuitCountdown})
	evt, done := e.WaitFor(
		spin.Message{ID: MessagePursuitAdvance},
		spin.Message{ID: MessagePursuitTimeout},
	)
	e.Do(spin.StopScript{ID: ScriptPursuitCountdown})
	if done {
		return
	}
	if evt == (spin.Message{ID: MessagePursuitAdvance}) {
		e.Do(spin.PlayScript{ID: ScriptPursuitCaught})
	} else {
		e.Do(spin.PlayScript{ID: ScriptPursuitEscape})
	}

	if _, done := e.WaitFor(spin.Message{ID: MessagePursuitAdvance}); done {
		return
	}
	e.Post(spin.ScriptFinishedEvent{ID: ScriptPursuitMode})
}

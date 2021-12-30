package jdx

import (
	"math/rand"
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/prog/builtin"
)

func meltdownCountdownFrame(e spin.Env) {
	r, g := e.Display("").Renderer("")
	vars := GetVars(e)

	r.Fill(spin.ColorBlack)
	g.Y = 2
	g.Font = builtin.FontPfArmaFive8
	r.Print(g, "MELTDOWN")
	g.Y = 12

	g.Font = builtin.Font14x10
	r.Print(g, "%v", vars.Timer)
}

func meltdownFailureFrame(e spin.Env) {
	r, g := e.Display("").Renderer("")
	vars := GetVars(e)

	r.Fill(spin.ColorBlack)
	g.Y = 2
	g.Font = builtin.FontPfArmaFive8
	r.Print(g, "MELTDOWN TOTAL")
	g.Y = 12

	g.Font = builtin.Font14x10
	r.Print(g, spin.FormatScore("%v", vars.MeltdownBonus))
}

func meltdownSequenceScript(e spin.Env) {
	vars := GetVars(e)
	if _, done := e.WaitFor(spin.SwitchEvent{ID: jd.SwitchCaptiveBall1}); done {
		return
	}

	vars.MeltdownBonus = ScoreMeltdown1
	e.Do(spin.PlaySpeech{ID: SpeechReactorOneStabilized})
	if _, done := e.WaitFor(spin.SwitchEvent{ID: jd.SwitchCaptiveBall2}); done {
		return
	}

	vars.MeltdownBonus = ScoreMeltdown2
	e.Do(spin.PlaySpeech{ID: SpeechReactorTwoStabilized})
	if _, done := e.WaitFor(spin.SwitchEvent{ID: jd.SwitchCaptiveBall3}); done {
		return
	}

	vars.MeltdownBonus = ScoreMeltdown3
	e.Post(spin.AdvanceEvent{ID: ScriptMeltdownMode})
}

func meltdownCountdownScript(e spin.Env) {
	ctx, cancel := e.Derive()
	defer cancel()

	e.NewCoroutine(ctx, meltdownCountdownAudioScript)
	e.NewCoroutine(ctx, meltdownSequenceScript)

	vars := GetVars(e)
	modeText := [3]string{"MELTDOWN", "SHOOT", "CAPTIVE BALLS"}
	if done := modeIntroVideo(e, modeText); done {
		return
	}

	vars.MeltdownBonus = ScoreMeltdown0
	vars.Timer = 30
	meltdownCountdownFrame(e)
	if done := e.Sleep(200 * time.Millisecond); done {
		return
	}

	for vars.Timer > 0 {
		if done := e.Sleep(1000 * time.Millisecond); done {
			return
		}
		vars.Timer -= 1
		meltdownCountdownFrame(e)

		if vars.Timer != 0 {
			e.Do(spin.PlaySound{ID: SoundMeltdownKlaxon})
		}

		switch vars.Timer {
		case 20:
			e.Do(spin.PlaySpeech{ID: SpeechAllReactorsApprochingCriticalMass})
		case 10:
			e.Do(spin.PlaySpeech{ID: SpeechMeltdownIsImminent})
		case 4:
			e.Do(spin.PlaySpeech{ID: SpeechFour})
		case 3:
			e.Do(spin.PlaySpeech{ID: SpeechThree})
		case 2:
			e.Do(spin.PlaySpeech{ID: SpeechTwo})
		case 1:
			e.Do(spin.PlaySpeech{ID: SpeechOne})
		}
	}
	e.Post(spin.TimeoutEvent{ID: ScriptMeltdownMode})
}

func meltdownCountdownAudioScript(e spin.Env) {
	e.Do(spin.PlayMusic{ID: MusicMode1})

	e.Do(spin.PlaySpeech{ID: SpeechContainmentFailureAtThreeMeterIsland, Notify: true})
	if _, done := e.WaitFor(spin.SpeechFinishedEvent{}); done {
		e.Do(spin.StopSpeech{ID: SpeechContainmentFailureAtThreeMeterIsland})
		return
	}

	for {
		wait := time.Duration(rand.Intn(3000) + 1500)
		if done := e.Sleep(wait * time.Millisecond); done {
			return
		}
		e.Do(spin.PlaySound{ID: SoundMeltdownCracking})
	}
}

func meltdownTimeoutScript(e spin.Env) {
	e.Do(spin.PlaySound{ID: SoundMeltdownExplosion})
	if done := e.Sleep(2000 * time.Millisecond); done {
		return
	}

	meltdownFailureFrame(e)
	e.Do(spin.PlaySound{ID: SoundSuccess})
	if done := e.Sleep(2500 * time.Millisecond); done {
		return
	}

	e.Post(spin.AdvanceEvent{ID: ScriptMeltdownMode})
}

func meltdownCompleteScript(e spin.Env) {
	e.Do(spin.PlaySpeech{ID: SpeechAllReactorsStabilized})
	if done := e.Sleep(3000 * time.Millisecond); done {
		return
	}

	e.Do(spin.PlaySpeech{ID: SpeechDreddToControl, Notify: true})
	if _, done := e.WaitFor(spin.SpeechFinishedEvent{}); done {
		e.Do(spin.StopSound{ID: SpeechDreddToControl})
		return
	}

	e.Do(spin.PlaySpeech{ID: SpeechThreeMeterIslandIsSecured, Notify: true})
	if _, done := e.WaitFor(spin.SpeechFinishedEvent{}); done {
		e.Do(spin.StopSpeech{ID: SpeechThreeMeterIslandIsSecured})
	}
	e.Post(spin.AdvanceEvent{ID: ScriptMeltdownMode})
}

func meltdownModeScript(e spin.Env) {
	e.Do(spin.PlayScript{ID: ScriptMeltdownCountdown})
	evt, done := e.WaitFor(
		spin.AdvanceEvent{ID: ScriptMeltdownMode},
		spin.TimeoutEvent{ID: ScriptMeltdownMode},
	)
	e.Do(spin.StopScript{ID: ScriptTankCountdown})
	if done {
		return
	}
	e.Do(spin.PlayMusic{ID: MusicMain})
	if evt == (spin.TimeoutEvent{ID: ScriptMeltdownMode}) {
		e.Do(spin.PlayScript{ID: ScriptMeltdownTimeout})
	} else {
		e.Do(spin.PlayScript{ID: ScriptMeltdownComplete})
	}
	if _, done := e.WaitFor(spin.AdvanceEvent{ID: ScriptMeltdownMode}); done {
		return
	}
	e.Post(spin.ScriptFinishedEvent{ID: ScriptMeltdownMode})
}

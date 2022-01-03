package jdx

import (
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/prog/builtin"
)

const (
	MessageOpenSafeAttempt = "jdx.MessageOpenSafeAttempt"
)

func safecrackerCountdown1Frame(e spin.Env) {
	vars := GetVars(e)
	r, g := e.Display("").Renderer("")

	r.Fill(spin.ColorBlack)
	g.Y = 2
	g.Font = builtin.FontPfArmaFive8
	r.Print(g, "SHOOT SAFECRACKER")
	g.Y = 12

	g.Font = builtin.Font14x10
	r.Print(g, spin.FormatScore("%v", vars.SafecrackerScore))
}

func safecrackereCountdown1VideoScript(e spin.Env) {
	vars := GetVars(e)

	modeText := [3]string{"SAFECRACKER", "SHOOT", "SUBWAY"}
	if done := modeIntroVideo(e, modeText); done {
		return
	}

	safecrackerCountdown1Frame(e)
	if done := e.Sleep(2000 * time.Millisecond); done {
		return
	}

	expires := time.Now().Add(17 * time.Second)
	for time.Now().Before(expires) {
		vars.SafecrackerScore -= ScoreSafecrackerDec
		safecrackerCountdown1Frame(e)
		if done := e.Sleep(250 * time.Millisecond); done {
			return
		}
	}

	vars.SafecrackerScore = ScoreSafecrackerEnd
	safecrackerCountdown1Frame(e)
	if done := e.Sleep(2000 * time.Millisecond); done {
		return
	}
	e.Post(spin.TimeoutEvent{ID: ScriptSafecrackerMode})
}

func safecrackerCountdown1AudioScript(e spin.Env) {
	e.Do(spin.PlayMusic{ID: MusicMode2})
	e.Do(spin.MusicVolume{Mul: 0.5})
	defer e.Do(spin.MusicVolume{Mul: 2})

	e.Do(spin.PlaySpeech{ID: SpeechWakeUpYouGeezer})
	if done := e.Sleep(2000 * time.Millisecond); done {
		e.Do(spin.StopSpeech{ID: SpeechWakeUpYouGeezer})
		return
	}

	e.Do(spin.PlaySound{ID: SoundSnore})
	if done := e.Sleep(1000 * time.Millisecond); done {
		e.Do(spin.StopSound{ID: SoundSnore})
		return
	}

	e.Do(spin.PlaySpeech{ID: SpeechIllBeBack})
	if done := e.Sleep(2000 * time.Millisecond); done {
		e.Do(spin.StopSpeech{ID: SpeechIllBeBack})
		return
	}
}

func safecrackerCountdown1Script(e spin.Env) {
	ctx, cancel := e.Derive()

	e.NewCoroutine(ctx, safecrackerCountdown1AudioScript)
	e.NewCoroutine(ctx, safecrackereCountdown1VideoScript)
	e.WaitFor(
		spin.AdvanceEvent{ID: ScriptSafecrackerMode},
		spin.TimeoutEvent{ID: ScriptSafecrackerMode},
	)
	cancel()
}

func safecrackerAwardedFrame(e spin.Env) {
	vars := GetVars(e)
	r, g := e.Display("").Renderer(spin.LayerPriority)

	r.Fill(spin.ColorBlack)
	g.Y = 5
	g.Font = builtin.Font14x10
	r.Print(g, spin.FormatScore("%v", vars.SafecrackerBonus))

	g.Y = 22
	g.Font = builtin.FontPfArmaFive8
	r.Print(g, "AWARDED")
}

func safecrackerOpenThatSafeScript(e spin.Env) {
	defer func() {
		e.Do(spin.MusicVolume{Mul: 2})
		e.Display("").Clear(spin.LayerPriority)
	}()

	vars := GetVars(e)
	vars.SafecrackerAttempts += 1
	vars.SafecrackerBonus = vars.SafecrackerScore * vars.SafecrackerAttempts

	e.Do(spin.MusicVolume{Mul: 0.5})
	e.Do(spin.PlaySpeech{ID: SpeechOpenThatSafe, Notify: true})
	safecrackerAwardedFrame(e)
	if _, done := e.WaitFor(spin.SpeechFinishedEvent{}); done {
		e.Do(spin.StopSpeech{ID: SpeechOpenThatSafe})
		return
	}
	switch vars.SafecrackerAttempts {
	case 1:
		e.Do(spin.PlaySound{ID: SoundSafecrackerGunFire1})
		if done := e.Sleep(750 * time.Millisecond); done {
			e.Do(spin.StopSound{ID: SoundSafecrackerGunFire1})
			return
		}
		e.Do(spin.PlaySound{ID: SoundSafecrackerGunFire2})
		if done := e.Sleep(2000 * time.Millisecond); done {
			e.Do(spin.StopSound{ID: SoundSafecrackerGunFire2})
			return
		}
		e.Do(spin.PlaySound{ID: SoundSnore})
		if done := e.Sleep(1000 * time.Millisecond); done {
			e.Do(spin.StopSound{ID: SoundSnore})
			return
		}
	case 2:
		e.Do(spin.PlaySound{ID: SoundSafecrackerTankFire})
		if done := e.Sleep(1000 * time.Millisecond); done {
			e.Do(spin.StopSound{ID: SoundSafecrackerTankFire})
			return
		}
		e.Do(spin.PlaySound{ID: SoundSafecrackerTankFire})
		if done := e.Sleep(1000 * time.Millisecond); done {
			e.Do(spin.StopSound{ID: SoundSafecrackerTankFire})
			return
		}
		e.Do(spin.PlaySound{ID: SoundSafecrackerExplosion})
		if done := e.Sleep(2000 * time.Millisecond); done {
			e.Do(spin.StopSound{ID: SoundSafecrackerExplosion})
			return
		}
		e.Do(spin.PlaySound{ID: SoundSnore})
		if done := e.Sleep(1000 * time.Millisecond); done {
			e.Do(spin.StopSound{ID: SoundSnore})
			return
		}
	case 3:
		for i := 0; i < 3; i++ {
			e.Do(spin.PlaySound{ID: SoundSafecrackerGunFire3})
			if done := e.Sleep(500 * time.Millisecond); done {
				e.Do(spin.StopSound{ID: SoundSafecrackerGunFire3})
				return
			}
		}
		e.Do(spin.PlaySound{ID: SoundSafecrackerExplosion})
		if done := e.Sleep(2000 * time.Millisecond); done {
			e.Do(spin.StopSound{ID: SoundSafecrackerExplosion})
			return
		}
		e.Do(spin.PlaySound{ID: SoundSnore})
		if done := e.Sleep(1000 * time.Millisecond); done {
			e.Do(spin.StopSound{ID: SoundSnore})
			return
		}
	}
}

func safecrackerCountdown2Script(e spin.Env) {
	vars := GetVars(e)
	vars.Timer = 30
	cancel := spin.CountdownScript(e, &vars.Timer, 1000, spin.TimeoutEvent{ID: ScriptSafecrackerMode})
	defer cancel()

	for {
		if _, done := e.WaitFor(spin.Message{ID: MessageOpenSafeAttempt}); done {
			return
		}
		cancel()
		if done := e.Sleep(2000 * time.Millisecond); done {
			return
		}
		cancel = spin.CountdownScript(e, &vars.Timer, 1000, spin.TimeoutEvent{ID: ScriptSafecrackerMode})
	}
}

func safecrackerWatchSubwayScript(e spin.Env) {
	for {
		if _, done := e.WaitFor(spin.SwitchEvent{ID: jd.SwitchSubwayEnter1}); done {
			return
		}
		e.Post(spin.Message{ID: MessageOpenSafeAttempt})
	}
}

func safecrackerCountdown2Frame(e spin.Env) {
	vars := GetVars(e)
	r, g := e.Display("").Renderer("")

	r.Fill(spin.ColorBlack)
	g.Y = 2
	g.Font = builtin.FontPfArmaFive8
	r.Print(g, "SHOOT SAFECRACKER")

	g.AnchorX = spin.AnchorLeft
	g.X = 5
	g.AnchorY = spin.AnchorMiddle
	g.Y = r.Height()/2 + 6
	g.Font = builtin.Font14x10
	r.Print(g, "%v", vars.Timer)

	g.X = r.Width() / 2
	g.AnchorX = spin.AnchorCenter
	g.Font = builtin.Font09x7
	r.Print(g, spin.FormatScore("%v", vars.SafecrackerScore))

	g.X = r.Width() - 2
	g.AnchorX = spin.AnchorRight
	g.Font = builtin.FontPfTempestaFiveBold8
	r.Print(g, spin.FormatScore("X%v", vars.SafecrackerAttempts+1))
}

func safecrackerTotalFrame(e spin.Env) {
	r, g := e.Display("").Renderer(spin.LayerPriority)
	vars := GetVars(e)

	r.Fill(spin.ColorBlack)
	g.Y = 2
	g.Font = builtin.FontPfArmaFive8
	r.Print(g, "SAFECRACKER TOTAL")
	g.Y = 12

	g.Font = builtin.Font14x10
	r.Print(g, spin.FormatScore("%v", vars.SafecrackerBonus))
}

func safecrackerIncompleteScript(e spin.Env) {
	e.Do(spin.PlayMusic{ID: MusicMain})
	safecrackerTotalFrame(e)
	if done := e.Sleep(3000 * time.Millisecond); done {
		return
	}
	e.Display("").Clear(spin.LayerPriority)
}

func safecrackerCompleteScript(e spin.Env) {
	e.Do(spin.PlayMusic{ID: MusicMain})
	e.Do(spin.MusicVolume{Mul: 0.5})
	defer e.Do(spin.MusicVolume{Mul: 2})

	safecrackerTotalFrame(e)
	e.Do(spin.PlaySpeech{ID: SpeechOpenThatSafe, Notify: true})
	if _, done := e.WaitFor(spin.SpeechFinishedEvent{}); done {
		e.Do(spin.StopSpeech{ID: SpeechOpenThatSafe})
		return
	}

	e.Do(spin.PlaySound{ID: SoundSafecrackerLaserFire})
	if done := e.Sleep(2000 * time.Millisecond); done {
		e.Do(spin.StopSound{ID: SoundSafecrackerLaserFire})
		return
	}

	e.Do(spin.PlaySound{ID: SoundSnore})
	if done := e.Sleep(1000 * time.Millisecond); done {
		e.Do(spin.StopSound{ID: SoundSnore})
		return
	}

	e.Do(spin.PlaySound{ID: SoundDing})
	if done := e.Sleep(250 * time.Millisecond); done {
		return
	}

	e.Do(spin.PlaySpeech{ID: SpeechDinnerTime})
	if done := e.Sleep(2250 * time.Millisecond); done {
		return
	}
	e.Display("").Clear(spin.LayerPriority)
}

func safecrackerModeScript(e spin.Env) {
	vars := GetVars(e)
	vars.SafecrackerAttempts = 0
	vars.SafecrackerBonus = 0
	vars.SafecrackerScore = ScoreSafecrackerStart

	ctx, cancel := e.Derive()
	defer cancel()
	e.NewCoroutine(ctx, safecrackerWatchSubwayScript)

	e.Do(spin.PlayScript{ID: ScriptSafecrackerCountdown1})
	for i := 0; i < 4; i++ {
		evt, done := e.WaitFor(
			spin.Message{ID: MessageOpenSafeAttempt},
			spin.TimeoutEvent{ID: ScriptSafecrackerMode},
		)
		if i == 0 {
			e.Do(spin.StopScript{ID: ScriptSafecrackerCountdown1})
		}
		if done {
			return
		}
		if evt == (spin.TimeoutEvent{ID: ScriptSafecrackerMode}) {
			e.Do(spin.PlayScript{ID: ScriptSafecrackerIncomplete})
			e.Post(spin.ScriptFinishedEvent{ID: ScriptSafecrackerMode})
			return
		}
		if i == 3 {
			break
		}
		e.Do(spin.PlayScript{ID: ScriptSafecrackerOpenThatSafe})
		if i == 0 {
			spin.RenderFrameScript(e, safecrackerCountdown2Frame)
			e.NewCoroutine(ctx, safecrackerCountdown2Script)
		}
	}
	cancel()
	e.Do(spin.PlayScript{ID: ScriptSafecrackerComplete})
	e.Post(spin.ScriptFinishedEvent{ID: ScriptSafecrackerMode})
}

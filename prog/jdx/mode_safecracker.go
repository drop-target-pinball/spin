package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/proc"
)

func safecrackerModeScript(e *spin.ScriptEnv) {
	vars := GetVars(e)

	e.Do(proc.DriverSchedule{ID: jd.LampAwardSafeCracker, Schedule: proc.BlinkSchedule})
	defer e.Do(spin.DriverOff{ID: jd.LampAwardSafeCracker})

	e.Do(spin.PlayScript{ID: ScriptSafecrackerMode1})
	e.NewCoroutine(watchCenterDropTargetLoop)

	evt, done := e.WaitFor(
		spin.AdvanceEvent{},
		spin.TimeoutEvent{},
	)
	if done {
		return
	}
	if evt == (spin.TimeoutEvent{}) {
		vars.SafecrackerBonus = vars.SafecrackerScore
		e.Do(spin.PlayScript{ID: ScriptSafecrackerIncomplete})
		return
	}
	e.Do(spin.PlayScript{ID: ScriptSafecrackerMode2})
	e.WaitFor(spin.ScriptFinishedEvent{ID: ScriptSafecrackerMode2})
}

func safecrackerMode1Script(e *spin.ScriptEnv) {
	r := e.Display("").Open(0)
	defer r.Close()

	e.Do(spin.PlayMusic{ID: MusicMode2})

	vars := GetVars(e)
	vars.Mode = ModeSafeCracker
	defer func() { vars.Mode = ModeNone }()
	vars.SafecrackerScore = ScoreSafecrackerStart

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.Do(spin.PlaySpeech{ID: SpeechWakeUpYouGeezer})
		s.Sleep(2_000)
		s.Do(spin.PlaySound{ID: SoundSnore})
		s.Sleep(1_250)
		s.Do(spin.PlaySpeech{ID: SpeechIllBeBack})
		s.Sleep(1_500)

		s.DoFunc(func() {
			e.NewCoroutine(func(e *spin.ScriptEnv) {
				if done := spin.ScoreHurryUpLoop(e,
					&vars.SafecrackerScore,
					250, // tick ms
					ScoreSafecrackerDec,
					ScoreSafecrackerEnd,
				); done {
					return
				}
				if done := e.Sleep(2_000); done {
					return
				}
				e.Post(spin.TimeoutEvent{})
			})
		})

		s.Run()
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		if done := ModeIntroScript(e, r, "SAFECRACKER", "SHOOT", "SUBWAY"); done {
			return
		}
		spin.RenderFrameLoop(e, func(e *spin.ScriptEnv) {
			ModeAndScorePanel(e, r, "SAFECRACKER", vars.SafecrackerScore)
		})
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)
		s.WaitFor(spin.SwitchEvent{ID: jd.SwitchSubwayEnter1})
		s.Post(spin.AdvanceEvent{})
		s.Run()
	})

	e.WaitFor(spin.AdvanceEvent{}, spin.TimeoutEvent{})
}

func safecrackerMode2Script(e *spin.ScriptEnv) {
	r := e.Display("").Open(0)
	defer r.Close()

	vars := GetVars(e)
	vars.Mode = ModeSafeCracker
	defer func() { vars.Mode = ModeNone }()
	vars.Timer = 30

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		e.NewCoroutine(func(e *spin.ScriptEnv) {
			spin.CountdownLoop(e, &vars.Timer, 1500, spin.TimeoutEvent{})
		})
		spin.RenderFrameLoop(e, func(e *spin.ScriptEnv) {
			safecrackerMode2Panel(e, r)
		})
	})

	e.Do(spin.PlayScript{ID: ScriptSafecrackerOpenThatSafe})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.WaitFor(spin.SwitchEvent{ID: jd.SwitchSubwayEnter1})
		s.Do(spin.PlayScript{ID: ScriptSafecrackerOpenThatSafe})
		s.WaitFor(spin.SwitchEvent{ID: jd.SwitchSubwayEnter1})
		s.Do(spin.PlayScript{ID: ScriptSafecrackerOpenThatSafe})
		s.WaitFor(spin.SwitchEvent{ID: jd.SwitchSubwayEnter1})
		s.Post(spin.AdvanceEvent{})

		s.Run()
	})

	evt, done := e.WaitFor(
		spin.AdvanceEvent{},
		spin.TimeoutEvent{},
	)
	if done {
		return
	}
	if evt == (spin.TimeoutEvent{}) {
		e.Do(spin.PlayScript{ID: ScriptSafecrackerIncomplete})
	} else {
		e.Do(spin.PlayScript{ID: ScriptSafecrackerComplete})
	}
}

func watchCenterDropTargetLoop(e *spin.ScriptEnv) {
	switches := spin.GetResourceVars(e).Switches
	for {
		if !switches[jd.SwitchDropTargetD].Active {
			if done := e.Sleep(500); done {
				return
			}
			e.Do(spin.DriverPulse{ID: jd.CoilDropTargetTrip})
		}
		if _, done := e.WaitFor(spin.SwitchEvent{ID: jd.SwitchDropTargetD, Released: true}); done {
			return
		}
	}
}

func safecrackerOpenThatSafeScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(spin.PriorityAnnounce)
	defer r.Close()

	vars := GetVars(e)
	vars.SafecrackerAttempts += 1
	vars.SafecrackerBonus = vars.SafecrackerScore * vars.SafecrackerAttempts

	ScoreAndLabelPanel(e, r, vars.SafecrackerBonus, "AWARDED")

	s := spin.NewSequencer(e)

	switch vars.SafecrackerAttempts {
	case 1:
		s.Do(spin.PlaySpeech{ID: SpeechOpenThatSafe, Notify: true, Duck: 0.5})
		s.WaitFor(spin.SpeechFinishedEvent{})
		s.Do(spin.PlaySound{ID: SoundSafecrackerGunFire1})
		s.Sleep(750)
		s.Do(spin.PlaySound{ID: SoundSafecrackerGunFire2})
		s.Sleep(2_000)
		s.Do(spin.PlaySound{ID: SoundSnore})
		s.Sleep(1_000)
	case 2:
		s.Do(spin.PlaySpeech{ID: SpeechOpenThatSafe, Notify: true, Duck: 0.5})
		s.WaitFor(spin.SpeechFinishedEvent{})
		s.Do(spin.PlaySound{ID: SoundSafecrackerTankFire})
		s.Sleep(1_000)
		s.Do(spin.PlaySound{ID: SoundSafecrackerTankFire})
		s.Sleep(1_000)
		s.Do(spin.PlaySound{ID: SoundSafecrackerExplosion})
		s.Sleep(2_000)
		s.Do(spin.PlaySound{ID: SoundSnore})
		s.Sleep(1_000)
	case 3:
		s.Do(spin.PlaySpeech{ID: SpeechOpenThatSafe, Notify: true, Duck: 0.5})
		s.WaitFor(spin.SpeechFinishedEvent{})
		s.Do(spin.PlaySound{ID: SoundSafecrackerGunFire3})
		s.Sleep(500)
		s.Do(spin.PlaySound{ID: SoundSafecrackerGunFire3})
		s.Sleep(500)
		s.Do(spin.PlaySound{ID: SoundSafecrackerGunFire3})
		s.Sleep(500)
		s.Do(spin.PlaySound{ID: SoundSafecrackerExplosion})
		s.Sleep(2_000)
		s.Do(spin.PlaySound{ID: SoundSnore})
		s.Sleep(1_000)
	}
	s.Run()
}

func safecrackerMode2Panel(e *spin.ScriptEnv, r spin.Renderer) {
	vars := GetVars(e)
	g := r.Graphics()

	r.Fill(spin.ColorOff)
	g.Y = 2
	g.Font = spin.FontPfArmaFive8
	r.Print(g, "SHOOT SAFECRACKER")

	g.AnchorX = spin.AnchorLeft
	g.X = 5
	g.AnchorY = spin.AnchorMiddle
	g.Y = r.Height()/2 + 6
	g.Font = spin.Font14x10
	r.Print(g, "%v", vars.Timer)

	g.X = r.Width() / 2
	g.AnchorX = spin.AnchorCenter
	g.Font = spin.Font09x7
	r.Print(g, spin.FormatScore("%v", vars.SafecrackerScore))

	g.X = r.Width() - 2
	g.AnchorX = spin.AnchorRight
	g.Font = spin.FontPfTempestaFiveBold8
	r.Print(g, spin.FormatScore("X%v", vars.SafecrackerAttempts+1))
}

func safecrackerIncompleteScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(0)
	defer r.Close()

	e.Do(spin.PlayMusic{ID: MusicMain})

	vars := GetVars(e)
	ModeAndScorePanel(e, r, "SAFECRACKER TOTAL", vars.SafecrackerBonus)
	e.Sleep(3_000)
}

func safecrackerCompleteScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(0)
	defer r.Close()

	e.Do(spin.PlayMusic{ID: MusicMain})

	vars := GetVars(e)
	ModeAndScorePanel(e, r, "SAFECRACKER TOTAL", vars.SafecrackerBonus)

	s := spin.NewSequencer(e)

	s.Do(spin.PlaySpeech{ID: SpeechOpenThatSafe, Notify: true, Duck: 0.5})
	s.WaitFor(spin.SpeechFinishedEvent{})
	s.Do(spin.PlaySound{ID: SoundSafecrackerLaserFire})
	s.Sleep(2_000)
	s.Do(spin.PlaySound{ID: SoundSnore})
	s.Sleep(1_000)
	s.Do(spin.PlaySound{ID: SoundDing})
	s.Sleep(250)
	s.Do(spin.PlaySpeech{ID: SpeechDinnerTime})
	s.Sleep(2_250)

	s.Run()
}

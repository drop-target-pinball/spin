package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/proc"
)

func pursuitModeScript(e *spin.ScriptEnv) {
	r := e.Display("").Open()
	defer r.Close()

	e.Do(spin.PlayMusic{ID: MusicMode2})

	vars := GetVars(e)
	player := spin.GetPlayerVars(e)

	vars.Mode = ModePursuit
	defer func() { vars.Mode = ModeNone }()
	vars.Timer = 30
	vars.PursuitBonus = ScorePursuit0

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.Do(spin.PlaySpeech{ID: SpeechDreddToControl, Notify: true})
		s.WaitFor(spin.SpeechFinishedEvent{})
		s.Do(spin.PlaySpeech{ID: SpeechImInPursuitOfAStolenVehicle, Notify: true})
		s.WaitFor(spin.SpeechFinishedEvent{})

		s.Sleep(1_000)

		s.Do(spin.PlaySound{ID: SoundPoliceSiren, Loop: true})
		s.Do(spin.StopSound{ID: SoundPursuitEngine})
		s.Defer(spin.StopSound{ID: SoundPoliceSiren})
		s.Defer(spin.StopSound{ID: SoundPursuitEngine})

		s.Sleep(3_000)

		s.Do(spin.PlaySound{ID: SoundMotorRev})
		s.Sleep(2_000)
		s.Do(spin.PlaySound{ID: SoundMotorRev})
		s.Sleep(4_000)
		s.Do(spin.PlaySound{ID: SoundTireSqueal1})
		s.Sleep(4_000)
		s.Do(spin.PlaySound{ID: SoundMotorRev})
		s.Sleep(4_000)
		s.Do(spin.PlaySound{ID: SoundTireSqueal2})
		s.Sleep(4_000)
		s.Do(spin.PlaySound{ID: SoundTireSqueal1})
		s.Sleep(4_000)
		s.Do(spin.PlaySound{ID: SoundMotorRev})

		s.Run()
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.Do(proc.DriverSchedule{ID: jd.FlasherRightPursuit, Schedule: proc.FlasherBlinkSchedule})
		s.WaitFor(spin.SwitchEvent{ID: jd.SwitchRightRampExit})
		s.Do(spin.PlaySound{ID: SoundPursuitMissile})
		s.DoFunc(func() { vars.PursuitBonus = ScorePursuit1 })

		s.Do(spin.DriverOff{ID: jd.FlasherRightPursuit})
		s.Do(proc.DriverSchedule{ID: jd.FlasherLeftPursuit, Schedule: proc.FlasherBlinkSchedule})
		s.WaitFor(spin.SwitchEvent{ID: jd.SwitchLeftRampExit})
		s.Do(spin.PlaySound{ID: SoundPursuitMissile})
		s.DoFunc(func() { vars.PursuitBonus = ScorePursuit2 })

		s.Do(spin.DriverOff{ID: jd.FlasherLeftPursuit})
		s.Do(proc.DriverSchedule{ID: jd.FlasherRightPursuit, Schedule: proc.FlasherBlinkSchedule})
		s.WaitFor(spin.SwitchEvent{ID: jd.SwitchRightRampExit})
		s.DoFunc(func() { vars.PursuitBonus = ScorePursuit3 })
		s.Post(spin.AdvanceEvent{})

		s.Defer(spin.DriverOff{ID: jd.FlasherLeftPursuit})
		s.Defer(spin.DriverOff{ID: jd.FlasherRightPursuit})

		s.Run()
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		spin.CountdownLoop(e, &vars.Timer, 1000, spin.TimeoutEvent{})
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		if done := ModeIntroScript(e, r, "PURSUIT", "SHOOT", "FLASHING RAMP"); done {
			return
		}
		spin.RenderFrameLoop(e, func(e *spin.ScriptEnv) {
			TimerAndScorePanel(e, r, "PURSUIT", vars.Timer, player.Score, "SHOOT FLASHING RAMP")
		})
	})

	evt, done := e.WaitFor(
		spin.AdvanceEvent{},
		spin.TimeoutEvent{},
	)
	if done {
		return
	}
	if evt == (spin.TimeoutEvent{}) {
		e.Do(spin.PlayScript{ID: ScriptPursuitIncomplete})
	} else {
		e.Do(spin.PlayScript{ID: ScriptPursuitComplete})
	}
}

func pursuitIncompleteScript(e *spin.ScriptEnv) {
	r := e.Display("").Open()
	defer r.Close()

	e.Do(spin.PlayMusic{ID: MusicMain})

	vars := GetVars(e)
	ModeAndScorePanel(e, r, "PURSUIT TOTAL", vars.PursuitBonus)

	s := spin.NewSequencer(e)

	s.Do(spin.PlaySpeech{ID: SpeechDreddToControl, Notify: true})
	s.WaitFor(spin.SpeechFinishedEvent{})
	s.Sleep(1_000)

	s.Do(spin.PlaySpeech{ID: SpeechSuspectGotAway, Notify: true})
	s.WaitFor(spin.SpeechFinishedEvent{})
	s.Run()
}

func pursuitCompleteScript(e *spin.ScriptEnv) {
	r := e.Display("").Open()
	defer r.Close()

	e.Do(spin.PlayMusic{ID: MusicMain})

	vars := GetVars(e)
	ModeAndScorePanel(e, r, "PURSUIT TOTAL", vars.PursuitBonus)

	s := spin.NewSequencer(e)

	s.Do(spin.PlaySound{ID: SoundPursuitMissile})
	s.Sleep(500)
	s.Do(spin.PlaySound{ID: SoundPursuitExplosion})
	s.Sleep(1_000)
	s.Do(spin.PlaySpeech{ID: SpeechYourDrivingDaysAreOverPunk, Notify: true})
	s.WaitFor(spin.SpeechFinishedEvent{})
	s.Sleep(1_000)
	s.Run()
}

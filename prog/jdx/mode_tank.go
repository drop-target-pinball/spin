package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/proc"
)

func tankModeScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(0)
	defer r.Close()

	e.Do(spin.PlayMusic{ID: MusicMode2})

	vars := GetVars(e)
	player := spin.GetPlayerVars(e)

	vars.Mode = ModeBattleTank
	defer func() { vars.Mode = ModeNone }()
	vars.Timer = 30
	vars.TankBonus = ScoreTank0

	e.Do(proc.DriverSchedule{ID: jd.LampLeftTank, Schedule: proc.BlinkSchedule})
	e.Do(proc.DriverSchedule{ID: jd.LampRightTank, Schedule: proc.BlinkSchedule})
	e.Do(proc.DriverSchedule{ID: jd.LampCenterTank, Schedule: proc.BlinkSchedule})
	defer func() {
		e.Do(spin.DriverOff{ID: jd.LampLeftTank})
		e.Do(spin.DriverOff{ID: jd.LampRightTank})
		e.Do(spin.DriverOff{ID: jd.LampCenterTank})
	}()

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.Do(spin.PlaySpeech{ID: SpeechBattleTankSightedInSectorSix, Notify: true, Duck: 0.5})
		s.WaitFor(spin.SpeechFinishedEvent{})
		s.Sleep(1_000)

		s.DoFunc(func() {
			s := spin.NewSequencer(e)
			s.Do(spin.PlaySound{ID: SoundTankFire, Vol: 100})
			s.Sleep(1_500)
			s.Loop()
			s.Run()
		})

		s.Run()
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		if done := ModeIntroScript(e, r, "BATTLE TANK", "SHOOT", "GREEN ARROWS"); done {
			return
		}
		spin.RenderFrameLoop(e, func(e *spin.ScriptEnv) {
			TimerAndScorePanel(e, r, "BATTLE TANK", vars.Timer, player.Score, "SHOOT GREEN ARROWS")
		})
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		spin.CountdownLoop(e, &vars.Timer, 1000, spin.TimeoutEvent{})
	})

	e.NewCoroutine(tankSequenceScript)

	evt, done := e.WaitFor(
		spin.AdvanceEvent{},
		spin.TimeoutEvent{},
	)
	if done {
		return
	}
	if evt == (spin.TimeoutEvent{}) {
		e.Do(spin.PlayScript{ID: ScriptTankIncomplete})
	} else {
		e.Do(spin.PlayScript{ID: ScriptTankComplete})
	}
}

func tankSequenceScript(e *spin.ScriptEnv) {
	vars := GetVars(e)

	shots := map[string]bool{
		jd.SwitchOuterLoopLeft:   false,
		jd.SwitchTopLeftRampExit: false,
		jd.SwitchBankTargets:     false,
	}
	lamps := map[string]string{
		jd.SwitchOuterLoopLeft:   jd.LampLeftTank,
		jd.SwitchTopLeftRampExit: jd.LampCenterTank,
		jd.SwitchBankTargets:     jd.LampRightTank,
	}

	vars.TankBonus = ScoreTank0
	hits := 0
	for hits < 3 {
		evt, done := e.WaitFor(
			spin.SwitchEvent{ID: jd.SwitchOuterLoopLeft},
			spin.SwitchEvent{ID: jd.SwitchTopLeftRampExit},
			spin.SwitchEvent{ID: jd.SwitchBankTargets},
		)
		if done {
			return
		}
		id := evt.(spin.SwitchEvent).ID
		if shots[id] {
			continue
		}
		hits += 1
		shots[id] = true
		e.Do(spin.DriverOff{ID: lamps[id]})
		e.Do(spin.PlayScript{ID: ScriptTankHit})
	}
	vars.TankBonus = ScoreTank3
	e.Post(spin.AdvanceEvent{})
}

func tankHitScript(e *spin.ScriptEnv) {
	vars := GetVars(e)
	vars.TankHits += 1

	var atPercent string
	switch vars.TankHits {
	case 1:
		atPercent = SpeechTwentyFivePercent
		vars.TankBonus = ScoreTank1
	case 2:
		atPercent = SpeechSixtyPercent
		vars.TankBonus = ScoreTank2
	}
	if atPercent == "" {
		return
	}

	s := spin.NewSequencer(e)

	s.Do(spin.PlaySpeech{ID: SpeechBattleTankDamageAt, Notify: true, Duck: 0.5})
	s.WaitFor(spin.SpeechFinishedEvent{})
	s.Do(spin.PlaySpeech{ID: atPercent, Notify: true, Duck: 0.5})
	s.WaitFor(spin.SpeechFinishedEvent{})

	s.Run()
}

func tankIncompleteScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(0)
	defer r.Close()

	vars := GetVars(e)
	e.Do(spin.PlayMusic{ID: MusicMain})
	ModeAndScorePanel(e, r, "BATTLE TANK TOTAL", vars.TankBonus)
	e.Sleep(3_000)
}

func tankCompleteScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(0)
	defer r.Close()

	vars := GetVars(e)
	vars.TankHits = 0

	e.Do(spin.PlayMusic{ID: MusicMain})
	ModeAndScorePanel(e, r, "BATTLE TANK TOTAL", vars.TankBonus)

	s := spin.NewSequencer(e)

	s.Do(spin.PlaySound{ID: SoundTankDestroyed})
	s.Sleep(1_000)
	s.Do(spin.PlaySpeech{ID: SpeechBattleTankDestroyed, Notify: true})
	s.Sleep(3_000)
	s.Run()
}

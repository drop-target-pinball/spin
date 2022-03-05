package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/proc"
)

/*
SwitchEvent ID=jd.SwitchRightFireButton

SwitchEvent ID=jd.SwitchDropTargetJ
SwitchEvent ID=jd.SwitchDropTargetU
SwitchEvent ID=jd.SwitchDropTargetD
SwitchEvent ID=jd.SwitchDropTargetG
SwitchEvent ID=jd.SwitchDropTargetE

SwitchEvent ID=jd.SwitchLeftRampExit
*/

func lightBallLockScript(e *spin.ScriptEnv) {
	vars := GetVars(e)

	defer func() {
		e.Do(spin.DriverOff{ID: jd.LampDropTargetJ})
		e.Do(spin.DriverOff{ID: jd.LampDropTargetU})
		e.Do(spin.DriverOff{ID: jd.LampDropTargetD})
		e.Do(spin.DriverOff{ID: jd.LampDropTargetG})
		e.Do(spin.DriverOff{ID: jd.LampDropTargetE})
	}()

	dropTargetSounds := []string{
		SoundDropTargetLitHit1,
		SoundDropTargetLitHit2,
		SoundDropTargetLitHit3,
		SoundDropTargetLitHit4,
	}

	for i := jd.MinDropTarget; i <= jd.MaxDropTarget; i++ {
		if i < vars.LitDropTarget {
			e.Do(spin.DriverOn{ID: jd.DropTargetLamps[i]})
		} else if i == vars.LitDropTarget {
			e.Do(proc.DriverSchedule{ID: jd.DropTargetLamps[i], Schedule: proc.BlinkSchedule})
		}
	}

	for {
		lit := vars.LitDropTarget
		e.Do(proc.DriverSchedule{ID: jd.DropTargetLamps[lit], Schedule: proc.BlinkSchedule})
		if _, done := e.WaitFor(
			spin.SwitchEvent{ID: jd.DropTargetSwitches[lit]},
			spin.SwitchEvent{ID: jd.SwitchSubwayEnter1}); done {
			return
		}

		e.Do(spin.AwardScore{Val: ScoreDropTargetLit * vars.Multiplier})
		if lit == jd.MaxDropTarget {
			vars.LitDropTarget = jd.MinDropTarget
			e.Do(spin.PlaySound{ID: SoundDropTargetLitHit5})
			e.Do(proc.DriverSchedule{ID: jd.LampDropTargetJ, Schedule: proc.BlinkSchedule})
			e.Do(spin.DriverOff{ID: jd.LampDropTargetU})
			e.Do(spin.DriverOff{ID: jd.LampDropTargetD})
			e.Do(spin.DriverOff{ID: jd.LampDropTargetG})
			e.Do(spin.DriverOff{ID: jd.LampDropTargetE})
			if vars.LocksReady == 0 {
				e.NewCoroutine(func(e *spin.ScriptEnv) {
					s := spin.NewSequencer(e)

					s.Sleep(1_250)
					// FIXME: Notify is required to duck -- should not be
					s.Do(spin.PlaySpeech{ID: SpeechDimensionalGlobeActivated, Notify: true, Duck: 0.5})

					s.Run()
				})

				if vars.Mode == ModeNone && vars.StartModeLeft {
					vars.StartModeLeft = false
					e.Do(spin.DriverOff{ID: jd.LampLeftModeStart})
					e.Do(proc.DriverSchedule{ID: jd.LampRightModeStart, Schedule: proc.BlinkSchedule})
				}
			}
			if vars.LocksReady < 3 {
				if vars.MultiballAttempted {
					vars.LocksReady += 1
					e.Do(proc.DriverSchedule{ID: jd.LockLamps[vars.LocksReady], Schedule: proc.BlinkSchedule})
				} else {
					vars.LocksReady = 3
					e.Do(proc.DriverSchedule{ID: jd.LockLamps[1], Schedule: proc.BlinkSchedule})
					e.Do(proc.DriverSchedule{ID: jd.LockLamps[2], Schedule: proc.BlinkSchedule})
					e.Do(proc.DriverSchedule{ID: jd.LockLamps[3], Schedule: proc.BlinkSchedule})
				}
			}
		} else {
			e.Do(spin.DriverOn{ID: jd.DropTargetLamps[lit]})
			e.Do(spin.PlaySound{ID: dropTargetSounds[lit]})
			vars.LitDropTarget += 1
		}

		e.Do(spin.PlayScript{ID: ScriptDropTargetHit})
	}
}

func ballLockScript(e *spin.ScriptEnv) {
	vars := GetVars(e)

	for i := 1; i <= 3; i++ {
		if i <= vars.BallsLocked {
			e.Do(spin.DriverOn{ID: jd.LockLamps[i]})
		} else if i <= vars.LocksReady {
			e.Do(proc.DriverSchedule{ID: jd.LockLamps[i], Schedule: proc.BlinkSchedule})
		}
	}

	defer func() {
		for i := 1; i <= 3; i++ {
			if i > vars.BallsLocked && i <= vars.LocksReady {
				e.Do(spin.DriverOff{ID: jd.LockLamps[i]})
			}
		}
	}()

	if vars.BallsLocked <= 2 {
		for {
			if _, done := e.WaitFor(spin.SwitchEvent{ID: jd.SwitchLeftRampExit}); done {
				return
			}
			if vars.LocksReady == vars.BallsLocked {
				continue
			}
			vars.BallsLocked += 1
			e.Do(spin.DriverOn{ID: jd.LockLamps[vars.BallsLocked]})
			e.Do(spin.PlayScript{ID: ScriptBallLocked})
			if vars.BallsLocked == 2 {
				break
			}
		}
	}

	for {
		if _, done := e.WaitFor(spin.SwitchEvent{ID: jd.SwitchLeftRampEnter}); done {
			return
		}
		// From: https://github.com/preble/JD-pyprocgame/blob/master/multiball.py#L56
		e.Do(proc.DriverSchedule{ID: jd.CoilDiverter, Schedule: 0xfff, CycleSeconds: 1, Now: true})
		evt, done := e.WaitFor(
			spin.SwitchEvent{ID: jd.SwitchLeftRampEnter},
			spin.SwitchEvent{ID: jd.SwitchLeftRampToLock},
		)
		if done {
			return
		}
		if evt == (spin.SwitchEvent{ID: jd.SwitchLeftRampToLock}) {
			break
		}
	}
	e.Do(spin.PlayScript{ID: ScriptMultiball})
}

func dropTargetJudgePanel(e *spin.ScriptEnv, r spin.Renderer, blinkOn bool) {
	g := r.Graphics()
	vars := GetVars(e)

	letters := []string{"J", "U", "D", "G", "E"}

	g.Font = spin.FontPfRondaSevenBold16
	g.Y = 0
	g.X = 32
	hit := vars.LitDropTarget - 1
	if hit < 0 {
		hit = jd.MaxDropTarget
	}
	for i, letter := range letters {
		if i < hit {
			g.Color = spin.ColorOn
		} else if i == hit {
			if blinkOn {
				g.Color = spin.ColorOn
			} else {
				g.Color = spin.ColorOn8
			}
		} else {
			g.Color = spin.ColorOn8
		}
		r.Print(g, letter)
		g.X += 16
	}

	g.Color = spin.ColorOn
	g.Font = spin.FontPfArmaFive8
	g.Y = 24
	g.X = r.Width() / 2
	if hit < jd.MaxDropTarget {
		r.Print(g, spin.FormatScore("%v", ScoreDropTargetLit))
	} else {
		r.Print(g, "LOCK LIT")
	}
}

func dropTargetHitScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(spin.PriorityAnnounce)
	defer r.Close()

	s := spin.NewSequencer(e)

	s.DoFunc(func() { dropTargetJudgePanel(e, r, true) })
	s.Sleep(100)
	s.DoFunc(func() { dropTargetJudgePanel(e, r, false) })
	s.Sleep(100)
	s.LoopN(10)

	s.Run()
}

func ballLockedPanel(e *spin.ScriptEnv, r spin.Renderer) {
	g := r.Graphics()
	vars := GetVars(e)

	g.Font = spin.FontPfRondaSevenBold8
	g.X = (r.Width() / 2) - 8
	g.Y = 8
	g.AnchorX = spin.AnchorCenter
	r.Print(g, "DIMENSIONAL")

	g.Y = 18
	r.Print(g, "LOCK")

	g.Font = spin.FontPfRondaSevenBold16
	g.X = 100
	g.Y = 5
	g.AnchorX = spin.AnchorLeft
	r.Print(g, "%v", vars.BallsLocked)
}

/*
SetVar Vars=jdx.1 ID=BallsLocked Val=1
PlayScript ID=jdx.ScriptBallLocked
*/

func ballLockedScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(spin.PriorityAnnounce)
	defer r.Close()

	vars := GetVars(e)

	lockSpeech := []string{
		"",
		SpeechDimensionalLockOne,
		SpeechDimensionalLockTwo,
	}

	ballLockedPanel(e, r)

	s := spin.NewSequencer(e)

	s.Do(spin.PlaySound{ID: SoundBallLock, Notify: true, Duck: 0.25})
	s.Sleep(500)
	s.Do(spin.PlaySpeech{ID: lockSpeech[vars.BallsLocked], Notify: true})
	s.WaitFor(spin.SoundFinishedEvent{ID: SoundBallLock})

	s.Run()
}

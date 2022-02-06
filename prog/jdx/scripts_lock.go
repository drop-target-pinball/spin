package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/proc"
)

func ballLockScript(e *spin.ScriptEnv) {
	vars := GetVars(e)

	e.NewCoroutine(ballLockRampRoutine)
	for i := jd.MinDropTarget; i <= jd.MaxDropTarget; i++ {
		if i < vars.LitDropTarget {
			e.Do(spin.DriverOn{ID: jd.DropTargetLamps[i]})
		} else if i == vars.LitDropTarget {
			e.Do(proc.DriverSchedule{ID: jd.DropTargetLamps[i], Schedule: proc.BlinkSchedule})
		}
	}

	for i := 1; i <= 3; i++ {
		if i <= vars.BallsLocked {
			e.Do(spin.DriverOn{ID: jd.LockLamps[i]})
		} else if i <= vars.LocksReady {
			e.Do(proc.DriverSchedule{ID: jd.LockLamps[i], Schedule: proc.BlinkSchedule})
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

		if lit == jd.MaxDropTarget {
			vars.LitDropTarget = jd.MinDropTarget
			e.Do(proc.DriverSchedule{ID: jd.LampDropTargetJ, Schedule: proc.BlinkSchedule})
			e.Do(spin.DriverOff{ID: jd.LampDropTargetU})
			e.Do(spin.DriverOff{ID: jd.LampDropTargetD})
			e.Do(spin.DriverOff{ID: jd.LampDropTargetG})
			e.Do(spin.DriverOff{ID: jd.LampDropTargetE})
			if vars.LocksReady < 3 {
				vars.LocksReady += 1
				e.Do(proc.DriverSchedule{ID: jd.LockLamps[vars.LocksReady], Schedule: proc.BlinkSchedule})
			}
		} else {
			e.Do(spin.DriverOn{ID: jd.DropTargetLamps[lit]})
			vars.LitDropTarget += 1
		}
	}
}

func ballLockRampRoutine(e *spin.ScriptEnv) {
	vars := GetVars(e)
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
}

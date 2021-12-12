package jdx

import (
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/prog/builtin"
)

func leftShooterLaneShotScript(e spin.Env) {
	builtin.ShotTrapScript(e, jd.SwitchLeftShooterLane, jd.ShotLeftShooterLane, 250*time.Millisecond)
}

func defaultLeftShooterLaneScript(e spin.Env) {
	for {
		if _, done := e.WaitFor(spin.ShotEvent{ID: jd.ShotLeftShooterLane}); done {
			return
		}
		e.Do(spin.PlayScript{ID: jd.ScriptRaiseDropTargets})
		if done := e.Sleep(1 * time.Second); done {
			return
		}
		e.Do(spin.DriverPulse{ID: jd.CoilLeftShooterLane})
	}
}

func leftPopperShotScript(e spin.Env) {
	builtin.ShotTrapScript(e, jd.SwitchLeftPopper, jd.ShotLeftPopper, 250*time.Millisecond)
}

func defaultLeftPopperScript(e spin.Env) {
	for {
		if _, done := e.WaitFor(spin.ShotEvent{ID: jd.ShotLeftPopper}); done {
			return
		}
		for i := 0; i < 3; i++ {
			e.Do(spin.DriverPulse{ID: jd.FlasherSubwayExit})
			if done := e.Sleep(250 * time.Millisecond); done {
				return
			}
		}
		e.Do(spin.DriverPulse{ID: jd.CoilLeftPopper})
	}
}

func rightPopperShotScript(e spin.Env) {
	builtin.ShotTrapScript(e, jd.SwitchRightPopper, jd.ShotRightPopper, 250*time.Millisecond)
}

func defaultRightPopperScript(e spin.Env) {
	for {
		if _, done := e.WaitFor(spin.ShotEvent{ID: jd.ShotRightPopper}); done {
			return
		}
		e.Do(spin.DriverPulse{ID: jd.CoilRightPopper})
	}
}

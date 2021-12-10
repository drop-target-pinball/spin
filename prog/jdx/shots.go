package jdx

import (
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/prog/builtin"
)

func leftShooterLaneShotScript(e spin.Env) {
	ctx, cancel := e.Derive()
	e.NewCoroutine(ctx, func(e spin.Env) {
		builtin.ShotTrapScript(e, jd.SwitchLeftShooterLane, jd.ShotLeftShooterLane, 250*time.Millisecond)
	})
	for {
		if _, done := e.WaitFor(spin.ShotEvent{ID: jd.ShotLeftShooterLane}); done {
			cancel()
			return
		}
		e.Do(spin.DriverPulse{ID: jd.CoilLeftShooterLane})
	}
}

func leftPopperShotScript(e spin.Env) {
	ctx, cancel := e.Derive()
	e.NewCoroutine(ctx, func(e spin.Env) {
		builtin.ShotTrapScript(e, jd.SwitchLeftPopper, jd.ShotLeftPopper, 250*time.Millisecond)
	})
	for {
		if _, done := e.WaitFor(spin.ShotEvent{ID: jd.ShotLeftPopper}); done {
			cancel()
			return
		}
		e.Do(spin.DriverPulse{ID: jd.CoilLeftPopper})
	}
}

func rightPopperShotScript(e spin.Env) {
	ctx, cancel := e.Derive()
	e.NewCoroutine(ctx, func(e spin.Env) {
		builtin.ShotTrapScript(e, jd.SwitchRightPopper, jd.ShotRightPopper, 250*time.Millisecond)
	})
	for {
		if _, done := e.WaitFor(spin.ShotEvent{ID: jd.ShotRightPopper}); done {
			cancel()
			return
		}
		e.Do(spin.DriverPulse{ID: jd.CoilRightPopper})
	}
}

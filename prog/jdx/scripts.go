package jdx

import (
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

const (
	ScriptBall                   = "jdx.ScriptBall"
	ScriptBasicMode              = "jdx.ScriptBasicMode"
	ScriptDefaultLeftShooterLane = "jdx.ScriptDefaultLeftShooterLane"
	ScriptDefaultLeftPopper      = "jdx.ScriptDefaultLeftPopper"
	ScriptDefaultRightPopper     = "jdx.ScriptDefaultRightPopper"
	ScriptDebugExtraBall         = "jdx.ScriptDebugExtraBall"
	ScriptGame                   = "jdx.ScriptGame"
	ScriptLeftShooterLaneShot    = "jdx.ScriptLeftShooterLaneShot"
	ScriptLeftPopperShot         = "jdx.ScriptLeftPopperShot"
	ScriptOutlane                = "jdx.ScriptOutlane"
	ScriptPlayerAnnounce         = "jdx.ScriptPlayerAnnounce"
	ScriptPlungeMode             = "jdx.ScriptPlungeMode"
	ScriptReturnLane             = "jdx.ScriptReturnLane"
	ScriptRightPopperShot        = "jdx.ScriptRightPopperShot"
	ScriptSling                  = "jdx.ScriptSling"
	ScriptSniperMode             = "jdx.ScriptSniperMode"
	ScriptSniperScoreCountdown   = "jdx.ScriptSniperScoreCountdown"
	ScriptSniperSplat            = "jdx.ScriptSniperSplat"
	ScriptSniperTakedown         = "jdx.ScriptSniperTakedown"
	ScriptSniperFallCountdown    = "jdx.ScriptSniperFallCountdown"
)

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

func defaultRightPopperScript(e spin.Env) {
	vars := ProgVars(e)
	for {
		if _, done := e.WaitFor(spin.ShotEvent{ID: jd.ShotRightPopper}); done {
			return
		}
		if vars.ManualRightPopper {
			continue
		}
		e.Do(spin.DriverPulse{ID: jd.CoilRightPopper})
	}
}

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{
		ID:     ScriptBall,
		Script: ballScript,
		Scope:  spin.ScopeBall,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptBasicMode,
		Script: basicModeScript,
		Scope:  spin.ScopeBall,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptDefaultLeftShooterLane,
		Script: defaultLeftShooterLaneScript,
		Scope:  spin.ScopeBall,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptDefaultLeftPopper,
		Script: defaultLeftPopperScript,
		Scope:  spin.ScopeBall,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptDefaultRightPopper,
		Script: defaultRightPopperScript,
		Scope:  spin.ScopeBall,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptDebugExtraBall,
		Script: debugExtraBallScript,
		Scope:  spin.ScopeBall,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptGame,
		Script: gameScript,
		Scope:  spin.ScopeGame,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptLeftShooterLaneShot,
		Script: leftShooterLaneShotScript,
		Scope:  spin.ScopeBall,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptLeftPopperShot,
		Script: leftPopperShotScript,
		Scope:  spin.ScopeBall,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptOutlane,
		Script: outlaneScript,
		Scope:  spin.ScopeBall,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptPlayerAnnounce,
		Script: playerAnnounceScript,
		Scope:  spin.ScopeGame,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptPlungeMode,
		Script: plungeScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptReturnLane,
		Script: returnLaneScript,
		Scope:  spin.ScopeBall,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptRightPopperShot,
		Script: rightPopperShotScript,
		Scope:  spin.ScopeBall,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSling,
		Script: slingScript,
		Scope:  spin.ScopeBall,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperTakedown,
		Script: sniperTakedownScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperScoreCountdown,
		Script: sniperScoreCountdownScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperSplat,
		Script: sniperSplatScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperFallCountdown,
		Script: sniperFallCountdownScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperMode,
		Script: sniperModeScript,
		Scope:  spin.ScopeMode,
	})
}

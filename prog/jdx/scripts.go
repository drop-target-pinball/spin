package jdx

import (
	"github.com/drop-target-pinball/spin"
)

const (
	ScriptBall                   = "jdx.ScriptBall"
	ScriptDefaultLeftShooterLane = "jdx.ScriptDefaultLeftShooterLane"
	ScriptDefaultLeftPopper      = "jdx.ScriptDefaultLeftPopper"
	ScriptDefaultRightPopper     = "jdx.ScriptDefaultRightPopper"
	ScriptDebugExtraBall         = "jdx.ScriptDebugExtraBall"
	ScriptGame                   = "jdx.ScriptGame"
	ScriptLeftShooterLaneShot    = "jdx.ScriptLeftShooterLaneShot"
	ScriptLeftPopperShot         = "jdx.ScriptLeftPopperShot"
	ScriptOutlane                = "jdx.ScriptOutlane"
	ScriptPlayerAnnounce         = "jdx.ScriptPlayerAnnounce"
	ScriptPlunge                 = "jdx.ScriptPlunge"
	ScriptReturnLane             = "jdx.ScriptReturnLane"
	ScriptRightPopperShot        = "jdx.ScriptRightPopperShot"
	ScriptSling                  = "jdx.ScriptSling"
	ScriptSniperMode             = "jdx.ScriptSniperMode"
	ScriptSniperScoreCountdown   = "jdx.ScriptSniperScoreCountdown"
	ScriptSniperSplat            = "jdx.ScriptSniperSplat"
	ScriptSniperTakedown         = "jdx.ScriptSniperTakedown"
	ScriptSniperFallCountdown    = "jdx.ScriptSniperFallCountdown"
)

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{
		ID:     ScriptBall,
		Script: ballScript,
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
		ID:     ScriptPlunge,
		Script: plungeScript,
		Scope:  spin.ScopeBall,
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

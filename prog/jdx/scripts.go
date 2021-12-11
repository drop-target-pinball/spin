package jdx

import (
	"github.com/drop-target-pinball/spin"
)

const (
	ScriptBall                 = "jdx.ScriptBall"
	ScriptGame                 = "jdx.ScriptGame"
	ScriptLeftShooterLaneShot  = "jdx.ScriptLeftShooterLaneShot"
	ScriptLeftPopperShot       = "jdx.ScriptLeftPopperShot"
	ScriptOutlane              = "jdx.ScriptOutlane"
	ScriptPlayerAnnounce       = "jdx.ScriptPlayerAnnounce"
	ScriptPlunge               = "jdx.ScriptPlunge"
	ScriptReturnLane           = "jdx.ScriptReturnLane"
	ScriptRightPopperShot      = "jdx.ScriptRightPopperShot"
	ScriptSling                = "jdx.ScriptSling"
	ScriptSniperMode           = "jdx.ScriptSniperMode"
	ScriptSniperScoreCountdown = "jdx.ScriptSniperScoreCountdown"
	ScriptSniperSplatTimeout   = "jdx.ScriptSniperSplatTimeout"
	ScriptSniperTakedown       = "jdx.ScriptSniperTakedown"
	ScriptSniperFallCountdown  = "jdx.ScriptSniperFallCountdown"
)

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{
		ID:     ScriptBall,
		Script: ballScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptGame,
		Script: gameScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptLeftShooterLaneShot,
		Script: leftShooterLaneShotScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptLeftPopperShot,
		Script: leftPopperShotScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptOutlane,
		Script: outlaneScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptPlayerAnnounce,
		Script: playerAnnounceScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptPlunge,
		Script: plungeScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptReturnLane,
		Script: returnLaneScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptRightPopperShot,
		Script: rightPopperShotScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSling,
		Script: slingScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperTakedown,
		Script: sniperTakedownScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperScoreCountdown,
		Script: sniperScoreCountdownScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperSplatTimeout,
		Script: sniperSplatTimeoutScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperFallCountdown,
		Script: sniperFallCountdownScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperMode,
		Script: sniperModeScript,
	})
}

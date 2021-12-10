package jdx

import (
	"github.com/drop-target-pinball/spin"
)

const (
	ScriptDirector             = "jdx.ScriptDirector"
	ScriptLeftShooterLaneShot  = "jdx.LeftShooterLaneShot"
	ScriptLeftPopperShot       = "jdx.ScriptLeftPopperShot"
	ScriptPlayerAnnounce       = "jdx.ScriptPlayerAnnounce"
	ScriptPlunge               = "jdx.ScriptPlunge"
	ScriptRightPopperShot      = "jdx.ScriptLeftPopperShot"
	ScriptSniperMode           = "jdx.ScriptSniperMode"
	ScriptSniperScoreCountdown = "jdx.ScriptSniperScoreCountdown"
	ScriptSniperSplatTimeout   = "jdx.ScriptSniperSplatTimeout"
	ScriptSniperTakedown       = "jdx.ScriptSniperTakedown"
	ScriptSniperFallCountdown  = "jdx.ScriptSniperFallCountdown"
)

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{
		ID:     ScriptDirector,
		Script: directorScript,
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
		ID:     ScriptPlayerAnnounce,
		Script: playerAnnounceScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptPlunge,
		Script: plungeScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptRightPopperShot,
		Script: rightPopperShotScript,
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

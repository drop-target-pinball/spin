package jdx

import (
	"github.com/drop-target-pinball/spin"
)

const (
	ScriptSniperMode           = "jdx.ScriptSniperMode"
	ScriptSniperScoreCountdown = "jdx.ScriptSniperScoreCountdown"
	ScriptSniperSplatTimeout   = "jdx.ScriptSniperSplatTimeout"
	ScriptSniperTakedown       = "jdx.ScriptSniperTakedown"
	ScriptSniperFallCountdown  = "jdx.ScriptSniperFallCountdown"
)

func RegisterScripts(eng *spin.Engine) {
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

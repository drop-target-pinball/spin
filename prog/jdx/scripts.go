package jdx

import (
	"github.com/drop-target-pinball/spin"
)

const (
	ScriptSniperMode                = "jdx.ScriptSniperMode"
	ScriptSniperScoreCountdown      = "jdx.ScriptSniperScoreCountdown"
	ScriptSniperScoreCountdownAudio = "jdx.ScriptSniperScoreCountdownAudio"
	ScriptSniperScoreCountdownVideo = "jdx.ScriptSniperScoreCountdownVideo"
	ScriptSniperSplatTimeout        = "jdx.ScriptSniperSplatTimeout"
	ScriptSniperTakedown            = "jdx.ScriptSniperTakedown"
	ScriptSniperTakedownAudio       = "jdx.ScriptSniperTakedownAudio"
	ScriptSniperTakedownVideo       = "jdx.ScriptSniperTakedownVideo"
	ScriptSniperFallCountdown       = "jdx.ScriptSniperFallCountdown"
	ScriptSniperFallCountdownAudio  = "jdx.ScriptSniperFallCountdownAudio"
	ScriptSniperFallCountdownVideo  = "jdx.ScriptSniperFallCountdownVideo"
)

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperTakedown,
		Script: sniperTakedownScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperTakedownAudio,
		Script: sniperTakedownAudioScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperTakedownVideo,
		Script: sniperTakedownVideoScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperScoreCountdown,
		Script: sniperScoreCountdownScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperScoreCountdownAudio,
		Script: sniperScoreCountdownAudioScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperScoreCountdownVideo,
		Script: sniperScoreCountdownVideoScript,
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
		ID:     ScriptSniperFallCountdownAudio,
		Script: sniperFallCountdownAudioScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperFallCountdownVideo,
		Script: sniperFallCountdownVideoScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperMode,
		Script: sniperModeScript,
	})
}

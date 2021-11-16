package jdx

import (
	"github.com/drop-target-pinball/spin"
)

const (
	SniperMode                = "SniperMode"
	SniperScoreCountdown      = "SniperScoreCountdown"
	SniperScoreCountdownAudio = "SniperScoreCountdownAudio"
	SniperScoreCountdownVideo = "SniperScoreCountdownVideo"
	SniperCaught              = "SniperCaught"
	SniperCaughtAudio         = "SniperCaughtAudio"
	SniperCaughtVideo         = "SniperCaughtVideo"
	SniperFallCountdown       = "SniperFallCountdown"
	SniperFallCountdownAudio  = "SniperFallCountdownAudio"
	SniperFallCountdownVideo  = "SniperFallCountdownVideo"
)

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{
		ID:     SniperCaught,
		Script: sniperCaught,
	})
	eng.Do(spin.RegisterScript{
		ID:     SniperCaughtAudio,
		Script: sniperCaughtAudio,
	})
	eng.Do(spin.RegisterScript{
		ID:     SniperCaughtVideo,
		Script: sniperCaughtVideo,
	})
	eng.Do(spin.RegisterScript{
		ID:     SniperScoreCountdown,
		Script: sniperScoreCountdown,
	})
	eng.Do(spin.RegisterScript{
		ID:     SniperScoreCountdownAudio,
		Script: sniperScoreCountdownAudio,
	})
	eng.Do(spin.RegisterScript{
		ID:     SniperScoreCountdownVideo,
		Script: sniperScoreCountdownVideo,
	})
	eng.Do(spin.RegisterScript{
		ID:     SniperFallCountdown,
		Script: sniperFallCountdown,
	})
	eng.Do(spin.RegisterScript{
		ID:     SniperFallCountdownAudio,
		Script: sniperFallCountdownAudio,
	})
	eng.Do(spin.RegisterScript{
		ID:     SniperFallCountdownVideo,
		Script: sniperFallCountdownVideo,
	})
	eng.Do(spin.RegisterScript{
		ID:     SniperMode,
		Script: sniperMode,
	})
}

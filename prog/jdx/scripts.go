package jdx

import (
	"github.com/drop-target-pinball/spin"
)

const (
	SniperCaught              = "SniperCaught"
	SniperCaughtAudio         = "SniperCaughtAudio"
	SniperCaughtVideo         = "SniperCaughtVideo"
	SniperScoreCountdown      = "SniperScoreCountdown"
	SniperScoreCountdownAudio = "SniperScoreCountdownAudio"
	SniperScoreCountdownVideo = "SniperScoreCountdownVideo"
	SniperFall                = "SniperFall"
	SniperMode                = "SniperMode"
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
		ID:     SniperFall,
		Script: sniperFall,
	})
	eng.Do(spin.RegisterScript{
		ID:     SniperMode,
		Script: sniperMode,
	})
}

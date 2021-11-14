package jdx

import (
	"github.com/drop-target-pinball/spin"
)

const (
	SniperCaught         = "SniperCaught"
	SniperScoreCountdown = "SniperScoreCountdown"
	SniperFall           = "SniperFall"
	SniperHunt           = "SniperHunt"
	SniperMode           = "SniperMode"
)

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{ID: SniperCaught, Script: sniperCaught})
	eng.Do(spin.RegisterScript{ID: SniperScoreCountdown, Script: sniperScoreCountdown})
	eng.Do(spin.RegisterScript{ID: SniperFall, Script: sniperFall})
	eng.Do(spin.RegisterScript{ID: SniperHunt, Script: sniperHunt})
	eng.Do(spin.RegisterScript{ID: SniperMode, Script: sniperMode})
}

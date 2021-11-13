package jdx

import "github.com/drop-target-pinball/spin"

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{ID: SniperCaught, Script: sniperCaught})
	eng.Do(spin.RegisterScript{ID: SniperModeIntroShow, Script: sniperModeIntroShow})
	eng.Do(spin.RegisterScript{ID: SniperFall, Script: sniperFall})
	eng.Do(spin.RegisterScript{ID: SniperHunt, Script: sniperHunt})
	eng.Do(spin.RegisterScript{ID: SniperMode, Script: sniperMode})
}

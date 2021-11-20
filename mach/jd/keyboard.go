package jd

import (
	"github.com/drop-target-pinball/spin"
)

func RegisterKeys(eng *spin.Engine) {
	eng.Do(spin.RegisterKey{
		Key:       "left",
		EventDown: spin.SwitchEvent{ID: SwitchLeftFlipperButton},
	})
	eng.Do(spin.RegisterKey{
		Key:       "left",
		Mod:       "shift",
		EventDown: spin.SwitchEvent{ID: SwitchLeftFireButton},
	})
	eng.Do(spin.RegisterKey{
		Key:       "right",
		EventDown: spin.SwitchEvent{ID: SwitchRightFlipperButton},
	})
	eng.Do(spin.RegisterKey{
		Key:       "right",
		Mod:       "shift",
		EventDown: spin.SwitchEvent{ID: SwitchRightFireButton},
	})
	eng.Do(spin.RegisterKey{
		Key:       "return",
		EventDown: spin.SwitchEvent{ID: SwitchStartButton},
	})
}

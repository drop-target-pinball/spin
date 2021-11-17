package jd

import (
	"github.com/drop-target-pinball/spin"
)

func RegisterKeys(eng *spin.Engine) {
	eng.Do(spin.RegisterKey{
		Key:       "left",
		EventDown: spin.SwitchEvent{ID: LeftFlipperButton},
	})
	eng.Do(spin.RegisterKey{
		Key:       "left",
		Mod:       "shift",
		EventDown: spin.SwitchEvent{ID: LeftFireButton},
	})
	eng.Do(spin.RegisterKey{
		Key:       "right",
		EventDown: spin.SwitchEvent{ID: RightFlipperButton},
	})
	eng.Do(spin.RegisterKey{
		Key:       "right",
		Mod:       "shift",
		EventDown: spin.SwitchEvent{ID: RightFireButton},
	})
	eng.Do(spin.RegisterKey{
		Key:       "return",
		EventDown: spin.SwitchEvent{ID: StartButton},
	})
}

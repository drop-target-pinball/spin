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
	eng.Do(spin.RegisterKey{
		Key:       "7",
		EventDown: spin.SwitchEvent{ID: SwitchExitButton},
	})
	eng.Do(spin.RegisterKey{
		Key:       "8",
		EventDown: spin.SwitchEvent{ID: SwitchPreviousButton},
	})
	eng.Do(spin.RegisterKey{
		Key:       "9",
		EventDown: spin.SwitchEvent{ID: SwitchNextButton},
	})
	eng.Do(spin.RegisterKey{
		Key:       "0",
		EventDown: spin.SwitchEvent{ID: SwitchEnterButton},
	})
}

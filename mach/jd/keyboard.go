package jd

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/sdl"
)

func RegisterKeys(eng *spin.Engine) {
	eng.Do(sdl.RegisterKey{
		Key:       "escape",
		EventDown: spin.SwitchEvent{ID: SwitchSuperGameButton},
	})
	eng.Do(sdl.RegisterKey{
		Key:       "left",
		EventDown: spin.SwitchEvent{ID: SwitchLeftFlipperButton},
	})
	eng.Do(sdl.RegisterKey{
		Key:       "left",
		Mod:       "shift",
		EventDown: spin.SwitchEvent{ID: SwitchLeftFireButton},
	})
	eng.Do(sdl.RegisterKey{
		Key:       "right",
		EventDown: spin.SwitchEvent{ID: SwitchRightFlipperButton},
	})
	eng.Do(sdl.RegisterKey{
		Key:       "right",
		Mod:       "shift",
		EventDown: spin.SwitchEvent{ID: SwitchRightFireButton},
	})
	eng.Do(sdl.RegisterKey{
		Key:       "return",
		EventDown: spin.SwitchEvent{ID: SwitchStartButton},
	})
	eng.Do(sdl.RegisterKey{
		Key:       "1",
		EventDown: spin.SwitchEvent{ID: SwitchTrough1},
	})
	eng.Do(sdl.RegisterKey{
		Key:       "2",
		EventDown: spin.SwitchEvent{ID: SwitchLeftOutlane},
	})
	eng.Do(sdl.RegisterKey{
		Key:       "7",
		EventDown: spin.SwitchEvent{ID: SwitchExitServiceButton},
	})
	eng.Do(sdl.RegisterKey{
		Key:       "8",
		EventDown: spin.SwitchEvent{ID: SwitchPreviousServiceButton},
	})
	eng.Do(sdl.RegisterKey{
		Key:       "9",
		EventDown: spin.SwitchEvent{ID: SwitchNextServiceButton},
	})
	eng.Do(sdl.RegisterKey{
		Key:       "0",
		EventDown: spin.SwitchEvent{ID: SwitchEnterServiceButton},
	})
	eng.Do(sdl.RegisterKey{
		Key:       "q",
		EventDown: spin.SwitchEvent{ID: SwitchLeftRampEnter},
	})
	eng.Do(sdl.RegisterKey{
		Key:       "w",
		EventDown: spin.SwitchEvent{ID: SwitchLeftRampExit},
	})
	eng.Do(sdl.RegisterKey{
		Key:       "e",
		EventDown: spin.SwitchEvent{ID: SwitchRightRampExit},
	})
	eng.Do(sdl.RegisterKey{
		Key:       "d",
		EventDown: spin.SwitchEvent{ID: SwitchRightPopper},
	})
}

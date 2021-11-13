package jd

import (
	"github.com/drop-target-pinball/spin"
	"github.com/veandco/go-sdl2/sdl"
)

func RegisterKeys(eng *spin.Engine) {
	eng.Do(spin.RegisterKeySDL{
		Key:       sdl.K_LEFT,
		EventDown: spin.SwitchEvent{ID: LeftFlipperButton},
	})
	eng.Do(spin.RegisterKeySDL{
		Key:       sdl.K_LEFT,
		Mod:       sdl.KMOD_SHIFT,
		EventDown: spin.SwitchEvent{ID: LeftFireButton},
	})
	eng.Do(spin.RegisterKeySDL{
		Key:       sdl.K_RIGHT,
		EventDown: spin.SwitchEvent{ID: RightFlipperButton},
	})
	eng.Do(spin.RegisterKeySDL{
		Key:       sdl.K_RIGHT,
		Mod:       sdl.KMOD_SHIFT,
		EventDown: spin.SwitchEvent{ID: RightFireButton},
	})
	eng.Do(spin.RegisterKeySDL{
		Key:       sdl.K_RETURN,
		EventDown: spin.SwitchEvent{ID: StartButton},
	})
}

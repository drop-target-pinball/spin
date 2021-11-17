package menu

import "github.com/drop-target-pinball/spin"

const (
	MenuAttractDropTargetPinball = "MenuAttractDropTargetPinball"
	MenuAttractFreePlay          = "MenuAttractFreePlay"
	MenuAttractGameOver          = "MenuAttractGameOver"
	MenuAttractMode              = "MenuAttractMode"
)

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{
		ID:     MenuAttractDropTargetPinball,
		Script: menuAttractDropTargetPinball,
	})
	eng.Do(spin.RegisterScript{
		ID:     MenuAttractFreePlay,
		Script: menuAttractFreePlay,
	})
	eng.Do(spin.RegisterScript{
		ID:     MenuAttractGameOver,
		Script: menuAttractGameOver,
	})
	eng.Do(spin.RegisterScript{
		ID:     MenuAttractMode,
		Script: menuAttractMode,
	})
}

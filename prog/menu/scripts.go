package menu

import "github.com/drop-target-pinball/spin"

const (
	MenuAttractDropTargetPinball = "MenuAttractDropTargetPinball"
	MenuAttractFreePlay          = "MenuAttractFreePlay"
	MenuAttractGameOver          = "MenuAttractGameOver"
	MenuAttractMode              = "MenuAttractMode"
	MenuSelect                   = "MenuSelect"
	MenuSelectGame               = "MenuSelectGame"
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
	eng.Do(spin.RegisterScript{
		ID:     MenuSelect,
		Script: menuSelect,
	})
	eng.Do(spin.RegisterScript{
		ID:     MenuSelectGame,
		Script: menuSelectGame,
	})
}

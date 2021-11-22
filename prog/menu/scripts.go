package menu

import "github.com/drop-target-pinball/spin"

const (
	ScriptAttractMode       = "menu.ScriptAttractMode"
	ScriptDropTargetPinball = "menu.ScriptDropTargetPinball"
	ScriptFreePlay          = "menu.ScriptFreePlay"
	ScriptGameOver          = "menu.ScriptGameOver"
	ScriptGameSelect        = "menu.ScriptGameSelect"
	ScriptSelectMode        = "menu.ScriptSelectMode"
)

func RegisterScripts(eng *spin.Engine) {
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptAttractMode,
	// 	Script: attractModeScript,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptDropTargetPinball,
	// 	Script: dropTargetPinballScript,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptFreePlay,
	// 	Script: freePlayScript,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptGameOver,
	// 	Script: gameOverScript,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptGameSelect,
	// 	Script: gameSelectScript,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptSelectMode,
	// 	Script: selectModeScript,
	// })
}

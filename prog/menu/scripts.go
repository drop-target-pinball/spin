package menu

import "github.com/drop-target-pinball/spin"

const (
	ScriptAttractMode = "menu.ScriptAttractMode"
	ScriptSelectMode  = "menu.ScriptSelectMode"
)

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{
		ID:     ScriptAttractMode,
		Script: attractModeScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSelectMode,
		Script: selectModeScript,
	})
}

package menu

import "github.com/drop-target-pinball/spin"

const (
	ScriptAttractMode = "menu.ScriptAttractMode"
	ScriptSelectGame  = "menu.ScriptSelectGame"
)

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{
		ID:     ScriptAttractMode,
		Script: attractModeScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSelectGame,
		Script: selectGameScript,
	})
}

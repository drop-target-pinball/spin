package menu

import "github.com/drop-target-pinball/spin"

const (
	ScriptAttractMode      = "menu.ScriptAttractMode"
	ScriptAttractModeSlide = "menu.ScriptAttractModeSlide"
	ScriptGameSelected     = "menu.ScriptGameSelected"
	ScriptSelectGame       = "menu.ScriptSelectGame"
)

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{
		ID:     ScriptAttractMode,
		Script: attractModeScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptAttractModeSlide,
		Script: attractModeSlideScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptGameSelected,
		Script: gameSelectedScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSelectGame,
		Script: selectGameScript,
	})
}

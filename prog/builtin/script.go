package builtin

import "github.com/drop-target-pinball/spin"

const (
	ScriptGameStartButton = "ScriptGameStartButton"
	ScriptScore           = "ScriptScore"
)

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{
		ID:     ScriptScore,
		Script: scoreScript,
		Scope:  spin.ScopeGame,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptGameStartButton,
		Script: gameStartButtonScript,
		Scope:  spin.ScopeGame,
	})
}

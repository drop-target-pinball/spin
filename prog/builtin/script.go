package builtin

import "github.com/drop-target-pinball/spin"

const (
	ScriptScore = "ScriptScore"
)

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{
		ID:     ScriptScore,
		Script: scoreScript,
	})
}

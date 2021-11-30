package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/prog/builtin"
)

func directorScript(e spin.Env) {
	e.Do(spin.AddPlayer{})
	e.Do(spin.AdvanceGame{})
	e.Do(spin.PlayScript{ID: builtin.ScriptScore})
	e.Do(spin.PlayScript{ID: ScriptPlunge})
}

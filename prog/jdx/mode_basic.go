package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/prog/builtin"
)

func basicModeScript(e spin.Env) {
	e.Do(spin.PlayScript{ID: builtin.ScriptScore})

	e.Do(spin.PlayScript{ID: ScriptDefaultLeftPopper})
	e.Do(spin.PlayScript{ID: ScriptDefaultLeftShooterLane})
	e.Do(spin.PlayScript{ID: ScriptDefaultRightPopper})
}

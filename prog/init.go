package prog

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/game"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/prog/boot"
	"github.com/drop-target-pinball/spin/prog/jdx"
	"github.com/drop-target-pinball/spin/prog/menu"
	"github.com/drop-target-pinball/spin/prog/sandbox"
)

const (
	ScriptInit = "ScriptInit"
)

func Load(eng *spin.Engine) {
	game.Load(eng)
	jd.Load(eng)
	boot.Load(eng)
	menu.Load(eng)
	jdx.Load(eng)
	sandbox.Load(eng)

	eng.Do(spin.RegisterScript{ID: ScriptInit, Script: scriptInit})
}

func scriptInit(e spin.Env) {
	e.Do(spin.PlayScript{ID: boot.ScriptSplashScreen})
	if _, done := e.WaitFor(spin.Message{ID: boot.MessageDone}); done {
		return
	}

	e.Do(spin.PlayScript{ID: menu.ScriptAttractMode})
	if _, done := e.WaitFor(spin.Message{ID: menu.MessageAttractDone}); done {
		return
	}

	e.Do(spin.PlayScript{ID: menu.ScriptSelectMode})
	// if _, done := e.WaitFor(spin.Message{ID: menu.MessageSelectDone}); done {
	// 	return
	// }
	e.WaitFor(spin.Message{ID: menu.MessageSelectDone})
}

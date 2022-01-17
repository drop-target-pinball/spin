package prog

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/prog/boot"
	"github.com/drop-target-pinball/spin/prog/jdx"
	"github.com/drop-target-pinball/spin/prog/menu"
)

const (
	ScriptInit = "ScriptInit"
)

func Load(eng *spin.Engine) {
	jd.Load(eng)
	boot.Load(eng)
	menu.Load(eng)
	jdx.Load(eng)
	//service.Load(eng)

	eng.Do(spin.RegisterScript{
		ID:     ScriptInit,
		Script: scriptInit,
	})
}

func scriptInit(e *spin.ScriptEnv) {
	defer panic("unexpected end of init")

	e.Do(spin.PlayScript{ID: boot.ScriptSplashScreen})
	if _, done := e.WaitFor(spin.ScriptFinishedEvent{ID: boot.ScriptSplashScreen}); done {
		return
	}

	for {
		e.Do(spin.PlayScript{ID: menu.ScriptAttractMode})
		if _, done := e.WaitFor(spin.ScriptFinishedEvent{ID: menu.ScriptAttractMode}); done {
			return
		}

		for {
			e.Do(spin.PlayScript{ID: menu.ScriptSelectGame})
			evt, done := e.WaitFor(
				spin.Message{ID: menu.MessageGameSelected},
				spin.Message{ID: menu.MessageExit},
			)
			if done {
				return
			}

			if evt == (spin.Message{ID: menu.MessageExit}) {
				break
			}
			e.Do(spin.PlayScript{ID: jdx.ScriptGame})
			if _, done := e.WaitFor(spin.ScriptFinishedEvent{ID: jdx.ScriptProgram}); done {
				return
			}
		}
	}

}

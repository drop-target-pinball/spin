package prog

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/prog/jdx"
)

// const (
// 	ScriptInit = "ScriptInit"
// )

func Load(eng *spin.Engine) {
	jd.Load(eng)
	jdx.Load(eng)
}

// func Load(eng *spin.Engine) {
// 	jd.Load(eng)
// 	boot.Load(eng)
// 	menu.Load(eng)
// 	jdx.Load(eng)
// 	service.Load(eng)

// 	eng.Do(spin.RegisterScript{
// 		ID:     ScriptInit,
// 		Script: scriptInit,
// 		Scope:  spin.ScopeInit,
// 	})
// }

// func scriptInit(e spin.Env) {
// 	defer panic("unexpected end of init")

// 	e.Do(spin.PlayScript{ID: boot.ScriptSplashScreen})
// 	if _, done := e.WaitFor(spin.Message{ID: boot.ScriptSplashScreen}); done {
// 		return
// 	}

// 	for {
// 		e.Do(spin.PlayScript{ID: menu.ScriptAttractMode})
// 		if _, done := e.WaitFor(spin.Message{ID: menu.MessageAttractDone}); done {
// 			return
// 		}

// 		for {
// 			e.Do(spin.PlayScript{ID: menu.ScriptSelectGame})
// 			evt, done := e.WaitFor(
// 				spin.Message{ID: menu.MessageGameSelected},
// 				spin.Message{ID: menu.MessageExit},
// 			)
// 			if done {
// 				return
// 			}

// 			if evt == (spin.Message{ID: menu.MessageExit}) {
// 				break
// 			}
// 			e.Do(spin.PlayScript{ID: jdx.ScriptGame})
// 			if _, done := e.WaitFor(spin.ScriptFinishedEvent{ID: jdx.ScriptProgram}); done {
// 				return
// 			}
// 		}
// 	}

// }

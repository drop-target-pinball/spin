package jd

import (
	"github.com/drop-target-pinball/spin"
)

const (
	ScriptInactiveGlobe               = "jd.ScriptInactiveGlobe"
	ScriptLeftRampShot                = "jd.ScriptLeftRampShot"
	ScriptRaiseDropTargets            = "jd.ScriptRaiseDropTargets"
	ScriptRaiseDropTargetsWhenAllDown = "jd.ScriptRaiseDropTargetsWhenAllDown"
)

func inactiveGlobeScript(e *spin.ScriptEnv) {
	rotations := 0
	running := false

	defer func() {
		if running {
			e.Do(spin.DriverOff{ID: MotorGlobe})
		}
	}()

	for {
		evt, done := e.WaitFor(
			spin.SwitchEvent{ID: SwitchLeftRampToLock},
			spin.SwitchEvent{ID: SwitchGlobePosition2},
		)
		if done {
			return
		}
		switch evt {
		case spin.SwitchEvent{ID: SwitchLeftRampToLock}:
			if !running {
				running = true
				e.Do(spin.DriverOn{ID: MotorGlobe})
			}
			rotations += 3
		case spin.SwitchEvent{ID: SwitchGlobePosition2}:
			if !running {
				continue
			}
			rotations -= 1
			if rotations == 0 {
				running = false
				e.Do(spin.DriverOff{ID: MotorGlobe})
			}
		}
	}
}

func raiseDropTargetsScript(e *spin.ScriptEnv) {
	rv := spin.GetResourceVars(e)
	raise := false
	for _, id := range []string{
		SwitchDropTargetJ,
		SwitchDropTargetU,
		SwitchDropTargetD,
		SwitchDropTargetG,
		SwitchDropTargetE,
	} {
		if rv.Switches[id].Active {
			raise = true
			break
		}
	}
	if raise {
		e.Do(spin.DriverPulse{ID: CoilDropTargetReset})
	}
}

func raiseDropTargetsWhenAllDownScript(e *spin.ScriptEnv) {
	rv := spin.GetResourceVars(e)

	for {
		if _, done := e.WaitFor(
			spin.SwitchEvent{ID: SwitchDropTargetJ},
			spin.SwitchEvent{ID: SwitchDropTargetU},
			spin.SwitchEvent{ID: SwitchDropTargetD},
			spin.SwitchEvent{ID: SwitchDropTargetG},
			spin.SwitchEvent{ID: SwitchDropTargetE},
		); done {
			return
		}

		down := 0
		for _, id := range []string{
			SwitchDropTargetJ,
			SwitchDropTargetU,
			SwitchDropTargetD,
			SwitchDropTargetG,
			SwitchDropTargetE,
		} {
			if rv.Switches[id].Active {
				down += 1
			}
		}

		if down == 5 {
			if done := e.Sleep(500); done {
				return
			}
			e.Do(spin.DriverPulse{ID: CoilDropTargetReset})
		}
	}

}

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{
		ID:     ScriptInactiveGlobe,
		Script: inactiveGlobeScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptRaiseDropTargets,
		Script: raiseDropTargetsScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptRaiseDropTargetsWhenAllDown,
		Script: raiseDropTargetsWhenAllDownScript,
	})
}

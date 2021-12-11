package jd

import (
	"github.com/drop-target-pinball/spin"
)

func inactiveGlobeScript(e spin.Env) {
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

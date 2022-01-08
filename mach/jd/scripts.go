package jd

const (
	ScriptInactiveGlobe               = "jd.ScriptInactiveGlobe"
	ScriptLeftPopperShot              = "jd.ScriptLeftPopperShot"
	ScriptLeftRampShot                = "jd.ScriptLeftRampShot"
	ScriptLeftShooterLaneShot         = "jd.ScriptLeftShooterLaneShot"
	ScriptRaiseDropTargets            = "jd.ScriptRaiseDropTargets"
	ScriptRaiseDropTargetsWhenAllDown = "jd.ScriptRaiseDropTargetsWhenAllDown"
	ScriptRightPopperShot             = "jd.ScriptRightPopperShot"
	ScriptRightRampShot               = "jd.ScriptRightRampShot"
	ScriptTopLeftRampShot             = "jd.ScriptTopLeftRampShot"
	ScriptTopRightRampShot            = "jd.ScriptTopRightRampShot"
)

const (
	ShotLeftPopper      = "jd.ShotLeftPopper"
	ShotLeftRamp        = "jd.ShotLeftRamp"
	ShotLeftShooterLane = "jd.ShotLeftShooterLane"
	ShotRightPopper     = "jd.ShotRightPopper"
	ShotRightRamp       = "jd.ShotRightRamp"
	ShotTopLeftRamp     = "jd.ShotTopLeftRamp"
	ShotTopRightRamp    = "jd.ShotTopRightRamp"
)

// func inactiveGlobeScript(e spin.Env) {
// 	rotations := 0
// 	running := false

// 	defer func() {
// 		if running {
// 			e.Do(spin.DriverOff{ID: MotorGlobe})
// 		}
// 	}()

// 	for {
// 		evt, done := e.WaitFor(
// 			spin.SwitchEvent{ID: SwitchLeftRampToLock},
// 			spin.SwitchEvent{ID: SwitchGlobePosition2},
// 		)
// 		if done {
// 			return
// 		}
// 		switch evt {
// 		case spin.SwitchEvent{ID: SwitchLeftRampToLock}:
// 			if !running {
// 				running = true
// 				e.Do(spin.DriverOn{ID: MotorGlobe})
// 			}
// 			rotations += 3
// 		case spin.SwitchEvent{ID: SwitchGlobePosition2}:
// 			if !running {
// 				continue
// 			}
// 			rotations -= 1
// 			if rotations == 0 {
// 				running = false
// 				e.Do(spin.DriverOff{ID: MotorGlobe})
// 			}
// 		}
// 	}
// }

// func raiseDropTargetsScript(e spin.Env) {
// 	rv := spin.GetResourceVars(e)
// 	raise := false
// 	for _, id := range []string{
// 		SwitchDropTargetJ,
// 		SwitchDropTargetU,
// 		SwitchDropTargetD,
// 		SwitchDropTargetG,
// 		SwitchDropTargetE,
// 	} {
// 		if rv.Switches[id].Active {
// 			raise = true
// 			break
// 		}
// 	}
// 	if raise {
// 		e.Do(spin.DriverPulse{ID: CoilDropTargetReset})
// 	}
// }

// func raiseDropTargetsWhenAllDownScript(e spin.Env) {
// 	rv := spin.GetResourceVars(e)

// 	for {
// 		if _, done := e.WaitFor(
// 			spin.SwitchEvent{ID: SwitchDropTargetJ},
// 			spin.SwitchEvent{ID: SwitchDropTargetU},
// 			spin.SwitchEvent{ID: SwitchDropTargetD},
// 			spin.SwitchEvent{ID: SwitchDropTargetG},
// 			spin.SwitchEvent{ID: SwitchDropTargetE},
// 		); done {
// 			return
// 		}

// 		down := 0
// 		for _, id := range []string{
// 			SwitchDropTargetJ,
// 			SwitchDropTargetU,
// 			SwitchDropTargetD,
// 			SwitchDropTargetG,
// 			SwitchDropTargetE,
// 		} {
// 			if rv.Switches[id].Active {
// 				down += 1
// 			}
// 		}

// 		if down == 5 {
// 			if done := e.Sleep(500 * time.Millisecond); done {
// 				return
// 			}
// 			e.Do(spin.DriverPulse{ID: CoilDropTargetReset})
// 		}
// 	}

// }

// func leftRampShotScript(e spin.Env) {
// 	builtin.ShotSequenceScript(e, []string{SwitchLeftRampEnter, SwitchLeftRampExit}, ShotLeftRamp, 1000*time.Millisecond)
// }

// func leftShooterLaneShotScript(e spin.Env) {
// 	builtin.ShotTrapScript(e, SwitchLeftShooterLane, ShotLeftShooterLane, 250*time.Millisecond)
// }

// func leftPopperShotScript(e spin.Env) {
// 	builtin.ShotTrapScript(e, SwitchLeftPopper, ShotLeftPopper, 250*time.Millisecond)
// }

// func rightPopperShotScript(e spin.Env) {
// 	builtin.ShotTrapScript(e, SwitchRightPopper, ShotRightPopper, 250*time.Millisecond)
// }

// func rightRampShotScript(e spin.Env) {
// 	builtin.ShotSwitchScript(e, SwitchRightRampExit, ShotRightRamp)
// }

// func topLeftRampShotScript(e spin.Env) {
// 	builtin.ShotSwitchScript(e, SwitchTopLeftRampExit, ShotTopLeftRamp)
// }

// func topRightRampShotScript(e spin.Env) {
// 	builtin.ShotSwitchScript(e, SwitchTopRightRampExit, ShotTopRightRamp)
// }

// func RegisterScripts(eng *spin.Engine) {
// 	eng.Do(spin.RegisterScript{
// 		ID:     ScriptInactiveGlobe,
// 		Script: inactiveGlobeScript,
// 		Scope:  spin.ScopeBall,
// 	})
// 	eng.Do(spin.RegisterScript{
// 		ID:     ScriptLeftShooterLaneShot,
// 		Script: leftShooterLaneShotScript,
// 		Scope:  spin.ScopeRoot,
// 	})
// 	eng.Do(spin.RegisterScript{
// 		ID:     ScriptLeftRampShot,
// 		Script: leftRampShotScript,
// 		Scope:  spin.ScopeRoot,
// 	})
// 	eng.Do(spin.RegisterScript{
// 		ID:     ScriptLeftPopperShot,
// 		Script: leftPopperShotScript,
// 		Scope:  spin.ScopeRoot,
// 	})
// 	eng.Do(spin.RegisterScript{
// 		ID:     ScriptRaiseDropTargets,
// 		Script: raiseDropTargetsScript,
// 		Scope:  spin.ScopeBall,
// 	})
// 	eng.Do(spin.RegisterScript{
// 		ID:     ScriptRaiseDropTargetsWhenAllDown,
// 		Script: raiseDropTargetsWhenAllDownScript,
// 		Scope:  spin.ScopeBall,
// 	})
// 	eng.Do(spin.RegisterScript{
// 		ID:     ScriptRightPopperShot,
// 		Script: rightPopperShotScript,
// 		Scope:  spin.ScopeRoot,
// 	})
// 	eng.Do(spin.RegisterScript{
// 		ID:     ScriptRightRampShot,
// 		Script: rightRampShotScript,
// 		Scope:  spin.ScopeRoot,
// 	})
// 	eng.Do(spin.RegisterScript{
// 		ID:     ScriptTopLeftRampShot,
// 		Script: topLeftRampShotScript,
// 		Scope:  spin.ScopeRoot,
// 	})
// 	eng.Do(spin.RegisterScript{
// 		ID:     ScriptTopRightRampShot,
// 		Script: topRightRampShotScript,
// 		Scope:  spin.ScopeRoot,
// 	})

// 	eng.Do(spin.PlayScript{ID: ScriptLeftPopperShot})
// 	eng.Do(spin.PlayScript{ID: ScriptLeftRampShot})
// 	eng.Do(spin.PlayScript{ID: ScriptLeftShooterLaneShot})
// 	eng.Do(spin.PlayScript{ID: ScriptRightPopperShot})
// 	eng.Do(spin.PlayScript{ID: ScriptRightRampShot})
// 	eng.Do(spin.PlayScript{ID: ScriptTopLeftRampShot})
// 	eng.Do(spin.PlayScript{ID: ScriptTopRightRampShot})
// }

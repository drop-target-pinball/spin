package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/proc"
)

const (
	MessageStartChainMode = "jdx.StartChainMode"
)

func chainScript(e *spin.ScriptEnv) {
	vars := GetVars(e)

	defer func() {
		if vars.SelectedMode == ModeNone {
			return
		}
		e.Do(spin.DriverOff{ID: ModeLamps[vars.SelectedMode]})
		if vars.StartModeLeft {
			e.Do(spin.DriverOff{ID: jd.LampLeftModeStart})
		} else {
			e.Do(spin.DriverOff{ID: jd.LampRightModeStart})
		}
	}()

	e.NewCoroutine(selectModeScript)
	nextChain(e)

	for _, mode := range Modes {
		if vars.AwardedModes&mode != 0 {
			e.Do(spin.DriverOn{ID: ModeLamps[mode]})
		}
		if vars.SelectedMode == mode {
			e.Do(proc.DriverSchedule{ID: ModeLamps[mode], Schedule: proc.BlinkSchedule})
		}
	}

	if vars.AwardedModes == AllChainModes {
		return
	}

	if vars.LocksReady > vars.BallsLocked {
		vars.StartModeLeft = false
	} else {
		vars.StartModeLeft = true
	}

	for {
		modeStartLamp := jd.LampRightModeStart
		if vars.StartModeLeft {
			modeStartLamp = jd.LampLeftModeStart
		}
		e.Do(proc.DriverSchedule{ID: modeStartLamp, Schedule: proc.BlinkSchedule})

		for {
			evt, done := e.WaitFor(
				spin.SwitchEvent{ID: jd.SwitchLeftRampExit},
				spin.SwitchEvent{ID: jd.SwitchRightPopper},
				spin.Message{ID: MessageStartChainMode},
			)
			if done {
				return
			}
			if evt == (spin.SwitchEvent{ID: jd.SwitchLeftRampExit}) && vars.StartModeLeft {
				break
			}
			if evt == (spin.SwitchEvent{ID: jd.SwitchRightPopper}) && !vars.StartModeLeft {
				break
			}
			if evt == (spin.Message{ID: MessageStartChainMode}) {
				break
			}
		}
		e.Do(spin.DriverOff{ID: modeStartLamp})

		vars.AwardedModes |= vars.SelectedMode

		if vars.SelectedMode == ModeBadImpersonator {
			e.Do(spin.StopScript{ID: ScriptLightBallLock})
		}
		if vars.SelectedMode == ModeBlackout {
			e.Do(spin.StopScript{ID: ScriptBallLock})
		}

		e.Do(spin.DriverOn{ID: ModeLamps[vars.SelectedMode]})
		e.Do(spin.PlayScript{ID: ModeScripts[vars.SelectedMode]})

		if _, done := e.WaitFor(spin.ScriptFinishedEvent{ID: ModeScripts[vars.SelectedMode]}); done {
			return
		}
		e.Do(spin.StopScriptGroup{ID: spin.ScriptGroupMode})

		if vars.SelectedMode == ModeBadImpersonator {
			e.Do(spin.PlayScript{ID: ScriptLightBallLock})
		}
		if vars.SelectedMode == ModeBlackout {
			e.Do(spin.PlayScript{ID: ScriptBallLock})
		}

		if vars.AwardedModes == AllChainModes {
			break
		}
		nextChain(e)
		vars.StartModeLeft = !vars.StartModeLeft
		if vars.LocksReady > vars.BallsLocked {
			vars.StartModeLeft = false
		}
	}
}

func selectModeScript(e *spin.ScriptEnv) {
	vars := GetVars(e)

	for {
		evt, done := e.WaitFor(
			spin.SwitchEvent{ID: jd.SwitchLeftFireButton},
			spin.SwitchEvent{ID: jd.SwitchRightFireButton},
		)
		if done {
			return
		}

		if evt == (spin.SwitchEvent{ID: jd.SwitchLeftFireButton}) {
			if vars.Mode == ModeAirRaid {
				continue
			}
			prevChain(e)
		}
		if evt == (spin.SwitchEvent{ID: jd.SwitchRightFireButton}) {
			if vars.Mode == ModePlunge {
				continue
			}
			nextChain(e)
		}
	}
}

func nextChain(e *spin.ScriptEnv) {
	vars := GetVars(e)
	previous := vars.SelectedMode
	next := previous
	for {
		next = next << 1
		if next > MaxChainMode {
			next = MinChainMode
		}
		if next&vars.AwardedModes == 0 {
			break
		}
	}
	if previous&vars.AwardedModes == 0 {
		e.Do(spin.DriverOff{ID: ModeLamps[previous]})
	}
	e.Do(proc.DriverSchedule{ID: ModeLamps[next], Schedule: proc.BlinkSchedule})
	vars.SelectedMode = next
}

func prevChain(e *spin.ScriptEnv) {
	vars := GetVars(e)
	previous := vars.SelectedMode
	next := previous
	for {
		next = next >> 1
		if next < MinChainMode {
			next = MaxChainMode
		}
		if next&vars.AwardedModes == 0 {
			break
		}
	}
	if previous&vars.AwardedModes == 0 {
		e.Do(spin.DriverOff{ID: ModeLamps[previous]})
	}
	e.Do(proc.DriverSchedule{ID: ModeLamps[next], Schedule: proc.BlinkSchedule})
	vars.SelectedMode = next
}

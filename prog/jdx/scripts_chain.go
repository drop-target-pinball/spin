package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

const (
	MessageStartChainMode = "jdx.StartChainMode"
)

func chainScript(e *spin.ScriptEnv) {
	vars := GetVars(e)

	e.NewCoroutine(selectModeScript)

	for _, mode := range Modes {
		if vars.AwardedModes&mode != 0 {
			e.Do(spin.DriverOn{ID: ModeLamps[mode]})
		}
		if vars.SelectedMode == mode {
			e.Do(spin.DriverBlink{ID: ModeLamps[mode]})
		}
	}

	if vars.AwardedModes == AllChainModes {
		return
	}
	vars.StartModeLeft = true

	for {
		modeStartLamp := jd.LampRightModeStart
		if vars.StartModeLeft {
			modeStartLamp = jd.LampLeftModeStart
		}
		e.Do(spin.DriverBlink{ID: modeStartLamp})

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
		e.Do(spin.DriverOn{ID: ModeLamps[vars.SelectedMode]})
		e.Do(spin.PlayScript{ID: ModeScripts[vars.SelectedMode]})

		if _, done := e.WaitFor(spin.ScriptFinishedEvent{ID: ModeScripts[vars.SelectedMode]}); done {
			return
		}
		e.Do(spin.StopScriptGroup{ID: spin.ScriptGroupMode})
		if vars.AwardedModes == AllChainModes {
			break
		}
		nextChain(e)
		vars.StartModeLeft = !vars.StartModeLeft
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
	next := 0
	for {
		next = previous << 1
		if next > MaxChainMode {
			next = MinChainMode
		}
		if next&vars.AwardedModes == 0 {
			break
		}
	}
	e.Do(spin.DriverOff{ID: ModeLamps[previous]})
	e.Do(spin.DriverBlink{ID: ModeLamps[next]})
	vars.SelectedMode = next
}

func prevChain(e *spin.ScriptEnv) {
	vars := GetVars(e)
	previous := vars.SelectedMode
	next := 0
	for {
		next = previous >> 1
		if next < MinChainMode {
			next = MaxChainMode
		}
		if next&vars.AwardedModes == 0 {
			break
		}
	}
	e.Do(spin.DriverOff{ID: ModeLamps[previous]})
	e.Do(spin.DriverBlink{ID: ModeLamps[next]})
	vars.SelectedMode = next
}

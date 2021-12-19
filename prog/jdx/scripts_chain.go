package jdx

import (
	"context"
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/prog/builtin"
)

const (
	ShotStartLeftMode  = "jdx.ShotStartLeftMode"
	ShotStartRightMode = "jdx.ShotStartRightMode"
)

type modeStartConfig struct {
	lamp       string
	shot       string
	shotScript func(spin.Env)
}

const (
	modeLeft  = true
	modeRight = false
)

var modeStartConfigs = map[bool]modeStartConfig{
	modeLeft: {
		lamp:       jd.LampLeftModeStart,
		shot:       ShotStartLeftMode,
		shotScript: startLeftModeShotScript,
	},
	modeRight: {
		lamp:       jd.LampRightModeStart,
		shot:       ShotStartRightMode,
		shotScript: startRightModeShotScript,
	},
}

func waitForModeStart(e spin.Env, parent context.Context, side bool) bool {
	ctx, cancel := e.Derive()
	defer cancel()

	m := modeStartConfigs[side]
	e.Do(spin.DriverPWM{ID: m.lamp, On: 127, Off: 127})
	e.NewCoroutine(ctx, m.shotScript)
	e.NewCoroutine(ctx, selectModeScript)
	_, done := e.WaitFor(spin.ShotEvent{ID: m.shot})
	e.Do(spin.DriverOff{ID: m.lamp})
	return done
}

func selectModeScript(e spin.Env) {
	rv := spin.ResourceVars(e)
	vars := GetVars(e)

	for {
		evt, done := e.WaitFor(
			spin.SwitchEvent{ID: jd.SwitchLeftFireButton},
			spin.SwitchEvent{ID: jd.SwitchRightFireButton},
		)
		if done {
			return
		}

		previous := vars.SelectedMode
		next := 0
		if evt == (spin.SwitchEvent{ID: jd.SwitchLeftFireButton}) {
			if rv.Switches[jd.SwitchLeftShooterLane].Active {
				continue
			}
			for {
				next = previous >> 1
				if next < MinMode {
					next = MaxMode
				}
				if next&vars.AwardedModes == 0 {
					break
				}
			}
		}
		if evt == (spin.SwitchEvent{ID: jd.SwitchRightFireButton}) {
			if rv.Switches[jd.SwitchRightShooterLane].Active {
				continue
			}
			for {
				next = previous << 1
				if next > MaxMode {
					next = MinMode
				}
				if next&vars.AwardedModes == 0 {
					break
				}
			}
		}
		e.Do(spin.DriverOff{ID: ModeLamps[previous]})
		e.Do(spin.DriverPWM{ID: ModeLamps[next], On: 127, Off: 127})
		vars.SelectedMode = next
	}
}

func startLeftModeShotScript(e spin.Env) {
	builtin.ShotSequenceScript(e,
		[]string{
			jd.SwitchLeftRampEnter,
			jd.SwitchLeftRampExit,
		},
		ShotStartLeftMode,
		2*time.Second)
}

func startRightModeShotScript(e spin.Env) {
	builtin.ShotSwitchScript(e,
		jd.SwitchRightRampExit,
		ShotStartRightMode)
}

func chainScript(e spin.Env) {
	ctx, cancel := e.Derive()
	defer cancel()

	vars := GetVars(e)

	for _, mode := range Modes {
		if vars.AwardedModes&mode != 0 {
			e.Do(spin.DriverOn{ID: ModeLamps[mode]})
		}
		if vars.SelectedMode == mode {
			e.Do(spin.DriverPWM{ID: ModeLamps[mode], On: 127, Off: 127})
		}
	}

	if vars.AwardedModes == AllModes {
		return
	}
	modeSide := modeLeft

	for {
		if done := waitForModeStart(e, ctx, modeSide); done {
			return
		}

		vars.AwardedModes |= vars.SelectedMode
		e.Do(spin.StopScope{ID: spin.ScopeMode})
		e.Do(spin.DriverOn{ID: ModeLamps[vars.SelectedMode]})
		e.Do(spin.PlayScript{ID: ModeScripts[vars.SelectedMode]})

		if _, done := e.WaitFor(spin.ScriptFinishedEvent{ID: ModeScripts[vars.SelectedMode]}); done {
			return
		}

		e.Do(spin.StopScope{ID: spin.ScopeMode})
		e.Do(spin.PlayScript{ID: ScriptBasicMode})
		if vars.AwardedModes == AllModes {
			break
		}
		modeSide = !modeSide
	}
}

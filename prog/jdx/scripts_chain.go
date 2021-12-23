package jdx

import (
	"context"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

const (
	ShotStartLeftMode  = "jdx.ShotStartLeftMode"
	ShotStartRightMode = "jdx.ShotStartRightMode"
)

type modeStartConfig struct {
	lamp string
	shot string
}

const (
	modeLeft  = true
	modeRight = false
)

var modeStartConfigs = map[bool]modeStartConfig{
	modeLeft: {
		lamp: jd.LampLeftModeStart,
		shot: jd.ShotLeftRamp,
	},
	modeRight: {
		lamp: jd.LampRightModeStart,
		shot: jd.ShotRightPopper,
	},
}

func waitForModeStart(e spin.Env, parent context.Context, side bool) bool {
	ctx, cancel := e.Derive()
	defer cancel()

	m := modeStartConfigs[side]
	e.Do(spin.DriverBlink{ID: m.lamp})
	e.NewCoroutine(ctx, selectModeScript)
	_, done := e.WaitFor(spin.ShotEvent{ID: m.shot})
	e.Do(spin.DriverOff{ID: m.lamp})
	return done
}

func selectModeScript(e spin.Env) {
	rv := spin.ResourceVars(e)

	for {
		evt, done := e.WaitFor(
			spin.SwitchEvent{ID: jd.SwitchLeftFireButton},
			spin.SwitchEvent{ID: jd.SwitchRightFireButton},
		)
		if done {
			return
		}

		if evt == (spin.SwitchEvent{ID: jd.SwitchLeftFireButton}) {
			if rv.Switches[jd.SwitchLeftShooterLane].Active {
				continue
			}
			prevChain(e)
		}
		if evt == (spin.SwitchEvent{ID: jd.SwitchRightFireButton}) {
			if rv.Switches[jd.SwitchRightShooterLane].Active {
				continue
			}
			nextChain(e)
		}
	}
}

func nextChain(e spin.Env) {
	vars := GetVars(e)
	previous := vars.SelectedMode
	next := 0
	for {
		next = previous << 1
		if next > MaxMode {
			next = MinMode
		}
		if next&vars.AwardedModes == 0 {
			break
		}
	}
	e.Do(spin.DriverOff{ID: ModeLamps[previous]})
	e.Do(spin.DriverBlink{ID: ModeLamps[next]})
	vars.SelectedMode = next
}

func prevChain(e spin.Env) {
	vars := GetVars(e)
	previous := vars.SelectedMode
	next := 0
	for {
		next = previous >> 1
		if next < MinMode {
			next = MaxMode
		}
		if next&vars.AwardedModes == 0 {
			break
		}
	}
	e.Do(spin.DriverOff{ID: ModeLamps[previous]})
	e.Do(spin.DriverBlink{ID: ModeLamps[next]})
	vars.SelectedMode = next
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
			e.Do(spin.DriverBlink{ID: ModeLamps[mode]})
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
		nextChain(e)
		modeSide = !modeSide
	}
}

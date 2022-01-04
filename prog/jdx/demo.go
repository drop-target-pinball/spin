package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

func demoScript(e spin.Env) {
	spin.NewSequencer().
		Do(spin.PlayScript{ID: ScriptGame}).
		Sleep(11_000).
		Post(spin.SwitchEvent{ID: jd.SwitchRightFireButton}).
		Sleep(3_000).
		Run(e)

	spin.NewSequencer().
		Post(spin.ShotEvent{ID: jd.ShotLeftRamp}).
		Sleep(7_000).
		Post(spin.ShotEvent{ID: jd.ShotRightRamp}).
		Sleep(3_000).
		Post(spin.ShotEvent{ID: jd.ShotLeftRamp}).
		Sleep(3_000).
		Post(spin.ShotEvent{ID: jd.ShotRightRamp}).
		Run(e)

}

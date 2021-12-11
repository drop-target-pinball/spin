package jd

import "github.com/drop-target-pinball/spin"

const (
	ScriptInactiveGlobe = "jd.ScriptInactiveGlobe"
)

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{
		ID:     ScriptInactiveGlobe,
		Script: inactiveGlobeScript,
	})
}

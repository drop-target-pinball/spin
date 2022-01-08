package jd

import "github.com/drop-target-pinball/spin"

func Load(eng *spin.Engine) {
	RegisterAuto(eng)
	RegisterCoils(eng)
	RegisterFlashers(eng)
	RegisterKeys(eng)
	RegisterLamps(eng)
	RegisterMagnets(eng)
	RegisterMotors(eng)
	//RegisterScripts(eng)
	RegisterSwitches(eng)
}

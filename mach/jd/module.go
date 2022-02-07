package jd

import "github.com/drop-target-pinball/spin"

func Load(eng *spin.Engine) {
	eng.Do(spin.RegisterConsole{
		ID: "",
		//Image: "console/judge-dredd-540x720.jpg",
		Image: "console/judge-dredd-648x864.jpg",
	})
	RegisterAuto(eng)
	RegisterCoils(eng)
	RegisterFlashers(eng)
	RegisterKeys(eng)
	RegisterLamps(eng)
	RegisterMagnets(eng)
	RegisterMotors(eng)
	RegisterScripts(eng)
	RegisterSwitches(eng)

}

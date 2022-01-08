package boot

import "github.com/drop-target-pinball/spin"

func Load(eng *spin.Engine) {
	RegisterMusic(eng)
	//RegisterScripts(eng)
}

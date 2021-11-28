package service

import "github.com/drop-target-pinball/spin"

func Load(eng *spin.Engine) {
	RegisterScripts(eng)
}

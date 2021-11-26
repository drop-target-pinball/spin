package game

import "github.com/drop-target-pinball/spin"

func Load(eng *spin.Engine) {
	Reset()
	RegisterFonts(eng)
	// 	RegisterMusic(eng)
	RegisterScripts(eng)
}

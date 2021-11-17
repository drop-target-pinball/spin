package menu

import "github.com/drop-target-pinball/spin"

func Load(eng *spin.Engine) {
	eng.Namespaces.Create(spin.System)
	RegisterFonts(eng)
	RegisterMusic(eng)
	RegisterScripts(eng)
	RegisterSounds(eng)
}

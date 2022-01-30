package gamepad

import "github.com/drop-target-pinball/spin"

func Load(eng *spin.Engine) {
	RegisterGamePad(eng)
}

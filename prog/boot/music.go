package boot

import "github.com/drop-target-pinball/spin"

const (
	BootTheme = "BootTheme"
)

func RegisterMusic(eng *spin.Engine) {
	eng.Do(spin.RegisterMusic{
		ID:   BootTheme,
		Path: "boot/boot.ogg",
	})
}

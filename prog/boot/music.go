package boot

import "github.com/drop-target-pinball/spin"

const (
	MusicSplashScreen = "boot.MusicSplashScreen"
)

func RegisterMusic(eng *spin.Engine) {
	eng.Do(spin.RegisterMusic{
		ID:   MusicSplashScreen,
		Path: "boot/boot.ogg",
	})
}

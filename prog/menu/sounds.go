package menu

import "github.com/drop-target-pinball/spin"

const (
	SoundSelectScroll = "menu.SoundSelectScroll"
)

func RegisterSounds(eng *spin.Engine) {
	eng.Do(spin.RegisterSound{
		ID:   SoundSelectScroll,
		Path: "smb/smb2_scroll.wav",
	})
}

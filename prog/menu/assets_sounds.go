package menu

import "github.com/drop-target-pinball/spin"

const (
	SoundScroll = "menu.SoundScroll"
	SoundSelect = "menu.SoundSelect"
)

func RegisterSounds(eng *spin.Engine) {
	eng.Do(spin.RegisterSound{
		ID:   SoundScroll,
		Path: "smb/smb2_scroll.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundSelect,
		Path: "smb/smb2_select.wav",
	})
}

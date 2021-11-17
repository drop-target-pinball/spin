package menu

import "github.com/drop-target-pinball/spin"

const (
	SMB2Scroll = "smb2_scroll"
)

func RegisterSounds(eng *spin.Engine) {
	eng.Do(spin.RegisterSound{
		ID:   SMB2Scroll,
		Path: "smb/smb2_scroll.wav",
	})
}

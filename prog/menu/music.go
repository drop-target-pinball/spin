package menu

import "github.com/drop-target-pinball/spin"

const (
	SMB2CharSelect = "SMB2CharSelect"
)

func RegisterMusic(eng *spin.Engine) {
	eng.Do(spin.RegisterMusic{
		ID:   SMB2CharSelect,
		Path: "smb/smb2_char_select.ogg",
	})
}

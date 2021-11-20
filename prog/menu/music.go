package menu

import "github.com/drop-target-pinball/spin"

const (
	MusicSelectMode = "menu.MusicSelectMode"
)

func RegisterMusic(eng *spin.Engine) {
	eng.Do(spin.RegisterMusic{
		ID:   MusicSelectMode,
		Path: "smb/smb2_char_select.ogg",
	})
}

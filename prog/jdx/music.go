package jdx

import "github.com/drop-target-pinball/spin"

const (
	MainTheme      = "MainTheme"
	MultiballTheme = "MultiballTheme"
	ModeTheme1     = "ModeTheme1"
)

func RegisterMusic(eng *spin.Engine) {
	eng.Do(spin.RegisterMusic{
		ID:   MainTheme,
		Path: "jd/pinsound/music/000002-main_theme/main_the__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterMusic{
		ID:   ModeTheme1,
		Path: "jd/pinsound/music/000004-air_raid/air_raid__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterMusic{
		ID:   MultiballTheme,
		Path: "jd/pinsound/music/000009-multi_ball/multi_ba__LEGACY_AUD.wav",
	})
}

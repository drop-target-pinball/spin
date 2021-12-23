package jdx

import "github.com/drop-target-pinball/spin"

const (
	MusicMain       = "jdx.MusicMain"
	MusicMatch      = "jdx.MusicMatch"
	MusicMatchHit   = "jdx.MusicMatchHit"
	MusicMultiball  = "jdx.MusicMultiball"
	MusicMode1      = "jdx.MusicMode1"
	MusicMode2      = "jdx.MusicMode2"
	MusicPlungeLoop = "jdx.MusicPlungeLoop"
)

func RegisterMusic(eng *spin.Engine) {
	eng.Do(spin.RegisterMusic{
		ID:   MusicMain,
		Path: "jd-pinsound/music/000002-main_theme/main_the__LEGACY_AUD-fixed.wav",
	})
	eng.Do(spin.RegisterMusic{
		ID:   MusicMatch,
		Path: "jd-pinsound/music/000010-end_of_game_match/end_of_g__LEGACY_AUD-fixed.wav",
	})
	eng.Do(spin.RegisterMusic{
		ID:   MusicMatchHit,
		Path: "jd-pinsound/sfx/000161-end_of_game/end_of_g__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterMusic{
		ID:   MusicMode1,
		Path: "jd-pinsound/music/000004-air_raid/air_raid__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterMusic{
		ID:   MusicMode2,
		Path: "jd-pinsound/music/000003-waiting_sound/waiting___LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterMusic{
		ID:   MusicMultiball,
		Path: "jd-pinsound/music/000009-multi_ball/multi_ba__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterMusic{
		ID:   MusicPlungeLoop,
		Path: "jd-pinsound/music/000001-back_ground/back_gro__LEGACY_AUD.wav",
	})
}

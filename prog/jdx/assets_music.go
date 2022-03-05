package jdx

import "github.com/drop-target-pinball/spin"

const (
	MusicBadImpersonator = "jdx.MusicBadImpersonator"
	MusicMain            = "jdx.MusicMain"
	MusicMatch           = "jdx.MusicMatch"
	MusicMatchHit        = "jdx.MusicMatchHit"
	MusicMultiballIntro  = "jdx.MusicMultiballIntro"
	MusicMultiballLoop   = "jdx.MusicMultiballLoop"
	MusicMode1           = "jdx.MusicMode1"
	MusicMode2           = "jdx.MusicMode2"
	MusicPlungeLoop      = "jdx.MusicPlungeLoop"
	MusicSuperGame       = "jdx.MusicSuperGame"
	MusicSuperGame2      = "jdx.MusicSuperGame2"
)

func RegisterMusic(eng *spin.Engine) {
	eng.Do(spin.RegisterMusic{
		ID:   MusicBadImpersonator,
		Path: "jd-pinsound/music/000005-bad_impersonator/bad_impe__LEGACY_AUD.wav",
	})
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
		ID:   MusicMultiballIntro,
		Path: "jd-pinsound/music/000009-multi_ball/multi_ba__LEGACY_AUD-intro.wav",
	})
	eng.Do(spin.RegisterMusic{
		ID:   MusicMultiballLoop,
		Path: "jd-pinsound/music/000009-multi_ball/multi_ba__LEGACY_AUD-loop.wav",
	})
	eng.Do(spin.RegisterMusic{
		ID:   MusicPlungeLoop,
		Path: "jd-pinsound/music/000001-back_ground/back_gro__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterMusic{
		ID:   MusicSuperGame,
		Path: "jd-pinsound/music/000006-weappon_sound/weappon___LEGACY_AUD-loop1.wav",
	})
	eng.Do(spin.RegisterMusic{
		ID:   MusicSuperGame2,
		Path: "jd-pinsound/music/000006-weappon_sound/weappon___LEGACY_AUD-loop2.wav",
	})
}

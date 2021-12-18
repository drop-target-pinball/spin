package jdx

import "github.com/drop-target-pinball/spin"

const (
	SoundBallLost        = "jdx.SoundBallLost"
	SoundGunLoadSniper   = "jdx.SoundGunLoadSniper"
	SoundGunFire         = "jdx.SoundGunFire"
	SoundMotorcycleStart = "jdx.SoundMotorcycleStart"
	SoundReturnLane      = "jdx.SoundReturnLane"
	SoundSling           = "jdx.SoundSling"
	SoundSniperSplat     = "jdx.SoundSniperSplat"
	SoundSuccess         = "jdx.SoundSuccess"
)

func RegisterSounds(eng *spin.Engine) {
	eng.Do(spin.RegisterSound{
		ID:   SoundBallLost,
		Path: "jd-pinsound/sfx/000112-ball_lost/ball_los__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundGunLoadSniper,
		Path: "jd-pinsound/sfx/000132-gun_loading_1/gun_load__LEGACY.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundGunFire,
		Path: "jd-pinsound/sfx/000133-gun_fire/gun_fire__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundMotorcycleStart,
		Path: "jd-pinsound/sfx/000103-motorcylce_start/motorcyl__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundReturnLane,
		Path: "jd-pinsound/sfx/000122-explosion_3/explosio__LEGACY.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundSling,
		Path: "jd-pinsound/sfx/000107-check_point/check_po__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundSniperSplat,
		Path: "jd-pinsound/sfx/000076-hurt_3/hurt_3____LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundSuccess,
		Path: "jd-pinsound/jingle/000153-accepted_sound/accepted__LEGACY_AUD.wav",
	})
}

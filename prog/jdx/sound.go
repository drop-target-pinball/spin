package jdx

import "github.com/drop-target-pinball/spin"

const (
	Success       = "success"
	GunLoadSniper = "gun-load-sniper"
	GunFire       = "gun-fire"
)

func RegisterSounds(eng *spin.Engine) {
	eng.Do(spin.RegisterSound{
		ID:   GunLoadSniper,
		Path: "jd/pinsound/sfx/000132-gun_loading_1/gun_load__LEGACY.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   GunFire,
		Path: "jd/pinsound/sfx/000133-gun_fire/gun_fire__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   Success,
		Path: "jd/pinsound/jingle/000153-accepted_sound/accepted__LEGACY_AUD.wav",
	})
}

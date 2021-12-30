package jdx

import "github.com/drop-target-pinball/spin"

const (
	SoundBallLost          = "jdx.SoundBallLost"
	SoundBonus             = "jdx.SoundBonus"
	SoundGunLoadSniper     = "jdx.SoundGunLoadSniper"
	SoundGunFire           = "jdx.SoundGunFire"
	SoundMeltdownCracking  = "jdx.SoundMeltdownCracking"
	SoundMeltdownExplosion = "jdx.SoundMeltdownExplosion"
	SoundMeltdownKlaxon    = "jdx.SoundMeltdownKlaxon"
	SoundMotorcycleStart   = "jdx.SoundMotorcycleStart"
	SoundMotorRev          = "jdx.SoundMotorRev"
	SoundPoliceSiren       = "jdx.SoundPoliceSiren"
	SoundPursuitEngine     = "jdx.SoundPursuitEngine"
	SoundPursuitExplosion  = "jdx.SoundPursuitExplosion"
	SoundPursuitMissile    = "jdx.SoundPursuitMissile"
	SoundReturnLane        = "jdx.SoundReturnLane"
	SoundSling             = "jdx.SoundSling"
	SoundSniperSplat       = "jdx.SoundSniperSplat"
	SoundSuccess           = "jdx.SoundSuccess"
	SoundTankDestroyed     = "jdx.SoundTankDestroyed"
	SoundTankFire          = "jdx.SoundTankFire"
	SoundTireSqueal1       = "jdx.SoundTireSqueal1"
	SoundTireSqueal2       = "jdx.SoundTireSqueal2"
)

func RegisterSounds(eng *spin.Engine) {
	eng.Do(spin.RegisterSound{
		ID:   SoundBallLost,
		Path: "jd-pinsound/sfx/000112-ball_lost/ball_los__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundBonus,
		Path: "jd-pinsound/sfx/000063-piano_1/piano_1___LEGACY_AUD.wav",
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
		ID:   SoundMeltdownCracking,
		Path: "jd-pinsound/sfx/000135-wood_broken/wood_bro__LEGACY.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundMeltdownExplosion,
		Path: "jd-pinsound/sfx/000128-very_big_explosion/very_big__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundMeltdownKlaxon,
		Path: "jd-pinsound/sfx/000134-alarm/alarm__L__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundMotorcycleStart,
		Path: "jd-pinsound/sfx/000103-motorcylce_start/motorcyl__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundMotorRev,
		Path: "jd-pinsound/sfx/000074-motor_noise/motor_no__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundPoliceSiren,
		Path: "jd-pinsound/sfx/000141-police_horn/police_h__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundPursuitEngine,
		Path: "jd-pinsound/sfx/000042-pursuit_mode/pursuit___LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundPursuitExplosion,
		Path: "jd-pinsound/sfx/000066-little_explosion/little_e__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundPursuitMissile,
		Path: "jd-pinsound/sfx/000100-flash_2/flash_2___LEGACY_AUD.wav",
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
	eng.Do(spin.RegisterSound{
		ID:   SoundTankDestroyed,
		Path: "jd-pinsound/sfx/000066-little_explosion/little_e__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundTankFire,
		Path: "jd-pinsound/sfx/000066-little_explosion/little_e__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundTireSqueal1,
		Path: "jd-pinsound/sfx/000072-rustling_tire/rustling__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundTireSqueal2,
		Path: "jd-pinsound/sfx/000073-rustling_tire_1/rustling_2_LEGACY_AUD.wav",
	})
}

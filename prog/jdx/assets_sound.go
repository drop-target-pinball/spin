package jdx

import "github.com/drop-target-pinball/spin"

const (
	SoundAnnounce               = "jdx.SoundAnnounce"
	SoundApplause               = "jdx.SoundApplause"
	SoundBadImpersonatorGunFire = "jdx.SoundBadImpersonatorGunFire"
	SoundBadImpersonatorThrow   = "jdx.SoundBadImpersonatorThrow"
	SoundBallLock               = "jdx.SoundBallLock"
	SoundBallLost               = "jdx.SoundBallLost"
	SoundBlackoutJackpot        = "jdx.SoundBlackoutJackpot"
	SoundBonus                  = "jdx.SoundBonus"
	SoundDeadworldExplosion     = "jdx.SoundDeadworldExplosion"
	SoundDing                   = "jdx.SoundDing"
	SoundDropTargetLitHit1      = "jdx.SoundDropTargetLitHit1"
	SoundDropTargetLitHit2      = "jdx.SoundDropTargetLitHit2"
	SoundDropTargetLitHit3      = "jdx.SoundDropTargetLitHit3"
	SoundDropTargetLitHit4      = "jdx.SoundDropTargetLitHit4"
	SoundDropTargetLitHit5      = "jdx.SoundDropTargetLitHit5"
	SoundGunLoadSniper          = "jdx.SoundGunLoadSniper"
	SoundGunFire                = "jdx.SoundGunFire"
	SoundLeftPopperLaser        = "jdx.SoundLeftPopperLaser"
	SoundLeftPost               = "jdx.SoundLeftPost"
	SoundLeftShooterLaneFire    = "jdx.SoundLeftShooterLaneFire"
	SoundManhuntAutoFire        = "jdx.SoundManhuntAutoFire"
	SoundManhuntSingleFire      = "jdx.SoundManhuntSingleFire"
	SoundMeltdownCracking       = "jdx.SoundMeltdownCracking"
	SoundMeltdownExplosion      = "jdx.SoundMeltdownExplosion"
	SoundMeltdownKlaxon         = "jdx.SoundMeltdownKlaxon"
	SoundMotorcycleRamp         = "jdx.SoundMotorcycleRamp"
	SoundMotorcycleStart        = "jdx.SoundMotorcycleStart"
	SoundMotorRev               = "jdx.SoundMotorRev"
	SoundMultiballJackpot       = "jdx.SoundMultiballJackpot"
	SoundMystery                = "jdx.SoundMystery"
	SoundPoint                  = "jdx.SoundPoint"
	SoundPoliceSiren            = "jdx.SoundPoliceSiren"
	SoundPursuitEngine          = "jdx.SoundPursuitEngine"
	SoundPursuitExplosion       = "jdx.SoundPursuitExplosion"
	SoundPursuitMissile         = "jdx.SoundPursuitMissile"
	SoundReturnLane             = "jdx.SoundReturnLane"
	SoundRightPost              = "jdx.SoundRightPost"
	SoundRightRamp              = "jdx.SoundRightRamp"
	SoundShock                  = "jdx.SoundShock"
	SoundSling                  = "jdx.SoundSling"
	SoundSafecrackerExplosion   = "jdx.SoundSafecrackerExplosion"
	SoundSafecrackerGunFire1    = "jdx.SoundSafecrackerGunFire1"
	SoundSafecrackerGunFire2    = "jdx.SoundSafecrackerGunFire2"
	SoundSafecrackerGunFire3    = "jdx.SoundSafecrackerGunFire3"
	SoundSafecrackerLaserFire   = "jdx.SoundSafecrackerLaserFire"
	SoundSafecrackerTankFire    = "jdx.SoundSafecrackerTankFire"
	SoundSniperSplat            = "jdx.SoundSniperSplat"
	SoundSniperTower            = "jdx.SoundSniperTower"
	SoundSnore                  = "jdx.SoundSnore"
	SoundSuccess                = "jdx.SoundSuccess"
	SoundTankDestroyed          = "jdx.SoundTankDestroyed"
	SoundTankFire               = "jdx.SoundTankFire"
	SoundTireSqueal1            = "jdx.SoundTireSqueal1"
	SoundTireSqueal2            = "jdx.SoundTireSqueal2"
	SoundTopLeftRamp            = "jdx.SoundTopLeftRamp"
	SoundWalking                = "jdx.SoundWalking"
)

func RegisterSounds(eng *spin.Engine) {
	eng.Do(spin.RegisterSound{
		ID:   SoundAnnounce,
		Path: "jd-pinsound/sfx/000157-wrong/wrong__L__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundApplause,
		Path: "afm-pinsound/sfx/0107231306-sameas-000954-applause_for_ruling_universe/applause__LEGACY.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundBadImpersonatorGunFire,
		Path: "jd-pinsound/sfx/000045-sniper/sniper_2___LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundBadImpersonatorThrow,
		Path: "jd-pinsound/sfx/000052-bad_inpersonator/bad_inpe__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundBallLock,
		Path: "jd-pinsound/sfx/000159-option_1/option_1__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundBallLost,
		Path: "jd-pinsound/sfx/000112-ball_lost/ball_los__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundBlackoutJackpot,
		Path: "jd-pinsound/jingle/000154-blackout_jackpot/000154-blackout_jackpot_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundBonus,
		Path: "jd-pinsound/sfx/000063-piano_1/piano_1___LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundDeadworldExplosion,
		Path: "afm-pinsound/sfx/1313658856-big_nice_explosion/big_nice__LEGACY.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundDing,
		Path: "jd-pinsound/sfx/000101-dring/dring__L__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundDropTargetLitHit1,
		Path: "jd-pinsound/sfx/000096-laser_3/laser_3___LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundDropTargetLitHit2,
		Path: "jd-pinsound/sfx/000097-laser_4/laser_4___LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundDropTargetLitHit3,
		Path: "jd-pinsound/sfx/000098-laser_4/laser_4__2_LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundDropTargetLitHit4,
		Path: "jd-pinsound/sfx/000099-laser_5/laser_5___LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundDropTargetLitHit5,
		Path: "jd-pinsound/sfx/000100-flash_2/flash_2___LEGACY_AUD.wav",
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
		ID:   SoundLeftPopperLaser,
		Path: "jd-pinsound/sfx/000068-laser_2/laser_2___LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundLeftPost,
		Path: "jd-pinsound/sfx/000151-wrong_signal/wrong_si__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundLeftShooterLaneFire,
		Path: "jd-pinsound/sfx/000041-fire/fire__LE__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundManhuntAutoFire,
		Path: "jd-pinsound/sfx/000040-manhunt_mode/manhunt___LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundManhuntSingleFire,
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
		ID:   SoundMotorcycleRamp,
		Path: "jd-pinsound/sfx/000102-motorcycle_noise_2/motorcyc__LEGACY_AUD.wav",
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
		ID:   SoundMultiballJackpot,
		Path: "jd-pinsound/jingle/000155-special_option/special___LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundMystery,
		Path: "jd-pinsound/sfx/000124-criquet_sound/criquet___LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundPoint,
		Path: "jd-pinsound/sfx/000106-point/point__L__LEGACY_AUD.wav",
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
		ID:   SoundRightPost,
		Path: "jd-pinsound/sfx/000152-wrong_signal_1/wrong_si__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundRightRamp,
		Path: "jd-pinsound/sfx/000095-special_gun_fire/special___LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundShock,
		Path: "jd-pinsound/sfx/000049-bad_inpersonator/bad_inpe__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundSling,
		Path: "jd-pinsound/sfx/000107-check_point/check_po__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundSafecrackerExplosion,
		Path: "jd-pinsound/sfx/000093-big_explosion/big_expl__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundSafecrackerGunFire1,
		Path: "jd-pinsound/sfx/000040-manhunt_mode/manhunt___LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundSafecrackerGunFire2,
		Path: "jd-pinsound/sfx/000085-laser_fire/laser_fi__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundSafecrackerGunFire3,
		Path: "jd-pinsound/sfx/000086-little_laser_fire/little_l__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundSafecrackerLaserFire,
		Path: "jd-pinsound/sfx/000095-special_gun_fire/special___LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundSafecrackerTankFire,
		Path: "jd-pinsound/sfx/000041-fire/fire__LE__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundSniperSplat,
		Path: "jd-pinsound/sfx/000076-hurt_3/hurt_3____LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundSniperTower,
		Path: "jd-pinsound/sfx/000069-guitar_2/guitar_2__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundSnore,
		Path: "jd-pinsound/sfx/000203-criquet_1/criquet___LEGACY_AUD.wav",
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
	eng.Do(spin.RegisterSound{
		ID:   SoundTopLeftRamp,
		Path: "jd-pinsound/sfx/000083-extra_sound/extra_so__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSound{
		ID:   SoundWalking,
		Path: "jd-pinsound/sfx/000146-walking_sound/walking___LEGACY_AUD.wav",
	})
}

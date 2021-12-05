package jdx

import "github.com/drop-target-pinball/spin"

const (
	SpeechAaaaah                                    = "jdx.SpeechAaaaah"
	SpeechICanSeeMyHouseFromHere                    = "jdx.SpeechICanSeeMyHouseFromHere"
	SpeechItsALongWayDown                           = "jdx.SpeechItsALongWayDown"
	SpeechLawMasterComputerOnlineWelcomeAboard      = "jdx.SpeechLawMasterComputerOnlineWelcomeAboard"
	SpeechPlayer2                                   = "jdx.SpeechPlayer2"
	SpeechPlayer3                                   = "jdx.SpeechPlayer3"
	SpeechPlayer4                                   = "jdx.SpeechPlayer4"
	SpeechShootSniperTower                          = "jdx.SpeechShootSniperTower"
	SpeechSniperEliminated                          = "jdx.SpeechSniperEliminated"
	SpeechSniperIsShootingIntoCrowdFromJohnsonTower = "jdx.SpeechSniperIsShootingIntoCrowdFromJohnsonTower"
	SpeechUseFireButtonToLaunchBall                 = "jdx.SpeechUseFireButtonToLaunchBall"
)

func RegisterSpeech(eng *spin.Engine) {
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechAaaaah,
		Path: "jd/pinsound/voice/000282-aaaaaaaaaaaaaaaaaaaah/aaaaaaaa__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechICanSeeMyHouseFromHere,
		Path: "jd/pinsound/voice/000284-i_can_see_my_house_from_here/i_can_se__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechItsALongWayDown,
		Path: "jd/pinsound/voice/000283-it_s_a_long_way_down/it_s_a_l__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechLawMasterComputerOnlineWelcomeAboard,
		Path: "jd/pinsound/voice/000322-law_master_computer_online_welcome_aboard/law_mast__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechPlayer2,
		Path: "jd/pinsound/voice/000210-player_two/player_t__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechPlayer3,
		Path: "jd/pinsound/voice/000211-player_three/player_three__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechPlayer4,
		Path: "jd/pinsound/voice/000214-player_four/player_f__LEGACY.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechShootSniperTower,
		Path: "jd/pinsound/voice/000356-shoot_sniper_tower/shoot_sn__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechSniperEliminated,
		Path: "jd/pinsound/voice/000253-sniper_eliminated/sniper_e__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechSniperIsShootingIntoCrowdFromJohnsonTower,
		Path: "jd/pinsound/voice/000346-sniper_is_shooting_into_crowd_from_johnson_tower/sniper_i__LEGACY.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechUseFireButtonToLaunchBall,
		Path: "jd/pinsound/voice/000326-use_fire_button_to_launch_ball/use_fire__LEGACY_AUD.wav",
	})
}

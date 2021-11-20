package jdx

import "github.com/drop-target-pinball/spin"

const (
	SpeechAaaaah                                    = "jdx.SpeechAaaaah"
	SpeechICanSeeMyHouseFromHere                    = "jdx.SpeechICanSeeMyHouseFromHere"
	SpeechItsALongWayDown                           = "jdx.SpeechItsALongWayDown"
	SpeechShootSniperTower                          = "jdx.SpeechShootSniperTower"
	SpeechSniperEliminated                          = "jdx.SpeechSniperEliminated"
	SpeechSniperIsShootingIntoCrowdFromJohnsonTower = "jdx.SpeechSniperIsShootingIntoCrowdFromJohnsonTower"
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
}

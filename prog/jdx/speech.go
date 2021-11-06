package jdx

import "github.com/drop-target-pinball/spin"

const (
	Ahhhhh                                    = "Ahhhhh"
	ICanSeeMyHouseFromHere                    = "i-can-see-my-house-from-here"
	ItsALongWayDown                           = "its-a-long-way-down"
	SniperIsShootingIntoCrowdFromJohnsonTower = "sniper-is-shooting-into-crowd-from-johnson-tower"
	ShootSniperTower                          = "shoot-sniper-tower"
)

func RegisterSpeech(eng *spin.Engine) {
	eng.Do(spin.RegisterSpeech{
		ID:   Ahhhhh,
		Path: "jd/pinsound/voice/000282-aaaaaaaaaaaaaaaaaaaah/aaaaaaaa__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   ICanSeeMyHouseFromHere,
		Path: "jd/pinsound/voice/000284-i_can_see_my_house_from_here/i_can_se__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   ItsALongWayDown,
		Path: "jd/pinsound/voice/000283-it_s_a_long_way_down/it_s_a_l__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   ShootSniperTower,
		Path: "jd/pinsound/voice/000356-shoot_sniper_tower/shoot_sn__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SniperIsShootingIntoCrowdFromJohnsonTower,
		Path: "jd/pinsound/voice/000346-sniper_is_shooting_into_crowd_from_johnson_tower/sniper_i__LEGACY.wav",
	})
}

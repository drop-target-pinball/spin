package jdx

import "github.com/drop-target-pinball/spin"

const (
	SpeechAaaaah                                        = "jdx.SpeechAaaaah"
	SpeechAllReactorsApprochingCriticalMass             = "jdx.SpeechAllReactorsApprochingCriticalMass"
	SpeechAllReactorsStabilized                         = "jdx.SpeechAllReactorsStabilized"
	SpeechBattleTankDamageAt                            = "jdx.SpeechBattleTankDamageAt"
	SpeechBattleTankDestroyed                           = "jdx.SpeechBattleTankDestroyed"
	SpeechBattleTankSightedInSectorSix                  = "jdx.SpeechBattleTankSightedInSectorSix"
	SpeechBoo                                           = "jdx.SpeechBoo"
	SpeechCivilDisorderHasEruptedInHeitschMusicHall     = "jdx.SpeechCivilDisorderHasEruptedInHeitschMusicHall"
	SpeechContainmentFailureAtThreeMeterIsland          = "jdx.SpeechContainmentFailureAtThreeMetereIsland"
	SpeechControlToDredd                                = "jdx.SpeechControlToDredd"
	SpeechDinnerTime                                    = "jdx.SpeechDinnerTime"
	SpeechDreddToControl                                = "jdx.SpeechDreddToControl"
	SpeechFour                                          = "jdx.SpeechFour"
	SpeechFreeze                                        = "jdx.SpeechFreeze"
	SpeechGoHome                                        = "jdx.SpeechGoHome"
	SpeechICanSeeMyHouseFromHere                        = "jdx.SpeechICanSeeMyHouseFromHere"
	SpeechIllBeBack                                     = "jdx.SpeechIllBeBack"
	SpeechImInPursuitOfAStolenVehicle                   = "jdx.SpeechImInPursuitOfAStolenVehicle"
	SpeechImStakingOutACrackHouseInSectorTwentyThree    = "jdx.SpeechImStakingOutACrackHouseInSectorTwentyThree"
	SpeechInteresting                                   = "jdx.SpeechInteresting"
	SpeechItsALongWayDown                               = "jdx.SpeechItsALongWayDown"
	SpeechIWonderWhatsDownThere                         = "jdx.SpeechIWonderWhatsDownThere"
	SpeechIWonderWhatsOverThere                         = "jdx.SpeechIWonderWhatsOverThere"
	SpeechLawMasterComputerOnlineWelcomeAboard          = "jdx.SpeechLawMasterComputerOnlineWelcomeAboard"
	SpeechMegaCityOneIsBlackedOutBeOnTheAlertForLooters = "jdx.SpeechMegaCityOneIsBlackedOutBeOnTheAlertForLooters"
	SpeechMeltdownIsImminent                            = "jdx.SpeechMeltdownIsImminent"
	SpeechOne                                           = "jdx.SpeechOne"
	SpeechOpenThatSafe                                  = "jdx.SpeechOpenThatSafe"
	SpeechOrIWillShoot                                  = "jdx.SpeechOrIWillShoot"
	SpeechPlayer2                                       = "jdx.SpeechPlayer2"
	SpeechPlayer3                                       = "jdx.SpeechPlayer3"
	SpeechPlayer4                                       = "jdx.SpeechPlayer4"
	SpeechReactorOneStabilized                          = "jdx.SpeechReactorOneStabilized"
	SpeechReactorTwoStabilized                          = "jdx.SpeechReactorTwoStabilized"
	SpeechSendBackupUnits                               = "jdx.SpeechSendBackupUnits"
	SpeechShootLeftRamp                                 = "jdx.SpeechShootLeftRamp"
	SpeechShootRightRamp                                = "jdx.SpeechShootRightRamp"
	SpeechShootSniperTower                              = "jdx.SpeechShootSniperTower"
	SpeechSixtyPercent                                  = "jdx.SpeechSixtyPercent"
	SpeechSniperEliminated                              = "jdx.SpeechSniperEliminated"
	SpeechSniperIsShootingIntoCrowdFromJohnsonTower     = "jdx.SpeechSniperIsShootingIntoCrowdFromJohnsonTower"
	SpeechStop                                          = "jdx.SpeechStop"
	SpeechSuspectGotAway                                = "jdx.SpeechSuspectGotAway"
	SpeechSuspiciousCharacterReportedInEugeneBlock      = "jdx.SpeechSuspiciousCharacterReportedInEugeneBlock"
	SpeechThree                                         = "jdx.SpeechThree"
	SpeechThreeMeterIslandIsSecured                     = "jdx.SpeechThreeMeterIslandIsSecured"
	SpeechTwentyFivePercent                             = "jdx.SpeechTwentyFivePercent"
	SpeechTwo                                           = "jdx.SpeechTwo"
	SpeechUseFireButtonToLaunchBall                     = "jdx.SpeechUseFireButtonToLaunchBall"
	SpeechWakeUpYouGeezer                               = "jdx.SpeechWakUpYouGeezer"
	SpeechYourDrivingDaysAreOverPunk                    = "jdx.SpeechYourDrivingDaysAreOverPunk"
	SpeechYouSuck                                       = "jdx.SpeechYouSuck"
)

func RegisterSpeech(eng *spin.Engine) {
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechAaaaah,
		Path: "jd-pinsound/voice/000282-aaaaaaaaaaaaaaaaaaaah/aaaaaaaa__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechAllReactorsApprochingCriticalMass,
		Path: "jd-pinsound/voice/000319-all_reactors_approaching_critical_mass/all_reac__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechAllReactorsStabilized,
		Path: "jd-pinsound/voice/000317-all_reactors_stabilized/all_reac__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechBattleTankDamageAt,
		Path: "jd-pinsound/voice/000323-battle_tank_damage_at/battle_t__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechBattleTankDestroyed,
		Path: "jd-pinsound/voice/000257-battle_tank_destroyed/battle_t__LEGACY.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechBattleTankSightedInSectorSix,
		Path: "jd-pinsound/voice/000345-battle_tank_sighted_in_sector_six/battle_t__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechBoo,
		Path: "jd-pinsound/sfx/000357-booooo_from_the_crowd/booooo_f__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechCivilDisorderHasEruptedInHeitschMusicHall,
		Path: "jd-pinsound/voice/000347-several_disorder_has_urepted_in_hight_school_musical/several___LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechContainmentFailureAtThreeMeterIsland,
		Path: "jd-pinsound/voice/000344-containment_failure_at_three_meter_island/containm__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechControlToDredd,
		Path: "jd-pinsound/voice/000342-control_to_dredd/control___LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechDinnerTime,
		Path: "jd-pinsound/voice/000208-aah_diner_time/aah_dine__LEGACY.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechDreddToControl,
		Path: "jd-pinsound/voice/000248-dredd_to_control/dredd_to__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechFour,
		Path: "jd-pinsound/voice/000310-4/4__LEGAC__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechFreeze,
		Path: "jd-pinsound/voice/000216-freeze/freeze____LEGACY.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechGoHome,
		Path: "jd-pinsound/voice/000359-go_home/go_home___LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechICanSeeMyHouseFromHere,
		Path: "jd-pinsound/voice/000284-i_can_see_my_house_from_here/i_can_se__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechIllBeBack,
		Path: "jd-pinsound/voice/000202-ill_be_back/000202-ill_be_back_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechImInPursuitOfAStolenVehicle,
		Path: "jd-pinsound/voice/000249-i_m_in_pursuit_of_a_stolen_vehicle/i_m_in_p__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechImStakingOutACrackHouseInSectorTwentyThree,
		Path: "jd-pinsound/voice/000259-i_m_staking_out_a_crack_house_in_sector_23/i_m_stak__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechInteresting,
		Path: "jd-pinsound/voice/000222-interesting/interest__LEGACY.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechItsALongWayDown,
		Path: "jd-pinsound/voice/000283-it_s_a_long_way_down/it_s_a_l__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechIWonderWhatsDownThere,
		Path: "jd-pinsound/voice/000221-i_wonder_what_s_down_there/i_wonder__LEGACY.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechIWonderWhatsOverThere,
		Path: "jd-pinsound/voice/000220-i_wonder_whats_over_there/000220-I_wonder_whats_over_there_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechLawMasterComputerOnlineWelcomeAboard,
		Path: "jd-pinsound/voice/000322-law_master_computer_online_welcome_aboard/law_mast__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechMegaCityOneIsBlackedOutBeOnTheAlertForLooters,
		Path: "jd-pinsound/voice/3609427385-megocity_one_is_blacked_out_be_on_the_alert_for_looters/megocity__LEGACY.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechMeltdownIsImminent,
		Path: "jd-pinsound/voice/000318-meltdown_is_imminent/meltdown__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechOne,
		Path: "jd-pinsound/voice/000313-1/1__LEGAC__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechOpenThatSafe,
		Path: "jd-pinsound/voice/000201-open_that_safe/000201-open_that_safe_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechOrIWillShoot,
		Path: "jd-pinsound/voice/000213-or_i_will_shoot/or_i_wil__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechPlayer2,
		Path: "jd-pinsound/voice/000210-player_two/player_t__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechPlayer3,
		Path: "jd-pinsound/voice/000211-player_three/player_three__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechPlayer4,
		Path: "jd-pinsound/voice/000214-player_four/player_f__LEGACY.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechReactorOneStabilized,
		Path: "jd-pinsound/voice/000315-reactor_one_stabilized/reactor___LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechReactorTwoStabilized,
		Path: "jd-pinsound/voice/000316-reactor_two_stabilized/reactor___LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechSendBackupUnits,
		Path: "jd-pinsound/voice/000260-send_backup_units/send_bac__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechShootLeftRamp,
		Path: "jd-pinsound/voice/000353-shoot_left_ramp/shoot_le__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechShootRightRamp,
		Path: "jd-pinsound/voice/000352-shoot_right_ramp/shoot_ri__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechShootSniperTower,
		Path: "jd-pinsound/voice/000356-shoot_sniper_tower/shoot_sn__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechSixtyPercent,
		Path: "jd-pinsound/voice/000324-60_percent/60_perce__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechSniperEliminated,
		Path: "jd-pinsound/voice/000253-sniper_eliminated/sniper_e__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechSniperIsShootingIntoCrowdFromJohnsonTower,
		Path: "jd-pinsound/voice/000346-sniper_is_shooting_into_crowd_from_johnson_tower/sniper_i__LEGACY.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechStop,
		Path: "jd-pinsound/voice/000212-stop/stop__LE__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechSuspectGotAway,
		Path: "jd-pinsound/voice/000252-suspect_got_away/000252-suspect_got_away_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechSuspiciousCharacterReportedInEugeneBlock,
		Path: "jd-pinsound/voice/000348-the_special_caracter_reported_are_using_block/the_spec__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechThree,
		Path: "jd-pinsound/voice/000311-3/3__LEGAC__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechThreeMeterIslandIsSecured,
		Path: "jd-pinsound/voice/000255-peview_around_is_secured/peview_a__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechTwentyFivePercent,
		Path: "jd-pinsound/voice/000306-twenty_five_percent/twenty_f__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechTwo,
		Path: "jd-pinsound/voice/000312-2/2__LEGAC__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechUseFireButtonToLaunchBall,
		Path: "jd-pinsound/voice/000326-use_fire_button_to_launch_ball/use_fire__LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechWakeUpYouGeezer,
		Path: "jd-pinsound/voice/000200-wake_up_you_geezer/wake_up___LEGACY_AUD.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechYourDrivingDaysAreOverPunk,
		Path: "jd-pinsound/voice/000209-your_driving_days_are_over_punk/your_dri__LEGACY.wav",
	})
	eng.Do(spin.RegisterSpeech{
		ID:   SpeechYouSuck,
		Path: "jd-pinsound/voice/000358-you_suck/you_suck__LEGACY_AUD.wav",
	})
}

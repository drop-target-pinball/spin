package jd

import "github.com/drop-target-pinball/spin"

var Config = spin.Config{
	CoilTrough: CoilTrough,

	SwitchEnterServiceButton:    SwitchEnterServiceButton,
	SwitchExitServiceButton:     SwitchExitServiceButton,
	SwitchLeftFlipperButton:     SwitchLeftFlipperButton,
	SwitchNextServiceButton:     SwitchNextServiceButton,
	SwitchPreviousServiceButton: SwitchPreviousServiceButton,
	SwitchRightFlipperButton:    SwitchRightFlipperButton,
	SwitchShooterLane:           SwitchRightShooterLane,
	SwitchStartButton:           SwitchStartButton,
	SwitchTroughJam:             SwitchTroughJam,
	SwitchDrain:                 SwitchTrough1,
	SwitchWillDrain:             []string{SwitchLeftOutlane, SwitchRightOutlane},

	LampStartButton: LampStartButton,

	GI: []string{
		GI1,
		GI2,
		GI3,
		GI4,
		GI5,
	},
	NumBalls: 6,
}

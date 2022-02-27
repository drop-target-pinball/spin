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
	PlayfieldSwitches:           PlayfieldSwitchEvents,
	LampStartButton:             LampStartButton,

	SwitchTrough: []string{
		SwitchTrough1,
		SwitchTrough2,
		SwitchTrough3,
		SwitchTrough4,
		SwitchTrough5,
		SwitchTrough6,
	},

	GI: []string{
		GI1,
		GI2,
		GI3,
		GI4,
		GI5,
	},
	NumBalls: 6,
}

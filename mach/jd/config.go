package jd

import "github.com/drop-target-pinball/spin"

var Config = spin.Config{
	SwitchEnterServiceButton:    SwitchEnterServiceButton,
	SwitchExitServiceButton:     SwitchExitServiceButton,
	SwitchLeftFlipperButton:     SwitchLeftFlipperButton,
	SwitchNextServiceButton:     SwitchNextServiceButton,
	SwitchPreviousServiceButton: SwitchPreviousServiceButton,
	SwitchRightFlipperButton:    SwitchRightFlipperButton,
	SwitchStartButton:           SwitchStartButton,
	SwitchDrain:                 SwitchTrough6,
	SwitchWillDrain:             []string{SwitchLeftOutlane, SwitchRightOutlane},
	LampStartButton:             LampStartButton,

	GI: []string{
		GI1,
		GI2,
		GI3,
		GI4,
		GI5,
	},
}

package jd

import (
	"github.com/drop-target-pinball/coroutine"
	"github.com/drop-target-pinball/spin"
)

const (
	DropTargetJ = 1 << iota
	DropTargetU
	DropTargetD
	DropTargetG
	DropTargetE

	MinDropTarget = 0
	MaxDropTarget = 4
)

var (
	DropTargetIndexes = map[string]int{
		SwitchDropTargetJ: 0,
		SwitchDropTargetU: 1,
		SwitchDropTargetD: 2,
		SwitchDropTargetG: 3,
		SwitchDropTargetE: 4,
		LampDropTargetJ:   0,
		LampDropTargetU:   1,
		LampDropTargetD:   2,
		LampDropTargetG:   3,
		LampDropTargetE:   4,
	}
)

const (
	CrimeSceneLeftLoop = 1 << iota
	CrimeSceneInnerLoop
	CrimeSceneRightLoop
	CrimeSceneRightPopper
	CrimeSceneRightRamp
)

var CrimeScenes = []int{
	CrimeSceneLeftLoop,
	CrimeSceneInnerLoop,
	CrimeSceneRightLoop,
	CrimeSceneRightPopper,
	CrimeSceneRightRamp,
}

var CrimeSceneSwitchEvents = []coroutine.Event{
	spin.SwitchEvent{ID: SwitchOuterLoopLeft},
	spin.SwitchEvent{ID: SwitchInnerLoop},
	spin.SwitchEvent{ID: SwitchOuterLoopRight},
	spin.SwitchEvent{ID: SwitchRightPopper},
	spin.SwitchEvent{ID: SwitchRightRampExit},
}

var CrimeSceneSwitches = map[string]int{
	SwitchOuterLoopLeft:  CrimeSceneLeftLoop,
	SwitchInnerLoop:      CrimeSceneInnerLoop,
	SwitchOuterLoopRight: CrimeSceneRightLoop,
	SwitchRightPopper:    CrimeSceneRightPopper,
	SwitchRightRampExit:  CrimeSceneRightRamp,
}

const (
	CrimeLevelNone = iota
	CrimeLevelWarning
	CrimeLevelMisdemeanor
	CrimeLevelFelony
	CrimeLevelClassXFelony
)

var CrimeSceneLamps = map[int]map[int]string{
	CrimeSceneLeftLoop: {
		CrimeLevelWarning:      LampLeftLoopCrimeSceneGreen,
		CrimeLevelMisdemeanor:  LampLeftLoopCrimeSceneYellow,
		CrimeLevelFelony:       LampLeftLoopCrimeSceneRed,
		CrimeLevelClassXFelony: LampLeftLoopCrimeSceneWhite,
	},
	CrimeSceneInnerLoop: {
		CrimeLevelWarning:      LampInnerLoopCrimeSceneGreen,
		CrimeLevelMisdemeanor:  LampInnerLoopCrimeSceneYellow,
		CrimeLevelFelony:       LampInnerLoopCrimeSceneRed,
		CrimeLevelClassXFelony: LampInnerLoopCrimeSceneWhite,
	},
	CrimeSceneRightLoop: {
		CrimeLevelWarning:      LampRightLoopCrimeSceneGreen,
		CrimeLevelMisdemeanor:  LampRightLoopCrimeSceneYellow,
		CrimeLevelFelony:       LampRightLoopCrimeSceneRed,
		CrimeLevelClassXFelony: LampRightLoopCrimeSceneWhite,
	},
	CrimeSceneRightPopper: {
		CrimeLevelWarning:      LampRightPopperCrimeSceneGreen,
		CrimeLevelMisdemeanor:  LampRightPopperCrimeSceneYellow,
		CrimeLevelFelony:       LampRightPopperCrimeSceneRed,
		CrimeLevelClassXFelony: LampRightPopperCrimeSceneWhite,
	},
	CrimeSceneRightRamp: {
		CrimeLevelWarning:      LampRightRampCrimeSceneGreen,
		CrimeLevelMisdemeanor:  LampRightRampCrimeSceneYellow,
		CrimeLevelFelony:       LampRightRampCrimeSceneRed,
		CrimeLevelClassXFelony: LampRightRampCrimeSceneWhite,
	},
}

var CrimeLevelLamps = []string{
	"",
	LampWarning,
	LampMisdemeanor,
	LampFelony,
	LampClassXFelony,
}

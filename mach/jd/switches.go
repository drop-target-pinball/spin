package jd

import (
	"github.com/drop-target-pinball/coroutine"
	"github.com/drop-target-pinball/go-pinproc/wpc"
	"github.com/drop-target-pinball/spin"
)

const (
	SwitchArmFarRight             = "jd.SwitchArmFarRight"
	SwitchBankTargets             = "jd.SwitchBankTargets"
	SwitchBuyInButton             = "jd.SwitchBuyInButton"
	SwitchCaptiveBall1            = "jd.SwitchCaptiveBall1"
	SwitchCaptiveBall2            = "jd.SwitchCaptiveBall2"
	SwitchCaptiveBall3            = "jd.SwitchCaptiveBall3"
	SwitchCenterCoinSlot          = "jd.SwitchCenterCoinSlot"
	SwitchCoinDoorOpen            = "jd.SwitchCoinDoorOpen"
	SwitchDropTargetJ             = "jd.SwitchDropTargetJ"
	SwitchDropTargetU             = "jd.SwitchDropTargetU"
	SwitchDropTargetD             = "jd.SwitchDropTargetD"
	SwitchDropTargetG             = "jd.SwitchDropTargetG"
	SwitchDropTargetE             = "jd.SwitchDropTargetE"
	SwitchEnterServiceButton      = "jd.SwitchEnterServiceButton"
	SwitchExitServiceButton       = "jd.SwitchExitServiceButton"
	SwitchGlobeExit               = "jd.SwitchGlobeExit"
	SwitchGlobePosition1          = "jd.SwitchGlobePosition1"
	SwitchGlobePosition2          = "jd.SwitchGlobePosition2"
	SwitchInnerLoop               = "jd.SwitchInnerLoop"
	SwitchInnerRightReturnLane    = "jd.SwitchInnerRightReturnLane"
	SwitchLeftCoinSlot            = "jd.SwitchLeftCoinSlot"
	SwitchLeftFireButton          = "jd.SwitchLeftFireButton"
	SwitchLeftFlipperButton       = "jd.SwitchLeftFlipperButton"
	SwitchLeftFlipperEOS          = "jd.SwitchLeftFlipperEOS"
	SwitchLeftOutlane             = "jd.SwitchLeftOutlane"
	SwitchLeftPopper              = "jd.SwitchLeftPopper"
	SwitchLeftPost                = "jd.SwitchLeftPost"
	SwitchLeftRampEnter           = "jd.SwitchLeftRampEnter"
	SwitchLeftRampExit            = "jd.SwitchLeftRampExit"
	SwitchLeftRampToLock          = "jd.SwitchLeftRampToLock"
	SwitchLeftReturnLane          = "jd.SwitchLeftReturnLane"
	SwitchLeftShooterLane         = "jd.SwitchLeftShooterLane"
	SwitchLeftSling               = "jd.SwitchLeftSling"
	SwitchMysteryTarget           = "jd.SwitchMysteryTarget"
	SwitchNextServiceButton       = "jd.SwitchNextServiceButton"
	SwitchOuterLoopLeft           = "jd.SwitchOuterLoopLeft"
	SwitchOuterLoopRight          = "jd.SwitchOuterLoopRight"
	SwitchOuterRightReturnLane    = "jd.SwitchOuterRightReturnLane"
	SwitchPreviousServiceButton   = "jd.SwitchPreviousServiceButton"
	SwitchRightCoinSlot           = "jd.SwitchRightCoinSlot"
	SwitchRightFireButton         = "jd.SwitchRightFireButton"
	SwitchRightFlipperButton      = "jd.SwitchRightFlipperButton"
	SwitchRightFlipperEOS         = "jd.SwitchRightFlipperEOS"
	SwitchRightOutlane            = "jd.SwitchRightOutlane"
	SwitchRightPopper             = "jd.SwitchRightPopper"
	SwitchRightPost               = "jd.SwitchRightPost"
	SwitchRightRampExit           = "jd.SwitchRightRampExit"
	SwitchRightShooterLane        = "jd.SwitchRightShooterLane"
	SwitchRightSling              = "jd.SwitchRightSling"
	SwitchSlamTilt                = "jd.SwitchSlamTilt"
	SwitchStartButton             = "jd.SwitchStartButton"
	SwitchSubwayEnter1            = "jd.SwitchSubwayEnter1"
	SwitchSubwayEnter2            = "jd.SwitchSubwayEnter2"
	SwitchSuperGameButton         = "jd.SwitchSuperGameButton"
	SwitchTilt                    = "jd.SwitchTilt"
	SwitchTopLeftRampExit         = "jd.SwitchTopLeftRampExit"
	SwitchTopRightRampExit        = "jd.SwitchTopRightRampExit"
	SwitchTrough1                 = "jd.SwitchTrough1"
	SwitchTrough2                 = "jd.SwitchTrough2"
	SwitchTrough3                 = "jd.SwitchTrough3"
	SwitchTrough4                 = "jd.SwitchTrough4"
	SwitchTrough5                 = "jd.SwitchTrough5"
	SwitchTrough6                 = "jd.SwitchTrough6"
	SwitchTroughJam               = "jd.SwitchTroughJam"
	SwitchUpperLeftFlipperButton  = "jd.SwitchUpperLeftFlipperButton"
	SwitchUpperLeftFlipperEOS     = "jd.SwitchUpperLeftFlipperEOS"
	SwitchUpperRightFlipperButton = "jd.SwitchUpperLeftFlipperButton"
	SwitchUpperRightFlipperEOS    = "jd.SwitchUpperRightFlipperEOS"
)

var DropTargetSwitches = []string{
	SwitchDropTargetJ,
	SwitchDropTargetU,
	SwitchDropTargetD,
	SwitchDropTargetG,
	SwitchDropTargetE,
}

var SwitchAnyDropTarget = []coroutine.Event{
	spin.SwitchEvent{ID: SwitchDropTargetJ},
	spin.SwitchEvent{ID: SwitchDropTargetU},
	spin.SwitchEvent{ID: SwitchDropTargetD},
	spin.SwitchEvent{ID: SwitchDropTargetG},
	spin.SwitchEvent{ID: SwitchDropTargetE},
}

func RegisterSwitches(eng *spin.Engine) {
	eng.Do(spin.RegisterSwitch{Addr: wpc.S71, ID: SwitchArmFarRight})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S18, ID: SwitchBankTargets})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S31, ID: SwitchBuyInButton})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S26, ID: SwitchCaptiveBall1})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S53, ID: SwitchCaptiveBall2})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S68, ID: SwitchCaptiveBall3})
	eng.Do(spin.RegisterSwitch{Addr: wpc.SD2, ID: SwitchCenterCoinSlot})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S22, ID: SwitchCoinDoorOpen, NC: true})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S54, ID: SwitchDropTargetJ, NC: true})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S55, ID: SwitchDropTargetU, NC: true})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S56, ID: SwitchDropTargetD, NC: true})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S57, ID: SwitchDropTargetG, NC: true})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S58, ID: SwitchDropTargetE, NC: true})
	eng.Do(spin.RegisterSwitch{Addr: wpc.SD8, ID: SwitchEnterServiceButton})
	eng.Do(spin.RegisterSwitch{Addr: wpc.SD5, ID: SwitchExitServiceButton})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S62, ID: SwitchGlobeExit, NC: true})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S61, ID: SwitchGlobePosition1})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S77, ID: SwitchGlobePosition2})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S35, ID: SwitchInnerLoop})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S34, ID: SwitchInnerRightReturnLane})
	eng.Do(spin.RegisterSwitch{Addr: wpc.SD1, ID: SwitchLeftCoinSlot})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S11, ID: SwitchLeftFireButton})
	eng.Do(spin.RegisterSwitch{Addr: wpc.SF4, ID: SwitchLeftFlipperButton})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S16, ID: SwitchLeftOutlane})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S73, ID: SwitchLeftPopper, NC: true})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S36, ID: SwitchLeftPost})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S67, ID: SwitchLeftRampEnter, NC: true})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S64, ID: SwitchLeftRampExit, NC: true})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S63, ID: SwitchLeftRampToLock, NC: true})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S17, ID: SwitchLeftReturnLane})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S15, ID: SwitchLeftShooterLane})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S51, ID: SwitchLeftSling})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S27, ID: SwitchMysteryTarget})
	eng.Do(spin.RegisterSwitch{Addr: wpc.SD7, ID: SwitchNextServiceButton})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S33, ID: SwitchOuterLoopLeft})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S72, ID: SwitchOuterLoopRight, NC: true})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S43, ID: SwitchOuterRightReturnLane})
	eng.Do(spin.RegisterSwitch{Addr: wpc.SD6, ID: SwitchPreviousServiceButton})
	eng.Do(spin.RegisterSwitch{Addr: wpc.SD3, ID: SwitchRightCoinSlot})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S12, ID: SwitchRightFireButton})
	eng.Do(spin.RegisterSwitch{Addr: wpc.SF2, ID: SwitchRightFlipperButton})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S42, ID: SwitchRightOutlane})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S74, ID: SwitchRightPopper, NC: true})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S25, ID: SwitchRightPost})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S76, ID: SwitchRightRampExit, NC: true})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S41, ID: SwitchRightShooterLane})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S52, ID: SwitchRightSling})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S21, ID: SwitchSlamTilt})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S13, ID: SwitchStartButton})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S37, ID: SwitchSubwayEnter1})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S38, ID: SwitchSubwayEnter2})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S44, ID: SwitchSuperGameButton})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S14, ID: SwitchTilt})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S66, ID: SwitchTopLeftRampExit, NC: true})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S75, ID: SwitchTopRightRampExit, NC: true})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S81, ID: SwitchTrough1, NC: true})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S82, ID: SwitchTrough2, NC: true})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S83, ID: SwitchTrough3, NC: true})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S84, ID: SwitchTrough4, NC: true})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S85, ID: SwitchTrough5, NC: true})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S86, ID: SwitchTrough6, NC: true})
	eng.Do(spin.RegisterSwitch{Addr: wpc.S87, ID: SwitchTroughJam, NC: true})
	eng.Do(spin.RegisterSwitch{Addr: wpc.SF8, ID: SwitchUpperLeftFlipperButton})

	if eng.Options.RegisterEOS {
		eng.Do(spin.RegisterSwitch{Addr: wpc.SF3, ID: SwitchLeftFlipperEOS})
		eng.Do(spin.RegisterSwitch{Addr: wpc.SF1, ID: SwitchRightFlipperEOS})
		eng.Do(spin.RegisterSwitch{Addr: wpc.SF6, ID: SwitchUpperRightFlipperButton})
		eng.Do(spin.RegisterSwitch{Addr: wpc.SF5, ID: SwitchUpperRightFlipperEOS})
	}
}

package jd

import (
	"github.com/drop-target-pinball/go-pinproc/wpc"
	"github.com/drop-target-pinball/spin"
)

const (
	LampAdvanceCrimeLevel           = "jd.LampAdvanceCrimeLevel"
	LampAirRaid                     = "jd.LampAirRaid"
	LampAwardBadImpersonator        = "jd.LampAwardBadImpersonator"
	LampAwardMeltdown               = "jd.LampAwardMeltdown"
	LampAwardSafeCracker            = "jd.LampAwardSafeCracker"
	LampAwardSniper                 = "jd.LampAwardSniper"
	LampAwardStakeout               = "jd.LampAwardStakeout"
	LampBadImpersonator             = "jd.LampBadImpersonator"
	LampBattleTank                  = "jd.LampBattleTank"
	LampBlackout                    = "jd.LampBlackout"
	LampBlackoutJackpot             = "jd.LampBlackoutJackpot"
	LampBuyInButton                 = "jd.LampBuyInButton"
	LampCenterTank                  = "jd.LampCenterTank"
	LampClassXFelony                = "jd.LampClassXFelony"
	LampDrainShield                 = "jd.LampDrainShield"
	LampDropTargetJ                 = "jd.LampDropTargetJ"
	LampDropTargetU                 = "jd.LampDropTargetU"
	LampDropTargetD                 = "jd.LampDropTargetD"
	LampDropTargetG                 = "jd.LampDropTargetG"
	LampDropTargetE                 = "jd.LampDropTargetE"
	LampExtraBall                   = "jd.LampExtraBall"
	LampFelony                      = "jd.LampFelony"
	LampInnerLoopCrimeSceneGreen    = "jd.LampInnerLoopCrimeSceneGreen"
	LampInnerLoopCrimeSceneRed      = "jd.LampInnerLoopCrimeSceneRed"
	LampInnerLoopCrimeSceneWhite    = "jd.LampInnerLoopCrimeSceneWhite"
	LampInnerLoopCrimeSceneYellow   = "jd.LampInnerLoopCrimeSceneYellow"
	LampJudgeAgain                  = "jd.LampJudgeAgain"
	LampLeftLoopCrimeSceneGreen     = "jd.LampLeftLoopCrimeSceneGreen"
	LampLeftLoopCrimeSceneRed       = "jd.LampLeftLoopCrimeSceneRed"
	LampLeftLoopCrimeSceneWhite     = "jd.LampLeftLoopCrimeSceneWhite"
	LampLeftLoopCrimeSceneYellow    = "jd.LampLeftLoopCrimeSceneYellow"
	LampLeftModeStart               = "jd.LampLeftModeStart"
	LampLeftTank                    = "jd.LampLeftTank"
	LampLock1                       = "jd.LampLock1"
	LampLock2                       = "jd.LampLock2"
	LampLock3                       = "jd.LampLock3"
	LampManhunt                     = "jd.LampManhunt"
	LampMeltdown                    = "jd.LampMeltdown"
	LampMisdemeanor                 = "jd.LampMisdemeanor"
	LampMultiballJackpot            = "jd.LampMultiballJackpot"
	LampMystery                     = "jd.LampMystery"
	LampPursuit                     = "jd.LampPursuit"
	LampRightLoopCrimeSceneGreen    = "jd.LampRightLoopCrimeSceneGreen"
	LampRightLoopCrimeSceneRed      = "jd.LampRightLoopCrimeSceneRed"
	LampRightLoopCrimeSceneWhite    = "jd.LampRightLoopCrimeSceneWhite"
	LampRightLoopCrimeSceneYellow   = "jd.LampRightLoopCrimeSceneYellow"
	LampRightModeStart              = "jd.LampRightModeStart"
	LampRightPopperCrimeSceneGreen  = "jd.LampRightPopperCrimeSceneGreen"
	LampRightPopperCrimeSceneRed    = "jd.LampRightPopperCrimeSceneRed"
	LampRightPopperCrimeSceneWhite  = "jd.LampRightPopperCrimeSceneWhite"
	LampRightPopperCrimeSceneYellow = "jd.LampRightPopperCrimeSceneYellow"
	LampRightRampCrimeSceneGreen    = "jd.LampRightRampCrimeSceneGreen"
	LampRightRampCrimeSceneRed      = "jd.LampRightRampCrimeSceneRed"
	LampRightRampCrimeSceneWhite    = "jd.LampRightRampCrimeSceneWhite"
	LampRightRampCrimeSceneYellow   = "jd.LampRightRampCrimeSceneYellow"
	LampRightTank                   = "jd.LampRightTank"
	LampSafeCracker                 = "jd.LampSafeCracker"
	LampSniper                      = "jd.LampSniper"
	LampStakeout                    = "jd.LampStakeout"
	LampStartButton                 = "jd.LampStartButton"
	LampSubwayCombo                 = "jd.LampSubwayCombo"
	LampSuperGameButton             = "jd.LampSuperGameButton"
	LampUltimateChallenge           = "jd.LampUltimateChallenge"
	LampWarning                     = "jd.LampWarning"
)

const (
	FlasherBlackout     = "jd.FlasherBlackout"
	FlasherCursedEarth  = "jd.FlasherCursedEarth"
	FlasherGlobe        = "jd.FlasherGlobe"
	FlasherInsert       = "jd.FlasherInsert"
	FlasherJudgeDeath   = "jd.FlasherJudgeDeath"
	FlasherJudgeFire    = "jd.FlasherJudgeFire"
	FlasherJudgeFear    = "jd.FlasherJudgeFear"
	FlasherJudgeMortis  = "jd.FlasherJudgeMortis"
	FlasherLeftPursuit  = "jd.FlasherLeftPursuit"
	FlasherRightPursuit = "jd.FlasherRightPursuit"
	FlasherRightRamp    = "jd.FlasherRightRamp"
	FlasherSubwayExit   = "jd.FlasherSubwayExit"
)

const (
	GI1 = "jd.GI1"
	GI2 = "jd.GI2"
	GI3 = "jd.GI3"
	GI4 = "jd.GI4"
	GI5 = "jd.GI5"
)

var DropTargetLamps = []string{
	LampDropTargetJ,
	LampDropTargetU,
	LampDropTargetD,
	LampDropTargetG,
	LampDropTargetE,
}

var LockLamps = []string{
	"",
	LampLock1,
	LampLock2,
	LampLock3,
}

func RegisterLamps(eng *spin.Engine) {
	eng.Do(spin.RegisterLamp{Addr: wpc.L85, ID: LampAdvanceCrimeLevel})
	eng.Do(spin.RegisterLamp{Addr: wpc.L65, ID: LampAirRaid})
	eng.Do(spin.RegisterLamp{Addr: wpc.L78, ID: LampAwardBadImpersonator})
	eng.Do(spin.RegisterLamp{Addr: wpc.L48, ID: LampAwardMeltdown})
	eng.Do(spin.RegisterLamp{Addr: wpc.L76, ID: LampAwardSafeCracker})
	eng.Do(spin.RegisterLamp{Addr: wpc.L64, ID: LampAwardSniper})
	eng.Do(spin.RegisterLamp{Addr: wpc.L81, ID: LampAwardStakeout})
	eng.Do(spin.RegisterLamp{Addr: wpc.L46, ID: LampBadImpersonator})
	eng.Do(spin.RegisterLamp{Addr: wpc.L47, ID: LampBattleTank})
	eng.Do(spin.RegisterLamp{Addr: wpc.L56, ID: LampBlackout})
	eng.Do(spin.RegisterLamp{Addr: wpc.L82, ID: LampBlackoutJackpot})
	eng.Do(spin.RegisterLamp{Addr: wpc.L87, ID: LampBuyInButton})
	eng.Do(spin.RegisterLamp{Addr: wpc.L63, ID: LampCenterTank})
	eng.Do(spin.RegisterLamp{Addr: wpc.L41, ID: LampClassXFelony})
	eng.Do(spin.RegisterLamp{Addr: wpc.L83, ID: LampDrainShield})
	eng.Do(spin.RegisterLamp{Addr: wpc.L71, ID: LampDropTargetJ})
	eng.Do(spin.RegisterLamp{Addr: wpc.L72, ID: LampDropTargetU})
	eng.Do(spin.RegisterLamp{Addr: wpc.L73, ID: LampDropTargetD})
	eng.Do(spin.RegisterLamp{Addr: wpc.L74, ID: LampDropTargetG})
	eng.Do(spin.RegisterLamp{Addr: wpc.L75, ID: LampDropTargetE})
	eng.Do(spin.RegisterLamp{Addr: wpc.L61, ID: LampExtraBall})
	eng.Do(spin.RegisterLamp{Addr: wpc.L42, ID: LampFelony})
	eng.Do(spin.RegisterLamp{Addr: wpc.L18, ID: LampInnerLoopCrimeSceneGreen})
	eng.Do(spin.RegisterLamp{Addr: wpc.L16, ID: LampInnerLoopCrimeSceneRed})
	eng.Do(spin.RegisterLamp{Addr: wpc.L15, ID: LampInnerLoopCrimeSceneWhite})
	eng.Do(spin.RegisterLamp{Addr: wpc.L17, ID: LampInnerLoopCrimeSceneYellow})
	eng.Do(spin.RegisterLamp{Addr: wpc.L84, ID: LampJudgeAgain})
	eng.Do(spin.RegisterLamp{Addr: wpc.L14, ID: LampLeftLoopCrimeSceneGreen})
	eng.Do(spin.RegisterLamp{Addr: wpc.L12, ID: LampLeftLoopCrimeSceneRed})
	eng.Do(spin.RegisterLamp{Addr: wpc.L11, ID: LampLeftLoopCrimeSceneWhite})
	eng.Do(spin.RegisterLamp{Addr: wpc.L13, ID: LampLeftLoopCrimeSceneYellow})
	eng.Do(spin.RegisterLamp{Addr: wpc.L66, ID: LampLeftModeStart})
	eng.Do(spin.RegisterLamp{Addr: wpc.L67, ID: LampLeftTank})
	eng.Do(spin.RegisterLamp{Addr: wpc.L35, ID: LampLock1})
	eng.Do(spin.RegisterLamp{Addr: wpc.L36, ID: LampLock2})
	eng.Do(spin.RegisterLamp{Addr: wpc.L37, ID: LampLock3})
	eng.Do(spin.RegisterLamp{Addr: wpc.L55, ID: LampManhunt})
	eng.Do(spin.RegisterLamp{Addr: wpc.L45, ID: LampMeltdown})
	eng.Do(spin.RegisterLamp{Addr: wpc.L43, ID: LampMisdemeanor})
	eng.Do(spin.RegisterLamp{Addr: wpc.L77, ID: LampMultiballJackpot})
	eng.Do(spin.RegisterLamp{Addr: wpc.L68, ID: LampMystery})
	eng.Do(spin.RegisterLamp{Addr: wpc.L53, ID: LampPursuit})
	eng.Do(spin.RegisterLamp{Addr: wpc.L24, ID: LampRightLoopCrimeSceneGreen})
	eng.Do(spin.RegisterLamp{Addr: wpc.L22, ID: LampRightLoopCrimeSceneRed})
	eng.Do(spin.RegisterLamp{Addr: wpc.L21, ID: LampRightLoopCrimeSceneWhite})
	eng.Do(spin.RegisterLamp{Addr: wpc.L23, ID: LampRightLoopCrimeSceneYellow})
	eng.Do(spin.RegisterLamp{Addr: wpc.L62, ID: LampRightModeStart})
	eng.Do(spin.RegisterLamp{Addr: wpc.L34, ID: LampRightPopperCrimeSceneGreen})
	eng.Do(spin.RegisterLamp{Addr: wpc.L32, ID: LampRightPopperCrimeSceneRed})
	eng.Do(spin.RegisterLamp{Addr: wpc.L31, ID: LampRightPopperCrimeSceneWhite})
	eng.Do(spin.RegisterLamp{Addr: wpc.L33, ID: LampRightPopperCrimeSceneYellow})
	eng.Do(spin.RegisterLamp{Addr: wpc.L28, ID: LampRightRampCrimeSceneGreen})
	eng.Do(spin.RegisterLamp{Addr: wpc.L26, ID: LampRightRampCrimeSceneRed})
	eng.Do(spin.RegisterLamp{Addr: wpc.L25, ID: LampRightRampCrimeSceneWhite})
	eng.Do(spin.RegisterLamp{Addr: wpc.L27, ID: LampRightRampCrimeSceneYellow})
	eng.Do(spin.RegisterLamp{Addr: wpc.L86, ID: LampRightTank})
	eng.Do(spin.RegisterLamp{Addr: wpc.L52, ID: LampSafeCracker})
	eng.Do(spin.RegisterLamp{Addr: wpc.L57, ID: LampSniper})
	eng.Do(spin.RegisterLamp{Addr: wpc.L51, ID: LampStakeout})
	eng.Do(spin.RegisterLamp{Addr: wpc.L88, ID: LampStartButton})
	eng.Do(spin.RegisterLamp{Addr: wpc.L58, ID: LampSubwayCombo})
	eng.Do(spin.RegisterLamp{Addr: wpc.L38, ID: LampSuperGameButton})
	eng.Do(spin.RegisterLamp{Addr: wpc.L54, ID: LampUltimateChallenge})
	eng.Do(spin.RegisterLamp{Addr: wpc.L44, ID: LampWarning})

	eng.Do(spin.RegisterLamp{Addr: wpc.G01, ID: GI1})
	eng.Do(spin.RegisterLamp{Addr: wpc.G02, ID: GI2})
	eng.Do(spin.RegisterLamp{Addr: wpc.G03, ID: GI3})
	eng.Do(spin.RegisterLamp{Addr: wpc.G04, ID: GI4})
	eng.Do(spin.RegisterLamp{Addr: wpc.G05, ID: GI5})
}

func RegisterFlashers(eng *spin.Engine) {
	eng.Do(spin.RegisterFlasher{Addr: wpc.C23, ID: FlasherBlackout})
	eng.Do(spin.RegisterFlasher{Addr: wpc.C24, ID: FlasherCursedEarth})
	eng.Do(spin.RegisterFlasher{Addr: wpc.C26, ID: FlasherGlobe})
	eng.Do(spin.RegisterFlasher{Addr: wpc.C28, ID: FlasherInsert})
	eng.Do(spin.RegisterFlasher{Addr: wpc.C19, ID: FlasherJudgeDeath})
	eng.Do(spin.RegisterFlasher{Addr: wpc.C17, ID: FlasherJudgeFire})
	eng.Do(spin.RegisterFlasher{Addr: wpc.C18, ID: FlasherJudgeFear})
	eng.Do(spin.RegisterFlasher{Addr: wpc.C20, ID: FlasherJudgeMortis})
	eng.Do(spin.RegisterFlasher{Addr: wpc.C21, ID: FlasherLeftPursuit})
	eng.Do(spin.RegisterFlasher{Addr: wpc.C22, ID: FlasherRightPursuit})
	eng.Do(spin.RegisterFlasher{Addr: wpc.C27, ID: FlasherRightRamp})
	eng.Do(spin.RegisterFlasher{Addr: wpc.C25, ID: FlasherSubwayExit})
}

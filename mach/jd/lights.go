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

const (
	DarkJudgeMortis = iota
	DarkJudgeFire
	DarkJudgeFear
	DarkJudgeDeath
)

var DarkJudgeFlashers = []string{
	FlasherJudgeMortis,
	FlasherJudgeFire,
	FlasherJudgeFear,
	FlasherJudgeDeath,
}

func RegisterLamps(eng *spin.Engine) {
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L85,
		ID:     LampAdvanceCrimeLevel,
		Layout: spin.NewLayoutRect(357, 354, 13, 28, spin.ColorLampWhite),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L65,
		ID:     LampAirRaid,
		Layout: spin.NewLayoutCircleFromRect(398, 220, 6, 7, spin.ColorLampRed),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L78,
		ID:     LampAwardBadImpersonator,
		Layout: spin.NewLayoutRect(289, 333, 13, 35, spin.ColorLampOrange),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L48,
		ID:     LampAwardMeltdown,
		Layout: spin.NewLayoutRect(426, 466, 15, 37, spin.ColorLampGreen),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L76,
		ID:     LampAwardSafeCracker,
		Layout: spin.NewLayoutRect(334, 312, 31, 13, spin.ColorLampOrange),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L64,
		ID:     LampAwardSniper,
		Layout: spin.NewLayoutCircleFromRect(374, 200, 12, 13, spin.ColorLampWhite),
	})
	eng.Do(spin.RegisterLamp{
		Addr: wpc.L81,
		ID:   LampAwardStakeout,
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L46,
		ID:     LampBadImpersonator,
		Layout: spin.NewLayoutRect(304, 575, 26, 12, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L47,
		ID:     LampBattleTank,
		Layout: spin.NewLayoutRect(265, 573, 26, 12, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L56,
		ID:     LampBlackout,
		Layout: spin.NewLayoutRect(226, 564, 26, 12, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L82,
		ID:     LampBlackoutJackpot,
		Layout: spin.NewLayoutRect(321, 134, 19, 22, spin.ColorLampWhite),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L87,
		ID:     LampBuyInButton,
		Layout: spin.NewLayoutCircleFromRect(597, 789, 12, 14, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L63,
		ID:     LampCenterTank,
		Layout: spin.NewLayoutRect(337, 209, 20, 29, spin.ColorLampGreen),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L41,
		ID:     LampClassXFelony,
		Layout: spin.NewLayoutRect(349, 599, 20, 42, spin.ColorLampWhite),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L83,
		ID:     LampDrainShield,
		Layout: spin.NewLayoutRect(284, 703, 63, 17, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L71,
		ID:     LampDropTargetJ,
		Layout: spin.NewLayoutRect(257, 320, 13, 13, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L72,
		ID:     LampDropTargetU,
		Layout: spin.NewLayoutRect(276, 310, 13, 13, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L73,
		ID:     LampDropTargetD,
		Layout: spin.NewLayoutRect(295, 301, 13, 13, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L74,
		ID:     LampDropTargetG,
		Layout: spin.NewLayoutRect(312, 289, 13, 13, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L75,
		ID:     LampDropTargetE,
		Layout: spin.NewLayoutRect(331, 278, 13, 13, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr: wpc.L61,
		ID:   LampExtraBall,
		Layout: spin.NewLayoutMulti(
			spin.NewLayoutCircleFromRect(180, 342, 6, 6, spin.ColorLampOrange),
			spin.NewLayoutCircleFromRect(408, 234, 6, 6, spin.ColorLampOrange),
		),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L42,
		ID:     LampFelony,
		Layout: spin.NewLayoutRect(320, 599, 20, 42, spin.ColorLampRed),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L18,
		ID:     LampInnerLoopCrimeSceneGreen,
		Layout: spin.NewLayoutRect(303, 192, 20, 8, spin.ColorLampGreen),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L16,
		ID:     LampInnerLoopCrimeSceneRed,
		Layout: spin.NewLayoutRect(303, 176, 20, 8, spin.ColorLampRed),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L15,
		ID:     LampInnerLoopCrimeSceneWhite,
		Layout: spin.NewLayoutRect(303, 168, 20, 8, spin.ColorLampWhite),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L17,
		ID:     LampInnerLoopCrimeSceneYellow,
		Layout: spin.NewLayoutRect(303, 184, 20, 8, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L84,
		ID:     LampJudgeAgain,
		Layout: spin.NewLayoutCircleFromRect(306, 745, 21, 19, spin.ColorLampOrange),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L14,
		ID:     LampLeftLoopCrimeSceneGreen,
		Layout: spin.NewLayoutRect(166, 384, 19, 8, spin.ColorLampGreen),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L12,
		ID:     LampLeftLoopCrimeSceneRed,
		Layout: spin.NewLayoutRect(166, 368, 19, 8, spin.ColorLampRed),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L11,
		ID:     LampLeftLoopCrimeSceneWhite,
		Layout: spin.NewLayoutRect(166, 360, 19, 8, spin.ColorLampWhite),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L13,
		ID:     LampLeftLoopCrimeSceneYellow,
		Layout: spin.NewLayoutRect(166, 376, 19, 8, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L66,
		ID:     LampLeftModeStart,
		Layout: spin.NewLayoutCircleFromRect(228, 407, 22, 17, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L67,
		ID:     LampLeftTank,
		Layout: spin.NewLayoutRect(183, 408, 20, 32, spin.ColorLampGreen),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L35,
		ID:     LampLock1,
		Layout: spin.NewLayoutCircleFromRect(223, 384, 8, 8, spin.ColorLampGreen),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L36,
		ID:     LampLock2,
		Layout: spin.NewLayoutCircleFromRect(217, 367, 8, 8, spin.ColorLampGreen),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L37,
		ID:     LampLock3,
		Layout: spin.NewLayoutCircleFromRect(212, 347, 8, 8, spin.ColorLampGreen),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L55,
		ID:     LampManhunt,
		Layout: spin.NewLayoutCircleFromRect(378, 560, 26, 12, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L45,
		ID:     LampMeltdown,
		Layout: spin.NewLayoutRect(342, 573, 25, 12, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L43,
		ID:     LampMisdemeanor,
		Layout: spin.NewLayoutRect(292, 599, 20, 42, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L77,
		ID:     LampMultiballJackpot,
		Layout: spin.NewLayoutRect(310, 323, 17, 28, spin.ColorLampRed),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L68,
		ID:     LampMystery,
		Layout: spin.NewLayoutRect(169, 471, 17, 19, spin.ColorLampBlue),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L53,
		ID:     LampPursuit,
		Layout: spin.NewLayoutRect(192, 546, 29, 12, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L24,
		ID:     LampRightLoopCrimeSceneGreen,
		Layout: spin.NewLayoutRect(394, 276, 34, 8, spin.ColorLampGreen),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L22,
		ID:     LampRightLoopCrimeSceneRed,
		Layout: spin.NewLayoutRect(394, 260, 34, 8, spin.ColorLampRed),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L21,
		ID:     LampRightLoopCrimeSceneWhite,
		Layout: spin.NewLayoutRect(394, 252, 34, 8, spin.ColorLampWhite),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L23,
		ID:     LampRightLoopCrimeSceneYellow,
		Layout: spin.NewLayoutRect(394, 268, 34, 8, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L62,
		ID:     LampRightModeStart,
		Layout: spin.NewLayoutCircleFromRect(369, 226, 17, 19, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L34,
		ID:     LampRightPopperCrimeSceneGreen,
		Layout: spin.NewLayoutRect(374, 183, 17, 8, spin.ColorLampGreen),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L32,
		ID:     LampRightPopperCrimeSceneRed,
		Layout: spin.NewLayoutRect(374, 167, 17, 8, spin.ColorLampRed),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L31,
		ID:     LampRightPopperCrimeSceneWhite,
		Layout: spin.NewLayoutRect(374, 159, 17, 8, spin.ColorLampWhite),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L33,
		ID:     LampRightPopperCrimeSceneYellow,
		Layout: spin.NewLayoutRect(374, 175, 17, 8, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L28,
		ID:     LampRightRampCrimeSceneGreen,
		Layout: spin.NewLayoutRect(388, 451, 18, 8, spin.ColorLampGreen),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L26,
		ID:     LampRightRampCrimeSceneRed,
		Layout: spin.NewLayoutRect(388, 435, 18, 8, spin.ColorLampRed),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L25,
		ID:     LampRightRampCrimeSceneWhite,
		Layout: spin.NewLayoutRect(388, 427, 18, 8, spin.ColorLampWhite),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L27,
		ID:     LampRightRampCrimeSceneYellow,
		Layout: spin.NewLayoutRect(388, 443, 18, 8, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L86,
		ID:     LampRightTank,
		Layout: spin.NewLayoutRect(316, 359, 31, 14, spin.ColorLampGreen),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L52,
		ID:     LampSafeCracker,
		Layout: spin.NewLayoutRect(352, 545, 26, 12, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L57,
		ID:     LampSniper,
		Layout: spin.NewLayoutRect(253, 547, 26, 12, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L51,
		ID:     LampStakeout,
		Layout: spin.NewLayoutRect(412, 542, 26, 12, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L88,
		ID:     LampStartButton,
		Layout: spin.NewLayoutCircleFromRect(36, 780, 12, 12, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterLamp{
		Addr: wpc.L58,
		ID:   LampSubwayCombo,
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L38,
		ID:     LampSuperGameButton,
		Layout: spin.NewLayoutCircleFromRect(34, 820, 15, 10, spin.ColorLampGreen),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L54,
		ID:     LampUltimateChallenge,
		Layout: spin.NewLayoutRect(296, 531, 39, 18, spin.ColorLampOrange),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.L44,
		ID:     LampWarning,
		Layout: spin.NewLayoutRect(263, 600, 20, 42, spin.ColorLampGreen),
	})

	eng.Do(spin.RegisterLamp{
		Addr:   wpc.G01,
		ID:     GI1,
		Layout: spin.NewLayoutRect(150, 5, 15, 15, spin.ColorLampWhite),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.G02,
		ID:     GI2,
		Layout: spin.NewLayoutRect(170, 5, 15, 15, spin.ColorLampWhite),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.G03,
		ID:     GI3,
		Layout: spin.NewLayoutRect(190, 5, 15, 15, spin.ColorLampWhite),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.G04,
		ID:     GI4,
		Layout: spin.NewLayoutRect(210, 5, 15, 15, spin.ColorLampWhite),
	})
	eng.Do(spin.RegisterLamp{
		Addr:   wpc.G05,
		ID:     GI5,
		Layout: spin.NewLayoutRect(230, 5, 15, 15, spin.ColorLampWhite),
	})
}

func RegisterFlashers(eng *spin.Engine) {
	eng.Do(spin.RegisterFlasher{
		Addr:   wpc.C23,
		ID:     FlasherBlackout,
		Layout: spin.NewLayoutRect(284, 673, 63, 20, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterFlasher{
		Addr: wpc.C24,
		ID:   FlasherCursedEarth,
	})
	eng.Do(spin.RegisterFlasher{
		Addr:   wpc.C26,
		ID:     FlasherGlobe,
		Layout: spin.NewLayoutCircleFromRect(226, 164, 43, 44, spin.ColorLampRed),
	})
	eng.Do(spin.RegisterFlasher{
		Addr: wpc.C28,
		ID:   FlasherInsert,
	})
	eng.Do(spin.RegisterFlasher{
		Addr:   wpc.C19,
		ID:     FlasherJudgeDeath,
		Layout: spin.NewLayoutCircleFromRect(313, 494, 15, 17, spin.ColorLampRed),
	})
	eng.Do(spin.RegisterFlasher{
		Addr:   wpc.C17,
		ID:     FlasherJudgeFire,
		Layout: spin.NewLayoutCircleFromRect(365, 494, 15, 17, spin.ColorLampYellow),
	})
	eng.Do(spin.RegisterFlasher{
		Addr:   wpc.C18,
		ID:     FlasherJudgeFear,
		Layout: spin.NewLayoutCircleFromRect(234, 496, 15, 17, spin.ColorLampBlue),
	})
	eng.Do(spin.RegisterFlasher{
		Addr:   wpc.C20,
		ID:     FlasherJudgeMortis,
		Layout: spin.NewLayoutCircleFromRect(272, 495, 15, 17, spin.ColorLampGreen),
	})
	eng.Do(spin.RegisterFlasher{
		Addr: wpc.C21,
		ID:   FlasherLeftPursuit,
		Layout: spin.NewLayoutMulti(
			spin.NewLayoutRect(171, 295, 25, 21, spin.ColorLampBlue),
			spin.NewLayoutRect(218, 286, 25, 21, spin.ColorLampRed),
		)})
	eng.Do(spin.RegisterFlasher{
		Addr: wpc.C22,
		ID:   FlasherRightPursuit,
		Layout: spin.NewLayoutMulti(
			spin.NewLayoutRect(381, 361, 22, 20, spin.ColorLampBlue),
			spin.NewLayoutRect(429, 375, 22, 20, spin.ColorLampRed),
		),
	})
	eng.Do(spin.RegisterFlasher{
		Addr: wpc.C27,
		ID:   FlasherRightRamp,
	})
	eng.Do(spin.RegisterFlasher{
		Addr:   wpc.C25,
		ID:     FlasherSubwayExit,
		Layout: spin.NewLayoutCircleFromRect(121, 676, 9, 10, spin.ColorLampWhite),
	})
}

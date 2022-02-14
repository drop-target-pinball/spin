package jdx

import (
	"fmt"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

// Scores
const (
	ScoreBadImpersonator0 = 3_000_000
	ScoreBadImpersonatorN = 5_000_000
	ScoreBlackoutJackpot  = 10_000_000
	ScoreDropTargetLit    = 500_000
	ScoreLeftRampN        = 500_000
	ScoreManhunt0         = 3_000_000
	ScoreManhuntN         = 6_000_000
	ScoreMeltdown0        = 3_000_000
	ScoreMeltdown1        = 13_000_000
	ScoreMeltdown2        = 23_000_000
	ScoreMeltdown3        = 33_000_000
	ScoreMystery          = 50_000
	ScoreOutlane          = 100_000
	ScorePost             = 5_210
	ScorePursuit0         = 3_000_000
	ScorePursuit1         = 6_000_000
	ScorePursuit2         = 12_000_000
	ScorePursuit3         = 36_000_000
	ScoreReturnLane       = 50_000
	ScoreRightRampN       = 500_000
	ScoreSafecrackerStart = 8_000_000
	ScoreSafecrackerEnd   = 3_000_000
	ScoreSafecrackerDec   = 75_400
	ScoreSling            = 1_930
	ScoreSniperStart      = 20_000_000
	ScoreSniperEnd        = 5_000_000
	ScoreSniperDec        = 78_330
	ScoreStakeout0        = 3_000_000
	ScoreStakeoutN        = 5_000_000
	ScoreTank0            = 3_000_000
	ScoreTank1            = 12_000_000
	ScoreTank2            = 24_000_000
	ScoreTank3            = 36_000_000
	ScoreTopLeftRampN     = 1_000_000

	ScoreMinimumBonus    = 1_000_000
	ScoreCrimeSceneBonus = 1_000_000
	MaxRampScore         = 5_000_000
)

const (
	ModePursuit = 1 << iota
	ModeBlackout
	ModeSniper
	ModeBattleTank
	ModeBadImpersonator
	ModeMeltdown
	ModeSafeCracker
	ModeManhunt
	ModeStakeout
	ModePlunge
	ModeAirRaid
	ModeNone
)

const (
	AllChainModes = ModePursuit | ModeBlackout | ModeSniper | ModeBattleTank | ModeBadImpersonator | ModeMeltdown | ModeSafeCracker | ModeManhunt | ModeStakeout
	MinChainMode  = ModePursuit
	MaxChainMode  = ModeStakeout
	MaxPlayers    = 4
)

var (
	Modes = []int{
		ModePursuit,
		ModeBlackout,
		ModeSniper,
		ModeBattleTank,
		ModeBadImpersonator,
		ModeMeltdown,
		ModeSafeCracker,
		ModeManhunt,
		ModeStakeout,
	}

	ModeLamps = map[int]string{
		ModePursuit:         jd.LampPursuit,
		ModeBlackout:        jd.LampBlackout,
		ModeSniper:          jd.LampSniper,
		ModeBattleTank:      jd.LampBattleTank,
		ModeBadImpersonator: jd.LampBadImpersonator,
		ModeMeltdown:        jd.LampMeltdown,
		ModeSafeCracker:     jd.LampSafeCracker,
		ModeManhunt:         jd.LampManhunt,
		ModeStakeout:        jd.LampStakeout,
	}

	ModeScripts = map[int]string{
		ModePursuit:         ScriptPursuitMode,
		ModeBlackout:        ScriptBlackoutMode,
		ModeSniper:          ScriptSniperMode,
		ModeBattleTank:      ScriptTankMode,
		ModeBadImpersonator: ScriptBadImpersonatorMode,
		ModeMeltdown:        ScriptMeltdownMode,
		ModeSafeCracker:     ScriptSafecrackerMode,
		ModeManhunt:         ScriptManhuntMode,
		ModeStakeout:        ScriptStakeoutMode,
	}
)

type Vars struct {
	AttractModeSlide       int
	AwardedModes           int
	BadImpersonatorBonus   int
	BadImpersonatorTargets int
	BallsLocked            int
	CrimeScenes            int
	LeftRampsMade          int
	LitDropTarget          int
	LocksReady             int
	ManhuntBonus           int
	MeltdownBonus          int
	Mode                   int
	Multiplier             int
	MultiballAttempted     bool
	PursuitBonus           int
	RightRampsMade         int
	SafecrackerAttempts    int
	SafecrackerBonus       int
	SafecrackerScore       int
	SelectedMode           int
	SniperBonus            int
	SniperScore            int
	StakeoutBonus          int
	StakeoutCallout        int
	StartModeLeft          bool
	TankBonus              int
	TankHits               int
	Timer                  int
	TopLeftRampsMade       int
}

func GetVarsFor(store spin.Store, player int) *Vars {
	name := fmt.Sprintf("jdx.%v", player)
	var vars *Vars

	v, ok := store.GetVars(name)
	if ok {
		vars = v.(*Vars)
	} else {
		vars = &Vars{}
		store.RegisterVars(name, vars)
	}
	return vars
}

func GetVars(store spin.Store) *Vars {
	game := spin.GetGameVars(store)
	return GetVarsFor(store, game.Player)
}

func Multiplier(store spin.Store) int {
	vars := GetVars(store)
	if vars.Multiplier > 0 {
		return vars.Multiplier
	}
	return 1
}

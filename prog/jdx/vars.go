package jdx

import (
	"fmt"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

// Scores
const (
	ScoreReturnLane = 50_000
	ScoreSling      = 1_930
	ScoreOutlane    = 100_000
	ScorePost       = 5_210

	ScorePursuit0 = 3_000_000
	ScorePursuit1 = 6_000_000
	ScorePursuit2 = 12_000_000
	ScorePursuit3 = 36_000_000
	ScoreTank0    = 3_000_000
	ScoreTank1    = 12_000_000
	ScoreTank2    = 24_000_000
	ScoreTank3    = 36_000_000

	ScoreMinimumBonus    = 1_000_000
	ScoreCrimeSceneBonus = 1_000_000
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
)

const (
	AllModes   = ModePursuit | ModeBlackout | ModeSniper | ModeBattleTank | ModeBadImpersonator | ModeMeltdown | ModeSafeCracker | ModeManhunt | ModeStakeout
	MinMode    = ModePursuit
	MaxMode    = ModeStakeout
	MaxPlayers = 4
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
		ModeBlackout:        ScriptSniperMode,
		ModeSniper:          ScriptSniperMode,
		ModeBattleTank:      ScriptTankMode,
		ModeBadImpersonator: ScriptSniperMode,
		ModeMeltdown:        ScriptSniperMode,
		ModeSafeCracker:     ScriptSniperMode,
		ModeManhunt:         ScriptSniperMode,
		ModeStakeout:        ScriptSniperMode,
	}
)

type Vars struct {
	AwardedModes int
	CrimeScenes  int
	PursuitBonus int
	SelectedMode int
	SniperBonus  int
	SniperScore  int
	TankBonus    int
	Timer        int
}

func startOfBallReset(store spin.Store) {
	vars := GetVars(store)
	vars.PursuitBonus = 0
	vars.SniperBonus = 0
	vars.SniperScore = 0
	if vars.SelectedMode == 0 {
		vars.SelectedMode = ModeStakeout // FIXME
	}
}

func GetVars(store spin.Store) *Vars {
	game := spin.GetGameVars(store)
	name := fmt.Sprintf("jdx.%v", game.Player)
	var vars *Vars

	v, ok := store.Vars(name)
	if ok {
		vars = v.(*Vars)
	} else {
		vars = &Vars{}
		store.RegisterVars(name, vars)
	}
	return vars
}

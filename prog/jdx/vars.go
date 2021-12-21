package jdx

import (
	"fmt"
	"log"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

// Scores
const (
	ScoreReturnLane = 50_000
	ScoreSling      = 1_930
	ScoreOutlane    = 100_000
	ScorePost       = 5_210

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
		ModePursuit:         ScriptSniperMode,
		ModeBlackout:        ScriptSniperMode,
		ModeSniper:          ScriptSniperMode,
		ModeBattleTank:      ScriptSniperMode,
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
	SelectedMode int
	SniperBonus  int
	SniperScore  int
}

func startOfBallReset(store spin.Store) {
	vars := GetVars(store)
	vars.SniperBonus = 0
	vars.SniperScore = 0
	if vars.SelectedMode == 0 {
		vars.SelectedMode = ModeSniper
	}
}

func GetVars(store spin.Store) *Vars {
	game := spin.GetGameVars(store)
	name := fmt.Sprintf("jdx.%v", game.Player)

	v, ok := store.Vars(name)
	if !ok {
		log.Panicf("no such vars: %v", name)
	}
	return v.(*Vars)
}

func RegisterVars(e *spin.Engine) {
	for i := 0; i <= MaxPlayers; i++ {
		id := fmt.Sprintf("jdx.%v", i)
		vars := &Vars{}
		e.RegisterVars(id, vars)
	}
}

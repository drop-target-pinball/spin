package jdx

import (
	"fmt"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
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
	AllModes = ModePursuit | ModeBlackout | ModeSniper | ModeBattleTank | ModeBadImpersonator | ModeMeltdown | ModeSafeCracker | ModeManhunt | ModeStakeout
	MinMode  = ModePursuit
	MaxMode  = ModeStakeout
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
	SelectedMode int
}

func GetVars(store spin.Store) *Vars {
	game := spin.GameVars(store)
	name := fmt.Sprintf("jdx.%v", game.Player)

	v, ok := store.Vars(name)
	var vars *Vars
	if ok {
		vars = v.(*Vars)
	} else {
		vars = &Vars{
			SelectedMode: ModeSniper,
		}
		store.RegisterVars(name, vars)
	}
	return vars
}

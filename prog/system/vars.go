package system

import "github.com/drop-target-pinball/spin"

const (
	GameJudgeDreddRemix = iota
	GameMegaMan3
	GameDrMario
	GamePinGolf
	GamePractice
)

type Vars struct {
	Game  int
	Games []string
}

func GetVars(store spin.Store) *Vars {
	v, ok := store.Vars("system")
	var vars *Vars
	if ok {
		vars = v.(*Vars)
	} else {
		vars = &Vars{
			Games: []string{
				"DREDD REMIX",
				"MEGAMAN 3",
				"DR MARIO",
				"PINGOLF",
				"PRACTICE",
			},
		}
		store.RegisterVars("system", vars)
	}
	return vars
}

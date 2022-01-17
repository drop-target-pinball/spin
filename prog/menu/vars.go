package menu

import "github.com/drop-target-pinball/spin"

const (
	GameJudgeDreddRemix = iota
	GameMegaMan3
	GameDrMario
	GamePinGolf
	GamePractice
	GameNone
)

type Vars struct {
	AttractModeSlide int
	Game             int
	Games            []string
}

func GetVars(store spin.Store) *Vars {
	v, ok := store.GetVars("menu")
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
		store.RegisterVars("menu", vars)
	}
	return vars
}

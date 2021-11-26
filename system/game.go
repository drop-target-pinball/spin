package system

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/game"
)

type Game struct{}

func RegisterGame(eng *spin.Engine) {
	s := &Game{}
	eng.RegisterActionHandler(s)
}

func (s *Game) HandleAction(action spin.Action) {
	switch act := action.(type) {
	case spin.Score:
		s.score(act)
	}
}

func (s *Game) score(act spin.Score) {
	game.VarPlayer.Score += act.Add
}

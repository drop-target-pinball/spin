package spin

type Player struct {
	Score int
}

var (
	VarBallsPerGame  int
	VarCurrentBall   int
	VarMaxPlayers    int
	VarCurrentPlayer int
	VarPlayers       []Player
	VarPlayer        *Player
)

func init() {
	VarMaxPlayers = 4
	VarPlayers = make([]Player, VarMaxPlayers)
	VarPlayer = &VarPlayers[0]
}

type gameSystem struct{}

func RegisterGameSystem(eng *Engine) {
	s := &gameSystem{}
	eng.RegisterActionHandler(s)
}

func (s gameSystem) HandleAction(action Action) {
	switch act := action.(type) {
	case Score:
		s.score(act)
	}
}

func (s *gameSystem) score(act Score) {
	VarPlayer.Score += act.Add
}

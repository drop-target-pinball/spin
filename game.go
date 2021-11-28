package spin

type gameSystem struct {
	eng  *Engine
	game *Game
}

func registerGameSystem(eng *Engine) {
	s := &gameSystem{
		eng:  eng,
		game: GameVars(eng),
	}
	eng.RegisterActionHandler(s)
}

func (s gameSystem) HandleAction(action Action) {
	switch act := action.(type) {
	case AwardScore:
		s.awardScore(act)
	}
}

func (s *gameSystem) awardScore(act AwardScore) {
	s.game.AddScore(act.Add)
}

type GameKeeper interface {
	CurrentBall() int
	ScoreFor(int) int
	Score() int
}

type Game struct {
	BallsPerGame  int
	CurrentBall   int
	MaxPlayers    int
	CurrentPlayer int
	Scores        []int
}

func (g *Game) Score() int {
	return g.Scores[g.CurrentPlayer]
}

func (g *Game) AddScore(s int) {
	g.Scores[g.CurrentPlayer] += s
}

func GameVars(store Store) *Game {
	v, ok := store.Vars("game")
	var vars *Game
	if ok {
		vars = v.(*Game)
	} else {
		vars = &Game{}
		vars.MaxPlayers = 4
		vars.Scores = make([]int, vars.MaxPlayers)
		store.RegisterVars("game", vars)
	}
	return vars
}

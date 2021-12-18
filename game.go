package spin

type Game struct {
	BallsPerGame int
	Player       int
	Ball         int
	MaxPlayers   int
	NumPlayers   int
	Scores       []int
	ExtraBalls   int
	IsExtraBall  bool
	BallActive   bool
}

func (g *Game) Score() int {
	return g.Scores[g.Player]
}

func (g *Game) AddScore(s int) {
	g.Scores[g.Player] += s
}

func (g *Game) SetScore(s int) {
	g.Scores[g.Player] = s
}

func GameVars(store Store) *Game {
	v, ok := store.Vars("game")
	var vars *Game
	if ok {
		vars = v.(*Game)
	} else {
		vars = &Game{}
		vars.BallsPerGame = 3
		vars.MaxPlayers = 4
		vars.Scores = make([]int, vars.MaxPlayers+1)
		store.RegisterVars("game", vars)
	}
	return vars
}

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
	case AddPlayer:
		s.addPlayer(act)
	case AdvanceGame:
		s.advanceGame(act)
	case AwardScore:
		s.awardScore(act)
	case SetScore:
		s.setScore(act)
	}
}

func (s *gameSystem) addPlayer(act AddPlayer) {
	if s.game.Ball > 1 {
		return
	}
	if s.game.NumPlayers == s.game.MaxPlayers {
		return
	}
	s.game.NumPlayers += 1
	s.eng.Post(PlayerAddedEvent{Player: s.game.NumPlayers})
}

func (s *gameSystem) advanceGame(act AdvanceGame) {
	g := s.game

	if g.NumPlayers == 0 {
		return
	}

	if g.Ball == 0 {
		for i := range g.Scores {
			g.Scores[i] = 0
		}
		g.Ball = 1
		g.Player = 1
		g.BallActive = true
		s.eng.Post(StartOfBallEvent{Player: 1, Ball: 1})
		return
	}

	if g.BallActive {
		s.eng.Post(EndOfBallEvent{Player: g.Player, Ball: g.Ball})
		g.BallActive = false
		return
	}

	if g.Ball == g.BallsPerGame && g.Player == g.NumPlayers {
		g.Ball = 0
		g.Player = 0
		g.NumPlayers = 0
		s.eng.Post(EndOfGameEvent{})
		return
	}

	shootAgain := false
	if g.ExtraBalls > 0 {
		g.ExtraBalls -= 1
		shootAgain = true
	} else {
		g.Player += 1
		if g.Player > g.NumPlayers {
			g.Player = 1
			g.Ball += 1
		}
	}
	g.BallActive = true
	s.eng.Post(StartOfBallEvent{
		Player:     g.Player,
		Ball:       g.Ball,
		ShootAgain: shootAgain,
	})
}

func (s *gameSystem) awardScore(act AwardScore) {
	s.game.AddScore(act.Val)
}

func (s *gameSystem) setScore(act SetScore) {
	s.game.SetScore(act.Val)
}

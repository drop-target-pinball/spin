package spin

type gameSystem struct {
	eng    *Engine
	config *Config
	game   *GameVars
}

func registerGameSystem(eng *Engine) {
	s := &gameSystem{
		eng:    eng,
		config: &eng.Config,
		game:   GetGameVars(eng),
	}
	eng.RegisterActionHandler(s)
}

func (s gameSystem) HandleAction(action Action) {
	switch act := action.(type) {
	case AddBall:
		s.addBall(act)
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

func (s gameSystem) addBall(act AddBall) {
	if s.game.BallsInPlay >= s.config.NumBalls {
		return
	}
	s.game.BallsInPlay += 1
	s.eng.Post(BallAddedEvent{BallsInPlay: s.game.BallsInPlay})
}

func (s *gameSystem) addPlayer(act AddPlayer) {
	game := s.game
	if game.Ball > 1 {
		return
	}
	if game.NumPlayers == game.MaxPlayers {
		return
	}
	game.NumPlayers += 1
	s.eng.Post(PlayerAddedEvent{Player: game.NumPlayers})
}

func (s *gameSystem) advanceGame(act AdvanceGame) {
	g := s.game

	if g.NumPlayers == 0 {
		return
	}

	if g.Ball == 0 {
		for i := 1; i < g.NumPlayers; i++ {
			player := GetPlayerVarsFor(s.eng, i)
			player.Score = 0
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
	player := GetPlayerVars(s.eng)
	player.Score += act.Val
}

func (s *gameSystem) setScore(act SetScore) {
	player := GetPlayerVars(s.eng)
	player.Score = act.Val
}

package game

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

func Reset() {
	VarBallsPerGame = 3
	VarMaxPlayers = 4
	VarCurrentPlayer = 0
	VarPlayers = make([]Player, VarMaxPlayers)
	VarPlayer = &VarPlayers[0]
}

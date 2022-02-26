package spin

import "fmt"

const (
	PriorityAudioModeCallout = 1
)

type GameVars struct {
	BallActive      bool
	BallsInPlay     int
	BallsPerGame    int
	BallSave        bool
	BallLaunchReady bool
	Player          int
	PlayfieldActive bool
	Ball            int
	MaxPlayers      int
	NumPlayers      int
	ExtraBalls      int
	IsExtraBall     bool
}

type PlayerVars struct {
	Score int
}

func GetGameVars(store Store) *GameVars {
	v, ok := store.GetVars("game")
	var vars *GameVars
	if ok {
		vars = v.(*GameVars)
	} else {
		vars = &GameVars{}
		vars.Player = 1
		vars.BallsPerGame = 3
		vars.MaxPlayers = 4
		store.RegisterVars("game", vars)
	}
	return vars
}

func GetPlayerVarsFor(store Store, player int) *PlayerVars {
	name := fmt.Sprintf("player.%v", player)
	v, ok := store.GetVars(name)
	var vars *PlayerVars
	if ok {
		vars = v.(*PlayerVars)
	} else {
		vars = &PlayerVars{}
		store.RegisterVars(name, vars)
	}
	return vars
}

func ResetPlayerVars(store Store) {
	game := GetGameVars(store)
	for i := 1; i < game.MaxPlayers; i++ {
		name := fmt.Sprintf("player.%v", i)
		store.RegisterVars(name, &PlayerVars{})
	}
}

func GetPlayerVars(store Store) *PlayerVars {
	game := GetGameVars(store)
	return GetPlayerVarsFor(store, game.Player)
}

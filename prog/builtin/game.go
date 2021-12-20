package builtin

import (
	"github.com/drop-target-pinball/spin"
)

func gameStartButtonScript(e spin.Env) {
	game := spin.GetGameVars(e)

	e.Do(spin.DriverOn{ID: e.Config.LampStartButton})

	for {
		evt, done := e.WaitFor(
			spin.SwitchEvent{ID: e.Config.SwitchStartButton},
			spin.GameOverEvent{},
		)
		if done || (evt == spin.GameOverEvent{}) {
			e.Do(spin.DriverOff{ID: e.Config.LampStartButton})
			return
		}
		if game.NumPlayers == game.MaxPlayers-1 {
			e.Do(spin.DriverOff{ID: e.Config.LampStartButton})
		}
		if game.NumPlayers < game.MaxPlayers {
			e.Do(spin.AddPlayer{})
		}
	}
}

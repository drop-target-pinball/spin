package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/prog/builtin"
)

func gameScript(e spin.Env) {
	for _, gi := range e.Config.GI {
		e.Do(spin.DriverOn{ID: gi})
	}
	e.Do(spin.AddPlayer{})
	e.Do(spin.AdvanceGame{})
	e.Do(spin.PlayScript{ID: ScriptPlayerAnnounce})
	e.Do(spin.PlayScript{ID: builtin.ScriptGameStartButton})
	e.Do(spin.PlayScript{ID: builtin.ScriptScore})
	e.Do(spin.PlayScript{ID: ScriptPlunge})

	e.Do(spin.PlayScript{ID: ScriptBall})
}

var playerSpeech = map[int]string{
	2: SpeechPlayer2,
	3: SpeechPlayer3,
	4: SpeechPlayer4,
}

func playerAnnounceScript(e spin.Env) {
	for {
		event, done := e.WaitFor(spin.PlayerAddedEvent{})
		if done {
			return
		}
		switch evt := event.(type) {
		case spin.PlayerAddedEvent:
			speechID, ok := playerSpeech[evt.Player]
			if ok {
				e.Do(spin.PlaySpeech{ID: speechID})
			}
		}
	}
}

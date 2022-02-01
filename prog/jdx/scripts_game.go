package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

func gameScript(e *spin.ScriptEnv) {
	for _, gi := range e.Config.GI {
		e.Do(spin.DriverOn{ID: gi})
	}

	for i := 1; i < 4; i++ {
		vars := spin.GetPlayerVarsFor(e, i)
		vars.Score = 0
	}

	e.Do(spin.AddPlayer{})
	e.NewCoroutine(playerAnnounceScript)
	e.NewCoroutine(startButtonRoutine)

	for {
		e.Do(spin.AdvanceGame{})
		evt, done := e.WaitFor(spin.StartOfBallEvent{}, spin.EndOfGameEvent{})
		if done {
			return
		}
		if evt == (spin.EndOfGameEvent{}) {
			break
		}
		e.Do(spin.PlayScript{ID: ScriptBall})
		if _, done := e.WaitFor(spin.EndOfBallEvent{}); done {
			return
		}
		e.Do(spin.StopAudio{})
		if done := e.Sleep(1000); done {
			return
		}
		e.Do(spin.PlayScript{ID: ScriptBonusMode})
		if _, done := e.WaitFor(spin.ScriptFinishedEvent{ID: ScriptBonusMode}); done {
			return
		}
		e.Do(spin.AllLampsOff{})
	}

	e.Do(spin.PlayScript{ID: ScriptMatchMode})
	if _, done := e.WaitFor(spin.ScriptFinishedEvent{ID: ScriptMatchMode}); done {
		return
	}
	e.Post(spin.GameOverEvent{})
}

var playerSpeech = map[int]string{
	2: SpeechPlayer2,
	3: SpeechPlayer3,
	4: SpeechPlayer4,
}

func playerAnnounceScript(e *spin.ScriptEnv) {
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

func startButtonRoutine(e *spin.ScriptEnv) {
	game := spin.GetGameVars(e)
	e.Do(spin.DriverOn{ID: jd.LampStartButton})
	defer e.Do(spin.DriverOff{ID: jd.LampStartButton})

	for {
		event, done := e.WaitFor(
			spin.SwitchEvent{ID: jd.SwitchStartButton},
			spin.StartOfBallEvent{})
		if done {
			return
		}
		switch evt := event.(type) {
		case spin.SwitchEvent:
			game.NumPlayers += 1
			e.Post(spin.PlayerAddedEvent{Player: game.NumPlayers})
			if game.NumPlayers == game.MaxPlayers {
				return
			}
		case spin.StartOfBallEvent:
			if evt.Ball == 2 {
				return
			}
		}
	}
}

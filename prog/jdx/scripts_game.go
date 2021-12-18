package jdx

import (
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/prog/builtin"
)

func gameScript(e spin.Env) {
	for _, gi := range e.Config.GI {
		e.Do(spin.DriverOn{ID: gi})
	}
	e.Do(spin.AddPlayer{})
	e.Do(spin.PlayScript{ID: ScriptPlayerAnnounce})
	e.Do(spin.PlayScript{ID: builtin.ScriptGameStartButton})

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
		e.Do(spin.StopScope{ID: spin.ScopeBall})
		e.Do(spin.StopAudio{})
		if done := e.Sleep(3 * time.Second); done {
			return
		}
	}

	e.Do(spin.StopScope{ID: spin.ScopeGame})
	e.Do(spin.PlayScript{ID: ScriptMatch})
	e.WaitFor(spin.GameOverEvent{})
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

func matchScript(e spin.Env) {
	r, g := e.Display("").Renderer("")

	r.Fill(spin.ColorBlack)
	g.Y = 2
	g.W = r.Width()
	g.H = r.Height()
	g.Font = builtin.FontPfRondaSevenBold8
	r.Print(g, "GAME OVER")

	e.Do(spin.PlayMusic{ID: MusicMatch, Loops: 1, Notify: true})
	if _, done := e.WaitFor(spin.MusicFinishedEvent{}); done {
		return
	}
	e.Do(spin.PlayMusic{ID: MusicMatchHit, Loops: 1, Notify: true})
	if _, done := e.WaitFor(spin.MusicFinishedEvent{}); done {
		return
	}
	e.Post(spin.GameOverEvent{})
}

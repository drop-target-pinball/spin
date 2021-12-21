package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/prog/builtin"
)

func matchModeScript(e spin.Env) {
	r, g := e.Display("").Renderer("")

	r.Fill(spin.ColorBlack)
	g.AnchorY = spin.AnchorMiddle
	g.Y = r.Height() / 2
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
	e.Post(spin.ScriptFinishedEvent{ID: ScriptMatchMode})
}

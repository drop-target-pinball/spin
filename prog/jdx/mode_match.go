package jdx

import "github.com/drop-target-pinball/spin"

/*
PlayScript ID=jdx.ScriptMatchMode
*/
func matchModeScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(0)
	defer r.Close()

	g := r.Graphics()
	r.Fill(spin.ColorBlack)
	g.AnchorY = spin.AnchorMiddle
	g.Y = r.Height() / 2
	g.Font = spin.FontPfRondaSevenBold8
	r.Print(g, "GAME OVER")

	e.Do(spin.PlayMusic{ID: MusicMatch, Loops: 1, Notify: true})
	if _, done := e.WaitFor(spin.MusicFinishedEvent{}); done {
		return
	}
	e.Do(spin.PlayMusic{ID: MusicMatchHit, Loops: 1, Notify: true})
	if _, done := e.WaitFor(spin.MusicFinishedEvent{}); done {
		return
	}
}

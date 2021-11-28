package builtin

import (
	"github.com/drop-target-pinball/spin"
)

// Sizing from
// https://github.com/preble/pyprocgame/blob/master/procgame/modes/scoredisplay.py#L104

func scoreFrame(e spin.Env) {
	r, g := e.Display("").Renderer()
	vars := spin.GameVars(e)

	r.Clear()
	g.Font = FontBm10

	g.W = r.Width()
	r.Print(g, spin.FormatScore("%d", vars.Score()))

	g.Font = FontBm3
	g.W = 0
	g.X, g.Y = 25, r.Height()-5
	r.Print(g, "BALL %v", vars.CurrentBall)
	g.X = 75
	r.Print(g, "FREE PLAY")
}

func scoreScript(e spin.Env) {
	for {
		scoreFrame(e)
		if done := e.Sleep(spin.FrameDuration); done {
			return
		}
	}
}

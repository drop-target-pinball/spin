package game

import "github.com/drop-target-pinball/spin"

func scoreFrame(e spin.Env) {
	r, g := e.Display("").Renderer()

	r.Clear()
	g.Font = FontBm10w

	g.W = r.Width()
	r.Print(g, formatScore())

	g.Font = FontBm3
	g.W = 0
	g.X, g.Y = 25, r.Height()-5
	r.Print(g, "BALL %v", VarCurrentBall)
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

func formatScore() string {
	score := spin.Sprintf("%d", VarPlayer.Score)
	if score == "0" {
		return "00"
	}
	return score
}

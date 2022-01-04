package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/prog/builtin"
)

func timerAndScorePanel(e spin.Env, r spin.Renderer, title string, instruction string) {
	vars := GetVars(e)
	player := spin.GetPlayerVars(e)
	g := r.Graphics()

	r.Fill(spin.ColorBlack)
	g.Y = 2
	g.Font = builtin.FontPfArmaFive8
	r.Print(g, title)

	g.AnchorY = spin.AnchorBottom
	g.Y = r.Height()
	r.Print(g, instruction)

	g.AnchorX = spin.AnchorLeft
	g.X = 5
	g.AnchorY = spin.AnchorMiddle
	g.Y = r.Height() / 2
	g.Font = builtin.Font14x10
	r.Print(g, "%v", vars.Timer)

	g.X = r.Width() - 2
	g.AnchorX = spin.AnchorRight
	g.Font = builtin.Font09x7
	r.Print(g, spin.FormatScore("%v", player.Score))
}

func modeAndScorePanel(e spin.Env, r spin.Renderer, title string, score int) {
	g := r.Graphics()

	r.Fill(spin.ColorBlack)
	g.Y = 2
	g.Font = builtin.FontPfArmaFive8
	r.Print(g, title)
	g.Y = 12

	g.Font = builtin.Font14x10
	r.Print(g, spin.FormatScore("%v", score))
}

func scoreAndLabelPanel(e spin.Env, r spin.Renderer, score int, label string) {
	g := r.Graphics()

	r.Fill(spin.ColorBlack)
	g.Y = 5
	g.Font = builtin.Font14x10
	r.Print(g, spin.FormatScore("%v", score))

	g.Y = 22
	g.Font = builtin.FontPfArmaFive8
	r.Print(g, label)
}

package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/prog/builtin"
)

func modeTotalPanel(e spin.Env, title string, total int) {
	r, g := e.Display("").Renderer(spin.LayerPriority)

	r.Fill(spin.ColorBlack)
	g.Y = 2
	g.Font = builtin.FontPfArmaFive8
	r.Print(g, title)
	g.Y = 12

	g.Font = builtin.Font14x10
	r.Print(g, spin.FormatScore("%v", total))
}

func scoreAwardedPanel(e spin.Env, score int) {
	r, g := e.Display("").Renderer(spin.LayerPriority)

	r.Fill(spin.ColorBlack)
	g.Y = 5
	g.Font = builtin.Font14x10
	r.Print(g, spin.FormatScore("%v", score))

	g.Y = 22
	g.Font = builtin.FontPfArmaFive8
	r.Print(g, "AWARDED")
}

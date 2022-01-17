package jdx

import (
	"github.com/drop-target-pinball/spin"
)

func ModeIntroPanel(e *spin.ScriptEnv, blinkOn bool, text [3]string) {
	r, g := e.Display("").Renderer("")

	r.Fill(spin.ColorBlack)
	g.Y = 2
	g.Font = spin.FontPfArmaFive8
	r.Print(g, text[0])
	if blinkOn {
		g.Y = 12
		g.Font = spin.FontPfRondaSevenBold8
		r.Print(g, text[1])
		g.Y = 22
		r.Print(g, text[2])
	}
}

func TimerAndScorePanel(e *spin.ScriptEnv, r spin.Renderer, title string, timer int, score int, instruction string) {
	g := r.Graphics()

	r.Fill(spin.ColorBlack)
	g.Y = 2
	g.Font = spin.FontPfArmaFive8
	r.Print(g, title)

	g.AnchorY = spin.AnchorBottom
	g.Y = r.Height()
	r.Print(g, instruction)

	yOffset := int32(0)
	if instruction == "" {
		yOffset = 5
	}

	g.AnchorX = spin.AnchorLeft
	g.X = 5
	g.AnchorY = spin.AnchorMiddle
	g.Y = r.Height()/2 + yOffset
	g.Font = spin.Font14x10
	r.Print(g, "%v", timer)

	g.X = r.Width() - 2
	g.AnchorX = spin.AnchorRight
	g.Font = spin.Font09x7
	r.Print(g, spin.FormatScore("%v", score))
}

func ModeAndScorePanel(e *spin.ScriptEnv, r spin.Renderer, title string, score int) {
	g := r.Graphics()

	r.Fill(spin.ColorBlack)
	g.Y = 2
	g.Font = spin.FontPfArmaFive8
	r.Print(g, title)
	g.Y = 12

	g.Font = spin.Font14x10
	r.Print(g, spin.FormatScore("%v", score))
}

func ModeAndBlinkingScorePanel(e *spin.ScriptEnv, r spin.Renderer, title string, score int, blinkOn bool) {
	g := r.Graphics()

	r.Fill(spin.ColorBlack)
	g.Y = 2
	g.Font = spin.FontPfArmaFive8
	r.Print(g, title)
	g.Y = 12

	if blinkOn {
		g.Font = spin.Font14x10
		score := spin.FormatScore("%10d", score)
		r.Print(g, score)
	}
}

func ScoreAndLabelPanel(e *spin.ScriptEnv, r spin.Renderer, score int, label string) {
	g := r.Graphics()

	r.Fill(spin.ColorBlack)
	g.Y = 5
	g.Font = spin.Font14x10
	r.Print(g, spin.FormatScore("%v", score))

	g.Y = 22
	g.Font = spin.FontPfArmaFive8
	r.Print(g, label)
}

func GameOverPanel(e *spin.ScriptEnv) {
}

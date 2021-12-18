package builtin

import (
	"github.com/drop-target-pinball/spin"
)

// Sizing from
// https://github.com/preble/pyprocgame/blob/master/procgame/modes/scoredisplay.py#L104

func singlePlayerPanel(e spin.Env) {
	r, g := e.Display("").Renderer("")
	vars := spin.GameVars(e)

	switch {
	case vars.Score() < 1_000_000_000:
		g.Font = Font18x12
	case vars.Score() < 10_000_000_000:
		g.Font = Font18x11
	default:
		g.Font = Font18x10
	}
	g.Y = 3
	g.W = r.Width()
	r.Print(g, spin.FormatScore("%d", vars.Score()))
}

func multiPlayerPanel(e spin.Env) {
	r, g := e.Display("").Renderer("")
	vars := spin.GameVars(e)

	sizedFont := func(player int) string {
		active := vars.Player == player
		switch {
		case active && vars.Scores[player] < 10_000_000:
			return Font14x10
		case active && vars.Scores[player] < 100_000_000:
			return Font14x9
		case active:
			return Font14x8
		case vars.Scores[player] < 10_000_000:
			return Font09x7
		case vars.Scores[player] < 10_000_000:
			return Font09x6
		default:
			return Font09x5
		}
	}

	g.X, g.Y = 0, 0
	g.Font = sizedFont(1)
	r.Print(g, spin.FormatScore("%d", vars.Scores[1]))

	g.X, g.Y = r.Width()+1, 0
	g.AnchorX = spin.AnchorRight
	g.Font = sizedFont(2)
	r.Print(g, spin.FormatScore("%d", vars.Scores[2]))

	if vars.NumPlayers >= 3 {
		g.X, g.Y = 0, r.Height()-6
		g.AnchorX = spin.AnchorLeft
		g.AnchorY = spin.AnchorBottom
		g.Font = sizedFont(3)
		r.Print(g, spin.FormatScore("%d", vars.Scores[3]))
	}

	if vars.NumPlayers == 4 {
		g.X, g.Y = r.Width(), r.Height()-6
		g.AnchorX = spin.AnchorRight
		g.AnchorY = spin.AnchorBottom
		g.Font = sizedFont(4)
		r.Print(g, spin.FormatScore("%d", vars.Scores[4]))
	}
}

func scoreFrame(e spin.Env) {
	r, g := e.Display("").Renderer("")
	game := spin.GameVars(e)

	r.Fill(spin.ColorBlack)
	if game.NumPlayers == 1 {
		singlePlayerPanel(e)
	} else if game.NumPlayers > 1 {
		multiPlayerPanel(e)
	}

	g.Font = Font04B_03_7px
	g.W = 0
	g.X, g.Y = 25, r.Height()-5
	r.Print(g, "BALL %v", game.Ball)
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

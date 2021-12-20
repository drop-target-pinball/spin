package builtin

import (
	"github.com/drop-target-pinball/spin"
)

// Sizing from
// https://github.com/preble/pyprocgame/blob/master/procgame/modes/scoredisplay.py#L104

func singlePlayerPanel(e spin.Env) {
	r, g := e.Display("").Renderer("")
	player := spin.GetPlayerVars(e)

	switch {
	case player.Score < 1_000_000_000:
		g.Font = Font18x12
	case player.Score < 10_000_000_000:
		g.Font = Font18x11
	default:
		g.Font = Font18x10
	}
	g.Y = 3
	g.W = r.Width()
	r.Print(g, spin.FormatScore("%d", player.Score))
}

func multiPlayerPanel(e spin.Env) {
	r, g := e.Display("").Renderer("")
	game := spin.GetGameVars(e)

	sizedFont := func(active bool, score int) string {
		//active := game.Player == playerNum
		//score := spin.GetPlayerVarsFor(e, playerNum).Score
		switch {
		case active && score < 10_000_000:
			return Font14x10
		case active && score < 100_000_000:
			return Font14x9
		case active:
			return Font14x8
		case score < 10_000_000:
			return Font09x7
		case score < 10_000_000:
			return Font09x6
		default:
			return Font09x5
		}
	}

	g.X, g.Y = 0, 0
	score := spin.GetPlayerVarsFor(e, 1).Score
	g.Font = sizedFont(game.Player == 1, score)
	r.Print(g, spin.FormatScore("%d", score))

	g.X, g.Y = r.Width()+1, 0
	g.AnchorX = spin.AnchorRight
	score = spin.GetPlayerVarsFor(e, 2).Score
	g.Font = sizedFont(game.Player == 2, score)
	score = spin.GetPlayerVarsFor(e, 2).Score
	r.Print(g, spin.FormatScore("%d", score))

	if game.NumPlayers >= 3 {
		g.X, g.Y = 0, r.Height()-6
		g.AnchorX = spin.AnchorLeft
		g.AnchorY = spin.AnchorBottom
		score = spin.GetPlayerVarsFor(e, 3).Score
		g.Font = sizedFont(game.Player == 3, score)
		r.Print(g, spin.FormatScore("%d", score))
	}

	if game.NumPlayers == 4 {
		g.X, g.Y = r.Width(), r.Height()-6
		g.AnchorX = spin.AnchorRight
		g.AnchorY = spin.AnchorBottom
		score = spin.GetPlayerVarsFor(e, 4).Score
		g.Font = sizedFont(game.Player == 4, score)
		r.Print(g, spin.FormatScore("%d", score))
	}
}

func scoreFrame(e spin.Env) {
	r, g := e.Display("").Renderer("")
	game := spin.GetGameVars(e)

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

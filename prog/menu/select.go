package menu

import (
	"context"
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

// Messages
const (
	GameUpdated = "GameUpdated"
)

// Variables
const (
	Game = "Game"
)

var games = []string{
	"DREDD REMIX",
	"MEGAMAN 3",
	"DR MARIO",
	"HIGH SPEED 3",
	"PINGOLF",
}

func gameSelectFrame(e spin.Env, blinkOn bool) {
	r, g := e.Display("").Renderer()
	defer r.Unlock()

	r.Clear()
	g.W = r.Width()
	g.Y = 7
	g.Font = PfTempestaFiveExtendedBold8
	r.Print(g, "GAME SELECT")
	g.Y = 18
	g.Font = PfTempestaFiveCompressedBold8
	game := games[e.Int(spin.System, Game)]
	r.Print(g, game)

	if blinkOn {
		g.W = 0
		g.X = 20
		g.Y = 18
		r.Print(g, ">>")
		g.X = 96
		r.Print(g, "<<")
	}
}

func menuSelectGame(ctx context.Context, e spin.Env) {
	e.SetInt(spin.System, Game, 0)
	for {
		gameSelectFrame(e, true)
		if _, done := spin.WaitForEventUntil(ctx, e, 256*time.Millisecond, spin.Message{ID: GameUpdated}); done {
			return
		}

		gameSelectFrame(e, false)
		if _, done := spin.WaitForEventUntil(ctx, e, 100*time.Millisecond, spin.Message{ID: GameUpdated}); done {
			return
		}
	}
}

func menuSelect(ctx context.Context, e spin.Env) {
	next := func() {
		game := e.Int(spin.System, Game)
		game += 1
		if game >= len(games) {
			game = 0
		}
		e.SetInt(spin.System, Game, game)
	}

	prev := func() {
		game := e.Int(spin.System, Game)
		game -= 1
		if game < 0 {
			game = len(games) - 1
		}
		e.SetInt(spin.System, Game, game)
	}

	e.Do(spin.PlayScript{ID: MenuSelectGame})
	e.Do(spin.PlayMusic{ID: SMB2CharSelect})
	for {
		evt, done := spin.WaitForEvents(ctx, e, []spin.Event{
			spin.Message{ID: MenuAttractAdvance},
			spin.SwitchEvent{ID: jd.LeftFlipperButton},
			spin.SwitchEvent{ID: jd.RightFlipperButton},
			spin.SwitchEvent{ID: jd.LeftFireButton},
			spin.SwitchEvent{ID: jd.RightFireButton},
		})
		if done {
			return
		}
		switch evt {
		case spin.Message{ID: MenuAttractAdvance}:
			next()
			e.Do(spin.PlaySound{ID: SMB2Scroll})
		case spin.SwitchEvent{ID: jd.LeftFlipperButton}:
			prev()
			e.Do(spin.PlaySound{ID: SMB2Scroll})
		case spin.SwitchEvent{ID: jd.RightFlipperButton}:
			next()
			e.Do(spin.PlaySound{ID: SMB2Scroll})
		case spin.SwitchEvent{ID: jd.LeftFireButton}:
			prev()
			e.Do(spin.PlaySound{ID: SMB2Scroll})
		case spin.SwitchEvent{ID: jd.RightFireButton}:
			next()
			e.Do(spin.PlaySound{ID: SMB2Scroll})
		}
		e.Post(spin.Message{ID: GameUpdated})
	}

}

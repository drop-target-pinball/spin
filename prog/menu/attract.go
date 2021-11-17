package menu

import (
	"context"
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

const (
	MenuAttractAdvance = "MenuAttractAdvance"
)

var attractScripts = []string{
	MenuAttractGameOver,
	MenuAttractDropTargetPinball,
	MenuAttractFreePlay,
}

func gameOverFrame(e spin.Env) {
	r, g := e.Display("").Renderer()
	defer r.Unlock()

	r.Clear()
	g.Y = 2
	g.W = r.Width()
	g.H = r.Height()
	g.Font = PfRondaSevenBold8
	r.Print(g, "GAME OVER")
}

func menuAttractGameOver(ctx context.Context, e spin.Env) {
	gameOverFrame(e)
	if done := spin.Wait(ctx, 4000*time.Millisecond); done {
		return
	}
	e.Post(spin.Message{ID: MenuAttractAdvance})
}

func dropTargetPinballFrame(e spin.Env) {
	r, g := e.Display("").Renderer()
	defer r.Unlock()

	r.Clear()
	g.W = r.Width()
	g.Y = 7
	g.Font = PfArmaFive8
	r.Println(g, "DROP TARGET PINBALL")
	g.Y = 18
	g.Font = PfRondaSevenBold8
	r.Println(g, "PRESENTS")
}

func menuAttractDropTargetPinball(ctx context.Context, e spin.Env) {
	dropTargetPinballFrame(e)
	if done := spin.Wait(ctx, 4000*time.Millisecond); done {
		return
	}
	e.Post(spin.Message{ID: MenuAttractAdvance})
}

func freePlayFrame(e spin.Env, blinkOn bool) {
	r, g := e.Display("").Renderer()
	defer r.Unlock()

	r.Clear()
	g.W = r.Width()
	g.Y = 7
	g.Font = PfRondaSevenBold8
	if blinkOn {
		r.Print(g, "PRESS START")
	}
	g.Y = 18
	r.Print(g, "FREE PLAY")
}

func menuAttractFreePlay(ctx context.Context, e spin.Env) {
	for i := 0; i < 5; i++ {
		freePlayFrame(e, true)
		if done := spin.Wait(ctx, 200*time.Millisecond); done {
			return
		}

		freePlayFrame(e, false)
		if done := spin.Wait(ctx, 100*time.Millisecond); done {
			return
		}
	}
	freePlayFrame(e, true)
	if done := spin.Wait(ctx, 2500*time.Millisecond); done {
		return
	}
	e.Post(spin.Message{ID: MenuAttractAdvance})
}

func menuAttractMode(ctx context.Context, e spin.Env) {
	script := 0

	next := func() {
		script += 1
		if script >= len(attractScripts) {
			script = 0
		}
	}

	prev := func() {
		script -= 1
		if script < 0 {
			script = len(attractScripts) - 1
		}
	}

	for {
		e.Do(spin.PlayScript{ID: attractScripts[script]})
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
		case spin.SwitchEvent{ID: jd.LeftFlipperButton}:
			e.Do(spin.StopScript{ID: attractScripts[script]})
			prev()
		case spin.SwitchEvent{ID: jd.RightFlipperButton}:
			e.Do(spin.StopScript{ID: attractScripts[script]})
			next()
		case spin.SwitchEvent{ID: jd.LeftFireButton}:
			e.Do(spin.StopScript{ID: attractScripts[script]})
			prev()
		case spin.SwitchEvent{ID: jd.RightFireButton}:
			e.Do(spin.StopScript{ID: attractScripts[script]})
			next()
		}
	}
}

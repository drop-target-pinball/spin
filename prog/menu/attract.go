package menu

import (
	"context"
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

const (
	MessageAttractAdvance = "menu.MessageAttractAdvance"
	MessageAttractDone    = "menu.MessageAttractDone"
)

var attractScripts = []spin.Script{
	gameOverScript,
	dropTargetPinballScript,
	freePlayScript,
}

func gameOverFrame(e spin.Env) {
	r, g := e.Display("").Renderer()

	r.Clear()
	g.Y = 2
	g.W = r.Width()
	g.H = r.Height()
	g.Font = FontPfRondaSevenBold8
	r.Print(g, "GAME OVER")
}

func gameOverScript(e spin.Env) {
	gameOverFrame(e)
	if done := e.Sleep(4000 * time.Millisecond); done {
		return
	}
	e.Post(spin.Message{ID: MessageAttractAdvance})
}

func dropTargetPinballFrame(e spin.Env) {
	r, g := e.Display("").Renderer()

	r.Clear()
	g.W = r.Width()
	g.Y = 7
	g.Font = FontPfArmaFive8
	r.Println(g, "DROP TARGET PINBALL")
	g.Y = 18
	g.Font = FontPfRondaSevenBold8
	r.Println(g, "PRESENTS")
}

func dropTargetPinballScript(e spin.Env) {
	dropTargetPinballFrame(e)
	if done := e.Sleep(4000 * time.Millisecond); done {
		return
	}
	e.Post(spin.Message{ID: MessageAttractAdvance})
}

func freePlayFrame(e spin.Env, blinkOn bool) {
	r, g := e.Display("").Renderer()

	r.Clear()
	g.W = r.Width()
	g.Y = 7
	g.Font = FontPfRondaSevenBold8
	if blinkOn {
		r.Print(g, "PRESS START")
	}
	g.Y = 18
	r.Print(g, "FREE PLAY")
}

func freePlayScript(e spin.Env) {
	for i := 0; i < 5; i++ {
		freePlayFrame(e, true)
		if done := e.Sleep(200 * time.Millisecond); done {
			return
		}

		freePlayFrame(e, false)
		if done := e.Sleep(100 * time.Millisecond); done {
			return
		}
	}
	freePlayFrame(e, true)
	if done := e.Sleep(2500 * time.Millisecond); done {
		return
	}
	e.Post(spin.Message{ID: MessageAttractAdvance})
}

func attractModeScript(e spin.Env) {
	script := 0
	var ctx context.Context
	var cancel context.CancelFunc

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
		ctx, cancel = e.Derive()
		e.NewCoroutine(ctx, attractScripts[script])
		evt, done := e.WaitFor(
			spin.Message{ID: MessageAttractAdvance},
			spin.SwitchEvent{ID: jd.SwitchLeftFlipperButton},
			spin.SwitchEvent{ID: jd.SwitchRightFlipperButton},
			spin.SwitchEvent{ID: jd.SwitchLeftFireButton},
			spin.SwitchEvent{ID: jd.SwitchRightFireButton},
			spin.SwitchEvent{ID: jd.SwitchStartButton},
		)
		if done {
			cancel()
			return
		}
		switch evt {
		case spin.Message{ID: MessageAttractAdvance}:
			next()
		case spin.SwitchEvent{ID: jd.SwitchLeftFlipperButton}:
			cancel()
			prev()
		case spin.SwitchEvent{ID: jd.SwitchRightFlipperButton}:
			cancel()
			next()
		case spin.SwitchEvent{ID: jd.SwitchLeftFireButton}:
			cancel()
			prev()
		case spin.SwitchEvent{ID: jd.SwitchRightFireButton}:
			cancel()
			next()
		case spin.SwitchEvent{ID: jd.SwitchStartButton}:
			cancel()
			e.Post(spin.Message{ID: MessageAttractDone})
			return
		}
	}
}

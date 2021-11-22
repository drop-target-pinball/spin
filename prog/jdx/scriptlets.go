package jdx

import (
	"time"

	"github.com/drop-target-pinball/spin"
)

func modeIntroFrame(e spin.Env, blinkOn bool, text [3]string) {
	r, g := e.Display("").Renderer()
	//defer r.Unlock()

	r.Clear()
	g.Y = 2
	g.W = r.Width()
	g.Font = FontPfArmaFive8
	r.Print(g, text[0])
	if blinkOn {
		g.Y = 12
		g.Font = FontPfRondaSevenBold8
		r.Print(g, text[1])
		g.Y = 22
		r.Print(g, text[2])
	}
}

func modeIntroVideo(e spin.Env, text [3]string) bool {
	for i := 0; i < 8; i++ {
		modeIntroFrame(e, true, text)
		if done := e.WaitFor(250 * time.Millisecond); done {
			return done
		}
		modeIntroFrame(e, false, text)
		if done := e.WaitFor(100 * time.Millisecond); done {
			return done
		}
	}
	return false
}

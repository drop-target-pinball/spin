package boot

import (
	"context"
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

// Scripts
const (
	SplashScreen = "SplashScreen"
)

// Messages
const (
	BootEnd = "BootEnd"
)

func splashScreenFrame(e spin.Env) {
	r, g := e.Display("").Renderer()
	defer r.Unlock()

	g.Font = PfTempestaFiveCompressedBold8
	g.W = r.Width()
	g.PaddingV = 2
	g.Y = 4
	r.Println(g, "SUPER PINBALL SYSTEM")
	r.Println(g, spin.Version)
	r.Println(g, spin.Date)
}

func splashScreen(ctx context.Context, e spin.Env) {

	e.Do(spin.StopAudio{})
	splashScreenFrame(e)
	e.Do(spin.PlayMusic{ID: BootTheme})
	if _, done := spin.WaitForEventsUntil(ctx, e, 8*time.Second, []spin.Event{
		spin.SwitchEvent{ID: jd.LeftFlipperButton},
		spin.SwitchEvent{ID: jd.RightFlipperButton},
	}); done {
		return
	}
	e.Post(spin.Message{ID: BootEnd})
}

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{ID: SplashScreen, Script: splashScreen})
}

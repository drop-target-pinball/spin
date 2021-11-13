package boot

import (
	"context"
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

const (
	SplashScreen = "SplashScreen"
)

func splashScreen(ctx context.Context, e spin.Env) {
	r := e.Display("").Renderer()
	defer func() {
		e.Do(spin.StopAudio{})
		r.Clear()
	}()
	g := &spin.Graphics{
		Color:    0x8,
		Font:     PfTempestaFiveCompressedBold8,
		W:        r.Width(),
		Y:        4,
		PaddingV: 2,
	}
	r.Println(g, "SUPER PINBALL SYSTEM")
	r.Println(g, spin.Version)
	r.Println(g, spin.Date)
	e.Do(spin.PlayMusic{ID: BootTheme})

	spin.WaitForEventsUntil(ctx, e, 8*time.Second, []spin.Event{
		spin.SwitchEvent{ID: jd.LeftFlipperButton},
		spin.SwitchEvent{ID: jd.RightFlipperButton},
	})
}

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{ID: SplashScreen, Script: splashScreen})
}

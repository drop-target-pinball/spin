package boot

import (
	"context"

	"github.com/drop-target-pinball/spin"
)

const (
	SplashScreen = "SplashScreen"
)

func splashScreen(ctx context.Context, e spin.Env) {
	r := e.Display("").Renderer()
	g := &spin.Graphics{
		Color:    0xffffffff,
		Font:     PfTempestaFiveCompressedBold8,
		W:        r.Width(),
		Y:        4,
		PaddingV: 2,
	}
	r.Println(g, "SUPER PINBALL SYSTEM")
	r.Println(g, spin.Version)
	r.Println(g, spin.Date)
	e.Do(spin.PlayMusic{ID: BootTheme})
}

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{ID: SplashScreen, Script: splashScreen})
}

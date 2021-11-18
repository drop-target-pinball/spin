package prog

import (
	"context"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/prog/boot"
	"github.com/drop-target-pinball/spin/prog/jdx"
	"github.com/drop-target-pinball/spin/prog/menu"
	"github.com/drop-target-pinball/spin/prog/sandbox"
)

const (
	SpinInit = "SpinInit"
)

func Load(eng *spin.Engine) {
	jd.Load(eng)
	boot.Load(eng)
	menu.Load(eng)
	jdx.Load(eng)
	sandbox.Load(eng)

	eng.Do(spin.RegisterScript{ID: SpinInit, Script: spinInit})
}

func spinInit(ctx context.Context, e spin.Env) {
	e.Do(spin.PlayScript{ID: boot.SplashScreen})
	if _, done := spin.WaitForEvent(ctx, e, spin.Message{ID: boot.BootEnd}); done {
		return
	}

	e.Do(spin.PlayScript{ID: menu.MenuAttractMode})
	if _, done := spin.WaitForEvent(ctx, e, spin.Message{ID: menu.MenuAttractEnd}); done {
		return
	}

	e.Do(spin.PlayScript{ID: menu.MenuSelect})
	if _, done := spin.WaitForEvent(ctx, e, spin.Message{ID: menu.MenuEnd}); done {
		return
	}
}

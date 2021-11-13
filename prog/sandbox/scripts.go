package sandbox

import (
	"context"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/prog/jdx"
)

const (
	SandboxFont = "SandboxFont"
)

func sandboxFont(ctx context.Context, e spin.Env) {
	r := e.Display("").Renderer()
	g := &spin.Graphics{
		Color: 0xffffffff,
		Font:  jdx.Bmsf,
		W:     r.Width(),
	}
	r.Print(g, "HELLO WORLD")
}

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{ID: SandboxFont, Script: sandboxFont})
}

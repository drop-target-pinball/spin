package sandbox

import (
	"context"

	"github.com/drop-target-pinball/spin"
)

const (
	SandboxFont = "SandboxFont"
)

func sandboxFont(ctx context.Context, e spin.Env) {
	r, g := e.Display("").Renderer()
	//defer r.Unlock()
	// g.Font = jdx.FontBm8
	r.Println(g, "0123456789")
	r.Println(g, "123,456,789")
}

func RegisterScripts(eng *spin.Engine) {
	//eng.Do(spin.RegisterScript{ID: SandboxFont, Script: sandboxFont})
}

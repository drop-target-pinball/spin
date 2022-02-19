package service

import (
	"github.com/drop-target-pinball/spin"
)

func gradientPanel(e *spin.ScriptEnv, r spin.Renderer) {
	g := r.Graphics()

	boxW := int32(8)
	g.X, g.Y = 0, 0
	g.W, g.H = boxW, r.Height()

	for i := 0; i <= 0xf; i++ {
		g.Color = spin.DotMatrixColors[i]
		r.FillRect(g)
		g.X += boxW
	}
}

func gradientScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(0)
	defer r.Close()

	spin.RenderFrameLoop(e, func(e *spin.ScriptEnv) {
		gradientPanel(e, r)
	})
}

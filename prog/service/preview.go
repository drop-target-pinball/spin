package service

import (
	"sort"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/prog/builtin"
)

type fontMode struct {
	fonts    []string
	selected int
}

func fontPreviewFrame(e spin.Env, fm *fontMode) {
	r, g := e.Display("").Renderer()

	r.Clear()
	g.Font = fm.fonts[fm.selected]
	r.Print(g, "0123456790")

	g.Font = builtin.FontBmsf
	g.Y = 27
	r.Print(g, fm.fonts[fm.selected])
}

func fontPreviewVideoScript(e spin.Env, fm *fontMode) {
	for {
		fontPreviewFrame(e, fm)
		if done := e.Sleep(spin.FrameDuration); done {
			return
		}
	}
}

func fontPreviewScript(e spin.Env) {
	fm := fontMode{
		fonts:    make([]string, 0),
		selected: 0,
	}
	rsrc := spin.ResourceVars(e)
	for name, _ := range rsrc.Fonts {
		fm.fonts = append(fm.fonts, name)
	}
	sort.Strings(fm.fonts)

	next := func() {
		fm.selected += 1
		if fm.selected >= len(fm.fonts) {
			fm.selected = 0
		}
	}

	prev := func() {
		fm.selected -= 1
		if fm.selected <= 0 {
			fm.selected = len(fm.fonts) - 1
		}
	}

	ctx, cancel := e.Derive()
	e.NewCoroutine(ctx, func(e spin.Env) { fontPreviewVideoScript(e, &fm) })

	for {
		evt, done := e.WaitFor(
			spin.SwitchEvent{ID: spin.SwitchExitServiceButton},
			spin.SwitchEvent{ID: spin.SwitchNextServiceButton},
			spin.SwitchEvent{ID: spin.SwitchPreviousServiceButton},
		)
		if done {
			cancel()
			return
		}
		switch evt {
		case spin.SwitchEvent{ID: spin.SwitchExitServiceButton}:
			cancel()
			return
		case spin.SwitchEvent{ID: spin.SwitchNextServiceButton}:
			next()
		case spin.SwitchEvent{ID: spin.SwitchPreviousServiceButton}:
			prev()
		}
	}
}

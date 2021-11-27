package app

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/prog/builtin"
	"github.com/drop-target-pinball/spin/sdl"
)

type Options struct {
	WithLogging    bool
	WithAudio      bool
	WithVirtualDMD bool
}

func DefaultOptions() Options {
	return Options{
		WithLogging:    true,
		WithAudio:      true,
		WithVirtualDMD: true,
	}
}

func NewEngine(opt Options) *spin.Engine {
	eng := spin.NewEngine()
	if opt.WithLogging {
		spin.RegisterLoggingSystem(eng)
	}
	if opt.WithAudio {
		sdl.RegisterAudioSystem(eng)
	}
	if opt.WithVirtualDMD {
		sdlOpts := sdl.DefaultOptionsDotMatrix()
		sdl.RegisterDotMatrixSystem(eng, sdlOpts)
	}
	sdl.RegisterDisplaySystem(eng, spin.DisplayOptions{Width: 128, Height: 32})
	sdl.RegisterInputSystem(eng)
	builtin.Load(eng)

	return eng
}

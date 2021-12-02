package app

import (
	"github.com/drop-target-pinball/go-pinproc/wpc"
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/proc"
	"github.com/drop-target-pinball/spin/prog/builtin"
	"github.com/drop-target-pinball/spin/sdl"
)

type Options struct {
	WithPROC       bool
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
	eng := spin.NewEngine(jd.Config)
	if opt.WithLogging {
		spin.RegisterLoggingSystem(eng)
	}
	if opt.WithAudio {
		sdl.RegisterAudioSystem(eng)
	}
	if opt.WithVirtualDMD {
		opts := sdl.DefaultOptionsDotMatrix()
		sdl.RegisterDotMatrixSystem(eng, opts)
	}
	if opt.WithPROC {
		opts := proc.Options{
			MachType:                wpc.MachType,
			DMDConfig:               wpc.DMDConfigDefault,
			SwitchConfig:            wpc.SwitchConfigDefault,
			DefaultCoilPulseTime:    25, // milliseconds
			DefaultFlasherPulseTime: 20, // milliseconds
		}
		proc.RegisterSystem(eng, opts)
	}
	sdl.RegisterDisplaySystem(eng, spin.DisplayOptions{Width: 128, Height: 32})
	sdl.RegisterInputSystem(eng)
	builtin.Load(eng)

	return eng
}

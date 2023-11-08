//go:build !sdl

package sdl

import "github.com/drop-target-pinball/spin/v2"

var (
	NewAudioDevice = spin.DeviceNotSupported
)

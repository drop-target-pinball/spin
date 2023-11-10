//go:build !sdl

package sdl

import (
	"fmt"

	"github.com/drop-target-pinball/spin/v2"
)

type AudioFactory struct{}

func (a AudioFactory) ID() string {
	return "sdl_mixer"
}

func (a AudioFactory) NewDevice(conf any) (spin.Device, error) {
	return nil, fmt.Errorf("device not supported, compile with 'sdl' tag")
}

package proc

import "github.com/drop-target-pinball/spin"

const (
	FlasherBlinkSchedule = uint32(0b10000000_00000000_10000000_00000000)
	BlinkSchedule        = 0xf0f0f0f0
	HurryUpBlinkSchedule = 0xcccccccc
)

type DriverSchedule struct {
	spin.Action
	ID       string
	Schedule uint32
	Now      bool
}

func registerActions(e *spin.Engine) {
	e.RegisterAction(DriverSchedule{})
}

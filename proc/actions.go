package proc

import "github.com/drop-target-pinball/spin"

const (
	Blink = 0xf0f0f0f0
)

type DriverSchedule struct {
	spin.Action
	ID       string
	Schedule uint32
}

func registerActions(e *spin.Engine) {
	e.RegisterAction(DriverSchedule{})
}

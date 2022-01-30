package spin

type DriverType int

const (
	Coil DriverType = iota
	Flasher
	Lamp
	Magnet
	Motor
)

type Driver struct {
	ID   string
	Type DriverType
	Addr interface{}
}

type Switch struct {
	ID     string
	NC     bool
	Active bool
}

type BallHold struct {
	Switch string
	Coil   string
}

type Flipper struct {
	ID            string
	SwitchAddr    interface{}
	PowerCoilAddr interface{}
	HoldCoilAddr  interface{}
	_             struct{}
}

type AutoPulse struct {
	ID         string
	SwitchAddr interface{}
	CoilAddr   interface{}
	Time       int // milliseconds
}

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
	ID string
	NC bool
}

type Flipper struct {
	ID            string
	SwitchAddr    interface{}
	PowerCoilAddr interface{}
	HoldCoilAddr  interface{}
	_             struct{}
}
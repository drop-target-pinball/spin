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
	ID      string
	Type    DriverType
	Address interface{}
}

type Switch struct {
	ID string
	NC bool
}

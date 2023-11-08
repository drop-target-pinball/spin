package spin

import "log"

type Device interface {
	Init(*Engine) bool
	Process(*Engine) bool
}

type NewDeviceFunc func(any) (Device, bool)

var (
	devices = map[string]NewDeviceFunc{}
)

func AddNewDeviceFunc(name string, fn NewDeviceFunc) {
	if _, exists := devices[name]; exists {
		log.Panicf("device handler already exists: %v", name)
	}
	devices[name] = fn
}

func NewDevice(conf any) (Device, bool) {
	for _, newFn := range devices {
		dev, ok := newFn(conf)
		if ok {
			return dev, true
		}
	}
	return nil, false
}

func DeviceNotSupported(_ any) (Device, bool) {
	return nil, false
}

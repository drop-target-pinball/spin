package spin

import (
	"fmt"
	"log"
)

type DeviceFactory interface {
	ID() string
	NewDevice(any) (Device, error)
}

type Device interface {
	ID() string
	Name() string
	Init(*Engine) bool
	Process(*Engine) bool
}

type NewDeviceFunc func(any) (Device, bool)

var (
	deviceFactories = map[string]DeviceFactory{}
)

func AddDeviceFactory(f DeviceFactory) {
	if _, exists := deviceFactories[f.ID()]; exists {
		log.Panicf("device factory '%v' already exists", f.ID())
	}
	deviceFactories[f.ID()] = f
}

func NewDevice(id string, conf any) (Device, error) {
	factory, ok := deviceFactories[id]
	if !ok {
		return nil, fmt.Errorf("unknown device type '%v'", id)
	}
	return factory.NewDevice(conf)
}

func DeviceNotSupported(_ any) (Device, bool) {
	return nil, false
}

func DeviceName(factoryId string, deviceID string) string {
	if deviceID == "" {
		return factoryId
	}
	return fmt.Sprintf("%v/%v", factoryId, deviceID)
}

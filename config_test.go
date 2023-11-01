package spin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDevice(t *testing.T) {

	minDevice := `
device "device_id" {
	namespace = "ns"
}`

	tests := []struct {
		name   string
		src    string
		device Device
	}{
		{"min_device.hcl", minDevice, Device{
			ID:        "device_id",
			Namespace: "ns",
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			conf := NewConfig()
			if err := conf.Include(test.name, []byte(test.src)); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			device := conf.Devices["device_id"]
			assert.Equal(t, test.device, device)
		})
	}
}

func TestDriver(t *testing.T) {
	minDriver := `
driver "driver_id" {
	address = "addr"
}`

	maxDriver := `
driver "driver_id" {
	address = "addr"
	type = "solenoid"
}`

	tests := []struct {
		name   string
		src    string
		driver Driver
	}{
		{"min_driver.hcl", minDriver, Driver{
			ID:      "driver_id",
			Address: "addr",
		}},
		{"max_driver.hcl", maxDriver, Driver{
			ID:      "driver_id",
			Address: "addr",
			Type:    "solenoid",
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			conf := NewConfig()
			if err := conf.Include(test.name, []byte(test.src)); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			driver := conf.Drivers["driver_id"]
			assert.Equal(t, test.driver, driver)
		})
	}
}

func TestInfo(t *testing.T) {
	minInfo := `
info "driver" "driver_id" {
}`

	maxInfo := `
info "driver" "driver_id" {
	name = "Test Driver"
	menu_name = "T.D."
	manual_name = "Driver for Testing"
	sort_name = "Driver, Test"
	wires = [ "red", "white", "blue" ]
	jumpers = [ "J1", "J2", "J3" ]
	transistor = "Q1"
}`

	tests := []struct {
		name string
		src  string
		info Info
	}{
		{"min_info.hcl", minInfo, Info{
			Type: "driver",
			ID:   "driver_id",
		}},
		{"max_driver.hcl", maxInfo, Info{
			Type:       "driver",
			ID:         "driver_id",
			Name:       "Test Driver",
			MenuName:   "T.D.",
			ManualName: "Driver for Testing",
			SortName:   "Driver, Test",
			Wires:      []string{"red", "white", "blue"},
			Jumpers:    []string{"J1", "J2", "J3"},
			Transistor: "Q1",
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			conf := NewConfig()
			if err := conf.Include(test.name, []byte(test.src)); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			info := conf.Info["driver/driver_id"]
			assert.Equal(t, test.info, info)
		})
	}
}

func TestSwitch(t *testing.T) {
	minSwitch := `
switch "switch_id" {
	address = "addr"
}`

	maxSwitch := `
switch "switch_id" {
	address = "addr"
	type = "opto"
}`

	tests := []struct {
		name    string
		src     string
		switch_ Switch
	}{
		{"min_switch.hcl", minSwitch, Switch{
			ID:      "switch_id",
			Address: "addr",
		}},
		{"max_switch.hcl", maxSwitch, Switch{
			ID:      "switch_id",
			Address: "addr",
			Type:    "opto",
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			conf := NewConfig()
			if err := conf.Include(test.name, []byte(test.src)); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			switch_ := conf.Switches["switch_id"]
			assert.Equal(t, test.switch_, switch_)
		})
	}
}

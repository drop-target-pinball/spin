package spin

import (
	"os"
	"path"
	"reflect"
	"strings"
	"testing"
)

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
			dir := t.TempDir()
			hcl := path.Join(dir, test.name)
			os.WriteFile(hcl, []byte(test.src), 0o644)

			conf := NewConfig()
			if err := conf.AddFile(hcl); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			driver := conf.Drivers["driver_id"]
			if !reflect.DeepEqual(driver, test.driver) {
				t.Errorf("\n have: %v \n want: %v", driver, test.driver)
			}
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
			dir := t.TempDir()
			hcl := path.Join(dir, test.name)
			os.WriteFile(hcl, []byte(test.src), 0o644)

			conf := NewConfig()
			if err := conf.AddFile(hcl); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			info := conf.Info["driver/driver_id"]
			if !reflect.DeepEqual(info, test.info) {
				t.Errorf("\n have: %v \n want: %v", info, test.info)
			}
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
			dir := t.TempDir()
			hcl := path.Join(dir, test.name)
			os.WriteFile(hcl, []byte(test.src), 0o644)

			conf := NewConfig()
			if err := conf.AddFile(hcl); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			switch_ := conf.Switches["switch_id"]
			if !reflect.DeepEqual(switch_, test.switch_) {
				t.Errorf("\n have: %v \n want: %v", switch_, test.switch_)
			}
		})
	}
}

func TestInclude(t *testing.T) {
	file1 := `
switch "switch_1" {
	address = "bad"
}

switch "switch_2" {
	address	= "sw2"
}
`

	file2 := `
include = [ "./file1.hcl" ]

switch "switch_1" {
	address = "sw1"
}

switch "switch_3" {
	address = "sw3"
}
`
	want := map[string]Switch{
		"switch_1": {
			ID:      "switch_1",
			Address: "sw1",
		},
		"switch_2": {
			ID:      "switch_2",
			Address: "sw2",
		},
		"switch_3": {
			ID:      "switch_3",
			Address: "sw3",
		},
	}

	dir := t.TempDir()
	hcl1 := path.Join(dir, "file1.hcl")
	os.WriteFile(hcl1, []byte(file1), 0o644)
	hcl2 := path.Join(dir, "file2.hcl")
	os.WriteFile(hcl2, []byte(file2), 0o644)

	conf := NewConfig()
	if err := conf.AddFile(hcl2); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(conf.Switches, want) {
		t.Errorf("\n have: %v \n want: %v", conf.Switches, want)
	}
}

func TestMissingInclude(t *testing.T) {
	file := `
include = [ "missing.hcl" ]

switch "switch_1" {
	address = "sw1"
}

switch "switch_3" {
	address = "sw3"
}
`
	dir := t.TempDir()
	hcl := path.Join(dir, "file.hcl")
	os.WriteFile(hcl, []byte(file), 0o644)

	conf := NewConfig()
	err := conf.AddFile(hcl)
	if err == nil {
		t.Fatal("expected error but got nil")
	}
	want := "does not exist."
	if !strings.HasSuffix(err.Error(), want) {
		t.Errorf("\n have: %v \n want: suffix with '%v'", err.Error(), want)
	}
}

func TestSettingsMerge(t *testing.T) {
	file1 := `
settings {
	redis_run_address = "localhost:1234"
}
	`

	file2 := `
settings {
	redis_var_address = "localhost:5678"
}
`

	dir := t.TempDir()
	hcl1 := path.Join(dir, "file1.hcl")
	os.WriteFile(hcl1, []byte(file1), 0o644)
	hcl2 := path.Join(dir, "file2.hcl")
	os.WriteFile(hcl2, []byte(file2), 0o644)

	conf := NewConfig()
	err := conf.AddFile(hcl1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	err = conf.AddFile(hcl2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := &Settings{
		RedisRunAddress: "localhost:1234",
		RedisVarAddress: "localhost:5678",
	}

	if *conf.Settings != *want {
		t.Errorf("\n have: %+v \n want: %+v", conf.Settings, want)
	}
}

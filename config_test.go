package spin

import (
	"path"
	"reflect"
	"strings"
	"testing"
)

func testSettings(t *testing.T) *Settings {
	return &Settings{
		Dir:          t.TempDir(),
		ConfigFile:   "test.hcl",
		RedisAddress: "localhost:1080",
	}
}

func TestDriver(t *testing.T) {
	tests := []struct {
		file   string
		driver Driver
	}{
		{"testdata/config/min_driver.hcl", Driver{
			ID:      "driver_id",
			Address: "addr",
		}},
		{"testdata/config/max_driver.hcl", Driver{
			ID:      "driver_id",
			Address: "addr",
			Type:    "solenoid",
		}},
	}

	for _, test := range tests {
		t.Run(path.Base(test.file), func(t *testing.T) {
			conf := NewConfig()
			if err := conf.AddFile(test.file); err != nil {
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
	tests := []struct {
		file string
		info Info
	}{
		{"testdata/config/min_info.hcl", Info{
			Type: "driver",
			ID:   "driver_id",
		}},
		{"testdata/config/max_info.hcl", Info{
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
		t.Run(path.Base(test.file), func(t *testing.T) {
			conf := NewConfig()
			if err := conf.AddFile(test.file); err != nil {
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
	tests := []struct {
		file    string
		switch_ Switch
	}{
		{"testdata/config/min_switch.hcl", Switch{
			ID:      "switch_id",
			Address: "addr",
		}},
		{"testdata/config/max_switch.hcl", Switch{
			ID:      "switch_id",
			Address: "addr",
			Type:    "opto",
		}},
	}

	for _, test := range tests {
		t.Run(path.Base(test.file), func(t *testing.T) {
			conf := NewConfig()
			if err := conf.AddFile(test.file); err != nil {
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

	conf := NewConfig()
	if err := conf.AddFile("testdata/config/include.hcl"); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(conf.Switches, want) {
		t.Errorf("\n have: %v \n want: %v", conf.Switches, want)
	}
}

func TestMissingInclude(t *testing.T) {
	conf := NewConfig()
	err := conf.AddFile("testdata/config/include_missing.hcl")
	if err == nil {
		t.Fatal("expected error but got nil")
	}
	want := "does not exist."
	if !strings.HasSuffix(err.Error(), want) {
		t.Errorf("\n have: %v \n want: suffix with '%v'", err.Error(), want)
	}
}

func TestSettingsMerge(t *testing.T) {
	conf := NewConfig()
	err := conf.AddFile("testdata/config/settings_1.hcl")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	err = conf.AddFile("testdata/config/settings_2.hcl")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := &Settings{
		RedisAddress: "localhost:1234",
	}

	if *conf.Settings != *want {
		t.Errorf("\n have: %+v \n want: %+v", conf.Settings, want)
	}
}

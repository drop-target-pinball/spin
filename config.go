package spin

import (
	"path"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

// ConfigFile represents the structure of a Spin HCL configuration file. This
// struct shouldn't be used directly--use Config instead.
type ConfigFile struct {
	Include      []string      `hcl:"include,optional"`
	Audio        []Audio       `hcl:"audio,block"`
	AudioDevices []AudioDevice `hcl:"audio_device,block"`
	Defaults     *Defaults     `hcl:"defaults,block"`
	Drivers      []Driver      `hcl:"driver,block"`
	Info         []Info        `hcl:"info,block"`
	Load         []string      `hcl:"load,optional"`
	Settings     *Settings     `hcl:"settings,block"`
	Switches     []Switch      `hcl:"switch,block"`
}

// Config is configuration that has been loaded from HCL configuration files.
type Config struct {
	Audio        map[string]Audio       `json:"audio,omitempty"`
	AudioDevices map[string]AudioDevice `json:"audio_devices,omitempty"`
	Drivers      map[string]Driver      `json:"drivers,omitempty"`
	Info         map[string]Info        `json:"info,omitempty"`
	Load         []string               `json:"load,omitempty"`
	Settings     *Settings              `json:"settings,omitempty"`
	Switches     map[string]Switch      `json:"switches,omitempty"`
}

type Defaults struct {
	Module string `hcl:"module,optional" json:"module,omitempty"`
}

type Settings struct {
	DevMode         bool
	Dir             string `json:"dir,omitempty"`
	ConfigFile      string `json:"config_file,omitempty"`
	RedisRunAddress string `hcl:"redis_run_address,optional" json:"redis_run_address,omitempty"`
	RedisVarAddress string `hcl:"redis_var_address,optional" json:"redis_var_address,omitempty"`
}

// Merge copies values from s2 into this struct where values in this struct
// contain a zero value. If s2 is nil, this method does nothing.
func (s *Settings) Merge(s2 *Settings) {
	if s2 == nil {
		return
	}
	if s.RedisRunAddress == "" {
		s.RedisRunAddress = s2.RedisRunAddress
	}
	if s.RedisVarAddress == "" {
		s.RedisVarAddress = s2.RedisVarAddress
	}
}

// NewConfig creates an empty configuration.
func NewConfig() *Config {
	return &Config{
		Audio:        make(map[string]Audio),
		AudioDevices: make(map[string]AudioDevice),
		Drivers:      make(map[string]Driver),
		Info:         make(map[string]Info),
		Settings:     &Settings{},
		Switches:     make(map[string]Switch),
	}
}

// AddFile loads the HCL configuration from the project path with the given
// filename and adds it to the current configuration. Configuration entities in
// the included file overwrite existing entries with the same key.
func (c *Config) AddFile(name string) error {
	var cf ConfigFile

	if err := hclsimple.DecodeFile(name, nil, &cf); err != nil {
		return err
	}
	for _, inc := range cf.Include {
		fullInc := resolveFrom(name, inc)
		if err := c.AddFile(fullInc); err != nil {
			return err
		}
	}

	if cf.Defaults != nil {
		for i, a := range cf.Audio {
			if a.Module == "" {
				a.Module = cf.Defaults.Module
			}
			cf.Audio[i] = a
		}
	}

	key[Audio](cf.Audio, c.Audio, func(a Audio) string { return a.ID })
	key[AudioDevice](cf.AudioDevices, c.AudioDevices, func(d AudioDevice) string { return d.ID })
	key[Driver](cf.Drivers, c.Drivers, func(d Driver) string { return d.ID })
	key[Info](cf.Info, c.Info, func(i Info) string { return i.Type + "/" + i.ID })
	key[Switch](cf.Switches, c.Switches, func(s Switch) string { return s.ID })

	c.Load = append(c.Load, cf.Load...)
	c.Settings.Merge(cf.Settings)

	return nil
}

// If name isn't an absolute path, return a path that is relative to the
// directory where from resides.
func resolveFrom(from string, name string) string {
	fullName := name
	if !strings.HasPrefix(name, "/") {
		fullName = path.Join(path.Dir(from), name)
	}
	return fullName
}

// add all items from the source slice to the target map by using the
// provided key function.
func key[T any](source []T, target map[string]T, keyfn func(T) string) {
	for _, s := range source {
		target[keyfn(s)] = s
	}
}

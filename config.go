package spin

import (
	"os"
	"path"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

var (
	ProjectDir string
)

func init() {
	ProjectDir = os.Getenv("SPIN_DIR")
	if ProjectDir == "" {
		ProjectDir = "./project"
	}
}

type ConfigFile struct {
	Include  []string `hcl:"include,optional"`
	Devices  []Device `hcl:"device,block"`
	Drivers  []Driver `hcl:"driver,block"`
	Info     []Info   `hcl:"info,block"`
	Switches []Switch `hcl:"switch,block"`
}

type Config struct {
	Devices  map[string]Device
	Drivers  map[string]Driver
	Info     map[string]Info
	Switches map[string]Switch
}

func NewConfig() *Config {
	return &Config{
		Devices:  make(map[string]Device),
		Drivers:  make(map[string]Driver),
		Info:     make(map[string]Info),
		Switches: make(map[string]Switch),
	}
}

/*
type Profile struct {
	ID           string `hcl:"id,label"`
	RedisRunPort int    `hcl:"redis_run_port,optional"`
	RedisVarPort int    `hcl:"redis_var_port,optional"`
}
*/

func (c *Config) AddFile(name string) error {
	var cf ConfigFile
	if err := hclsimple.DecodeFile(name, nil, &cf); err != nil {
		return err
	}
	for _, inc := range cf.Include {
		if err := c.AddFile(path.Join(path.Dir(name), inc)); err != nil {
			return err
		}
	}

	key[Device](cf.Devices, c.Devices, func(d Device) string { return d.ID })
	key[Driver](cf.Drivers, c.Drivers, func(d Driver) string { return d.ID })
	key[Info](cf.Info, c.Info, func(i Info) string { return i.Type + "/" + i.ID })
	key[Switch](cf.Switches, c.Switches, func(s Switch) string { return s.ID })

	return nil
}

func key[T any](source []T, target map[string]T, keyfn func(T) string) {
	for _, s := range source {
		target[keyfn(s)] = s
	}
}

func LoadConfig() (*Config, error) {
	c := NewConfig()
	if err := c.AddFile(path.Join(ProjectDir, "spin.hcl")); err != nil {
		return nil, err
	}
	return c, nil
}

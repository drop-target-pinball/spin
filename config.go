package spin

import (
	"io/fs"
	"os"
	"path"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

var (
	// ProjectDir is the directory containing all configuration and asset
	// files for a given pinball project. This can be set by using the
	// SPIN_DIR environment variable, otherwise this defaults to ./project.
	// When Spin starts, the entry point is found by checking the
	// project.hcl file in this directory.
	ProjectDir string

	ProjectFS fs.FS
)

func init() {
	ProjectDir = os.Getenv("SPIN_DIR")
	if ProjectDir == "" {
		ProjectDir = "./project"
	}
	ProjectFS = os.DirFS(ProjectDir)
}

// ConfigFile represents the structure of a Spin HCL configuration file. This
// struct shouldn't be used directly--use Config instead.
type ConfigFile struct {
	Include  []string `hcl:"include,optional"`
	Devices  []Device `hcl:"device,block"`
	Drivers  []Driver `hcl:"driver,block"`
	Info     []Info   `hcl:"info,block"`
	Switches []Switch `hcl:"switch,block"`
}

// Config is configuration that has been loaded from HCL configuration files.
type Config struct {
	Devices    map[string]Device `json:"devices,omitempty"`
	Drivers    map[string]Driver `json:"drivers,omitempty"`
	Info       map[string]Info   `json:"info,omitempty"`
	Switches   map[string]Switch `json:"switches,omitempty"`
	FileSystem fs.FS
}

// NewConfig creates an empty configuration.
func NewConfig() *Config {
	return &Config{
		Devices:    make(map[string]Device),
		Drivers:    make(map[string]Driver),
		Info:       make(map[string]Info),
		Switches:   make(map[string]Switch),
		FileSystem: os.DirFS("/"),
	}
}

// Include loads the HCL configuration a file from within src and
// adds it to the current configuration. Configuration entities in the included
// file overwrite existing entries with the same key.
func (c *Config) Include(name string, src []byte) error {
	var cf ConfigFile
	if err := hclsimple.Decode(name, src, nil, &cf); err != nil {
		return err
	}
	for _, inc := range cf.Include {
		if err := c.IncludeFile(path.Join(path.Dir(name), inc)); err != nil {
			return err
		}
	}

	key[Device](cf.Devices, c.Devices, func(d Device) string { return d.ID })
	key[Driver](cf.Drivers, c.Drivers, func(d Driver) string { return d.ID })
	key[Info](cf.Info, c.Info, func(i Info) string { return i.Type + "/" + i.ID })
	key[Switch](cf.Switches, c.Switches, func(s Switch) string { return s.ID })

	return nil
}

// add all items from the source slice to the target map by using the
// provided key function.
func key[T any](source []T, target map[string]T, keyfn func(T) string) {
	for _, s := range source {
		target[keyfn(s)] = s
	}
}

// IncludeFile loads the HCL configuration a file with the given filename and
// adds it to the current configuration. Configuration entities in the included
// file overwrite existing entries with the same key.
func (c *Config) IncludeFile(name string) error {
	src, err := fs.ReadFile(c.FileSystem, name)
	if err != nil {
		return err
	}
	return c.Include(name, src)
}

// LoadConfig reads the spin.hcl configuration file found in the
// ProjectDir and returns the parsed configuration structure.
func LoadConfig() (*Config, error) {
	c := NewConfig()
	if err := c.IncludeFile(path.Join(ProjectDir, "spin.hcl")); err != nil {
		return nil, err
	}
	return c, nil
}

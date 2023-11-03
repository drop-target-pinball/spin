package spin

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

// ConfigFile represents the structure of a Spin HCL configuration file. This
// struct shouldn't be used directly--use Config instead.
type ConfigFile struct {
	Include  []string  `hcl:"include,optional"`
	Devices  []Device  `hcl:"device,block"`
	Drivers  []Driver  `hcl:"driver,block"`
	Info     []Info    `hcl:"info,block"`
	Settings *Settings `hcl:"settings,block"`
	Switches []Switch  `hcl:"switch,block"`
}

// Project represents directories where configurations and assets are located.
// Functions are provided for finding and reading files from these directories.
type Project struct {
	Path []string // Search path to use when finding files.
}

// Creates a new project. If the environment variable SPINDIR is not blank, the
// initial project path is populated with colon-separated directories found in
// that variable. If no SPINDIR is set, the path is initialized to check
// "./project" first, and then to check "./lib". To use a custom path,
// create the object and then modify the value of Path directly.
func NewProject() *Project {
	proj := &Project{}

	envPath := os.Getenv("SPINDIR")
	if envPath == "" {
		proj.Path = []string{"./project", "./lib"}
	} else {
		proj.Path = strings.Split(envPath, ":")
	}
	return proj
}

// Creates a new project with a path containing a single directory.
func NewProjectWithDir(dir string) *Project {
	proj := NewProject()
	proj.Path = []string{dir}
	return proj
}

// FindFileFrom locates a file with the given name in the project path
// or relative to the source. If the name starts with "./", the file is
// searched relative to the source's directory. Otherwise, the path is searched
// in order to find the file with the given name.
//
// If the file is found, the full path to the file is returned. Otherwise
// an error is returned.
func (p *Project) FindFileFrom(source string, name string) (string, error) {
	if strings.HasPrefix(name, "./") {
		fullName := path.Join(path.Dir(source), name)
		fmt.Printf("**** SOURCE: %v\n", source)
		fmt.Printf("**** FULLNAME: %v\n", fullName)
		_, err := os.Stat(fullName)
		return fullName, err
	}
	for _, dir := range p.Path {
		fullName := path.Join(dir, name)
		_, err := os.Stat(fullName)
		if !errors.Is(err, os.ErrNotExist) {
			return fullName, err
		}
	}
	return "", fmt.Errorf("%v: file does not exist", name)
}

// FindFile locates a file with the given name in the project path. If the file
// is found, the full path to the file is returned. Otherwise an error is
// returned.
func (p *Project) FindFile(name string) (string, error) {
	return p.FindFileFrom("", name)
}

// ReadFileFrom reads in a file with the given name in the project path
// or relative to the source. This function is same as calling FindFileFrom
// and then using os.ReadFile.
func (p *Project) ReadFileFrom(src string, name string) ([]byte, error) {
	fullName, err := p.FindFileFrom(src, name)
	if err != nil {
		return nil, err
	}
	return os.ReadFile(fullName)
}

// ReadFile reads in a file with the given name in the project path.
// This function is same as calling FindFile and then using os.ReadFile.
func (p *Project) ReadFile(name string) ([]byte, error) {
	return p.ReadFileFrom("", name)
}

// Config is configuration that has been loaded from HCL configuration files.
type Config struct {
	Devices  map[string]Device `json:"devices,omitempty"`
	Drivers  map[string]Driver `json:"drivers,omitempty"`
	Info     map[string]Info   `json:"info,omitempty"`
	Settings Settings          `json:"settings,omitempty"`
	Switches map[string]Switch `json:"switches,omitempty"`
	project  *Project
}

type Settings struct {
	RedisRunAddress string `hcl:"redis_run_address,optional" json:"redis_run_address,omitempty"`
	RedisVarAddress string `hcl:"redis_var_address,optional" json:"redis_var_address,omitempty"`
}

// Merge copies any non-zero values from s2 into this struct. If s2 is nil,
// this method does nothing.
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
func NewConfig(project *Project) *Config {
	return &Config{
		Devices:  make(map[string]Device),
		Drivers:  make(map[string]Driver),
		Info:     make(map[string]Info),
		Switches: make(map[string]Switch),
		project:  project,
	}
}

// AddFile loads the HCL configuration from the project path with the given
// filename and adds it to the current configuration. Configuration entities in
// the included file overwrite existing entries with the same key.
func (c *Config) AddFile(name string) error {
	var cf ConfigFile

	fullName, err := c.project.FindFile(name)
	if err != nil {
		return err
	}
	if err := hclsimple.DecodeFile(fullName, nil, &cf); err != nil {
		return err
	}
	for _, inc := range cf.Include {
		fullInc, err := c.project.FindFileFrom(fullName, inc)
		if err != nil {
			return err
		}
		if err := c.AddFile(fullInc); err != nil {
			return err
		}
	}

	key[Device](cf.Devices, c.Devices, func(d Device) string { return d.ID })
	key[Driver](cf.Drivers, c.Drivers, func(d Driver) string { return d.ID })
	key[Info](cf.Info, c.Info, func(i Info) string { return i.Type + "/" + i.ID })
	key[Switch](cf.Switches, c.Switches, func(s Switch) string { return s.ID })
	c.Settings.Merge(cf.Settings)

	return nil
}

// add all items from the source slice to the target map by using the
// provided key function.
func key[T any](source []T, target map[string]T, keyfn func(T) string) {
	for _, s := range source {
		target[keyfn(s)] = s
	}
}

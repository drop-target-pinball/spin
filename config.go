package spin

import (
	"os"
	"path"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
)

type Config struct {
	Modules map[string]*Module `json:"modules"`
}

func NewConfig() *Config {
	return &Config{Modules: make(map[string]*Module)}
}

type System struct {
}

type Module struct {
	Dir   string  `json:"dir"`
	Audio []Audio `hcl:"audio,block" json:"audio"`
}

type Audio struct {
	ID   string `hcl:"id,label" json:"id"`
	Chan string `hcl:"chan,optional" json:"chan"`
	File string `hcl:"file" json:"file"`
}

func LoadHCL(dir string) (hcl.Body, hcl.Diagnostics) {
	var diags hcl.Diagnostics

	p := hclparse.NewParser()
	body := hcl.EmptyBody()
	files, err := os.ReadDir(dir)
	if err != nil {
		diags.Append(&hcl.Diagnostic{Severity: hcl.DiagError, Summary: err.Error()})
		return body, diags
	}
	for _, file := range files {
		switch {
		case file.IsDir():
			body2, diags2 := LoadHCL(path.Join(dir, file.Name()))
			body = hcl.MergeBodies([]hcl.Body{body, body2})
			diags.Extend(diags2)
		case strings.HasSuffix(file.Name(), "hcl"):
			name := path.Join(dir, file.Name())
			f, diags2 := p.ParseHCLFile(name)
			if !diags2.HasErrors() {
				body = hcl.MergeBodies([]hcl.Body{body, f.Body})
			}
			diags.Extend(diags2)
		}
	}
	return body, diags
}

func LoadConfig(config *Config, dir string) hcl.Diagnostics {
	config.Modules = make(map[string]*Module)
	body, diags := LoadHCL(dir)
	diags2 := gohcl.DecodeBody(body, nil, config)
	return diags.Extend(diags2)
}

func LoadModule(config *Config, dir string) hcl.Diagnostics {
	name := path.Base(dir)
	if _, ok := config.Modules[name]; ok {
		return hcl.Diagnostics{}
	}
	var module Module
	body, diags := LoadHCL(dir)
	diags2 := gohcl.DecodeBody(body, nil, &module)
	module.Dir = dir
	config.Modules[name] = &module
	return diags.Extend(diags2)
}

/*
type Base struct {
	A string `hcl:"a,optional"`
	B string `hcl:"b,optional"`
}

type ItemC struct {
	Base `hcl:",remain"`
	C    string `hcl:"c,optional"`
}

type ItemD struct {
	Base `hcl:",remain"`
	D    string `hcl:"d,optional"`
}
*/

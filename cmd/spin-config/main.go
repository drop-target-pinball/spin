package main

import (
	"fmt"
	"os"

	"github.com/drop-target-pinball/spin/v2"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
)

var data1 = []byte(`
item_c {
	a = "foo1"
	c = "baz1"
}
`)

var data2 = []byte(`
item_d {
	a = "foo2"
	d = "baz2"
}

item_c {
	a = "other2"
	c = "other2"
}
`)

func printDiag(diag hcl.Diagnostics) {
	for _, err := range diag.Errs() {
		fmt.Fprintln(os.Stderr, err)
	}
}

func main() {
	var c spin.ConfigFile

	p := hclparse.NewParser()
	f1, diag := p.ParseHCL(data1, "data1.hcl")
	if diag != nil {
		printDiag(diag)
		return
	}

	f2, diag := p.ParseHCL(data2, "data2.hcl")
	if diag != nil {
		printDiag(diag)
		return
	}

	m := hcl.MergeBodies([]hcl.Body{f1.Body, f2.Body})
	diag = gohcl.DecodeBody(m, nil, &c)
	if diag != nil {
		printDiag(diag)
		return
	}

	fmt.Printf("%+v\n", c)
	//fmt.Println(c.ItemC[0].A)
	//fmt.Println(c.ItemC[0].B)
	//fmt.Println(c.ItemC[0].C)

	// fmt.Println(c.ItemC[1].A)
	// fmt.Println(c.ItemC[1].B)
	// fmt.Println(c.ItemC[1].C)
}

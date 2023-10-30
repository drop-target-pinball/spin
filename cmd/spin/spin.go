package main

import (
	"fmt"
	"os"

	"github.com/drop-target-pinball/spin/v2"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

func main() {
	var config spin.Config
	err := hclsimple.DecodeFile(os.Args[1], nil, &config)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("%+v\n", config)
}

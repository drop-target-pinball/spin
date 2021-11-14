package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

type metrics struct {
	Top          int                 `json:"top"`
	MarginTop    int                 `json:"margin_top"`
	Bottom       int                 `json:"bottom"`
	Left         int                 `json:"left"`
	LeftOverride map[string]int      `json:"left_override"`
	Widths       [][]interface{}     `json:"widths"`
	Chars        map[string][]string `json:"chars"`
}

type tile struct {
	X int
	Y int
	W int
	H int
}

func main() {
	log.SetFlags(0)
	flag.Parse()
	if flag.NArg() != 3 {
		log.Fatalf("Usage: <source_json> <size> <target_json>")
	}

	source := flag.Arg(0)
	size, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatalf("invalid size: %v", size)
	}
	target := flag.Arg(2)

	data, err := ioutil.ReadFile(source)
	if err != nil {
		log.Fatalf("unable to read source file: %v", err)
	}

	var met metrics
	if err := json.Unmarshal(data, &met); err != nil {
		if err != nil {
			log.Fatalf("unable to parse json file: %v", err)
		}
	}

	tiles := make(map[string]tile)

	for i, widths := range met.Widths {
		var t tile
		ch := widths[0].(string)
		top := met.Top
		left, ok := met.LeftOverride[ch]
		if !ok {
			left = met.Left
		} else {
			fmt.Printf("FOUND OVERRIDE: %v\n", ch)
			top += 1
		}

		w := widths[1].(float64)
		t.X = (i % 10 * size) + left
		t.Y = (i / 10 * size) + top
		t.W = int(w)
		t.H = size
		tiles[ch] = t
	}

	out, err := json.Marshal(&tiles)
	if err != nil {
		log.Fatalf("unable to create JSON file: %v", err)
	}
	if err := ioutil.WriteFile(target, out, 0o644); err != nil {
		log.Fatalf("unable to write JSON file: %v", err)
	}
}

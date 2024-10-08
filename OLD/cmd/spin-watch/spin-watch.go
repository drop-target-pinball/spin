package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"reflect"

	"github.com/drop-target-pinball/spin/v2"
)

var (
	addr   string
	pretty bool
)

func main() {
	log.SetFlags(0)
	flag.StringVar(&addr, "addr", "localhost:1080", "redis address to run database")
	flag.BoolVar(&pretty, "pretty", false, "pretty print messages")
	flag.Parse()

	pin, err := spin.NewClient(addr)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}

	for {
		msg, err := pin.Read()
		if err != nil {
			log.Fatal(err)
		}
		var payload []byte
		if pretty {
			fmt.Println()
			if payload, err = json.MarshalIndent(msg, "", "  "); err != nil {
				log.Fatal(err)
			}
		} else {
			if payload, err = json.Marshal(msg); err != nil {
				log.Fatal(err)
			}
		}
		fmt.Printf("%v: %v\n", reflect.TypeOf(msg).Name(), string(payload))
	}
}

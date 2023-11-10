package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"reflect"

	"github.com/drop-target-pinball/spin/v2"
	"github.com/redis/go-redis/v9"
)

var (
	addr   string
	pretty bool
)

func abort(err error) {
	fmt.Printf("error: %v\n", err)
	os.Exit(1)
}

func main() {
	flag.StringVar(&addr, "addr", "localhost:1080", "redis address to run database")
	flag.BoolVar(&pretty, "pretty", false, "pretty print messages")
	flag.Parse()

	db := redis.NewClient(&redis.Options{Addr: addr})
	cli := spin.NewStreamClient(db)

	for {
		msg, err := cli.Read()
		if err != nil {
			abort(err)
		}
		var payload []byte
		if pretty {
			fmt.Println()
			if payload, err = json.MarshalIndent(msg, "", "  "); err != nil {
				abort(err)
			}
		} else {
			if payload, err = json.Marshal(msg); err != nil {
				abort(err)
			}
		}
		fmt.Printf("%v: %v\n", reflect.TypeOf(msg).Name(), string(payload))
	}
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/drop-target-pinball/spin/v2"
)

var (
	RedisAddr string
	Channels  string
	Version   bool
)

func main() {
	log.SetFlags(0)

	flag.Usage = func() {
		fmt.Printf("Usage: %v [options]\n", os.Args[0])
		fmt.Print(`
Reads messages from the queue and prints out their contents to standard
output in JSON format. 	

`)
		flag.PrintDefaults()
	}

	flag.StringVar(&Channels, "c", "*", "subscribe to channels")
	flag.StringVar(&RedisAddr, "r", spin.DefaultRedisAddr, "redis address")
	flag.BoolVar(&Version, "v", false, "version")
	flag.Parse()

	if Version {
		fmt.Println(spin.Banner())
		return
	}

	c := spin.NewClient(RedisAddr)
	chans := strings.Split(Channels, ":")
	c.Subscribe(chans...)

	for {
		msg, err := c.Receive()
		if err != nil {
			log.Fatalf("unable to receive message: %v", err)
		}
		fmt.Printf("- %v\n", spin.FormatBody(msg.Body))
	}
}

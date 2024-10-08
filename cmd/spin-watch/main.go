package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/drop-target-pinball/spin/v2"
	"github.com/redis/go-redis/v9"
)

var (
	RedisAddr string
	Channels  string
)

func main() {
	log.SetFlags(0)
	flag.StringVar(&Channels, "c", "*", "subscribe to channels")
	flag.StringVar(&RedisAddr, "r", spin.DefaultRedisAddr, "redis address")
	flag.Parse()

	c := spin.NewClient(redis.Options{Addr: RedisAddr})
	chans := strings.Split(Channels, ":")
	c.Subscribe(chans...)

	for {
		msg, err := c.Receive()
		if err != nil {
			log.Fatalf("unable to receive message: %v", err)
		}
		fmt.Printf("[%v] %v\n", msg.Header.To, msg.Body)
	}
}

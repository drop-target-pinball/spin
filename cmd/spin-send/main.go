package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io"
	"log"
	"os"

	"github.com/drop-target-pinball/spin/v2"
	"github.com/redis/go-redis/v9"
)

var (
	RedisAddr string
)

func main() {
	log.SetFlags(0)
	flag.StringVar(&RedisAddr, "r", spin.DefaultRedisAddr, "redis address")
	flag.Parse()

	c := spin.NewClient(redis.Options{Addr: RedisAddr})
	dec := json.NewDecoder(os.Stdin)

	for {
		msg, err := spin.ParseMessage(dec)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatalf("unable to parse message: %v", err)
		}
		if err := c.Send(msg); err != nil {
			log.Fatalf("unable to send message: %v", err)
		}
	}
}

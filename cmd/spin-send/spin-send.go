package main

import (
	"flag"
	"io"
	"log"
	"os"
	"strings"

	"github.com/drop-target-pinball/spin/v2"
	"github.com/redis/go-redis/v9"
)

var (
	addr string
)

func main() {
	log.SetFlags(0)
	flag.StringVar(&addr, "addr", "localhost:1080", "redis address to runtime database")
	flag.Parse()

	cli := redis.NewClient(&redis.Options{Addr: addr})
	stream := spin.NewStreamClient(cli)

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	text := string(data)
	typ, payload, found := strings.Cut(text, " ")
	if !found {
		log.Fatal("invalid message, missing type")
	}
	msg, err := spin.ParseMessage(typ, []byte(payload))
	if err != nil {
		log.Fatalf("unable to parse message: %v", err)
	}
	if err := stream.Send(msg); err != nil {
		log.Fatalf("unable to send message: %v", err)
	}
}

package main

import (
	"flag"
	"io"
	"log"
	"os"
	"strings"

	"github.com/drop-target-pinball/spin/v2"
)

var (
	addr string
)

func main() {
	log.SetFlags(0)
	flag.StringVar(&addr, "addr", "localhost:1080", "redis address to runtime database")
	flag.Parse()

	pin, err := spin.NewClient(addr)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}

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
	if err := pin.Send(msg); err != nil {
		log.Fatalf("unable to send message: %v", err)
	}
}
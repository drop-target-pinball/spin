package main

import (
	"flag"
	"log"

	"github.com/drop-target-pinball/spin/v2"
	"github.com/drop-target-pinball/spin/v2/app/cli"
	"github.com/redis/go-redis/v9"
)

var (
	addr string
)

func main() {
	log.SetFlags(0)
	flag.StringVar(&addr, "addr", "localhost:1080", "redis address to runtime database")
	flag.Parse()

	db := redis.NewClient(&redis.Options{Addr: addr})
	stream := spin.NewStreamClient(db)

	app := cli.NewApp(stream)
	app.Run()
}

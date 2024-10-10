package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/drop-target-pinball/spin/v2"
	"github.com/redis/go-redis/v9"
)

var (
	RedisAddr string
	Version   bool
	Help      bool
)

func main() {
	log.SetFlags(0)

	flag.Usage = func() {
		fmt.Printf("Usage: %v [options] [file...]\n", os.Args[0])
		fmt.Print(`
Sends the JSON contents found in 'file' to the message queue. If no
files are specified or if a file is '-', JSON is read from standard input. 	

`)
		flag.PrintDefaults()
	}

	flag.StringVar(&RedisAddr, "r", spin.DefaultRedisAddr, "redis address")
	flag.BoolVar(&Version, "v", false, "version")
	flag.Parse()

	if Version {
		fmt.Println(spin.Banner())
		return
	}

	c := spin.NewClient(redis.Options{Addr: RedisAddr})
	args := flag.Args()
	if len(args) == 0 {
		args = []string{"-"}
	}
	for _, arg := range args {
		processFile(c, arg)
	}
}

func processFile(c *spin.Client, name string) {
	var dec *json.Decoder
	if name == "-" {
		dec = json.NewDecoder(os.Stdin)
	} else {
		f, err := os.Open(name)
		if err != nil {
			log.Fatalf("unable to open file: %v", err)
		}
		defer f.Close()
		dec = json.NewDecoder(f)
	}

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

package spin

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	DefaultRedisAddr = ":6379"
)

type Client struct {
	ctx context.Context
	rdb *redis.Client
	sub *redis.PubSub
	ch  <-chan *redis.Message
}

func NewClient(addr string) *Client {
	opts := redis.Options{
		Addr:     addr,
		Protocol: 3,
	}
	c := &Client{
		ctx: context.Background(),
		rdb: redis.NewClient(&opts),
	}
	return c
}

func (c *Client) Send(msg Packet) error {
	header, err := json.Marshal(msg.Header)
	if err != nil {
		return fmt.Errorf("invalid header: %v", err)
	}
	body, err := json.Marshal(msg.Body)
	if err != nil {
		return fmt.Errorf("invalid body: %v", err)
	}

	data := append(header, 0)
	data = append(data, body...)

	if err := c.rdb.Publish(c.ctx, msg.Header.Chan, data).Err(); err != nil {
		return err
	}
	return nil
}

func (c *Client) Subscribe(chans ...string) {
	if c.sub != nil {
		c.sub.PUnsubscribe(c.ctx, "")
		c.sub.Close()
	}
	c.sub = c.rdb.PSubscribe(c.ctx, chans...)
	c.ch = c.sub.Channel()
}

func (c *Client) Receive() (Packet, error) {
	return c.ReceiveWithTimeout(nil)
}

var ErrTimeout = errors.New("timeout")

func (c *Client) ReceiveWithTimeout(timeout <-chan time.Time) (Packet, error) {
	if c.sub == nil {
		return Packet{}, fmt.Errorf("not subscribed to any channels")
	}

	var rm *redis.Message
	select {
	case rm = <-c.ch:
		return c.decodeMessage(rm)
	case <-timeout:
		return Packet{}, ErrTimeout
	}
}

func (c *Client) decodeMessage(rm *redis.Message) (Packet, error) {
	var m Packet

	data := []byte(rm.Payload)
	parts := bytes.Split(data, []byte{0})
	if len(parts) != 2 {
		return m, fmt.Errorf("invalid separation")
	}
	header, body := parts[0], parts[1]

	if err := ParseHeader(header, &m); err != nil {
		return m, fmt.Errorf("parse error: %v", err)
	}
	if err := ParseBody(body, &m); err != nil {
		return m, fmt.Errorf("parse error: %v", err)
	}
	return m, nil
}

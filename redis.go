package spin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

const (
	DefaultRedisAddr = ":6379"
)

type Client struct {
	ctx context.Context
	rdb *redis.Client
	sub *redis.PubSub
}

func NewClient(opts redis.Options) *Client {
	opts.Protocol = 3
	c := &Client{
		ctx: context.Background(),
		rdb: redis.NewClient(&opts),
	}
	return c
}

func (c *Client) Send(msg Message) error {
	header, err := json.Marshal(msg.Header)
	if err != nil {
		return fmt.Errorf("invalid header: %v", err)
	}
	body, err := json.Marshal(msg.Body)
	if err != nil {
		return fmt.Errorf("invalid body: %v", err)
	}
	if err := c.rdb.Publish(c.ctx, msg.Header.Chan, header).Err(); err != nil {
		return err
	}
	if err := c.rdb.Publish(c.ctx, msg.Header.Chan, body).Err(); err != nil {
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
}

func (c *Client) Receive() (Message, error) {
	var msg Message
	if c.sub == nil {
		return msg, fmt.Errorf("not subscribed to any channels")
	}
	header, err := c.sub.ReceiveMessage(c.ctx)
	if err != nil {
		return msg, err
	}
	if err := ParseHeader([]byte(header.Payload), &msg); err != nil {
		return msg, fmt.Errorf("parse error: %v", err)
	}

	body, err := c.sub.ReceiveMessage(c.ctx)
	if err != nil {
		return msg, err
	}
	if err := ParseBody([]byte(body.Payload), &msg); err != nil {
		return msg, fmt.Errorf("parse error: %v", err)
	}
	return msg, nil
}

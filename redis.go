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
		return fmt.Errorf("unable to encode message header: %v", err)
	}
	body, err := json.Marshal(msg.Body)
	if err != nil {
		return fmt.Errorf("unable to encode message body: %v", err)
	}
	if err := c.rdb.Publish(c.ctx, msg.Header.To, header).Err(); err != nil {
		return fmt.Errorf("unable to send message header: %v", err)
	}
	if err := c.rdb.Publish(c.ctx, msg.Header.To, body).Err(); err != nil {
		return fmt.Errorf("unable to send message body: %v", err)
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
		return msg, fmt.Errorf("unable to receive message header: %v", err)
	}
	if err := ParseHeader([]byte(header.Payload), &msg); err != nil {
		return msg, fmt.Errorf("unable to parse message header: %v", err)
	}

	body, err := c.sub.ReceiveMessage(c.ctx)
	if err != nil {
		return msg, fmt.Errorf("unable to receive message body: %v", err)
	}
	if err := ParseBody([]byte(body.Payload), &msg); err != nil {
		return msg, fmt.Errorf("unable to parse message body: %v", err)
	}
	return msg, nil
}

package spin

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	Addr   string
	db     *redis.Client
	buf    []any
	lastID string
}

func NewClient(addr string) (*Client, error) {
	c := &Client{
		Addr: addr,
	}

	ctx := context.Background()

	c.db = redis.NewClient(&redis.Options{Addr: addr})
	resp := c.db.Conn().Hello(ctx, 3, "", "", "client-name-test")
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	return c, nil
}

func (c *Client) Reset() {
	c.lastID = "0-0"
}

func (c *Client) Read() (any, error) {
	if len(c.buf) > 0 {
		return c.pop(), nil
	}
	if err := c.fillBuf(c.lastID); err != nil {
		return nil, err
	}
	return c.pop(), nil
}

func (c *Client) Send(messages ...any) error {
	for _, message := range messages {
		payload, err := json.Marshal(message)
		if err != nil {
			return err
		}
		ctx := context.Background()
		typ := reflect.TypeOf(message)
		name := typ.Name()
		if name == "" && typ.Kind() == reflect.Ptr {
			name = typ.Elem().Name()
		}

		result := c.db.XAdd(ctx, &redis.XAddArgs{
			Stream: "mq",
			Values: []any{
				"type", name,
				"payload", payload,
			},
		})
		if result.Err() != nil {
			return result.Err()
		}
	}
	return nil
}

func (c *Client) pop() any {
	var msg any
	msg, c.buf = c.buf[0], c.buf[1:]
	return msg
}

func (c *Client) fillBuf(fromID string) error {
	ctx := context.Background()

	if c.lastID != "" {
		fromID = c.lastID
	}
	res := c.db.XRead(ctx, &redis.XReadArgs{
		Streams: []string{"mq", fromID},
	})
	if res.Err() != nil {
		return res.Err()
	}
	streams, err := res.Result()
	if err != nil {
		return err
	}
	for _, stream := range streams {
		for _, msg := range stream.Messages {
			anyType, ok := msg.Values["type"]
			if !ok {
				return fmt.Errorf("message did not contain a 'type' field")
			}
			typ, ok := anyType.(string)
			if !ok {
				return fmt.Errorf("message type field is not a string")
			}
			anyPayload, ok := msg.Values["payload"]
			if !ok {
				return fmt.Errorf("message did not contain a 'payload' field")
			}
			payload, ok := anyPayload.(string)
			if !ok {
				return fmt.Errorf("message payload field is not a string")
			}
			val, err := ParseMessage(typ, []byte(payload))
			if err != nil {
				return err
			}
			c.buf = append(c.buf, val)
			c.lastID = msg.ID
		}
	}
	return nil
}

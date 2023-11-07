package spin

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

const (
	MessageQueueKey = "mq"
)

type QueueClient struct {
	db     *redis.Client
	buf    []any
	lastID string
}

func NewQueueClient(db *redis.Client) *QueueClient {
	return &QueueClient{db: db}
}

func (c *QueueClient) Reset() error {
	return c.fillBuf("0-0")
}

func (c *QueueClient) Read() (any, error) {
	if len(c.buf) > 0 {
		return c.pop(), nil
	}
	if err := c.fillBuf("$"); err != nil {
		return nil, err
	}
	return c.pop(), nil
}

func (c *QueueClient) pop() any {
	var msg any
	msg, c.buf = c.buf[0], c.buf[1:]
	return msg
}

func (c *QueueClient) fillBuf(fromID string) error {
	ctx := context.Background()

	if c.lastID != "" {
		fromID = c.lastID
	}
	res := c.db.XRead(ctx, &redis.XReadArgs{
		Streams: []string{MessageQueueKey, fromID},
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

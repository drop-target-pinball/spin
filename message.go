package spin

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Play struct {
	ID       string `json:"id"`
	Loops    int    `json:"loops,omitempty"`
	Repeat   bool   `json:"repeat,omitempty"`
	Priority int    `json:"priority,omitempty"`
	Notify   bool   `json:"notify,omitempty"`
}

type Stop struct {
}

type ParseFunc func(data []byte) (any, error)

var parsers = map[string]ParseFunc{
	"Play": func(data []byte) (any, error) { m := Play{}; err := json.Unmarshal(data, &m); return m, err },
}

func ParseMessage(typ string, data []byte) (any, error) {
	parser, ok := parsers[typ]
	if !ok {
		return nil, fmt.Errorf("unable to parse message type '%v'", typ)
	}
	return parser(data)
}

const (
	MessageQueueKey = "mq"
)

type MessageClient struct {
	db *redis.Client
}

func NewMessageClient(db *redis.Client) *MessageClient {
	return &MessageClient{db: db}
}

func (c *MessageClient) Read() ([]any, error) {
	var messages []any
	ctx := context.Background()

	res := c.db.XRead(ctx, &redis.XReadArgs{
		Streams: []string{MessageQueueKey, "$"},
	})
	if res.Err() != nil {
		return nil, res.Err()
	}
	streams, err := res.Result()
	if err != nil {
		return nil, err
	}
	for _, stream := range streams {
		for _, msg := range stream.Messages {
			anyType, ok := msg.Values["type"]
			if !ok {
				return nil, fmt.Errorf("message did not contain a 'type' field")
			}
			typ, ok := anyType.(string)
			if !ok {
				return nil, fmt.Errorf("message type field is not a string")
			}
			anyPayload, ok := msg.Values["payload"]
			if !ok {
				return nil, fmt.Errorf("message did not contain a 'payload' field")
			}
			payload, ok := anyPayload.(string)
			if !ok {
				return nil, fmt.Errorf("message payload field is not a string")
			}
			val, err := ParseMessage(typ, []byte(payload))
			if err != nil {
				return nil, err
			}
			messages = append(messages, val)
		}
	}
	return messages, nil
}

package spin

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/drop-target-pinball/spin/v2/msg/sys"
)

type Header struct {
	To   string `json:"to"`
	Addr string `json:"addr"`
	Type string `json:"type"`
}

type Body interface {
	Type() string
	Interface() string
}

type Message struct {
	Header Header
	Body   Body
}

type parseFunc func([]byte) (Body, error)

var parsers = map[string]parseFunc{
	sys.TypePing: func(b []byte) (Body, error) { m := sys.Ping{}; err := json.Unmarshal(b, &m); return m, err },
}

func ParseHeader(data []byte, dest *Message) error {
	if err := json.Unmarshal(data, &dest.Header); err != nil {
		return err
	}
	return nil
}

func ParseBody(data []byte, dest *Message) error {
	if dest.Header.To == "" {
		return fmt.Errorf("header does not contain To address")
	}
	parser, ok := parsers[dest.Header.Type]
	if !ok {
		return fmt.Errorf("unknown message type: %v", dest.Header.Type)
	}
	body, err := parser(data)
	if err != nil {
		return err
	}
	dest.Body = body
	return nil
}

func ParseMessage(d *json.Decoder) (Message, error) {
	var msg Message
	var obj1, obj2 any

	if err := d.Decode(&obj1); err != nil {
		if errors.Is(err, io.EOF) {
			return msg, io.EOF
		}
		return msg, err
	}
	if err := d.Decode(&obj2); err != nil {
		if errors.Is(err, io.EOF) {
			return msg, io.EOF
		}
		return msg, err
	}

	header, err := json.Marshal(obj1)
	if err != nil {
		return msg, err
	}
	body, err := json.Marshal(obj2)
	if err != nil {
		return msg, err
	}

	if err := ParseHeader(header, &msg); err != nil {
		return msg, err
	}
	if err := ParseBody(body, &msg); err != nil {
		return msg, err
	}

	return msg, nil
}

package spin

import (
	"encoding/json"
	"fmt"
)

type Load struct {
	ID string `json:"id"`
}

type Play struct {
	ID       string `json:"id"`
	Loops    int    `json:"loops,omitempty"`
	Repeat   bool   `json:"repeat,omitempty"`
	Priority int    `json:"priority,omitempty"`
	Notify   bool   `json:"notify,omitempty"`
}

type Stop struct {
	ID string `json:"id"`
}

type ParseFunc func(data []byte) (any, error)

var parsers = map[string]ParseFunc{
	"Load": func(data []byte) (any, error) { m := Load{}; err := json.Unmarshal(data, &m); return m, err },
	"Play": func(data []byte) (any, error) { m := Play{}; err := json.Unmarshal(data, &m); return m, err },
}

func ParseMessage(typ string, data []byte) (any, error) {
	parser, ok := parsers[typ]
	if !ok {
		return nil, fmt.Errorf("unable to parse message type '%v'", typ)
	}
	return parser(data)
}

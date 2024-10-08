package sys

import "fmt"

const Interface = "sys"

const (
	TypePing = "ping"
)

type Ping struct {
	From string `json:"from"`
	ID   string `json:"id"`
}

func (m Ping) String() string {
	return fmt.Sprintf("%v: from=%v id=%v", m.Type(), m.From, m.ID)
}

func (m Ping) Type() string      { return TypePing }
func (m Ping) Interface() string { return Interface }

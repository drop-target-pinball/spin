package spin

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"slices"
	"strconv"
	"strings"

	"github.com/drop-target-pinball/spin/v2/msg"
)

type Header struct {
	Chan string `json:"chan"`
	Name string `json:"name"`
}

type Message interface {
	Name() string
	Chan() string
}

type Packet struct {
	Header Header
	Body   Message
}

type parseFunc func([]byte) (Message, error)

var parsers = map[string]parseFunc{
	msg.PingName: func(b []byte) (Message, error) { m := msg.Ping{}; err := json.Unmarshal(b, &m); return m, err },
	msg.PongName: func(b []byte) (Message, error) { m := msg.Pong{}; err := json.Unmarshal(b, &m); return m, err },
}

func ParseHeader(data []byte, dest *Packet) error {
	if err := json.Unmarshal(data, &dest.Header); err != nil {
		return err
	}
	return nil
}

func ParseBody(data []byte, dest *Packet) error {
	if dest.Header.Chan == "" {
		return fmt.Errorf("header does not contain channel name")
	}
	parser, ok := parsers[dest.Header.Name]
	if !ok {
		return fmt.Errorf("unknown message name: %v", dest.Header.Name)
	}
	body, err := parser(data)
	if err != nil {
		return err
	}
	dest.Body = body
	return nil
}

func ParseMessage(d *json.Decoder) (Packet, error) {
	var msg Packet
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

func FormatBody(b Message) string {
	var s strings.Builder
	s.WriteString(b.Name())
	t := reflect.TypeOf(b)
	v := reflect.ValueOf(b)
	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i)
		fv := v.Field(i)
		j, ok := ft.Tag.Lookup("json")
		if !ok {
			continue
		}
		jf := strings.Split(j, ",")
		if len(jf) == 0 {
			continue
		}
		name := jf[0]
		optional := slices.Contains(jf[1:], "omitempty")

		var sval string
		var prefix string
		var flag bool
		val := fv.Interface()

		switch a := val.(type) {
		case int:
			if !optional || a != 0 {
				sval = strconv.Itoa(a)
			}
		case float64:
			if !optional || a != 0 {
				sval = strconv.FormatFloat(a, 'f', -1, 64)
			}
		case string:
			if !optional || a != "" {
				sval = Quote(a)
			}
		case bool:
			if a {
				flag = true
			} else {
				if !optional {
					flag = true
					prefix = "no-"
				}
			}
		}
		if sval == "" && !flag {
			continue
		}
		s.WriteString(" ")
		s.WriteString(prefix)
		s.WriteString(name)
		if !flag {
			s.WriteString("=")
			s.WriteString(sval)
		}
	}
	return s.String()
}

func Quote(s string) string {
	if s == "" {
		return "''"
	}
	if !strings.ContainsAny(s, ` "`) {
		return s
	}
	if strings.Contains(s, "'") {
		return `"` + strings.ReplaceAll(s, "\"", `\"`) + `"`
	}
	return "'" + s + "'"
}

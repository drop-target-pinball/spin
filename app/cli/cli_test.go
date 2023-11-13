package cli

import (
	"testing"

	"github.com/drop-target-pinball/spin/v2"
)

func TestFormatMessage(t *testing.T) {
	tests := []struct {
		name string
		str  string
		msg  any
	}{
		{"1pos", "play service_enter", spin.Play{
			ID: "service_enter",
		}},
		{"option", "play service_enter --loops=2", spin.Play{
			ID:    "service_enter",
			Loops: 2,
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			str := formatMessage(test.msg)
			if str != test.str {
				t.Errorf("\n have: %v \n want: %v", str, test.str)
			}
		})
	}
}

func TestCamelToSnake(t *testing.T) {
	tests := []struct {
		camel string
		snake string
	}{
		{"Play", "play"},
		{"play", "play"},
		{"AddBall", "add_ball"},
	}

	for _, test := range tests {
		t.Run(test.camel, func(t *testing.T) {
			snake := camelToSnake(test.camel)
			if snake != test.snake {
				t.Errorf("\n have: %v \n want: %v", snake, test.snake)
			}
		})
	}
}

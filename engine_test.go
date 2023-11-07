package spin

import (
	"log"
	"strings"
	"testing"
)

func TestPathTo(t *testing.T) {
	e := NewEngine(TestSettings(t))
	e.Settings.Dir = "/tmp/example-project"
	have := e.PathTo("myconfig/config.hcl")
	want := "/tmp/example-project/myconfig/config.hcl"
	if have != want {
		t.Errorf("\n have: %v \n want: %v", have, want)
	}
}

func TestLogging(t *testing.T) {
	e := NewEngine(TestSettings(t))

	tests := []struct {
		name    string
		devMode bool
		panic   bool
		fn      func(...any)
	}{
		{"error", false, true, e.Error},
		{"warn", false, false, e.Warn},
		{"warn", true, false, e.Warn},
		{"note", false, false, e.Log},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var buf strings.Builder
			log.SetFlags(0)
			log.SetOutput(&buf)
			e.DevMode = test.devMode

			want := "this is a test"
			defer func() {
				if r := recover(); r == nil {
					if test.panic {
						t.Errorf("expected panic")
					}
				}
				have := strings.TrimSpace(buf.String())
				if have != want {
					t.Errorf("\n have: %v \n want: %v", have, want)
				}
			}()
			test.fn(want)
		})
	}

}

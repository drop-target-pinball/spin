package spin

import (
	"encoding/json"
	"testing"
)

// Error calls testing.Errorf with a formatting message that shows the
// contents of both have and want as JSON.
func TestError(t *testing.T, have any, want any) {
	jsonHave, err := json.Marshal(have)
	if err != nil {
		panic(err)
	}
	jsonWant, err := json.Marshal(want)
	if err != nil {
		panic(err)
	}
	t.Errorf("\nhave: %v\nwant: %v", string(jsonHave), string(jsonWant))
}

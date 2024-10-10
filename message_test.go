package spin

import "testing"

type testStr struct {
	A string `json:"a"`
	B string `json:"b,omitempty"`
}

func (t testStr) Chan() string { return "chan" }
func (t testStr) Name() string { return "test-str" }

type testInt struct {
	A int `json:"a"`
	B int `json:"b,omitempty"`
}

func (t testInt) Chan() string { return "chan" }
func (t testInt) Name() string { return "test-int" }

type testFloat struct {
	A float64 `json:"a"`
	B float64 `json:"b,omitempty"`
}

func (t testFloat) Chan() string { return "chan" }
func (t testFloat) Name() string { return "test-float" }

type testBool struct {
	A bool `json:"a"`
	B bool `json:"b,omitempty"`
}

func (t testBool) Chan() string { return "chan" }
func (t testBool) Name() string { return "test-bool" }

func TestFormatBody(t *testing.T) {
	tests := []struct {
		msg Body
		str string
	}{
		{testStr{A: "foo"}, "test-str a=foo"},
		{testStr{A: "foo bar"}, "test-str a='foo bar'"},
		{testStr{A: "foo", B: "bar"}, "test-str a=foo b=bar"},
		{testStr{B: "bar"}, "test-str a='' b=bar"},

		{testInt{A: 11}, "test-int a=11"},
		{testInt{A: 11, B: 22}, "test-int a=11 b=22"},
		{testInt{B: 22}, "test-int a=0 b=22"},

		{testFloat{A: 1.1}, "test-float a=1.1"},
		{testFloat{A: 1.1, B: 2.2}, "test-float a=1.1 b=2.2"},
		{testFloat{B: 2.2}, "test-float a=0 b=2.2"},

		{testBool{A: true}, "test-bool a"},
		{testBool{A: true, B: true}, "test-bool a b"},
		{testBool{B: true}, "test-bool no-a b"},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			str := FormatBody(test.msg)
			if str != test.str {
				t.Fatalf("\n have: %v \n want: %v", str, test.str)
			}
		})
	}
}

func TestQuote(t *testing.T) {
	tests := []struct {
		str    string
		quoted string
	}{
		{"foo", "foo"},
		{"foo bar", "'foo bar'"},
		{"foo's bar", `"foo's bar"`},
		{"", "''"},
		{" ", "' '"},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			quoted := Quote(test.str)
			if quoted != test.quoted {
				t.Fatalf("\n have: %v \n want: %v", quoted, test.quoted)
			}
		})
	}
}

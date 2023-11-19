package spin

import (
	"reflect"
	"testing"
)

func TestMessageStream(t *testing.T) {
	want1 := Play{ID: "sound1"}
	want2 := Play{ID: "sound2"}
	want3 := Play{ID: "sound3"}

	s := testSettings(t)
	pin, err := NewClient(s.RedisAddress)
	if err != nil {
		t.Fatal(err)
	}

	pin.Reset()
	if err := pin.Send(want1, want2); err != nil {
		t.Fatal(err)
	}

	have1, err := pin.Read()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(have1, want1) {
		t.Fatalf("\n have: %v \n want: %v", have1, want1)
	}
	have2, err := pin.Read()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(have2, want2) {
		t.Fatalf("\n have: %v \n want: %v", have2, want2)
	}

	if err := pin.Send(want3); err != nil {
		t.Fatal(err)
	}
	have3, err := pin.Read()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(have3, want3) {
		t.Errorf("\n have: %v \n want: %v", have3, want3)
	}
}

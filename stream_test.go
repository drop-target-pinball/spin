package spin

import (
	"os"
	"reflect"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestMessageStream(t *testing.T) {
	want1 := Play{ID: "sound1"}
	want2 := Play{ID: "sound2"}
	want3 := Play{ID: "sound3"}

	e := NewEngine(testSettings(t))
	if err := os.WriteFile(e.PathTo("test.hcl"), []byte{}, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := e.Init(); err != nil {
		t.Fatal(err)
	}

	db := redis.NewClient(&redis.Options{Addr: e.Settings.RedisRunAddress})
	stream := NewStreamClient(db)

	if err := stream.Send(want1, want2); err != nil {
		t.Fatal(err)
	}
	stream.Reset()
	stream.Read() // Discard the first load message

	have1, err := stream.Read()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(have1, want1) {
		t.Fatalf("\n have: %v \n want: %v", have1, want1)
	}
	have2, err := stream.Read()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(have2, want2) {
		t.Fatalf("\n have: %v \n want: %v", have2, want2)
	}

	if err := stream.Send(want3); err != nil {
		t.Fatal(err)
	}
	have3, err := stream.Read()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(have3, want3) {
		t.Errorf("\n have: %v \n want: %v", have3, want3)
	}
}

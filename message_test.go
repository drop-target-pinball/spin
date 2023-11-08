package spin

import (
	"os"
	"reflect"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestMessageClient(t *testing.T) {
	want1 := Play{ID: "sound1"}
	want2 := Play{ID: "sound2"}
	want3 := Play{ID: "sound3"}

	e := NewEngine(TestSettings(t))
	if err := os.WriteFile(e.PathTo("test.hcl"), []byte{}, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := e.Init(); err != nil {
		t.Fatal(err)
	}

	db := redis.NewClient(&redis.Options{Addr: e.Settings.RedisRunAddress})
	cli := NewQueueClient(db)

	e.Send(want1)
	e.Send(want2)
	cli.Reset()

	have1, err := cli.Read()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(have1, want1) {
		t.Errorf("\n have: %v \n want: %v", have1, want1)
	}
	have2, err := cli.Read()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(have2, want2) {
		t.Errorf("\n have: %v \n want: %v", have2, want2)
	}

	e.Send(want3)
	have3, err := cli.Read()
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(have3, want3) {
		t.Errorf("\n have: %v \n want: %v", have3, want3)
	}
}

package spin

import (
	"os"
	"reflect"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestMessageClient(t *testing.T) {
	e := NewEngine(TestSettings(t))
	if err := os.WriteFile(e.PathTo("test.hcl"), []byte{}, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := e.Init(); err != nil {
		t.Fatal(err)
	}

	db := redis.NewClient(&redis.Options{Addr: e.Settings.RedisRunAddress})
	cli := NewMessageClient(db)

	e.Send(Play{ID: "sound1"})
	e.Send(Play{ID: "sound2"})

	msgs, err := cli.Read()
	if err != nil {
		t.Fatal(err)
	}
	want := []Play{{ID: "sound1"}, {ID: "sound2"}}

	if !reflect.DeepEqual(msgs, want) {
		t.Errorf("\n have: %v \n want: %v", msgs, want)
	}
}

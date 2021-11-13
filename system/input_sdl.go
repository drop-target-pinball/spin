package system

import (
	"fmt"
	"os"

	"github.com/drop-target-pinball/spin"
	"github.com/veandco/go-sdl2/sdl"
)

type InputSDL struct{}

func RegisterInputSDL(eng *spin.Engine) {
	s := &InputSDL{}
	eng.RegisterServer(s)
}

func (s *InputSDL) Service() {
	for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
		switch event := e.(type) {
		case *sdl.KeyboardEvent:
			fmt.Println(event)
		case *sdl.QuitEvent:
			os.Exit(0)
		}
	}
}

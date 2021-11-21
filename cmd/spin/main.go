package main

import (
	"context"
	"fmt"
	"time"

	"github.com/drop-target-pinball/spin/coroutine"
)

type Advance struct{}

func (a Advance) Key() interface{} {
	return a
}

func a(ctx *coroutine.Context) {
	for {
		fmt.Println("a")
		if done := ctx.WaitFor(1 * time.Second); done {
			fmt.Println("a done")
			return
		}
	}
}

func b(ctx *coroutine.Context) {
	for {
		fmt.Println("b")
		if done := ctx.WaitFor(1 * time.Second); done {
			fmt.Println("b done")
			return
		}
	}
}

func c(ctx *coroutine.Context) {
loop1:
	for {
		fmt.Println("c-1")
		evt := ctx.WaitForUntil(1*time.Second, Advance{})
		switch evt {
		case coroutine.Cancel{}:
			fmt.Println("c done")
			return
		case Advance{}:
			fmt.Println("c advance")
			break loop1
		}
	}

	for {
		fmt.Println("c-2")
		if done := ctx.WaitFor(1 * time.Second); done {
			fmt.Println("c done")
			return
		}
	}
}

func main() {
	//panic("these aren't the droids you are looking for")
	e := coroutine.NewEnv()
	ctx, cancel := context.WithCancel(context.Background())

	e.Create(ctx, a)
	e.Create(ctx, b)
	e.Create(ctx, c)
	for i := 0; i < 6*10; i++ {
		fmt.Println(".")
		e.Service()
		if i == 20 {
			fmt.Println("request advance")
			e.Post(Advance{})
		}
		if i == 40 {
			fmt.Println("cancel")
			cancel()
		}
		time.Sleep(100 * time.Millisecond)
	}
}

package main

import "fmt"

func integers() chan bool {
	yield := make(chan bool)
	count := 0
	go func() {
		for i := 0; i < 2; i++ {
			yield <- true
			count++
		}
		yield <- false
	}()
	return yield
}

func main() {
	resume := integers()
	generateInteger := func() bool {
		return <-resume
	}

	i := 0
	for {
		fmt.Println(i)
		if !generateInteger() {
			break
		}
		i++
	}
}

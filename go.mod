module github.com/drop-target-pinball/spin

go 1.16

require (
	github.com/chzyer/logex v1.2.0 // indirect
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e
	github.com/chzyer/test v0.0.0-20210722231415-061457976a23 // indirect
	github.com/drop-target-pinball/coroutine v0.2.1
	github.com/drop-target-pinball/go-pinproc v0.1.0
	github.com/veandco/go-sdl2 v0.4.10
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
	golang.org/x/text v0.3.7
)

// replace github.com/drop-target-pinball/go-pinproc => /Users/mcgann/Code/drop-target-pinball/go-pinproc
replace github.com/drop-target-pinball/coroutine => /Users/mcgann/Code/drop-target-pinball/coroutine

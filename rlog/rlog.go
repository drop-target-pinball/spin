package rlog

import (
	"fmt"
	"log"
	"time"

	"github.com/benbjohnson/clock"
)

type entry struct {
	expiration time.Time
	message    string
}

type Logger struct {
	clock    clock.Clock
	duration time.Duration
	logger   *log.Logger
	entries  []entry
}

func New(duration time.Duration, logger *log.Logger) *Logger {
	return &Logger{
		clock:    clock.New(),
		duration: duration,
		entries:  make([]entry, 0),
		logger:   logger,
	}
}

func (l *Logger) addEntry(message string) {
	l.entries = append(l.entries, entry{l.clock.Now().Add(l.duration), message})
	l.reap()
	if l.logger != nil {
		l.logger.Print(message)
	}
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.addEntry(fmt.Sprintf(format, v...))
}

func (l *Logger) Print(v ...interface{}) {
	l.addEntry(fmt.Sprint(v...))
}

func (l *Logger) Messages() []string {
	l.reap()
	m := make([]string, len(l.entries))
	for i := 0; i < len(l.entries); i++ {
		m[i] = l.entries[i].message
	}
	return m
}

func (l *Logger) reap() {
	now := l.clock.Now()

	if len(l.entries) == 0 {
		return
	}

	if !now.After(l.entries[0].expiration) {
		return
	}

	for i := 1; i < len(l.entries); i++ {
		if !now.After(l.entries[i].expiration) {
			l.entries = l.entries[i:]
			return
		}
	}
	l.entries = nil
}

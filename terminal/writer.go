package terminal

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"sync"
	"time"
)

type Writer struct {
	RefreshFunc     func()
	w               io.Writer
	line            bytes.Buffer
	backlog         bytes.Buffer
	timer           *time.Timer
	mutex           sync.Mutex
	firstInterval   time.Duration // wait this long before emitting the first line
	backlogInterval time.Duration // wait this long between backlog processing
	maxUpdate       int           // maximum number of characters per update
}

func NewWriter(w io.Writer) *Writer {
	cw := &Writer{
		RefreshFunc:     func() {},
		w:               w,
		firstInterval:   time.Millisecond * 10,
		backlogInterval: time.Millisecond * 100,
		maxUpdate:       2000,
	}
	return cw
}

func (c *Writer) Write(p []byte) (int, error) {
	for _, b := range p {
		c.line.WriteByte(b)
		if b == '\n' {
			c.backlog.Write(c.line.Bytes())
			if c.timer == nil {
				c.timer = time.AfterFunc(c.firstInterval, c.emit)
			}
			c.line.Reset()
		}
	}
	return len(p), nil
}

func (c *Writer) Flush() {
	c.emit()
}

func (c *Writer) emit() {
	c.mutex.Lock()
	defer func() {
		c.mutex.Unlock()
	}()

	if c.backlog.Len() == 0 {
		c.timer = nil
		return
	}
	update := c.backlog.String()
	lines := strings.Count(update, "\n")
	omission := false
	if lines > c.maxUpdate {
		// Count backwards to find start of the first line in the
		// maximum lines allowed per update
		omission = true
		seen := 0
		for i := len(update) - 1; i >= 0; i-- {
			if update[i] == '\n' {
				seen++
				if seen == c.maxUpdate {
					update = update[i+1 : len(update)-1]
					break
				}
			}
		}
	}
	io.WriteString(c.w, AnsiClearLine)
	if omission {
		text := fmt.Sprintf("... omitted %v lines\n", lines-c.maxUpdate)
		io.WriteString(c.w, text)
	}
	io.WriteString(c.w, update)
	c.RefreshFunc()
	c.backlog.Reset()
	c.timer = time.AfterFunc(c.backlogInterval, c.emit)
}

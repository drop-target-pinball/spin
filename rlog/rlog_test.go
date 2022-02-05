package rlog

import (
	"reflect"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
)

func TestNoExpirations(t *testing.T) {
	l := New(time.Duration(10*time.Second), nil)
	mock := clock.NewMock()
	l.clock = mock

	l.Print("one")
	mock.Add(2 * time.Second)
	l.Print("two")
	mock.Add(2 * time.Second)
	l.Print("three")

	have := l.Messages()
	want := []string{"one", "two", "three"}
	if !reflect.DeepEqual(have, want) {
		t.Errorf("\n have: %v \n want: %v", have, want)
	}
}

func TestOneExpiration(t *testing.T) {
	l := New(time.Duration(10*time.Second), nil)
	mock := clock.NewMock()
	l.clock = mock

	l.Print("one")
	mock.Add(6 * time.Second)
	l.Print("two")
	mock.Add(6 * time.Second)
	l.Print("three")

	have := l.Messages()
	want := []string{"two", "three"}
	if !reflect.DeepEqual(have, want) {
		t.Errorf("\n have: %v \n want: %v", have, want)
	}
}

func TestTwoExpiration(t *testing.T) {
	l := New(time.Duration(10*time.Second), nil)
	mock := clock.NewMock()
	l.clock = mock

	l.Print("one")
	mock.Add(12 * time.Second)
	l.Print("two")
	mock.Add(12 * time.Second)
	l.Print("three")

	have := l.Messages()
	want := []string{"three"}
	if !reflect.DeepEqual(have, want) {
		t.Errorf("\n have: %v \n want: %v", have, want)
	}
}

func TestAllExpiration(t *testing.T) {
	l := New(time.Duration(10*time.Second), nil)
	mock := clock.NewMock()
	l.clock = mock

	l.Print("one")
	l.Print("two")
	l.Print("three")
	mock.Add(12 * time.Second)

	have := l.Messages()
	want := []string{}
	if !reflect.DeepEqual(have, want) {
		t.Errorf("\n have: %v \n want: %v", have, want)
	}
}

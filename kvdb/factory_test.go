package kvdb

import (
	"strings"
	"testing"
)

func TestPassthroughFactory(t *testing.T) {
	UnregisterAll()
	defer func() {
		UnregisterAll()
		Register("passthrough", PassthroughFactory)
	}()
	f := Factories()
	if len(f) != 0 {
		t.Fatal(f)
	}
	Register("passthrough", PassthroughFactory)
	f = Factories()
	if len(f) != 1 {
		t.Fatal(f)
	}
}

func TestNotexistedDriver(t *testing.T) {
	d, err := NewDriver("notexist", nil)
	if d != nil {
		t.Fatal(d)
	}
	if err == nil || !strings.Contains(err.Error(), "unknown driver") {
		t.Fatal(err)
	}
}

func TestRegisterExistedDriver(t *testing.T) {
	UnregisterAll()
	defer func() {
		UnregisterAll()
		Register("passthrough", PassthroughFactory)
	}()
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal(r)
		}
		err := r.(error)
		if err == nil || !strings.Contains(err.Error(), "twice") {
			t.Fatal(err)
		}
	}()
	Register("passthrough", PassthroughFactory)
	Register("passthrough", PassthroughFactory)
}

func TestRegisterNilDriver(t *testing.T) {
	UnregisterAll()
	defer func() {
		UnregisterAll()
		Register("simpleid", PassthroughFactory)
	}()
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal(r)
		}
		err := r.(error)
		if err == nil || !strings.Contains(err.Error(), "nil") {
			t.Fatal(err)
		}
	}()
	Register("nil", nil)
}

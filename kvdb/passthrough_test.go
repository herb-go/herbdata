package kvdb

import (
	"testing"

	"github.com/herb-go/herbdata"
)

func TestPassthrough(t *testing.T) {
	UnregisterAll()
	Register("passthrough", PassthroughFactory)
	defer func() {
		UnregisterAll()
		Register("passthrough", PassthroughFactory)
	}()

	var err error
	var ok bool
	var key = []byte("key")
	var value = []byte("value")
	c := &Config{
		Driver: "passthrough",
	}
	p := New()
	err = c.ApplyTo(p)
	if err != nil {
		panic(err)
	}
	err = p.Start()
	if err != nil {
		panic(err)
	}
	if !p.Features().SupportAll(passthroughFeatures) {
		t.Fatal()
	}
	err = p.Set(key, value)
	if err != nil {
		t.Fatal()
	}
	_, err = p.Get(key)
	if err != herbdata.ErrNotFound {
		t.Fatal()
	}
	err = p.SetWithTTL(key, value, 3600)
	if err != nil {
		t.Fatal()
	}
	_, err = p.Get(key)
	if err != herbdata.ErrNotFound {
		t.Fatal()
	}
	ok, err = p.Insert(key, value)
	if ok != false || err != nil {
		t.Fatal()
	}
	_, err = p.Get(key)
	if err != herbdata.ErrNotFound {
		t.Fatal()
	}

	ok, err = p.InsertWithTTL(key, value, 3600)
	if ok != false || err != nil {
		t.Fatal()
	}
	_, err = p.Get(key)
	if err != herbdata.ErrNotFound {
		t.Fatal()
	}

	ok, err = p.Update(key, value)
	if ok != false || err != nil {
		t.Fatal()
	}
	_, err = p.Get(key)
	if err != herbdata.ErrNotFound {
		t.Fatal()
	}

	ok, err = p.UpdateWithTTL(key, value, 3600)
	if ok != false || err != nil {
		t.Fatal()
	}
	_, err = p.Get(key)
	if err != herbdata.ErrNotFound {
		t.Fatal()
	}
	err = p.Delete(key)
	if err != nil {
		t.Fatal()
	}
	err = p.Stop()
	if err != nil {
		panic(err)
	}
}

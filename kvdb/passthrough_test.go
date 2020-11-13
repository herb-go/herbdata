package kvdb

import (
	"testing"

	"github.com/herb-go/herbdata"
)

func TestPassthrough(t *testing.T) {
	var err error
	var ok bool
	var key = []byte("key")
	var value = []byte("value")
	if !Passthrough.Features().SupportAll(passthroughFeatures) {
		t.Fatal()
	}
	err = Passthrough.Set(key, value)
	if err != nil {
		t.Fatal()
	}
	_, err = Passthrough.Get(key)
	if err != herbdata.ErrNotFound {
		t.Fatal()
	}
	err = Passthrough.SetWithTTL(key, value, 3600)
	if err != nil {
		t.Fatal()
	}
	_, err = Passthrough.Get(key)
	if err != herbdata.ErrNotFound {
		t.Fatal()
	}
	ok, err = Passthrough.Insert(key, value)
	if ok != false || err != nil {
		t.Fatal()
	}
	_, err = Passthrough.Get(key)
	if err != herbdata.ErrNotFound {
		t.Fatal()
	}

	ok, err = Passthrough.InsertWithTTL(key, value, 3600)
	if ok != false || err != nil {
		t.Fatal()
	}
	_, err = Passthrough.Get(key)
	if err != herbdata.ErrNotFound {
		t.Fatal()
	}

	ok, err = Passthrough.Update(key, value)
	if ok != false || err != nil {
		t.Fatal()
	}
	_, err = Passthrough.Get(key)
	if err != herbdata.ErrNotFound {
		t.Fatal()
	}

	ok, err = Passthrough.UpdateWithTTL(key, value, 3600)
	if ok != false || err != nil {
		t.Fatal()
	}
	_, err = Passthrough.Get(key)
	if err != herbdata.ErrNotFound {
		t.Fatal()
	}
	err = Passthrough.Delete(key)
	if err != nil {
		t.Fatal()
	}
}

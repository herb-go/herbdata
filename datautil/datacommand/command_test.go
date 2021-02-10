package datacommand

import (
	"bytes"
	"testing"
)

func equal(src, target *Command) bool {
	return src.Type == target.Type &&
		bytes.Compare(src.Key, target.Key) == 0 &&
		bytes.Compare(src.Data, target.Data) == 0 &&
		src.Expired == target.Expired
}
func TestCommand(t *testing.T) {
	var err error
	var c *Command
	var result *Command
	var bs []byte
	var key = []byte("testkey")
	var data = []byte("testdata")
	var expired = int64(123456)
	c = NewDeleteCommand(key)
	bs, err = c.EncodeData()
	if err != nil {
		t.Fatal(err)
	}
	result = New()
	err = result.DecodeData(bs)
	if err != nil {
		t.Fatal(err)
	}
	if !equal(c, result) {
		t.Fatal(c, result)
	}
	c = NewSetCommand(key, data)
	bs, err = c.EncodeData()
	if err != nil {
		t.Fatal(err)
	}
	result = New()
	err = result.DecodeData(bs)
	if err != nil {
		t.Fatal(err)
	}
	if !equal(c, result) {
		t.Fatal(c, result)
	}
	c = NewSetWithExpiredCommand(key, data, expired)
	bs, err = c.EncodeData()
	if err != nil {
		t.Fatal(err)
	}
	result = New()
	err = result.DecodeData(bs)
	if err != nil {
		t.Fatal(err)
	}
	if !equal(c, result) {
		t.Fatal(c, result)
	}
}

func TestError(t *testing.T) {
	c, err := ParseCommand(bytes.NewBuffer([]byte{100}))
	if c != nil || err != ErrInvalidCommandType {
		t.Fatal(c, err)
	}
	c = New()
	c.Type = 100
	_, err = c.EncodeData()
	if err != ErrInvalidCommandType {
		t.Fatal(c, err)
	}
}

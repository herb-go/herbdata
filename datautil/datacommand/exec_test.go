package datacommand

import "testing"

type testStore struct {
	Command *Command
}

func (s *testStore) Set(key []byte, value []byte) error {
	s.Command = NewSetCommand(key, value)
	return nil
}
func (s *testStore) SetWithExpired(key []byte, value []byte, expired int64) error {
	s.Command = NewSetWithExpiredCommand(key, value, expired)
	return nil
}
func (s *testStore) Delete(key []byte) error {
	s.Command = NewDeleteCommand(key)
	return nil
}

func newTestStore() *testStore {
	return &testStore{
		Command: New(),
	}
}
func TestExec(t *testing.T) {
	var key = []byte("testkey")
	var value = []byte("testvalue")
	var expired = int64(123456)
	var c *Command
	var err error
	var s *testStore
	c = NewDeleteCommand(key)
	s = newTestStore()
	err = Exec(c, s)
	if err != nil {
		panic(err)
	}
	if !equal(c, s.Command) {
		t.Fatal(c, s)
	}
	c = NewDeleteCommand(key)
	s = newTestStore()
	err = ExecWithExpired(c, s)
	if err != nil {
		panic(err)
	}
	if !equal(c, s.Command) {
		t.Fatal(c, s)
	}

	c = NewSetCommand(key, value)
	s = newTestStore()
	err = Exec(c, s)
	if err != nil {
		panic(err)
	}
	if !equal(c, s.Command) {
		t.Fatal(c, s)
	}

	c = NewSetCommand(key, value)
	s = newTestStore()
	err = ExecWithExpired(c, s)
	if err != nil {
		panic(err)
	}
	if !equal(c, s.Command) {
		t.Fatal(c, s)
	}

	c = NewSetWithExpiredCommand(key, value, expired)
	s = newTestStore()
	err = Exec(c, s)
	if err != ErrInvalidCommandType {
		t.Fatal(err)
	}

	c = NewSetWithExpiredCommand(key, value, expired)
	s = newTestStore()
	err = ExecWithExpired(c, s)
	if err != nil {
		panic(err)
	}
	if !equal(c, s.Command) {
		t.Fatal(c, s)
	}

	c = New()
	c.Type = byte(100)
	err = Exec(c, s)
	if err != ErrInvalidCommandType {
		t.Fatal(err)
	}
	err = ExecWithExpired(c, s)
	if err != ErrInvalidCommandType {
		t.Fatal(err)
	}
}

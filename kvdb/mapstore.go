package kvdb

import "sync"

type MapStore struct {
	Nop
	store sync.Map
}

//Set set value by given key
func (s *MapStore) Set(key []byte, value []byte) error {
	s.store.Store(key, value)
	return nil
}

//Get get value by given key
func (s *MapStore) Get(key []byte) ([]byte, error) {

}

//Del delete value by given key
func (s *MapStore) Del(key []byte) error {
	s.store.Delete(key)
	return nil
}

func NewMapStore() *MapStore {
	return &MapStore{}
}

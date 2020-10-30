package kvdb

import "sync"

var mapstoreFeatures = FeatureStore

type MapStore struct {
	Nop
	store sync.Map
}

//Set set value by given key
func (s *MapStore) Set(key []byte, value []byte) error {
	s.store.Store(string(key), value)
	return nil
}

//Get get value by given key
func (s *MapStore) Get(key []byte) ([]byte, error) {
	d, ok := s.store.Load(string(key))
	if !ok {
		return nil, ErrNotFound
	}
	return d.([]byte), nil
}

//Del delete value by given key
func (s *MapStore) Del(key []byte) error {
	s.store.Delete(string(key))
	return nil
}
func (s *MapStore) Features() Feature {
	return mapstoreFeatures
}
func NewMapStore() *MapStore {
	return &MapStore{}
}

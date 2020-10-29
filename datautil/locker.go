package datautil

import "sync"

//Locker cache locker
type Locker struct {
	sync.RWMutex
	store *sync.Map
	Key   []byte
}

//Unlock unlock and delete locker
func (l *Locker) Unlock() {
	l.RWMutex.Unlock()
	l.store.Delete(l.Key)
}

//LockerStore lock store
type LockerStore struct {
	store *sync.Map
}

//Locker create new locker with given key.
//Return locker and if locker is locked.
func (s *LockerStore) Locker(key []byte) (*Locker, bool) {

	v, ok := s.store.LoadOrStore(key, s.newLocker(key))
	return v.(*Locker), ok
}

func (s *LockerStore) newLocker(key []byte) *Locker {
	return &Locker{
		store: s.store,
		Key:   key,
	}
}

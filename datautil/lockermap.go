package datautil

import "sync"

type lockerentity struct {
	count   int64
	rwcount int64
	sync.RWMutex
}
type LockerMap struct {
	locker sync.Mutex
	store  map[string]*lockerentity
}

type locker struct {
	m   *LockerMap
	key string
}

func (l *locker) Lock() {
	l.m.Lock(l.key)
}
func (l *locker) Unlock() {
	l.m.Unlock(l.key)
}

type rlocker struct {
	m   *LockerMap
	key string
}

func (l *rlocker) Lock() {
	l.m.RLock(l.key)
}
func (l *rlocker) Unlock() {
	l.m.RUnlock(l.key)
}
func (m *LockerMap) Lock(k string) {
	m.locker.Lock()
	l := m.store[k]
	if l == nil {
		l = &lockerentity{}
		m.store[k] = l
	}
	l.count++
	m.locker.Unlock()
	l.Lock()
}

func (m *LockerMap) Unlock(k string) {
	m.locker.Lock()
	l := m.store[k]
	l.count--
	if l.count == 0 && l.rwcount == 0 {
		delete(m.store, k)
	}
	m.locker.Unlock()
	l.Unlock()
}
func (m *LockerMap) RLock(k string) {
	m.locker.Lock()
	l := m.store[k]
	if l == nil {
		l = &lockerentity{}
		m.store[k] = l
	}
	l.rwcount++
	m.locker.Unlock()
	l.RLock()

}

func (m *LockerMap) RUnlock(k string) {
	m.locker.Lock()
	l := m.store[k]
	l.rwcount--
	if l.count == 0 && l.rwcount == 0 {
		delete(m.store, k)
	}
	m.locker.Unlock()
	l.RUnlock()
}

func (m *LockerMap) Locker(k string) sync.Locker {
	return &locker{
		m:   m,
		key: k,
	}
}
func (m *LockerMap) RLocker(k string) sync.Locker {
	return &rlocker{
		m:   m,
		key: k,
	}
}
func NewLockerMap() *LockerMap {
	return &LockerMap{
		store: map[string]*lockerentity{},
	}
}

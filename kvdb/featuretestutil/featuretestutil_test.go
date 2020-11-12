package featuretestutil

import (
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/herb-go/herbdata/kvdb"
)

type entiy struct {
	data    []byte
	expired *time.Time
}

type counter struct {
	data    int64
	expired *time.Time
}

type testStore struct {
	locker sync.Mutex
	kvdb.Nop
	m sync.Map
	c sync.Map
}

//Close close database
func (t *testStore) Close() error {
	return nil
}

//Set set value by given key
func (t *testStore) Set(key []byte, value []byte) error {
	t.locker.Lock()
	defer t.locker.Unlock()
	t.m.Store(string(key), &entiy{data: value})
	return nil
}

//Get get value by given key
func (t *testStore) Get(key []byte) ([]byte, error) {
	t.locker.Lock()
	defer t.locker.Unlock()
	return t.get(key)
}
func (t *testStore) get(key []byte) ([]byte, error) {
	v, ok := t.m.Load(string(key))
	if !ok {
		return nil, kvdb.ErrNotFound
	}
	e := v.(*entiy)
	if e.expired != nil && e.expired.UnixNano() < time.Now().UnixNano() {
		return nil, kvdb.ErrNotFound
	}
	return e.data, nil

}

//Del delete value by given key
func (t *testStore) Del(key []byte) error {
	t.locker.Lock()
	defer t.locker.Unlock()
	t.m.Delete(string(key))
	return nil
}

//Next return values after key not more than given limit
func (t *testStore) Next(iter []byte, limit int) (keys [][]byte, newiter []byte, err error) {
	if limit <= 0 {
		return nil, iter, nil
	}
	iterstr := string(iter)
	result := [][]byte{}
	keylist := []string{}
	t.m.Range(func(key interface{}, data interface{}) bool {
		keylist = append(keylist, key.(string))
		return true
	})
	sort.Strings(keylist)
	for _, v := range keylist {
		if v > iterstr {
			result = append(result, []byte(v))
			if limit <= len(result) {
				return result, result[len(result)-1], nil
			}
		}
	}
	return result, nil, nil
}

//SetWithTTL set value by given key and ttl
func (t *testStore) SetWithTTL(key []byte, value []byte, ttl time.Duration) error {
	t.locker.Lock()
	defer t.locker.Unlock()
	expired := time.Now().Add(ttl)
	t.m.Store(string(key), &entiy{data: value, expired: &expired})
	return nil
}

//Begin begin new transaction
func (t *testStore) Begin() (kvdb.Transaction, error) {
	return nil, kvdb.ErrFeatureNotSupported
}

//Features return supported features
func (t *testStore) Features() kvdb.Feature {
	return kvdb.FeatureStore |
		kvdb.FeatureTTLStore |
		kvdb.FeatureInsert |
		kvdb.FeatureTTLInsert |
		kvdb.FeatureUpdate |
		kvdb.FeatureTTLUpdate |
		kvdb.FeatureCounter |
		kvdb.FeatureTTLCounter |
		kvdb.FeatureNext
}

func (t *testStore) SetCounter(key []byte, value int64) error {
	t.locker.Lock()
	defer t.locker.Unlock()
	t.c.Store(string(key), &counter{data: value})
	return nil
}
func (t *testStore) SetCounterWithTTL(key []byte, value int64, ttl time.Duration) error {
	t.locker.Lock()
	defer t.locker.Unlock()
	expired := time.Now().Add(ttl)
	t.c.Store(string(key), &counter{data: value, expired: &expired})
	return nil
}
func (t *testStore) IncreaseCounter(key []byte, incr int64) (int64, error) {
	t.locker.Lock()
	defer t.locker.Unlock()
	v := t.getCounter(key)
	final := v + incr
	t.c.Store(string(key), &counter{data: final})
	return final, nil
}
func (t *testStore) IncreaseCounterWithTTL(key []byte, incr int64, ttl time.Duration) (int64, error) {
	t.locker.Lock()
	defer t.locker.Unlock()
	v := t.getCounter(key)
	expired := time.Now().Add(ttl)
	final := v + incr
	t.c.Store(string(key), &counter{data: final, expired: &expired})
	return final, nil
}

func (t *testStore) GetCounter(key []byte) (int64, error) {
	t.locker.Lock()
	defer t.locker.Unlock()
	return t.getCounter(key), nil
}
func (t *testStore) getCounter(key []byte) int64 {
	v, ok := t.c.Load(string(key))
	if !ok {
		return 0
	}
	e := v.(*counter)
	if e.expired != nil && e.expired.UnixNano() < time.Now().UnixNano() {
		return 0
	}
	return e.data
}
func (t *testStore) DelCounter(key []byte) error {
	t.locker.Lock()
	defer t.locker.Unlock()
	t.c.Delete(string(key))
	return nil
}

func (t *testStore) Insert(key []byte, value []byte) (bool, error) {
	t.locker.Lock()
	defer t.locker.Unlock()
	_, err := t.get(key)
	if err == nil {
		return false, nil
	} else if err != kvdb.ErrNotFound {
		return false, err
	}
	t.m.Store(string(key), &entiy{data: value})
	return true, nil
}

func (t *testStore) InsertWithTTL(key []byte, value []byte, ttl time.Duration) (bool, error) {
	t.locker.Lock()
	defer t.locker.Unlock()
	_, err := t.get(key)
	if err == nil {
		return false, nil
	} else if err != kvdb.ErrNotFound {
		return false, err
	}
	expired := time.Now().Add(ttl)
	t.m.Store(string(key), &entiy{data: value, expired: &expired})
	return true, nil
}
func (t *testStore) Update(key []byte, value []byte) (bool, error) {
	t.locker.Lock()
	defer t.locker.Unlock()
	_, err := t.get(key)
	if err == kvdb.ErrNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}
	t.m.Store(string(key), &entiy{data: value})
	return true, nil
}
func (t *testStore) UpdateWithTTL(key []byte, value []byte, ttl time.Duration) (bool, error) {
	t.locker.Lock()
	defer t.locker.Unlock()
	_, err := t.get(key)
	if err == kvdb.ErrNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}
	expired := time.Now().Add(ttl)
	t.m.Store(string(key), &entiy{data: value, expired: &expired})
	return true, nil
}
func TestTester(t *testing.T) {
	var result interface{}
	tester := &Tester{
		Hanlder: func(v ...interface{}) {
			result = v[0]
		},
	}
	tester.Assert(true, "fatal")
	if result != nil {
		t.Fatal()
	}
	tester.Assert(false, "fatal")
	if result != "fatal" {
		t.Fatal()
	}
}

func TestUtil(t *testing.T) {
	TestDriver(func() kvdb.Driver { return &testStore{} }, t.Fatal)
}

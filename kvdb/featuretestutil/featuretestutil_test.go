package featuretestutil

import (
	"fmt"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/herb-go/herbdata"
	"github.com/herb-go/herbdata/kvdb"
)

type entiy struct {
	data    []byte
	expired *int64
}

type counter struct {
	data    int64
	expired *int64
}

type transactionEnity struct {
	entiy
	deleted bool
}

type testTransaction struct {
	driver *testStore
	locker sync.Mutex
	data   map[string]*transactionEnity
}

func (t *testTransaction) Rollback() error {
	t.locker.Lock()
	defer t.locker.Unlock()
	t.data = map[string]*transactionEnity{}
	return nil
}
func (t *testTransaction) Commit() error {
	t.locker.Lock()
	defer t.locker.Unlock()
	t.driver.locker.Lock()
	defer t.driver.locker.Unlock()
	for k, v := range t.data {
		if v.deleted {
			t.driver.m.Delete(k)
		} else {
			t.driver.m.Store(k, &v.entiy)
		}
	}
	return nil
}

func (t *testTransaction) IsolationLevel() kvdb.IsolationLevel {
	return kvdb.IsolationLevelBatch
}

//Set set value by given key
func (t *testTransaction) Set(key []byte, value []byte) error {
	t.locker.Lock()
	defer t.locker.Unlock()
	t.data[string(key)] = &transactionEnity{
		entiy: entiy{
			data: value,
		},
	}
	return nil
}

//Get get value by given key
func (t *testTransaction) Get(key []byte) ([]byte, error) {
	t.locker.Lock()
	defer t.locker.Unlock()
	v, ok := t.data[string(key)]
	if !ok {
		return t.driver.Get(key)
	}
	if v.deleted {
		return nil, herbdata.ErrNotFound
	}
	if v.entiy.expired != nil && *v.entiy.expired < time.Now().Unix() {
		return nil, herbdata.ErrNotFound
	}
	return v.entiy.data, nil
}

//Delete delete value by given key
func (t *testTransaction) Delete(key []byte) error {
	t.locker.Lock()
	defer t.locker.Unlock()
	t.data[string(key)] = &transactionEnity{deleted: true}
	return nil
}
func (t *testTransaction) SetWithTTL(key []byte, value []byte, ttl int64) error {
	t.locker.Lock()
	defer t.locker.Unlock()
	if ttl <= 0 {
		return herbdata.ErrInvalidatedTTL
	}
	expired := time.Now().Unix() + ttl
	t.data[string(key)] = &transactionEnity{
		entiy: entiy{
			data:    value,
			expired: &expired,
		},
	}
	return nil

}

func newTestTransaction() *testTransaction {
	return &testTransaction{
		data: map[string]*transactionEnity{},
	}
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
		return nil, herbdata.ErrNotFound
	}
	e := v.(*entiy)
	if e.expired != nil && *e.expired < time.Now().Unix() {
		return nil, herbdata.ErrNotFound
	}
	return e.data, nil

}

//Delete delete value by given key
func (t *testStore) Delete(key []byte) error {
	t.locker.Lock()
	defer t.locker.Unlock()
	t.m.Delete(string(key))
	return nil
}

//Next return values after key not more than given limit
func (t *testStore) Next(iter []byte, limit int) (kv []*herbdata.KeyValue, newiter []byte, err error) {
	if limit <= 0 {
		return nil, nil, kvdb.ErrUnsupportedNextLimit
	}
	t.locker.Lock()
	defer t.locker.Unlock()
	iterstr := string(iter)
	result := []*herbdata.KeyValue{}
	keylist := []string{}
	t.m.Range(func(key interface{}, data interface{}) bool {
		keylist = append(keylist, key.(string))
		return true
	})
	sort.Strings(keylist)
	for _, v := range keylist {
		if iterstr == "" || v > iterstr {
			result = append(result, &herbdata.KeyValue{Key: []byte(v), Value: []byte(v)})
			if limit <= len(result) {
				return result, result[len(result)-1].Key, nil
			}
		}
	}
	return result, nil, nil
}

//Prev return values before key not more than given limit
func (t *testStore) Prev(iter []byte, limit int) (kv []*herbdata.KeyValue, newiter []byte, err error) {
	if limit <= 0 {
		return nil, nil, kvdb.ErrUnsupportedNextLimit
	}
	t.locker.Lock()
	defer t.locker.Unlock()
	iterstr := string(iter)
	result := []*herbdata.KeyValue{}
	keylist := []string{}
	t.m.Range(func(key interface{}, data interface{}) bool {
		keylist = append(keylist, key.(string))
		return true
	})
	sort.Sort(sort.Reverse(sort.StringSlice(keylist)))
	for _, v := range keylist {
		if iterstr == "" || v < iterstr {
			result = append(result, &herbdata.KeyValue{Key: []byte(v), Value: []byte(v)})
			if limit <= len(result) {
				return result, result[len(result)-1].Key, nil
			}
		}
	}
	return result, nil, nil
}

//SetWithTTL set value by given key and ttl
func (t *testStore) SetWithTTL(key []byte, value []byte, ttl int64) error {
	t.locker.Lock()
	defer t.locker.Unlock()
	if ttl <= 0 {
		return herbdata.ErrInvalidatedTTL
	}
	expired := time.Now().Unix() + ttl
	t.m.Store(string(key), &entiy{data: value, expired: &expired})
	return nil
}

//Begin begin new transaction
func (t *testStore) Begin() (kvdb.Transaction, error) {
	trans := newTestTransaction()
	trans.driver = t
	return trans, nil
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
		kvdb.FeatureNext |
		kvdb.FeaturePrev |
		kvdb.FeatureTransaction
}

func (t *testStore) SetCounter(key []byte, value int64) error {
	t.locker.Lock()
	defer t.locker.Unlock()
	t.c.Store(string(key), &counter{data: value})
	return nil
}
func (t *testStore) SetCounterWithTTL(key []byte, value int64, ttl int64) error {
	if ttl <= 0 {
		return herbdata.ErrInvalidatedTTL
	}
	t.locker.Lock()
	defer t.locker.Unlock()
	expired := time.Now().Unix() + ttl

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
func (t *testStore) IncreaseCounterWithTTL(key []byte, incr int64, ttl int64) (int64, error) {
	if ttl <= 0 {
		return 0, herbdata.ErrInvalidatedTTL
	}
	t.locker.Lock()
	defer t.locker.Unlock()
	v := t.getCounter(key)
	expired := time.Now().Unix() + ttl
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
	if e.expired != nil && *e.expired < time.Now().Unix() {
		return 0
	}
	return e.data
}
func (t *testStore) DeleteCounter(key []byte) error {
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
	} else if err != herbdata.ErrNotFound {
		return false, err
	}
	t.m.Store(string(key), &entiy{data: value})
	return true, nil
}

func (t *testStore) InsertWithTTL(key []byte, value []byte, ttl int64) (bool, error) {
	if ttl <= 0 {
		return false, herbdata.ErrInvalidatedTTL
	}
	t.locker.Lock()
	defer t.locker.Unlock()
	_, err := t.get(key)
	if err == nil {
		return false, nil
	} else if err != herbdata.ErrNotFound {
		return false, err
	}
	expired := time.Now().Unix() + ttl
	t.m.Store(string(key), &entiy{data: value, expired: &expired})
	return true, nil
}
func (t *testStore) Update(key []byte, value []byte) (bool, error) {
	t.locker.Lock()
	defer t.locker.Unlock()
	_, err := t.get(key)
	if err == herbdata.ErrNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}
	t.m.Store(string(key), &entiy{data: value})
	return true, nil
}
func (t *testStore) UpdateWithTTL(key []byte, value []byte, ttl int64) (bool, error) {
	if ttl <= 0 {
		return false, herbdata.ErrInvalidatedTTL
	}
	t.locker.Lock()
	defer t.locker.Unlock()
	_, err := t.get(key)
	if err == herbdata.ErrNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}
	expired := time.Now().Unix() + ttl
	t.m.Store(string(key), &entiy{data: value, expired: &expired})
	return true, nil
}

func (t *testStore) IsolationLevel() kvdb.IsolationLevel {
	return kvdb.IsolationLevelBatch
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
	TestDriver(func() kvdb.Driver { return &testStore{} }, func(args ...interface{}) { fmt.Println(args...); panic("fatal") })
}

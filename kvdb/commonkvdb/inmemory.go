package commonkvdb

import (
	"sort"
	"sync"

	"github.com/herb-go/herbdata"
	"github.com/herb-go/herbdata/kvdb"
)

//InMemory in memory key-value database driver
type InMemory struct {
	kvdb.Nop
	values        sync.Map
	counterlocker sync.Mutex
	counters      map[string]int64
}

//Set set value by given key
func (i *InMemory) Set(key []byte, value []byte) error {
	i.values.Store(string(key), value)
	return nil
}

//Get get value by given key
func (i *InMemory) Get(key []byte) ([]byte, error) {
	v, ok := i.values.Load(string(key))
	if !ok {
		return nil, herbdata.ErrNotFound
	}
	return v.([]byte), nil
}

//Delete delete value by given key
func (i *InMemory) Delete(key []byte) error {
	i.values.Delete(string(key))
	return nil
}

//Features return supported features
func (i *InMemory) Features() kvdb.Feature {
	return kvdb.FeatureStable |
		kvdb.FeatureStore |
		kvdb.FeatureCounter |
		kvdb.FeatureNext |
		kvdb.FeatureEmbedded
}

//SetCounter set counter value with given key
func (i *InMemory) SetCounter(key []byte, value int64) error {
	i.counterlocker.Lock()
	defer i.counterlocker.Unlock()
	i.counters[string(key)] = value
	return nil
}

//IncreaseCounter increace counter value with given key and increasement.
//Value not existed coutn as 0.
//Return final value and any error if raised.
func (i *InMemory) IncreaseCounter(key []byte, incr int64) (int64, error) {
	i.counterlocker.Lock()
	defer i.counterlocker.Unlock()
	v := i.counters[string(key)]
	v = v + incr
	i.counters[string(key)] = v
	return v, nil
}

//GetCounter get counter value with given key
//Value not existed coutn as 0.
func (i *InMemory) GetCounter(key []byte) (int64, error) {
	i.counterlocker.Lock()
	defer i.counterlocker.Unlock()
	return i.counters[string(key)], nil
}

//DeleteCounter delete counter value with given key
func (i *InMemory) DeleteCounter(key []byte) error {
	i.counterlocker.Lock()
	defer i.counterlocker.Unlock()
	delete(i.counters, string(key))
	return nil
}

//Next return values after key not more than given limit
func (i *InMemory) Next(iter []byte, limit int) (kv []*herbdata.KeyValue, newiter []byte, err error) {
	if limit <= 0 {
		return nil, nil, kvdb.ErrUnsupportedNextLimit
	}
	iterstr := string(iter)
	result := []*herbdata.KeyValue{}
	keylist := []string{}
	i.values.Range(func(key interface{}, data interface{}) bool {
		keylist = append(keylist, key.(string))
		return true
	})
	sort.Strings(keylist)
	for _, v := range keylist {
		if v > iterstr {
			data, _ := i.values.Load(string(v))
			result = append(result, &herbdata.KeyValue{Key: []byte(v), Value: data.([]byte)})
			if limit <= len(result) {
				return result, result[len(result)-1].Key, nil
			}
		}
	}
	return result, nil, nil
}

//NewInMemory create new in memory key-value database driver
func NewInMemory() *InMemory {
	return &InMemory{
		counters: map[string]int64{},
	}
}

//InMemoryFactory in-memory driver factory
func InMemoryFactory(loader func(v interface{}) error) (kvdb.Driver, error) {
	return NewInMemory(), nil
}

func init() {
	kvdb.Register("inmemory", InMemoryFactory)
}

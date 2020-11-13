package featuretestutil

import (
	"bytes"
	"sort"
	"time"

	"github.com/herb-go/herbdata"
	"github.com/herb-go/herbdata/kvdb"
)

//KeyNotfound test key not found
var KeyNotfound = []byte("not found")

//KeySuccess test key success
var KeySuccess = []byte("key success")

//DataSuccess test data success
var DataSuccess = []byte("data success")

//DataUpdated test data updated
var DataUpdated = []byte("data updated")

//DataCounterSuccess test counter data success
const DataCounterSuccess = int64(1)

//DataCounterStep test counter data step
const DataCounterStep = int64(2)

//DataCounterUpdated test counter data updated
const DataCounterUpdated = int64(3)

//KeyListForNext test key list for next
var KeyListForNext = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

//CounterKeyListForNext test counter key list for next
var CounterKeyListForNext = []string{"1", "2", "3", "4", "5"}

//TransactionKeySet test transaction key for Set
var TransactionKeySet = []byte("set")

// TransactionKeySetTTL test transaction key for SetTTL
var TransactionKeySetTTL = []byte("setttl")

// TransactionKeyDelete test transaction key for Delete
var TransactionKeyDelete = []byte("delete")

//Tester tester struct
type Tester struct {
	Hanlder func(...interface{})
}

//Assert assert of ok is true.
//Call hanlder with given args
func (t *Tester) Assert(ok bool, args ...interface{}) {
	if !ok {
		t.Hanlder(args...)
	}
}

//TestDriver test kvdb driver
func TestDriver(creator func() kvdb.Driver, fatal func(...interface{})) {
	t := &Tester{Hanlder: fatal}
	TestFeatureStore(creator(), t)
	TestFeatureTTLStore(creator(), t)
	TestFeatureStoreAndFeatureTTLStore(creator(), t)
	TestFeatureCounter(creator(), t)
	TestFeatureTTLCounter(creator(), t)
	TestFeatureCounterAndFeatureTTLCounter(creator(), t)
	TestFeatureStoreAndFeatureCounter(creator(), t)
	TestFeatureTTLStoreAndFeatureTTLCounter(creator(), t)
	TestFeatureNext(creator(), t)
	TestFeatureInsert(creator(), t)
	TestFeatureTTLInsert(creator(), t)
	TestFeatureUpdate(creator(), t)
	TestFeatureTTLUpdate(creator(), t)
	TestFeatureTransaction(creator(), t)
}

//TestFeatureStore test driver FeatureStore
func TestFeatureStore(driver kvdb.Driver, t *Tester) {
	var err error
	var data []byte
	if driver.Features().SupportAll(kvdb.FeatureStore) {
		_, err = driver.Get(KeyNotfound)
		t.Assert(err == herbdata.ErrNotFound, err)
		err = driver.Delete(KeyNotfound)
		t.Assert(err == nil, err)
		err = driver.Set(KeySuccess, DataSuccess)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		err = driver.Set(KeySuccess, DataUpdated)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataUpdated) == 0, data, err)
		err = driver.Delete(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == herbdata.ErrNotFound, err)
	}
}

//TestFeatureTTLStore test driver FeatureTTLStore
func TestFeatureTTLStore(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureTTLStore) {
		var err error
		var data []byte
		_, err = driver.Get(KeyNotfound)
		t.Assert(err == herbdata.ErrNotFound, err)
		err = driver.Delete(KeyNotfound)
		t.Assert(err == nil, err)
		err = driver.SetWithTTL(KeySuccess, DataSuccess, 3600)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		time.Sleep(time.Second)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		err = driver.SetWithTTL(KeySuccess, DataSuccess, 1)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		time.Sleep(2 * time.Second)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == herbdata.ErrNotFound, err)
		err = driver.SetWithTTL(KeySuccess, DataSuccess, 3600)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		time.Sleep(time.Second)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		err = driver.Delete(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == herbdata.ErrNotFound, err)
		err = driver.SetWithTTL(KeySuccess, DataSuccess, -1)
		t.Assert(err == herbdata.ErrInvalidatedTTL, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == herbdata.ErrNotFound, err)
	}
}

//TestFeatureStoreAndFeatureTTLStore test driver FeatureTTLStore and FeatureStore
func TestFeatureStoreAndFeatureTTLStore(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureStore | kvdb.FeatureTTLStore) {
		var err error
		var data []byte
		_, err = driver.Get(KeyNotfound)
		t.Assert(err == herbdata.ErrNotFound, err)
		err = driver.Delete(KeyNotfound)
		t.Assert(err == nil, err)
		err = driver.Set(KeySuccess, DataSuccess)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		err = driver.SetWithTTL(KeySuccess, DataUpdated, 1)
		t.Assert(err == nil, err)
		time.Sleep(time.Second)
		data, err = driver.Get(KeySuccess)
		_, err = driver.Get(KeyNotfound)
		t.Assert(err == herbdata.ErrNotFound, err)
	}
}

//TestFeatureCounter test driver FeatureCounter
func TestFeatureCounter(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureCounter) {
		var err error
		var data int64
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == 0 && err == nil, data, err)
		err = driver.DeleteCounter(KeySuccess)
		t.Assert(err == nil, err)
		err = driver.SetCounter(KeySuccess, DataCounterSuccess)
		t.Assert(err == nil, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(err == nil && data == DataCounterSuccess, data, err)
		err = driver.DeleteCounter(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == 0 && err == nil, data, err)
		data, err = driver.IncreaseCounter(KeySuccess, DataCounterStep)
		t.Assert(data == DataCounterStep && err == nil, data, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == DataCounterStep && err == nil, data, err)
		err = driver.DeleteCounter(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == 0 && err == nil, data, err)
		err = driver.SetCounter(KeySuccess, DataCounterSuccess)
		t.Assert(err == nil, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(err == nil && data == DataCounterSuccess, data, err)
		data, err = driver.IncreaseCounter(KeySuccess, DataCounterStep)
		t.Assert(data == DataCounterUpdated && err == nil, data, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == DataCounterUpdated && err == nil, data, err)

	}
}

//TestFeatureTTLCounter test driver FeatureTTLCounter
func TestFeatureTTLCounter(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureTTLCounter) {
		var err error
		var data int64
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == 0 && err == nil, data, err)
		err = driver.SetCounterWithTTL(KeySuccess, DataCounterSuccess, -1)
		t.Assert(err == herbdata.ErrInvalidatedTTL, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == 0 && err == nil, data, err)
		data, err = driver.IncreaseCounterWithTTL(KeySuccess, DataCounterSuccess, -1)
		t.Assert(data == 0 && err == herbdata.ErrInvalidatedTTL, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == 0 && err == nil, data, err)
		err = driver.DeleteCounter(KeySuccess)
		t.Assert(err == nil, err)
		err = driver.SetCounterWithTTL(KeySuccess, DataCounterSuccess, 3600)
		t.Assert(err == nil, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(err == nil && data == DataCounterSuccess, data, err)
		err = driver.DeleteCounter(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == 0 && err == nil, data, err)
		data, err = driver.IncreaseCounterWithTTL(KeySuccess, DataCounterStep, 3600)
		t.Assert(data == DataCounterStep && err == nil, data, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == DataCounterStep && err == nil, data, err)
		err = driver.DeleteCounter(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == 0 && err == nil, data, err)
		err = driver.SetCounterWithTTL(KeySuccess, DataCounterSuccess, 3600)
		t.Assert(err == nil, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(err == nil && data == DataCounterSuccess, data, err)
		data, err = driver.IncreaseCounterWithTTL(KeySuccess, DataCounterStep, 3600)
		t.Assert(data == DataCounterUpdated && err == nil, data, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == DataCounterUpdated && err == nil, data, err)
		err = driver.DeleteCounter(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == 0 && err == nil, data, err)
		err = driver.SetCounterWithTTL(KeySuccess, DataCounterSuccess, 1)
		t.Assert(err == nil, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == DataCounterSuccess && err == nil, data, err)
		time.Sleep(2 * time.Second)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == 0 && err == nil, data, err)
		err = driver.SetCounterWithTTL(KeySuccess, DataCounterSuccess, 1)
		t.Assert(err == nil, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == DataCounterSuccess && err == nil, data, err)
		data, err = driver.IncreaseCounterWithTTL(KeySuccess, DataCounterStep, 1)
		t.Assert(data == DataCounterUpdated && err == nil, data, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == DataCounterUpdated && err == nil, data, err)
		time.Sleep(2 * time.Second)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == 0 && err == nil, data, err)
	}
}

//TestFeatureCounterAndFeatureTTLCounter test driver FeatureCounter and FeatureTTLCounter
func TestFeatureCounterAndFeatureTTLCounter(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureCounter | kvdb.FeatureTTLCounter) {
		var err error
		var data int64
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == 0 && err == nil, data, err)
		err = driver.SetCounter(KeySuccess, DataCounterSuccess)
		t.Assert(err == nil, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == DataCounterSuccess && err == nil, data, err)
		data, err = driver.IncreaseCounterWithTTL(KeySuccess, DataCounterStep, 3600)
		t.Assert(data == DataCounterUpdated && err == nil, data, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == DataCounterUpdated && err == nil, data, err)
		err = driver.DeleteCounter(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == 0 && err == nil, data, err)
		err = driver.SetCounterWithTTL(KeySuccess, DataCounterSuccess, 3600)
		t.Assert(err == nil, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == DataCounterSuccess && err == nil, data, err)
		data, err = driver.IncreaseCounter(KeySuccess, DataCounterStep)
		t.Assert(data == DataCounterUpdated && err == nil, data, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == DataCounterUpdated && err == nil, data, err)

	}
}

//TestFeatureStoreAndFeatureCounter test driver FeatureStore and FeatureCounter
func TestFeatureStoreAndFeatureCounter(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureStore | kvdb.FeatureCounter) {
		var data []byte
		var datacounter int64
		var err error
		data, err = driver.Get(KeySuccess)
		t.Assert(err == herbdata.ErrNotFound, err)
		err = driver.Set(KeySuccess, DataSuccess)
		t.Assert(err == nil, err)
		err = driver.SetCounter(KeySuccess, DataCounterSuccess)
		t.Assert(err == nil, err)
		datacounter, err = driver.IncreaseCounter(KeySuccess, DataCounterStep)
		t.Assert(datacounter == DataCounterUpdated, err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		datacounter, err = driver.GetCounter(KeySuccess)
		t.Assert(datacounter == DataCounterUpdated, err == nil, err)
	}
}

//TestFeatureTTLStoreAndFeatureTTLCounter test driver FeatureTTLStore and FeatureTTLCounter
func TestFeatureTTLStoreAndFeatureTTLCounter(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureTTLStore | kvdb.FeatureTTLCounter) {
		var data []byte
		var datacounter int64
		var err error
		data, err = driver.Get(KeySuccess)
		t.Assert(err == herbdata.ErrNotFound, err)
		err = driver.SetWithTTL(KeySuccess, DataSuccess, 3600)
		t.Assert(err == nil, err)
		err = driver.SetCounterWithTTL(KeySuccess, DataCounterSuccess, 3600)
		t.Assert(err == nil, err)
		datacounter, err = driver.IncreaseCounterWithTTL(KeySuccess, DataCounterStep, 3600)
		t.Assert(datacounter == DataCounterUpdated, err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		datacounter, err = driver.GetCounter(KeySuccess)
		t.Assert(datacounter == DataCounterUpdated, err == nil, err)

	}
}

//TestFeatureNext test driver FeatureNext
func TestFeatureNext(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureNext) && driver.Features().SupportAny(kvdb.FeatureStore|kvdb.FeatureTTLStore) {
		var err error
		var keys [][]byte
		_, _, err = driver.Next([]byte{}, 0)
		t.Assert(err == kvdb.ErrUnsupportedNextLimit, err)
		_, _, err = driver.Next([]byte{}, -1)
		t.Assert(err == kvdb.ErrUnsupportedNextLimit, err)
		for _, v := range KeyListForNext {
			if driver.Features().SupportAny(kvdb.FeatureStore) {
				err = driver.Set([]byte(v), DataSuccess)
				t.Assert(err == nil, err)
			}
			if driver.Features().SupportAny(kvdb.FeatureTTLStore) {
				err = driver.SetWithTTL([]byte(v), DataSuccess, 3600)
				t.Assert(err == nil, err)
			}
		}
		for _, v := range CounterKeyListForNext {
			if driver.Features().SupportAny(kvdb.FeatureCounter) {
				err = driver.SetCounter([]byte(v), DataCounterSuccess)
				t.Assert(err == nil, err)
			}
			if driver.Features().SupportAny(kvdb.FeatureTTLCounter) {
				err = driver.SetCounterWithTTL([]byte(v), DataCounterSuccess, 3600)
				t.Assert(err == nil, err)
			}
		}
		result := []string{}
		iter := []byte{}
		for {
			keys, iter, err = driver.Next(iter, 3)
			t.Assert(err == nil, err)
			t.Assert(len(iter) == 0 || len(keys) == 3, keys, iter)
			for _, v := range keys {
				result = append(result, string(v))
			}
			if len(iter) == 0 {
				break
			}
		}
		target := make([]string, len(KeyListForNext))
		copy(target, KeyListForNext)
		sort.Strings(target)
		sort.Strings(result)
		t.Assert(len(target) == len(result), target, result)
		for k := range result {
			t.Assert(result[k] == target[k])
		}
	}
}

//TestFeatureInsert test driver FeatureInsert
func TestFeatureInsert(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureInsert) {
		var err error
		var data []byte
		var ok bool
		err = driver.Set(KeySuccess, DataSuccess)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		ok, err = driver.Insert(KeySuccess, DataUpdated)
		t.Assert(ok == false && err == nil, ok, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		err = driver.Delete(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == herbdata.ErrNotFound, data, err)
		ok, err = driver.Insert(KeySuccess, DataUpdated)
		t.Assert(ok == true && err == nil, ok, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataUpdated) == 0, data, err)
	}
}

//TestFeatureTTLInsert test driver FeatureInsert
func TestFeatureTTLInsert(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureTTLInsert) {
		var err error
		var data []byte
		var ok bool

		err = driver.SetWithTTL(KeySuccess, DataSuccess, 3600)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		ok, err = driver.InsertWithTTL(KeySuccess, DataUpdated, 3600)
		t.Assert(ok == false && err == nil, ok, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		err = driver.Delete(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == herbdata.ErrNotFound, data, err)
		ok, err = driver.InsertWithTTL(KeySuccess, DataUpdated, 3600)
		t.Assert(ok == true && err == nil, ok, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataUpdated) == 0, data, err)
		err = driver.Delete(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == herbdata.ErrNotFound, data, err)
		ok, err = driver.InsertWithTTL(KeySuccess, DataUpdated, 1)
		t.Assert(ok == true && err == nil, ok, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataUpdated) == 0, data, err)
		t.Assert(err == nil && bytes.Compare(data, DataUpdated) == 0, data, err)
		time.Sleep(2 * time.Second)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == herbdata.ErrNotFound, data, err)

		ok, err = driver.InsertWithTTL(KeySuccess, DataSuccess, -1)
		t.Assert(ok == false && err == herbdata.ErrInvalidatedTTL, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == herbdata.ErrNotFound, data, err)

	}
}

//TestFeatureUpdate test driver FeatureUpdate
func TestFeatureUpdate(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureUpdate) {
		var err error
		var data []byte
		var ok bool
		err = driver.Set(KeySuccess, DataSuccess)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		ok, err = driver.Update(KeySuccess, DataUpdated)
		t.Assert(ok == true && err == nil, ok, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataUpdated) == 0, data, err)
		err = driver.Delete(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == herbdata.ErrNotFound, data, err)
		ok, err = driver.Update(KeySuccess, DataUpdated)
		t.Assert(ok == false && err == nil, ok, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == herbdata.ErrNotFound, data, err)
	}
}

//TestFeatureTTLUpdate test driver FeatureTTLUpdate
func TestFeatureTTLUpdate(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureUpdate) {
		var err error
		var data []byte
		var ok bool
		err = driver.SetWithTTL(KeySuccess, DataSuccess, 3600)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		ok, err = driver.UpdateWithTTL(KeySuccess, DataUpdated, 3600)
		t.Assert(ok == true && err == nil, ok, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataUpdated) == 0, data, err)
		err = driver.Delete(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == herbdata.ErrNotFound, data, err)
		ok, err = driver.UpdateWithTTL(KeySuccess, DataUpdated, 3600)
		t.Assert(ok == false && err == nil, ok, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == herbdata.ErrNotFound, data, err)
		err = driver.SetWithTTL(KeySuccess, DataSuccess, 3600)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		ok, err = driver.UpdateWithTTL(KeySuccess, DataUpdated, 1)
		t.Assert(ok == true && err == nil, ok, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataUpdated) == 0, data, err)
		time.Sleep(2 * time.Second)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == herbdata.ErrNotFound, data, err)

		ok, err = driver.UpdateWithTTL(KeySuccess, DataSuccess, -1)
		t.Assert(ok == false && err == herbdata.ErrInvalidatedTTL, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == herbdata.ErrNotFound, data, err)

	}
}

//TestFeatureTransaction test driver FeatureTransaction
func TestFeatureTransaction(driver kvdb.Driver, t *Tester) {
	var err error
	var data []byte
	var trans kvdb.Transaction
	if driver.Features().SupportAll(kvdb.FeatureTransaction) && driver.Features().SupportAny(kvdb.FeatureStore|kvdb.FeatureTTLStore) {
		t.Assert(driver.IsolationLevel() != 0, driver.IsolationLevel())
		err = driver.Set(TransactionKeyDelete, DataSuccess)
		t.Assert(err == nil, err)
		if driver.Features().SupportAny(kvdb.FeatureStore) {
			data, err = driver.Get(TransactionKeySet)
			t.Assert(err == herbdata.ErrNotFound, data, err)
		}
		if driver.Features().SupportAny(kvdb.FeatureTTLStore) {
			data, err = driver.Get(TransactionKeySetTTL)
			t.Assert(err == herbdata.ErrNotFound, data, err)
		}
		data, err = driver.Get(TransactionKeyDelete)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		trans, err = driver.Begin()
		t.Assert(err == nil, err)
		t.Assert(driver.IsolationLevel() == trans.IsolationLevel(), driver.IsolationLevel(), trans.IsolationLevel())
		if driver.Features().SupportAny(kvdb.FeatureStore) {
			err = trans.Set(TransactionKeySet, DataSuccess)
			t.Assert(err == nil, err)
		}
		if driver.Features().SupportAny(kvdb.FeatureTTLStore) {
			err = trans.SetWithTTL(TransactionKeySetTTL, DataSuccess, 3600)
			t.Assert(err == nil, err)
		}
		err = trans.Delete(TransactionKeyDelete)
		t.Assert(err == nil, err)
		err = trans.Commit()
		t.Assert(err == nil, err)
		if driver.Features().SupportAny(kvdb.FeatureStore) {
			data, err = driver.Get(TransactionKeySet)
			t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		}
		if driver.Features().SupportAny(kvdb.FeatureTTLStore) {
			data, err = driver.Get(TransactionKeySetTTL)
			t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		}
		data, err = driver.Get(TransactionKeyDelete)
		t.Assert(err == herbdata.ErrNotFound, data, err)

		err = driver.Delete(TransactionKeySet)
		t.Assert(err == nil, err)
		err = driver.Delete(TransactionKeySetTTL)
		t.Assert(err == nil, err)
		err = driver.Set(TransactionKeyDelete, DataSuccess)
		t.Assert(err == nil, err)

		err = driver.Set(TransactionKeyDelete, DataSuccess)
		t.Assert(err == nil, err)
		if driver.Features().SupportAny(kvdb.FeatureStore) {
			data, err = driver.Get(TransactionKeySet)
			t.Assert(err == herbdata.ErrNotFound, data, err)
		}
		if driver.Features().SupportAny(kvdb.FeatureTTLStore) {
			data, err = driver.Get(TransactionKeySetTTL)
			t.Assert(err == herbdata.ErrNotFound, data, err)
		}
		data, err = driver.Get(TransactionKeyDelete)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		trans, err = driver.Begin()
		t.Assert(err == nil, err)
		if driver.Features().SupportAny(kvdb.FeatureStore) {
			err = trans.Set(TransactionKeySet, DataSuccess)
			t.Assert(err == nil, err)
		}
		if driver.Features().SupportAny(kvdb.FeatureTTLStore) {
			err = trans.SetWithTTL(TransactionKeySetTTL, DataSuccess, 3600)
			t.Assert(err == nil, err)
		}
		err = trans.Delete(TransactionKeyDelete)
		t.Assert(err == nil, err)
		err = trans.Rollback()
		t.Assert(err == nil, err)
		if driver.Features().SupportAny(kvdb.FeatureStore) {
			data, err = driver.Get(TransactionKeySet)
			t.Assert(err == herbdata.ErrNotFound, data, err)
		}
		if driver.Features().SupportAny(kvdb.FeatureTTLStore) {
			data, err = driver.Get(TransactionKeySetTTL)
			t.Assert(err == herbdata.ErrNotFound, data, err)
		}
		data, err = driver.Get(TransactionKeyDelete)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)

	}
}

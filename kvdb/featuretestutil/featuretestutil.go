package featuretestutil

import (
	"bytes"
	"fmt"
	"time"

	"github.com/herb-go/herbdata/kvdb"
)

var KeyNotfound = []byte("not found")
var KeySuccess = []byte("key success")
var DataSuccess = []byte("data success")
var DataUpdated = []byte("data updated")

const DataCounterSuccess = int64(1)
const DataCounterStep = int64(2)
const DataCounterUpdated = int64(3)

type Tester struct {
	Hanlder func(...interface{})
}

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
		t.Assert(err == kvdb.ErrNotFound, err)
		err = driver.Del(KeyNotfound)
		t.Assert(err == nil, err)
		err = driver.Set(KeySuccess, DataSuccess)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		err = driver.Set(KeySuccess, DataUpdated)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataUpdated) == 0, data, err)
		err = driver.Del(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == kvdb.ErrNotFound, err)
	}
}

//TestFeatureTTLStore test driver FeatureTTLStore
func TestFeatureTTLStore(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureTTLStore) {
		var err error
		var data []byte
		_, err = driver.Get(KeyNotfound)
		t.Assert(err == kvdb.ErrNotFound, err)
		err = driver.Del(KeyNotfound)
		t.Assert(err == nil, err)
		err = driver.SetWithTTL(KeySuccess, DataSuccess, time.Second)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		time.Sleep(time.Microsecond)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		err = driver.SetWithTTL(KeySuccess, DataSuccess, time.Millisecond)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		fmt.Println(time.Now().UnixNano())
		time.Sleep(2 * time.Millisecond)
		fmt.Println(time.Now().UnixNano())
		data, err = driver.Get(KeySuccess)
		t.Assert(err == kvdb.ErrNotFound, err)
		err = driver.SetWithTTL(KeySuccess, DataSuccess, time.Second)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		time.Sleep(time.Microsecond)
		err = driver.SetWithTTL(KeySuccess, DataSuccess, -time.Millisecond)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == kvdb.ErrNotFound, err)
	}
}

//TestFeatureStoreAndFeatureTTLStore test driver FeatureTTLStore and FeatureStore
func TestFeatureStoreAndFeatureTTLStore(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureStore | kvdb.FeatureTTLStore) {
		var err error
		var data []byte
		_, err = driver.Get(KeyNotfound)
		t.Assert(err == kvdb.ErrNotFound, err)
		err = driver.Del(KeyNotfound)
		t.Assert(err == nil, err)
		err = driver.Set(KeySuccess, DataSuccess)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		err = driver.SetWithTTL(KeySuccess, DataUpdated, time.Millisecond)
		t.Assert(err == nil, err)
		time.Sleep(time.Microsecond)
		data, err = driver.Get(KeySuccess)
		_, err = driver.Get(KeyNotfound)
		t.Assert(err == kvdb.ErrNotFound, err)
	}
}

//TestFeatureCounter test driver FeatureCounter
func TestFeatureCounter(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureCounter) {
		var err error
		var data int64
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == 0 && err == nil, data, err)
		err = driver.DelCounter(KeySuccess)
		t.Assert(err == nil, err)
		err = driver.SetCounter(KeySuccess, DataCounterSuccess)
		t.Assert(err == nil, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(err == nil && data == DataCounterSuccess, data, err)
		err = driver.DelCounter(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == 0 && err == nil, data, err)
		data, err = driver.IncreaseCounter(KeySuccess, DataCounterStep)
		t.Assert(data == DataCounterStep && err == nil, data, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == DataCounterStep && err == nil, data, err)
		err = driver.DelCounter(KeySuccess)
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
		err = driver.DelCounter(KeySuccess)
		t.Assert(err == nil, err)
		err = driver.SetCounterWithTTL(KeySuccess, DataCounterSuccess, time.Second)
		t.Assert(err == nil, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(err == nil && data == DataCounterSuccess, data, err)
		err = driver.DelCounter(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == 0 && err == nil, data, err)
		data, err = driver.IncreaseCounterWithTTL(KeySuccess, DataCounterStep, time.Second)
		t.Assert(data == DataCounterStep && err == nil, data, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == DataCounterStep && err == nil, data, err)
		err = driver.DelCounter(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == 0 && err == nil, data, err)
		err = driver.SetCounterWithTTL(KeySuccess, DataCounterSuccess, time.Second)
		t.Assert(err == nil, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(err == nil && data == DataCounterSuccess, data, err)
		data, err = driver.IncreaseCounterWithTTL(KeySuccess, DataCounterStep, time.Second)
		t.Assert(data == DataCounterUpdated && err == nil, data, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == DataCounterUpdated && err == nil, data, err)
		err = driver.DelCounter(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == 0 && err == nil, data, err)
		err = driver.SetCounterWithTTL(KeySuccess, DataCounterSuccess, time.Millisecond)
		t.Assert(err == nil, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == DataCounterSuccess && err == nil, data, err)
		time.Sleep(2 * time.Millisecond)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == 0 && err == nil, data, err)
		err = driver.SetCounterWithTTL(KeySuccess, DataCounterSuccess, time.Millisecond)
		t.Assert(err == nil, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == DataCounterSuccess && err == nil, data, err)
		data, err = driver.IncreaseCounterWithTTL(KeySuccess, DataCounterStep, time.Millisecond)
		t.Assert(data == DataCounterUpdated && err == nil, data, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == DataCounterUpdated && err == nil, data, err)
		time.Sleep(2 * time.Millisecond)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == 0 && err == nil, data, err)
	}
}

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
		data, err = driver.IncreaseCounterWithTTL(KeySuccess, DataCounterStep, time.Second)
		t.Assert(data == DataCounterUpdated && err == nil, data, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == DataCounterUpdated && err == nil, data, err)
		err = driver.DelCounter(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.GetCounter(KeySuccess)
		t.Assert(data == 0 && err == nil, data, err)
		err = driver.SetCounterWithTTL(KeySuccess, DataCounterSuccess, time.Second)
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
	if driver.Features().SupportAll(kvdb.FeatureStore & kvdb.FeatureCounter) {
		var data []byte
		var datacounter int64
		var err error
		data, err = driver.Get(KeySuccess)
		t.Assert(err == kvdb.ErrNotFound, err)
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
	if driver.Features().SupportAll(kvdb.FeatureTTLStore & kvdb.FeatureTTLCounter) {
		var data []byte
		var datacounter int64
		var err error
		data, err = driver.Get(KeySuccess)
		t.Assert(err == kvdb.ErrNotFound, err)
		err = driver.SetWithTTL(KeySuccess, DataSuccess, time.Second)
		t.Assert(err == nil, err)
		err = driver.SetCounterWithTTL(KeySuccess, DataCounterSuccess, time.Second)
		t.Assert(err == nil, err)
		datacounter, err = driver.IncreaseCounterWithTTL(KeySuccess, DataCounterStep, time.Second)
		t.Assert(datacounter == DataCounterUpdated, err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		datacounter, err = driver.GetCounter(KeySuccess)
		t.Assert(datacounter == DataCounterUpdated, err == nil, err)

	}
}

//TestFeatureNext test driver FeatureNext
func TestFeatureNext(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureNext) {

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
		err = driver.Del(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == kvdb.ErrNotFound, data, err)
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
		err = driver.SetWithTTL(KeySuccess, DataSuccess, time.Second)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		ok, err = driver.InsertWithTTL(KeySuccess, DataUpdated, time.Second)
		t.Assert(ok == false && err == nil, ok, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		err = driver.Del(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == kvdb.ErrNotFound, data, err)
		ok, err = driver.InsertWithTTL(KeySuccess, DataUpdated, time.Second)
		t.Assert(ok == true && err == nil, ok, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataUpdated) == 0, data, err)
		err = driver.Del(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == kvdb.ErrNotFound, data, err)
		ok, err = driver.InsertWithTTL(KeySuccess, DataUpdated, time.Millisecond)
		t.Assert(ok == true && err == nil, ok, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataUpdated) == 0, data, err)
		t.Assert(err == nil && bytes.Compare(data, DataUpdated) == 0, data, err)
		time.Sleep(2 * time.Millisecond)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == kvdb.ErrNotFound, data, err)
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
		err = driver.Del(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == kvdb.ErrNotFound, data, err)
		ok, err = driver.Update(KeySuccess, DataUpdated)
		t.Assert(ok == false && err == nil, ok, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == kvdb.ErrNotFound, data, err)
	}
}

//TestFeatureTTLUpdate test driver FeatureTTLUpdate
func TestFeatureTTLUpdate(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureUpdate) {
		var err error
		var data []byte
		var ok bool
		err = driver.SetWithTTL(KeySuccess, DataSuccess, time.Second)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		ok, err = driver.UpdateWithTTL(KeySuccess, DataUpdated, time.Second)
		t.Assert(ok == true && err == nil, ok, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataUpdated) == 0, data, err)
		err = driver.Del(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == kvdb.ErrNotFound, data, err)
		ok, err = driver.UpdateWithTTL(KeySuccess, DataUpdated, time.Second)
		t.Assert(ok == false && err == nil, ok, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == kvdb.ErrNotFound, data, err)
		err = driver.SetWithTTL(KeySuccess, DataSuccess, time.Second)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataSuccess) == 0, data, err)
		ok, err = driver.UpdateWithTTL(KeySuccess, DataUpdated, time.Millisecond)
		t.Assert(ok == true && err == nil, ok, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == nil && bytes.Compare(data, DataUpdated) == 0, data, err)
		time.Sleep(2 * time.Millisecond)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == kvdb.ErrNotFound, data, err)
	}
}

//TestFeatureTransaction test driver FeatureTransaction
func TestFeatureTransaction(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureTransaction) {

	}
}

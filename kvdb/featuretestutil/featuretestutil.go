package featuretestutil

import (
	"bytes"

	"github.com/herb-go/herbdata/kvdb"
)

var KeyNotfound = []byte("not found")
var KeySuccess = []byte("key success")
var DataSuccess = []byte("data success")

type Tester struct {
	Hanlder func(...interface{})
}

func (t *Tester) Assert(ok bool, args ...interface{}) {
	if !ok {
		t.Hanlder(args...)
	}
}

//TestDriver test kvdb driver
func TestDriver(driver kvdb.Driver, fatal func(...interface{})) {
	t := &Tester{Hanlder: fatal}
	TestFeatureStore(driver, t)
	TestFeatureTTLStore(driver, t)
	TestFeatureCounter(driver, t)
	TestFeatureTTLCounter(driver, t)
	TestFeatureNext(driver, t)
	TestFeaturePrev(driver, t)
	TestFeatureInsert(driver, t)
	TestFeatureUpdate(driver, t)
	TestFeatureTransaction(driver, t)
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
		err = driver.Del(KeySuccess)
		t.Assert(err == nil, err)
		data, err = driver.Get(KeySuccess)
		t.Assert(err == kvdb.ErrNotFound, err)
	}
}

//TestFeatureTTLStore test driver FeatureTTLStore
func TestFeatureTTLStore(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureTTLStore) {

	}
}

//TestFeatureCounter test driver FeatureCounter
func TestFeatureCounter(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureCounter) {

	}
}

//TestFeatureTTLCounter test driver FeatureTTLCounter
func TestFeatureTTLCounter(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureTTLCounter) {

	}
}

//TestFeatureNext test driver FeatureNext
func TestFeatureNext(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureNext) {

	}
}

//TestFeaturePrev test driver FeaturePrev
func TestFeaturePrev(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeaturePrev) {

	}
}

//TestFeatureInsert test driver FeatureInsert
func TestFeatureInsert(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureInsert) {

	}
}

//TestFeatureUpdate test driver FeatureUpdate
func TestFeatureUpdate(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureUpdate) {

	}
}

//TestFeatureTransaction test driver FeatureTransaction
func TestFeatureTransaction(driver kvdb.Driver, t *Tester) {
	if driver.Features().SupportAll(kvdb.FeatureTransaction) {

	}
}

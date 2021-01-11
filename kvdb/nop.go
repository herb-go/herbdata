package kvdb

import "github.com/herb-go/herbdata"

//Nop key-value database driver
//All method except "Close" will raise an ErrFeatureNotSupported error.
//All key-value database driver should implement Nop driver
type Nop struct{}

//Start start database
func (n Nop) Start() error {
	return nil
}

//Stop stop database
func (n Nop) Stop() error {
	return nil
}

//Set set value by given key
func (n Nop) Set(key []byte, value []byte) error {
	return ErrFeatureNotSupported
}

//Get get value by given key
func (n Nop) Get(key []byte) ([]byte, error) {
	return nil, ErrFeatureNotSupported
}

//Delete delete value by given key
func (n Nop) Delete(key []byte) error {
	return ErrFeatureNotSupported
}

//Next return values after key not more than given limit
func (n Nop) Next(iter []byte, limit int) (kv []herbdata.KeyValue, newiter []byte, err error) {
	return nil, nil, ErrFeatureNotSupported
}

//Prev return values after key not more than given limit
func (n Nop) Prev(iter []byte, limit int) (kv []herbdata.KeyValue, newiter []byte, err error) {
	return nil, nil, ErrFeatureNotSupported
}

//SetWithTTL set value by given key and ttl in second
func (n Nop) SetWithTTL(key []byte, value []byte, ttlInSecond int64) error {
	return ErrFeatureNotSupported
}

//Begin begin new transaction
func (n Nop) Begin() (Transaction, error) {
	return nil, ErrFeatureNotSupported
}

//Features return supported features
func (n Nop) Features() Feature {
	return 0
}

//SetCounter set counter value with given key
func (n Nop) SetCounter(key []byte, value int64) error {
	return ErrFeatureNotSupported
}

//IncreaseCounter increace counter value with given key and increasement.
//Value not existed coutn as 0.
//Return final value and any error if raised.
func (n Nop) IncreaseCounter(key []byte, incr int64) (int64, error) {
	return 0, ErrFeatureNotSupported
}

//IncreaseCounterWithTTL increace counter value with given key ,increasement,and ttl in second
//Value not existed coutn as 0.
//Return final value and any error if raised.
func (n Nop) IncreaseCounterWithTTL(key []byte, incr int64, ttlInSecond int64) (int64, error) {
	return 0, ErrFeatureNotSupported
}

//SetCounterWithTTL set counter value with given key and ttl in second
func (n Nop) SetCounterWithTTL(key []byte, value int64, ttlInSecond int64) error {
	return ErrFeatureNotSupported
}

//GetCounter get counter value with given key
//Value not existed coutn as 0.
func (n Nop) GetCounter(key []byte) (int64, error) {
	return 0, ErrFeatureNotSupported
}

//DeleteCounter delete counter value with given key
func (n Nop) DeleteCounter(key []byte) error {
	return ErrFeatureNotSupported
}

//Insert insert value with given key.
//Insert will fail if data with given key exists.
//Return if operation success and any error if raised
func (n Nop) Insert(Key []byte, value []byte) (bool, error) {
	return false, ErrFeatureNotSupported
}

//InsertWithTTL insert value with given key and ttl in second.
//Insert will fail if data with given key exists.
//Return if operation success and any error if raised
func (n Nop) InsertWithTTL(Key []byte, value []byte, ttlInSecond int64) (bool, error) {
	return false, ErrFeatureNotSupported
}

//Update update value with given key.
//Update will fail if data with given key does nto exist.
//Return if operation success and any error if raised
func (n Nop) Update(key []byte, value []byte) (bool, error) {
	return false, ErrFeatureNotSupported
}

//UpdateWithTTL update value with given key and ttl in second.
//Update will fail if data with given key does nto exist.
//Return if operation success and any error if raised
func (n Nop) UpdateWithTTL(key []byte, value []byte, ttlInSecond int64) (bool, error) {
	return false, ErrFeatureNotSupported
}

//IsolationLevel transaction isolation level
func (n Nop) IsolationLevel() IsolationLevel {
	return 0
}

//SetErrorHanlder set error hanlder
func (n Nop) SetErrorHanlder(func(error)) {
}

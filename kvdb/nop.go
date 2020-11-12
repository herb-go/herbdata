package kvdb

import "time"

type Nop struct{}

//Close close database
func (n Nop) Close() error {
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
func (n Nop) Next(iter []byte, limit int) (keys [][]byte, newiter []byte, err error) {
	return nil, nil, ErrFeatureNotSupported
}

//SetWithTTL set value by given key and ttl
func (n Nop) SetWithTTL(key []byte, value []byte, ttl time.Duration) error {
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

func (n Nop) SetCounter(key []byte, value int64) error {
	return ErrFeatureNotSupported
}
func (n Nop) IncreaseCounter(key []byte, incr int64) (int64, error) {
	return 0, ErrFeatureNotSupported
}
func (n Nop) IncreaseCounterWithTTL(key []byte, incr int64, ttl time.Duration) (int64, error) {
	return 0, ErrFeatureNotSupported
}

func (n Nop) SetCounterWithTTL(key []byte, value int64, ttl time.Duration) error {
	return ErrFeatureNotSupported
}
func (n Nop) GetCounter(key []byte) (int64, error) {
	return 0, ErrFeatureNotSupported
}
func (n Nop) DeleteCounter(key []byte) error {
	return ErrFeatureNotSupported
}

func (n Nop) Insert(Key []byte, value []byte) (bool, error) {
	return false, ErrFeatureNotSupported
}
func (n Nop) InsertWithTTL(Key []byte, value []byte, ttl time.Duration) (bool, error) {
	return false, ErrFeatureNotSupported
}
func (n Nop) UpdateWithTTL(key []byte, value []byte, ttl time.Duration) (bool, error) {
	return false, ErrFeatureNotSupported
}
func (n Nop) Update(key []byte, value []byte) (bool, error) {
	return false, ErrFeatureNotSupported
}

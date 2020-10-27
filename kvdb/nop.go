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

//Del delete value by given key
func (n Nop) Del(key []byte) error {
	return ErrFeatureNotSupported
}

//Next return values after key not more than given limit
func (n Nop) Next(key []byte, limit int) ([][]byte, error) {
	return nil, ErrFeatureNotSupported
}

//Prev return values before key not more than given limit
func (n Nop) Prev(key []byte, limit int) ([][]byte, error) {
	return nil, ErrFeatureNotSupported
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

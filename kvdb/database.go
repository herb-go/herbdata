package kvdb

import (
	"time"
)

type Database struct {
	Driver
}

func (d *Database) ShouldSupport(features Feature) error {
	if !d.Driver.Features().SupportAll(features) {
		return ErrFeatureNotSupported
	}
	return nil
}
func (d *Database) ShouldNotSupport(features Feature) error {
	if d.Driver.Features().SupportAny(features) {
		return ErrFeatureSupported
	}
	return nil
}
func New() *Database {
	return &Database{}
}

type Driver interface {
	//Close close database
	Close() error
	//Set set value by given key
	Set(key []byte, value []byte) error
	//Get get value by given key
	Get(key []byte) ([]byte, error)
	//Delete delete value by given key
	Delete(key []byte) error
	//Next return keys after key not more than given limit
	Next(iter []byte, limit int) (keys [][]byte, newiter []byte, err error)
	//SetWithTTL set value by given key and ttl
	SetWithTTL(key []byte, value []byte, ttl time.Duration) error
	//Begin begin new transaction
	Begin() (Transaction, error)
	Insert(Key []byte, value []byte) (bool, error)
	InsertWithTTL(Key []byte, value []byte, ttl time.Duration) (bool, error)
	Update(key []byte, value []byte) (bool, error)
	UpdateWithTTL(key []byte, value []byte, ttl time.Duration) (bool, error)
	SetCounter(key []byte, value int64) error
	SetCounterWithTTL(key []byte, value int64, ttl time.Duration) error
	IncreaseCounter(key []byte, incr int64) (int64, error)
	IncreaseCounterWithTTL(key []byte, incr int64, ttl time.Duration) (int64, error)
	GetCounter(key []byte) (int64, error)
	DeleteCounter(key []byte) error
	//Features return supported features
	Features() Feature
}

type Transaction interface {
	Rollback() error
	Commit() error
	//Set set value by given key
	Set(key []byte, value []byte) error
	//Get get value by given key
	Get(key []byte) ([]byte, error)
	//Delete delete value by given key
	Del(key []byte) error
	SetWithTTL(key []byte, value []byte, ttl time.Duration) error
}

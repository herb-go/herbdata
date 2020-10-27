package kvdb

import (
	"time"
)

type Database interface {
	//Close close database
	Close() error
	//Set set value by given key
	Set(key []byte, value []byte) error
	//Get get value by given key
	Get(key []byte) ([]byte, error)
	//Del delete value by given key
	Del(key []byte) error
	//Next return values after key not more than given limit
	Next(key []byte, limit int) ([][]byte, error)
	//Prev return values before key not more than given limit
	Prev(key []byte, limit int) ([][]byte, error)
	//SetWithTTL set value by given key and ttl
	SetWithTTL(key []byte, value []byte, ttl time.Duration) error
	//Begin begin new transaction
	Begin() (Transaction, error)
	//Features return supported features
	Features() Feature
}

type Transaction interface {
	Rollback() error
	Commint() error
	//Set set value by given key
	Set(key []byte, value []byte) error
	//Get get value by given key
	Get(key []byte) ([]byte, error)
	//Delete delete value by given key
	Del(key []byte) error
	//Next return values after key not more than given limit
	Next(key []byte, limit int) ([][]byte, error)
	//Prev return values before key not more than given limit
	Prev(key []byte, limit int) ([][]byte, error)
	//SetWithTTL set value by given key and ttl
	SetWithTTL(key []byte, value []byte, ttl time.Duration) error
}

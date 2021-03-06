package kvdb

import "github.com/herb-go/herbdata"

//Driver key value database driver interface
type Driver interface {
	//Start start database
	Start() error
	//Stop stop database
	Stop() error
	//Set set value by given key
	Set(key []byte, value []byte) error
	//Get get value by given key
	Get(key []byte) ([]byte, error)
	//Delete delete value by given key
	Delete(key []byte) error
	//Next return keys after iter not more than given limit
	//Empty iter (nil or 0 length []byte) will start a new search
	//Return keyvalue ,newiter and any error if raised.
	//Empty iter (nil or 0 length []byte) will be returned if no more keys
	Next(iter []byte, limit int) (kv []*herbdata.KeyValue, newiter []byte, err error)
	//Prev return keys before iter not more than given limit
	//Empty iter (nil or 0 length []byte) will start a new search
	//Return keys ,newiter and any error if raised.
	//Empty iter (nil or 0 length []byte) will be returned if no more keys
	Prev(iter []byte, limit int) (kv []*herbdata.KeyValue, newiter []byte, err error)
	//SetWithTTL set value by given key and ttl in second.
	SetWithTTL(key []byte, value []byte, ttlInSecond int64) error
	//SetWithExpired set value by given key and expired timestamp.
	SetWithExpired(key []byte, value []byte, expired int64) error
	//Begin begin new transaction
	Begin() (Transaction, error)
	//Insert insert value with given key.
	//Insert will fail if data with given key exists.
	//Return if operation success and any error if raised
	Insert(Key []byte, value []byte) (bool, error)
	//InsertWithTTL insert value with given key and ttl in second.
	//Insert will fail if data with given key exists.
	//Return if operation success and any error if raised
	InsertWithTTL(Key []byte, value []byte, ttlInSecond int64) (bool, error)
	//Update update value with given key.
	//Update will fail if data with given key does nto exist.
	//Return if operation success and any error if raised
	Update(key []byte, value []byte) (bool, error)
	//UpdateWithTTL update value with given key and ttl in second.
	//Update will fail if data with given key does nto exist.
	//Return if operation success and any error if raised
	UpdateWithTTL(key []byte, value []byte, ttlInSecond int64) (bool, error)
	//SetCounter set counter value with given key
	SetCounter(key []byte, value int64) error
	//SetCounterWithTTL set counter value with given key and ttl
	SetCounterWithTTL(key []byte, value int64, ttlInSecond int64) error
	//IncreaseCounter increace counter value with given key and increasement.
	//Value not existed coutn as 0.
	//Return final value and any error if raised.
	IncreaseCounter(key []byte, incr int64) (int64, error)
	//IncreaseCounterWithTTL increace counter value with given key ,increasement,and ttl.
	//Value not existed coutn as 0.
	//Return final value and any error if raised.
	IncreaseCounterWithTTL(key []byte, incr int64, ttlInSecond int64) (int64, error)
	//GetCounter get counter value with given key
	//Value not existed coutn as 0.
	GetCounter(key []byte) (int64, error)
	//DeleteCounter delete counter value with given key
	DeleteCounter(key []byte) error
	//Features return supported features
	Features() Feature
	//IsolationLevel transaction isolation level
	//Return 0 if transaction is not supported
	IsolationLevel() IsolationLevel
	//SetErrorHanlder set error hanlder
	SetErrorHanlder(func(error))
}

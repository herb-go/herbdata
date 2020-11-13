package kvdb

//Database key-value database struct
type Database struct {
	Driver
}

//ShouldSupport return ErrFeatureNotSupported if any given feature not be supported by driver
func (d *Database) ShouldSupport(features Feature) error {
	if !d.Driver.Features().SupportAll(features) {
		return ErrFeatureNotSupported
	}
	return nil
}

//ShouldNotSupport return ErrFeatureSupported if any given feature  be supported by driver
func (d *Database) ShouldNotSupport(features Feature) error {
	if d.Driver.Features().SupportAny(features) {
		return ErrFeatureSupported
	}
	return nil
}

//New create new database
func New() *Database {
	return &Database{}
}

//Driver key value database driver interface
type Driver interface {
	//Close close database
	Close() error
	//Set set value by given key
	Set(key []byte, value []byte) error
	//Get get value by given key
	Get(key []byte) ([]byte, error)
	//Delete delete value by given key
	Delete(key []byte) error
	//Next return keys after iter not more than given limit
	//Empty iter (nil or 0 length []byte) will start a new search
	//Return keys ,newiter and any error if raised.
	//Empty iter (nil or 0 length []byte) will be returned if no more keys
	Next(iter []byte, limit int) (keys [][]byte, newiter []byte, err error)
	//SetWithTTL set value by given key and ttl in second.
	SetWithTTL(key []byte, value []byte, ttlInSecond int64) error
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

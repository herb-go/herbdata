package kvdb

//Transaction datanase transaction interface
type Transaction interface {
	//Rollback rollback transaction.
	//Discard all operations in transaction.
	Rollback() error
	//Commit commit transaction
	//Apply all operaions in transaction.
	Commit() error
	//Set set value by given key
	Set(key []byte, value []byte) error
	//Get get value by given key
	Get(key []byte) ([]byte, error)
	//Delete delete value by given key
	Delete(key []byte) error
	//SetWithTTL set value by given key and ttl in Second
	SetWithTTL(key []byte, value []byte, ttlInSecond int64) error
	//IsolationLevel transaction isolation level
	IsolationLevel() IsolationLevel
}

//IsolationLevel transaction isolation level
type IsolationLevel int64

const (
	//IsolationLevelBatch isolation-level batch.Batch insert data,get data in transaction will return data in databse directly.
	IsolationLevelBatch = IsolationLevel(1 << iota)
	//IsolationLevelReadUncommitted isolation-level read uncommitted
	IsolationLevelReadUncommitted
	//IsolationLevelReadCommitted isolation-level read committed
	IsolationLevelReadCommitted
	//IsolationLevelRepeatableRead isolation-level repeatable read
	IsolationLevelRepeatableRead
	//IsolationLevelSerializable isolation-level serializable
	IsolationLevelSerializable
)

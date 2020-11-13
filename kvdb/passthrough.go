package kvdb

import "github.com/herb-go/herbdata"

var passthroughFeatures = FeatureTTLStore | FeatureStore | FeatureInsert | FeatureTTLInsert | FeatureUpdate | FeatureTTLUpdate

type passthrough struct {
	Nop
}

//Set set value by given key
func (p *passthrough) Set(key []byte, value []byte) error {
	return nil
}

//Get get value by given key
func (p *passthrough) Get(key []byte) ([]byte, error) {
	return nil, herbdata.ErrNotFound
}

//Delete delete value by given key
func (p *passthrough) Delete(key []byte) error {
	return nil
}

//SetWithTTL set value by given key and ttl in second
func (p *passthrough) SetWithTTL(key []byte, value []byte, ttl int64) error {
	return nil
}

//Features return supported features
func (p *passthrough) Features() Feature {
	return passthroughFeatures
}

//Insert insert value with given key.
//Insert will fail if data with given key exists.
//Return if operation success and any error if raised
func (p *passthrough) Insert(Key []byte, value []byte) (bool, error) {
	return false, nil
}

//InsertWithTTL insert value with given key and ttl in second.
//Insert will fail if data with given key exists.
//Return if operation success and any error if raised
func (p *passthrough) InsertWithTTL(Key []byte, value []byte, ttl int64) (bool, error) {
	return false, nil
}

//Update update value with given key.
//Update will fail if data with given key does nto exist.
//Return if operation success and any error if raised
func (p *passthrough) Update(key []byte, value []byte) (bool, error) {
	return false, nil
}

//UpdateWithTTL update value with given key and ttl in second.
//Update will fail if data with given key does nto exist.
//Return if operation success and any error if raised
func (p *passthrough) UpdateWithTTL(key []byte, value []byte, ttl int64) (bool, error) {
	return false, nil
}

//Passthrough key value database which do not store any data
var Passthrough = &passthrough{}

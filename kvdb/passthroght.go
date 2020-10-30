package kvdb

import "time"

var passthroghtFeatures = FeatureTTLStore | FeatureStore

type passthroght struct {
	Nop
}

//Set set value by given key
func (p *passthroght) Set(key []byte, value []byte) error {
	return nil
}

//Get get value by given key
func (p *passthroght) Get(key []byte) ([]byte, error) {
	return nil, ErrNotFound
}

//Del delete value by given key
func (p *passthroght) Del(key []byte) error {
	return nil
}

//SetWithTTL set value by given key and ttl
func (p *passthroght) SetWithTTL(key []byte, value []byte, ttl time.Duration) error {
	return nil
}

//Features return supported features
func (p *passthroght) Features() Feature {
	return passthroghtFeatures
}

var Passthroght = &passthroght{}

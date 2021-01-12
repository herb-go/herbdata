package herbdata

type KeyValue struct {
	Key   []byte
	Value []byte
}

func (k *KeyValue) Clone() *KeyValue {
	kv := &KeyValue{}
	kv.Key = make([]byte, len(k.Key))
	copy(kv.Key, k.Key)
	kv.Value = make([]byte, len(k.Value))
	copy(kv.Value, k.Value)
	return kv
}

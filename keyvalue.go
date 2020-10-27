package herbdata

type KeyValueData interface {
	DataKey() []byte
	DataValue() []byte
}

type KeyValueDataSetter interface {
	KeyValueData
	SetDataKey([]byte)
	SetDataValue([]byte)
}

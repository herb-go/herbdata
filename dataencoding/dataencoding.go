package dataencoding

type Encoding struct {
	Marshal   func(v interface{}) ([]byte, error)
	Unmarshal func(data []byte, v interface{}) error
}

func (e *Encoding) MarshalData(v interface{}) ([]byte, error) {
	return e.Marshal(v)
}

func (e *Encoding) UnmarshalData(data []byte, v interface{}) error {
	return e.Unmarshal(data, v)
}

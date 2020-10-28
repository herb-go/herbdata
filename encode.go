package herbdata

type DataEncoder interface {
	EncodeData() ([]byte, error)
}

type DataDecoder interface {
	DecodeData([]byte) error
}

type DataEncoderFunc func() ([]byte, error)

func (f DataEncoderFunc) EncodeData() ([]byte, error) {
	return f()
}

type DataDecoderFunc func([]byte) error

func (f DataDecoderFunc) DecodeData(data []byte) error {
	return f(data)
}

type DataMarshaler interface {
	MarshalData(interface{}) ([]byte, error)
}

type DataUnmarshaler interface {
	UnmarshalData([]byte, interface{}) error
}

type DataEncoding interface {
	DataMarshaler
	DataUnmarshaler
}

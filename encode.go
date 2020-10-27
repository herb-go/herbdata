package herbdata

import "io"

type DataEncoder interface {
	EncodeData(w io.Writer) error
}

type DataDecoder interface {
	DecodeData(r io.Reader) error
}

type DataEncoderFunc func(w io.Writer) error

func (f DataEncoderFunc) EncodeData(w io.Writer) error {
	return f(w)
}

type DataDecoderFunc func(r io.Reader) error

func (f DataDecoderFunc) DecodeData(r io.Reader) error {
	return f(r)
}

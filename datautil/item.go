package datautil

import (
	"io"

	"github.com/herb-go/herbdata"
)

type item struct {
	v interface{}
}

func (i *item) EncodeData(w io.Writer) error {
	data, err := Encode(i.v)
	if err != nil {
		return err
	}
	return PackTo(w, nil, data)
}

func (i *item) DecodeData(r io.Reader) error {
	data, err := UnpackFrom(r, nil)
	if err != nil {
		return err
	}
	return Decode(data, i.v)
}

type encoders struct {
	items []*item
}

func (e *encoders) EncodeData(w io.Writer) error {
	var err error
	for k := range e.items {
		err = e.items[k].EncodeData(w)
		if err != nil {
			return err
		}
	}
	return nil
}

type decoders struct {
	items []*item
}

func (d *decoders) DecodeData(r io.Reader) error {
	var err error
	for k := range d.items {
		err = d.items[k].DecodeData(r)
		if err != nil {
			return err
		}
	}
	return nil
}

func Decoder(values ...interface{}) herbdata.DataDecoder {
	ds := &decoders{
		items: make([]*item, len(values)),
	}
	for k, v := range values {
		ds.items[k] = &item{v}
	}
	return ds
}

func Encoder(values ...interface{}) herbdata.DataEncoder {
	es := &encoders{
		items: make([]*item, len(values)),
	}
	for k, v := range values {
		es.items[k] = &item{v}
	}
	return es
}

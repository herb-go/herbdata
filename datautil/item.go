package datautil

import (
	"bytes"

	"github.com/herb-go/herbdata"
)

type item struct {
	v interface{}
}

func (i *item) EncodeData() ([]byte, error) {
	return Encode(i.v)
}

func (i *item) DecodeData(data []byte) error {
	return Decode(data, i.v)
}

type encoders struct {
	items []*item
}

func (e *encoders) EncodeData() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	for k := range e.items {
		data, err := e.items[k].EncodeData()
		if err != nil {
			return nil, err
		}
		err = PackTo(buf, nil, data)
		if err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

type decoders struct {
	items []*item
}

func (d *decoders) DecodeData(data []byte) error {
	buf := bytes.NewBuffer(data)
	for k := range d.items {
		bs, err := UnpackFrom(buf, nil)
		if err != nil {
			return err
		}
		err = d.items[k].DecodeData(bs)
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

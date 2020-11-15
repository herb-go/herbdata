package datautil

import (
	"bytes"
	"io"

	"github.com/herb-go/herbdata"
)

const PackLengthCodeNextByte = byte(255)
const PackLengthCodeNext2Byte = byte(254)
const PackLengthCodeNext4Byte = byte(253)
const PackLengthCodeNext8Byte = byte(252)
const PackLengthCodeReserved = byte(251)
const maxPackLengthByte = int64(250)
const maxPackLengthNextByte = int64(0xff)
const maxPackLengthNext2Byte = int64(0xffff)
const maxPackLengthNext4Byte = int64(0xffffffff)
const maxPackLengthNext8Byte = int64(uint64(1 << 62))

var DelimiterZero = []byte{0}

func getLengthBytes(data []byte) []byte {
	l := int64(len(data))
	var result []byte
	if l <= maxPackLengthByte {
		return []byte{byte(l)}
	}
	if l < maxPackLengthNextByte {
		return []byte{PackLengthCodeNextByte, byte(l)}
	}
	if l < maxPackLengthNext2Byte {
		result = make([]byte, 3)
		result[0] = PackLengthCodeNext2Byte
		herbdata.DataOrder.PutUint16(result[1:], uint16(l))
		return result
	}
	if l < maxPackLengthNext4Byte {
		result = make([]byte, 5)
		result[0] = PackLengthCodeNext4Byte
		herbdata.DataOrder.PutUint32(result[1:], uint32(l))
		return result
	}
	result = make([]byte, 9)
	result[0] = PackLengthCodeNext8Byte
	herbdata.DataOrder.PutUint64(result[1:], uint64(l))
	return result
}

func PackTo(w io.Writer, delimiter []byte, data ...[]byte) error {
	var err error
	var writeDelimiter = len(delimiter) > 0
	for k := range data {
		_, err = w.Write(getLengthBytes(data[k]))
		if err != nil {
			return err
		}
		_, err = w.Write(data[k])
		if err != nil {
			return err
		}
		if writeDelimiter {
			_, err = w.Write(delimiter)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func Join(delimiter []byte, data ...[]byte) []byte {
	return Append(nil, delimiter, data...)
}

func Append(dst []byte, delimiter []byte, data ...[]byte) []byte {
	buf := bytes.NewBuffer(nil)
	_, err := buf.Write(dst)
	if err != nil {
		panic(err)
	}
	err = PackTo(buf, delimiter, data...)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()

}
func readLength(r io.Reader) (length int64, err error) {
	defer func() {
		if err == nil {
			if length < 0 {
				length = 0
				err = ErrDataLengthOverflow
			}
		} else if err == io.EOF {
			err = ErrUnpackDataFail
		}
	}()
	var buf = make([]byte, 1)
	_, err = r.Read(buf)
	if err != nil {
		return 0, err
	}
	code := buf[0]
	if int64(code) <= maxPackLengthByte {
		length = int64(code)
		return
	}
	switch code {
	case PackLengthCodeNextByte:
		_, err = r.Read(buf)
		if err != nil {
			return 0, err
		}
		length = int64(buf[0])
		return
	case PackLengthCodeNext2Byte:
		buf = make([]byte, 2)
		_, err = r.Read(buf)
		if err != nil {
			return 0, err
		}
		length = int64(herbdata.DataOrder.Uint16(buf))
		return

	case PackLengthCodeNext4Byte:
		buf = make([]byte, 4)
		_, err = r.Read(buf)
		if err != nil {
			return 0, err
		}
		length = int64(herbdata.DataOrder.Uint32(buf))
		return
	case PackLengthCodeNext8Byte:
		buf = make([]byte, 8)
		_, err = r.Read(buf)
		if err != nil {
			return 0, err
		}
		l8 := herbdata.DataOrder.Uint64(buf)
		if l8 != uint64(int64(l8)) {
			err = ErrDataLengthOverflow
			return
		}
		length = int64(l8)
		return
	}
	return 0, ErrDataLengthOverflow
}

func failOnEOF(p int, err error) error {
	if err == io.EOF {
		return ErrUnpackDataFail
	}
	return nil
}

func UnpackFrom(r io.Reader, delimiter []byte) ([]byte, error) {
	var err error
	l, err := readLength(r)
	if err != nil {
		return nil, err
	}
	data := make([]byte, l)
	err = failOnEOF(r.Read(data))
	if err != nil {
		return nil, err
	}
	buf := make([]byte, len(delimiter))
	if len(delimiter) != 0 {
		err = failOnEOF(r.Read(buf))
		if err != nil {
			return nil, err
		}
		if bytes.Compare(delimiter, buf) != 0 {
			return nil, ErrUnpackDataFail

		}
	}
	return data, nil
}

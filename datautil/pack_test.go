package datautil

import (
	"bytes"
	"testing"
)

func testInLength(t *testing.T, length int, datalength int) {
	var data []byte
	var buf *bytes.Buffer
	var result []byte
	var err error
	data = bytes.Repeat([]byte{1}, length)
	buf = bytes.NewBuffer(nil)
	err = PackTo(buf, DelimiterZero, data)
	if err != nil {
		t.Fatal(err)
	}
	if buf.Len() != 1+length+1+datalength {
		t.Fatal(buf.Len())
	}
	buf = bytes.NewBuffer(buf.Bytes())
	result, err = UnpackFrom(buf, DelimiterZero)
	if err != nil {
		if bytes.Compare(result, data) != 0 {
			t.Fatal(result, len(result), err)
		}
	}

}
func TestPack(t *testing.T) {
	testInLength(t, 250, 0)
	testInLength(t, 251, 1)
	testInLength(t, 1<<9, 2)
	testInLength(t, 1<<17, 4)
	testInLength(t, 1<<32, 8)
}

func TestJoin(t *testing.T) {
	var data = []byte("abcdef")
	var result = Join(nil, data[:3], data[3:])
	if bytes.Compare(data, result) == 0 {
		t.Fatal(result)
	}
}

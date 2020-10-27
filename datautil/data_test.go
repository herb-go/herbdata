package datautil

import (
	"bytes"
	"testing"
)

func testLengthNotMatch(t *testing.T, d []byte, values ...interface{}) {
	for _, v := range values {
		err := Decode(d, v)
		if err != ErrDataLengthNotMatch {
			t.Fatal(err)
		}
	}
}
func test(v interface{}, v2 interface{}) error {
	data, err := Encode(v)
	if err != nil {
		return err
	}
	return Decode(data, v2)
}
func TestData(t *testing.T) {
	var err error

	var vbool = true
	var rbool bool = false
	err = test(vbool, &rbool)
	if err != nil || vbool != rbool {
		t.Fatal(err, vbool, rbool)
	}
	rbool = false
	err = test(&vbool, &rbool)
	if err != nil || vbool != rbool {
		t.Fatal(err, vbool, rbool)
	}
	var vbool2 = false
	var rbool2 = true
	err = test(vbool2, &rbool2)
	if err != nil || vbool2 != rbool2 {
		t.Fatal(err, vbool2, rbool2)
	}
	rbool2 = true
	err = test(&vbool2, &rbool2)
	if err != nil || vbool2 != rbool2 {
		t.Fatal(err, vbool2, rbool2)
	}
	var vbyte = byte(1)
	var rbyte byte = 0
	err = test(vbyte, &rbyte)
	if err != nil || vbyte != rbyte {
		t.Fatal(err, vbyte, rbyte)
	}
	rbyte = 0
	err = test(&vbyte, &rbyte)
	if err != nil || vbyte != rbyte {
		t.Fatal(err, vbyte, rbyte)
	}

	var vint = int(-1)
	var rint int = 0
	err = test(vint, &rint)
	if err != nil || vint != rint {
		t.Fatal(err, vint, rint)
	}
	rint = 0
	err = test(&vint, &rint)
	if err != nil || vint != rint {
		t.Fatal(err, vint, rint)
	}
	var vuint = uint(1)
	var ruint uint = 0
	err = test(vuint, &ruint)
	if err != nil || vuint != ruint {
		t.Fatal(err, vuint, ruint)
	}
	ruint = 0
	err = test(&vuint, &ruint)
	if err != nil || vuint != ruint {
		t.Fatal(err, vuint, ruint)
	}

	var vint8 = int8(-1)
	var rint8 int8 = 0
	err = test(vint8, &rint8)
	if err != nil || vint8 != rint8 {
		t.Fatal(err, vint8, rint8)
	}
	rint8 = 0
	err = test(&vint8, &rint8)
	if err != nil || vint8 != rint8 {
		t.Fatal(err, vint8, rint8)
	}
	var vuint8 = uint8(1)
	var ruint8 uint8 = 0
	err = test(vuint8, &ruint8)
	if err != nil || vuint8 != ruint8 {
		t.Fatal(err, vuint8, ruint8)
	}
	ruint8 = 0
	err = test(&vuint8, &ruint8)
	if err != nil || vuint8 != ruint8 {
		t.Fatal(err, vuint8, ruint8)
	}

	var vint16 = int16(-1)
	var rint16 int16 = 0
	err = test(vint16, &rint16)
	if err != nil || vint16 != rint16 {
		t.Fatal(err, vint16, rint16)
	}
	rint16 = 0
	err = test(&vint16, &rint16)
	if err != nil || vint16 != rint16 {
		t.Fatal(err, vint16, rint16)
	}
	var vuint16 = uint16(1)
	var ruint16 uint16 = 0
	err = test(vuint16, &ruint16)
	if err != nil || vuint16 != ruint16 {
		t.Fatal(err, vuint16, ruint16)
	}
	ruint16 = 0
	err = test(&vuint16, &ruint16)
	if err != nil || vuint16 != ruint16 {
		t.Fatal(err, vuint16, ruint16)
	}

	var vint32 = int32(-1)
	var rint32 int32 = 0
	err = test(vint32, &rint32)
	if err != nil || vint32 != rint32 {
		t.Fatal(err, vint32, rint32)
	}
	rint32 = 0
	err = test(&vint32, &rint32)
	if err != nil || vint32 != rint32 {
		t.Fatal(err, vint32, rint32)
	}
	var vuint32 = uint32(1)
	var ruint32 uint32 = 0
	err = test(vuint32, &ruint32)
	if err != nil || vuint32 != ruint32 {
		t.Fatal(err, vuint32, ruint32)
	}
	ruint32 = 0
	err = test(&vuint32, &ruint32)
	if err != nil || vuint32 != ruint32 {
		t.Fatal(err, vuint32, ruint32)
	}

	var vint64 = int64(-1)
	var rint64 int64 = 0
	err = test(vint64, &rint64)
	if err != nil || vint64 != rint64 {
		t.Fatal(err, vint64, rint64)
	}
	rint64 = 0
	err = test(&vint64, &rint64)
	if err != nil || vint64 != rint64 {
		t.Fatal(err, vint64, rint64)
	}
	var vuint64 = uint64(1)
	var ruint64 uint64 = 0
	err = test(vuint64, &ruint64)
	if err != nil || vuint64 != ruint64 {
		t.Fatal(err, vuint64, ruint64)
	}
	ruint64 = 0
	err = test(&vuint64, &ruint64)
	if err != nil || vuint64 != ruint64 {
		t.Fatal(err, vuint64, ruint64)
	}
	var vf32 = float32(1)
	var rf32 float32 = 0
	err = test(vf32, &rf32)
	if err != nil || vf32 != rf32 {
		t.Fatal(err, vf32, rf32)
	}
	rf32 = 0
	err = test(&vf32, &rf32)
	if err != nil || vf32 != rf32 {
		t.Fatal(err, vf32, rf32)
	}

	var vf64 = float64(1)
	var rf64 float64 = 0
	err = test(vf64, &rf64)
	if err != nil || vf64 != rf64 {
		t.Fatal(err, vf64, rf64)
	}
	rf64 = 0
	err = test(&vf64, &rf64)
	if err != nil || vf64 != rf64 {
		t.Fatal(err, vf64, rf64)
	}

	var vs = "1"
	var rs = "0"
	err = test(vs, &rs)
	if err != nil || vs != rs {
		t.Fatal(err, vs, rs)
	}
	rs = "0"
	err = test(&vs, &rs)
	if err != nil || rs != rs {
		t.Fatal(err, rs, rs)
	}

	var vbs = []byte("1")
	var rbs = []byte("0")
	err = test(vbs, &rbs)
	if err != nil || bytes.Compare(vbs, rbs) != 0 {
		t.Fatal(err, vbs, rbs)
	}
	rbs = []byte("0")
	err = test(&vbs, &rbs)
	if err != nil || bytes.Compare(vbs, rbs) != 0 {
		t.Fatal(err, vbs, rbs)
	}

	vbs = []byte("1")
	rbs = []byte("0")

	err = test(Encoder(vbs), Decoder(&rbs))
	if err != nil || bytes.Compare(vbs, rbs) != 0 {
		t.Fatal(err, vbs, rbs)
	}

	vbs = []byte("1")
	rbs = []byte("0")

	err = test(Encoder(vbs).EncodeData, Decoder(&rbs).DecodeData)
	if err != nil || bytes.Compare(vbs, rbs) != 0 {
		t.Fatal(err, vbs, rbs)
	}

	var vfun = func() {}
	_, err = Encode(vfun)
	if err != ErrDataTypeNotSupported {
		t.Fatal(err)
	}
	d := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	err = Decode(d, &vfun)
	if err != ErrDataTypeNotSupported {
		t.Fatal(err)
	}

	testLengthNotMatch(t, d, &rbool, &rbyte, &rint8, &ruint8, &rint16, &ruint16, &rint32, &ruint32, &rint64, &ruint64, &rint, &ruint, &rf32, &rf64)
	testLengthNotMatch(t, []byte{}, &rbyte)
}

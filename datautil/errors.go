package datautil

import "errors"

var ErrDataLengthNotMatch = errors.New("binary data length not match")
var ErrDataTypeNotSupported = errors.New("binary data type not supported")
var ErrDataLengthOverflow = errors.New("binary data length overflow")
var ErrUnpackDataFail = errors.New("unpack binary data fail")

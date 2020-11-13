package kvdb

import "errors"

//ErrFeatureNotSupported error raised if given feature is not supported
var ErrFeatureNotSupported = errors.New("feature not supported")

//ErrFeatureSupported error raised if given feature is supported
var ErrFeatureSupported = errors.New("feature supported")

//ErrUnsupportedNextLimit error raised when next limit unsupported
var ErrUnsupportedNextLimit = errors.New("unsuported next limit")

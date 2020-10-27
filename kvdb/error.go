package kvdb

import "errors"

var ErrFeatureNotSupported = errors.New("feature not supported")
var ErrInvalidateKey = errors.New("invalidate key")
var ErrKeyNotFound = errors.New("key not found")

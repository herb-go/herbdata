package kvdb

import "errors"

var ErrFeatureNotSupported = errors.New("feature not supported")
var ErrFeatureSupported = errors.New("feature supported")
var ErrInvalidateKey = errors.New("invalidate key")
var ErrNotFound = errors.New("data not found")

//ErrEntryTooLarge raised when data is too large to store.
var ErrEntryTooLarge = errors.New("data entry too large")

//ErrKeyTooLarge raised when key is too large to store.
var ErrKeyTooLarge = errors.New("data key too large")

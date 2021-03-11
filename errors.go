package herbdata

import "errors"

// ErrInvalidatedKey error raised if given key invalidated
var ErrInvalidatedKey = errors.New("invalidated key")

// ErrNotFound error raised if given key not found form key value database
var ErrNotFound = errors.New("data not found")

//ErrEntryTooLarge error raised when data is too large to store.
var ErrEntryTooLarge = errors.New("data entry too large")

//ErrKeyTooLarge error raised when key is too large to store.
var ErrKeyTooLarge = errors.New("data key too large")

// ErrInvalidatedTTL error raised if given key invalidated
var ErrInvalidatedTTL = errors.New("invalidated ttl")

//ErrIrrevocable error  raised if irrevocable
var ErrIrrevocable = errors.New("irrevocable")

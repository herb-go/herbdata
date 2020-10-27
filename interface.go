package herbdata

import "time"

type Store interface {
	Set([]byte, []byte) error
	Get([]byte) ([]byte, error)
	Del([]byte) error
}

type Deleter interface {
	Del([]byte) error
}
type Setter interface {
	Set([]byte, []byte) error
}
type Getter interface {
	Get([]byte) ([]byte, error)
}

type Closer interface {
	Close() error
}
type Cache interface {
	Get([]byte) ([]byte, error)
	SetWithTTL(key []byte, value []byte, ttl time.Duration) error
	Del([]byte) error
}

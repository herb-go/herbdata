package herbdata

type Server interface {
	Start() error
	Stop() error
}
type Store interface {
	Set([]byte, []byte) error
	Get([]byte) ([]byte, error)
	Delete([]byte) error
}

type StoreServer interface {
	Store
	Server
}
type Deleter interface {
	Delete([]byte) error
}
type Setter interface {
	Set([]byte, []byte) error
}
type Getter interface {
	Get([]byte) ([]byte, error)
}

type Cache interface {
	Get([]byte) ([]byte, error)
	SetWithTTL(key []byte, value []byte, ttl int64) error
	Delete([]byte) error
}

type CacheServer interface {
	Cache
	Server
}

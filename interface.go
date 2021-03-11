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

type Revocable interface {
	Revoke() error
}

type RevocableCache interface {
	Revocable
	Cache
}

type NestedCache interface {
	GetNested(key []byte, path ...[]byte) ([]byte, error)
	SetWithTTLNested(key []byte, value []byte, ttl int64, path ...[]byte) error
	DeleteNested(key []byte, path ...[]byte) error
}

type RevocableNestedCache interface {
	NestedCache
	RevokeNested(path ...[]byte) error
}
type CacheServer interface {
	Cache
	Server
}

type SetterDeleter interface {
	Set([]byte, []byte) error
	Delete([]byte) error
}

type ExpiredSetter interface {
	Set([]byte, []byte) error
	SetWithExpired(key []byte, value []byte, expired int64) error
}

type ExpiredSetterDeleter interface {
	Set([]byte, []byte) error
	SetWithExpired(key []byte, value []byte, expired int64) error
	Delete([]byte) error
}

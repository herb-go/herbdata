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

type NamespacedCache interface {
	GetNamespaced(namespace []byte, key []byte) ([]byte, error)
	SetWithTTLNamespaced(namespace []byte, key []byte, value []byte, ttl int64) error
	DeleteNamespaced(namespace []byte, key []byte) error
}

type RevocableNamespacedCache interface {
	NamespacedCache
	RevokeNamespaced(namespace []byte) error
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

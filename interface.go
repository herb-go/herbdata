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

type SetterGetter interface {
	Setter
	Getter
}
type SetterGetterServer interface {
	Setter
	Getter
	Server
}
type Cache interface {
	Get([]byte) ([]byte, error)
	SetWithTTL(key []byte, value []byte, ttl int64) error
	Delete([]byte) error
}

type RevocableCache interface {
	Revoke() error
	Cache
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

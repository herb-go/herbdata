package datauri

import (
	"net/url"
	"sync"
)

type DataURILoader interface {
	LoadDataURI(u *url.URL) ([]byte, error)
}

type DataURILoaderFunc func(u *url.URL) ([]byte, error)

func (f DataURILoaderFunc) LoadDataURI(u *url.URL) ([]byte, error) {
	return f(u)
}

var lock sync.RWMutex
var loaders = map[string]DataURILoader{}

func Register(scheme string, loader DataURILoader) {
	lock.Lock()
	defer lock.Unlock()
	loaders[scheme] = loader
}

func Schemes() []string {
	lock.RLock()
	defer lock.RUnlock()
	results := make([]string, 0, len(loaders))
	for k := range loaders {
		results = append(results, k)
	}
	return results
}

func Load(uri string) ([]byte, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	lock.Lock()
	defer lock.Unlock()
	loader, ok := loaders[u.Scheme]
	if !ok {
		return nil, NewSchemaNotReigsteredError(u.Scheme)
	}
	return loader.LoadDataURI(u)
}

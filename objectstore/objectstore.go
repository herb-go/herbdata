package objectstore

import (
	"io"
)

type ObjectStore interface {
	Stat(path string) (*Stat, error)
	List(path string) ([]*Stat, string, error)
	Remove(path string) error
	Rename(from string, to string) error
	Copy(from string, to string) error
	LoadObject(path string, w io.Writer) (int64, error)
	LoadObjectPart(path string, from int, to int, w io.Writer) (int64, error)
	SaveObject(path string, r io.Reader) (int64, error)
}

func NewObject(o ObjectStore, p string) *Object {
	return &Object{
		Store: o,
		Path:  p,
	}
}

type NopObjectStore struct{}

func (s NopObjectStore) List(path string, iter string, limit int64) ([]*Stat, string, error) {
	return nil, "", ErrFeatureNotSupported
}
func (s NopObjectStore) Stat(path string) (*Stat, error) {
	return nil, ErrFeatureNotSupported
}

func (s NopObjectStore) Remove(path string) error {
	return ErrFeatureNotSupported
}
func (s NopObjectStore) Rename(from string, to string) error {
	return ErrFeatureNotSupported
}
func (s NopObjectStore) Copy(from string, to string) error {
	return ErrFeatureNotSupported
}
func (s NopObjectStore) LoadObject(path string, w io.Writer) (int64, error) {
	return 0, ErrFeatureNotSupported
}
func (s NopObjectStore) LoadObjectPart(path string, from int64, to int64, w io.Writer) (int64, error) {
	return 0, ErrFeatureNotSupported
}
func (s NopObjectStore) SaveObject(path string, r io.Reader) (int64, error) {
	return 0, ErrFeatureNotSupported
}
func (s NopObjectStore) MakeFolder(path string) error {
	return ErrFeatureNotSupported
}

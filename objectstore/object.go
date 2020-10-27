package objectstore

import "io"

type Object struct {
	Path  string
	Store ObjectStore
}

func (o *Object) Load(path string, w io.Writer) (int64, error) {
	return o.Store.LoadObject(path, w)
}
func (o *Object) LoadPart(path string, from int, to int, w io.Writer) (int64, error) {
	return o.Store.LoadObjectPart(path, from, to, w)

}
func (o *Object) Save(path string, r io.Reader) (int64, error) {
	return o.Store.SaveObject(path, r)
}

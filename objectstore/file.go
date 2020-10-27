package objectstore

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type FileObjectStore struct {
	NopObjectStore
	Path string
	Mode os.FileMode
}

func (s *FileObjectStore) abs(path string) (string, error) {
	filepath := filepath.Join(s.Path, path)
	if strings.HasPrefix(filepath, s.Path) {
		return filepath, nil
	}
	return "", NewErrPathInvalid(path)
}
func (s *FileObjectStore) List(path string, iter string, limit int) (result []*Stat, newiter string, err error) {
	defer func() { err = convertFileObjectStoreError(path, err) }()
	var abs string
	var files []os.FileInfo
	abs, err = s.abs(path)
	if err != nil {
		return
	}
	files, err = ioutil.ReadDir(abs)
	if err != nil {
		return
	}
	data := make(Stats, len(files))
	for k, v := range files {
		data[k] = &Stat{
			Name:         v.Name(),
			IsFolder:     v.IsDir(),
			Size:         v.Size(),
			ModifiedTime: v.ModTime(),
		}
	}
	sort.Sort(data)
	var i int
	for i = 0; i < len(data); i++ {
		if data[i].Name > iter {
			break
		}
	}
	var to int
	if limit < 0 {
		to = len(data) - 1
	} else {
		to = i + limit
		if to >= len(data) {
			to = len(data) - 1
		}
	}
	if to == i {
		return
	}
	result = make(Stats, to-i)
	copy(result, data[i:to])
	if to == len(data)-1 {
		return
	}
	iter = data[to].Name
	return
}
func (s *FileObjectStore) Stat(path string) (stat *Stat, err error) {
	defer func() { err = convertFileObjectStoreError(path, err) }()
	var abs string
	abs, err = s.abs(path)
	if err != nil {
		return
	}
	i, err := os.Stat(abs)
	if err != nil {
		return
	}
	return convertStat(i), nil
}

func (s *FileObjectStore) Remove(path string) (err error) {
	defer func() { err = convertFileObjectStoreError(path, err) }()
	var abs string
	abs, err = s.abs(path)
	if err != nil {
		return
	}
	err = os.Remove(abs)
	return
}
func (s *FileObjectStore) Rename(from string, to string) (err error) {
	defer func() { err = convertFileObjectStoreError(from, err) }()
	var absfrom, absto string
	absfrom, err = s.abs(from)
	if err != nil {
		return
	}
	absto, err = s.abs(from)
	if err != nil {
		return
	}
	err = os.Rename(absfrom, absto)
	return
}
func (s *FileObjectStore) Copy(from string, to string) (err error) {
	defer func() { err = convertFileObjectStoreError(from, err) }()
	var absfrom, absto string
	var filefrom, fileto *os.File
	absfrom, err = s.abs(from)
	if err != nil {
		return
	}
	absto, err = s.abs(from)
	if err != nil {
		return
	}
	filefrom, err = os.Open(absfrom)
	if err != nil {
		return
	}
	defer filefrom.Close()
	fileto, err = os.Open(absto)
	if err != nil {
		return
	}
	defer fileto.Close()
	_, err = io.Copy(fileto, filefrom)
	return
}
func (s *FileObjectStore) LoadObject(path string, w io.Writer) (writen int64, err error) {
	defer func() { err = convertFileObjectStoreError(path, err) }()
	var abs string
	var f *os.File
	abs, err = s.abs(path)
	if err != nil {
		return
	}
	f, err = os.Open(abs)
	if err != nil {
		return
	}
	defer f.Close()
	writen, err = io.Copy(w, f)
	return
}
func (s *FileObjectStore) LoadObjectPart(path string, from int64, to int64, w io.Writer) (writen int64, err error) {
	defer func() { err = convertFileObjectStoreError(path, err) }()
	var abs string
	var f *os.File
	abs, err = s.abs(path)
	if err != nil {
		return
	}
	_, err = f.Seek(from, 0)
	if err != nil {
		return
	}
	f, err = os.Open(abs)
	if err != nil {
		return
	}
	defer f.Close()
	writen, err = io.CopyN(w, f, to-from)
	return
}
func (s *FileObjectStore) SaveObject(path string, r io.Reader) (writen int64, err error) {
	defer func() { err = convertFileObjectStoreError(path, err) }()
	var abs string
	var f *os.File
	abs, err = s.abs(path)
	if err != nil {
		return
	}
	f, err = os.Open(abs)
	if err != nil {
		return
	}
	defer f.Close()
	writen, err = io.Copy(f, r)
	return
}

func convertFileObjectStoreError(path string, err error) error {
	if err == nil {
		return nil
	} else if os.IsNotExist(err) {
		return NewErrObjectNotExist(path)
	} else if os.IsExist(err) {
		return NewErrObjectExist(path)
	} else if os.IsPermission(err) {
		return NewErrPermissionDenied(path)
	}
	return err
}

func convertStat(i os.FileInfo) *Stat {
	return &Stat{
		Name:         i.Name(),
		IsFolder:     i.IsDir(),
		Size:         i.Size(),
		ModifiedTime: i.ModTime(),
	}
}

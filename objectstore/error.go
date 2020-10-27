package objectstore

import (
	"errors"
	"fmt"
)

var ErrObjectNotExist = errors.New("object not found")

func NewErrObjectNotExist(path string) error {
	return fmt.Errorf("%w [%s]", ErrObjectNotExist, path)
}
func IsErrObjectNotExist(err error) bool {
	return errors.Is(err, ErrObjectNotExist)
}

var ErrObjectExist = errors.New("object found")

func NewErrObjectExist(path string) error {
	return fmt.Errorf("%w [%s]", ErrObjectExist, path)
}

func IsErrObjectExist(err error) bool {
	return errors.Is(err, ErrObjectExist)
}

var ErrPathInvalid = errors.New("path invalid")

func NewErrPathInvalid(path string) error {
	return fmt.Errorf("%w [%s]", ErrPathInvalid, path)

}
func IsErrPathInvalid(err error) bool {
	return errors.Is(err, ErrPathInvalid)

}

var ErrPermissionDenied = errors.New("permission denied")

func NewErrPermissionDenied(path string) error {
	return fmt.Errorf("%w [%s]", ErrPermissionDenied, path)

}
func IsErrPermissionDenied(err error) bool {
	return errors.Is(err, ErrPermissionDenied)

}

var ErrFeatureNotSupported = errors.New("feature not supported")

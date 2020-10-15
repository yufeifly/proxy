package cusErr

import "errors"

var (
	ErrNotFound            = errors.New("no such record")
	ErrServiceNotAvailable = errors.New("service not available")
	ErrBadAddress          = errors.New("bad address")
)

package cusErr

import "errors"

var (
	ErrNotFound            = errors.New("no such record")
	ErrServiceNotFound     = errors.New("service not found")
	ErrServiceNotAvailable = errors.New("service not available")
	ErrBadAddress          = errors.New("bad address")
)

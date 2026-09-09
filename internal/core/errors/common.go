package core_errors

import "errors"

var (
	ErrNotFound        = errors.New("Not Found")
	ErrinvalidArgument = errors.New("invalid argument")
	ErrorConflict      = errors.New("conflict")
)

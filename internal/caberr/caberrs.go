package caberr

import (
	"errors"
	"fmt"
)

var (
	ErrInternal      = errors.New("internal error")
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)

type CabErr struct {
	Code    error
	Message string
	Err     error
}

func (e *CabErr) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *CabErr) Is(target error) bool {
	return e.Code == target
}

func (e *CabErr) Unwrap() error {
	return e.Err
}

func Internal(message string, err error) error {
	return &CabErr{
		Code:    ErrInternal,
		Message: message,
		Err:     err,
	}
}

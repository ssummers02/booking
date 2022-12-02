package domain

import (
	"errors"
	"fmt"
	"net"

	"github.com/gocraft/dbr/v2"
	"github.com/lib/pq"
)

type ErrCode uint8

const (
	ErrCodeInternal        ErrCode = iota + 1 // 1
	ErrCodeAlreadyExists                      // 2
	ErrCodeNotFound                           // 3
	ErrCodeValidation                         // 4
	ErrCodeDatabaseFailure                    // 5
	ErrCodeDatabaseError                      // 6
	ErrCodeNotAuthorized                      // 7
	ErrCodeForbidden                          // 8
	ErrCodeUploadFailed                       // 9
	ErrCodeExternalFailure                    // 10
	ErrCodeDeleteFailed                       // 11
	ErrCodeTokenExpired                       // 12
)

var (
	ErrInternal        = &Error{Code: ErrCodeInternal}
	ErrAlreadyExists   = &Error{Code: ErrCodeAlreadyExists}
	ErrNotFound        = &Error{Code: ErrCodeNotFound}
	ErrValidation      = &Error{Code: ErrCodeValidation}
	ErrDatabaseFailure = &Error{Code: ErrCodeDatabaseFailure}
	ErrDatabaseError   = &Error{Code: ErrCodeDatabaseError}
	ErrNotAuthorized   = &Error{Code: ErrCodeNotAuthorized}
	ErrForbidden       = &Error{Code: ErrCodeForbidden}
	ErrUploadFailed    = &Error{Code: ErrCodeUploadFailed}
	ErrExternalFailure = &Error{Code: ErrCodeExternalFailure}
	ErrDeleteFailed    = &Error{Code: ErrCodeDeleteFailed}
	ErrTokenExpired    = &Error{Code: ErrCodeTokenExpired}
)

type Error struct {
	UUID         string  `json:"uuid"`
	UserID       int64   `json:"user_id"`
	Code         ErrCode `json:"code"`
	ErrorMessage string  `json:"error_message"`
	parent       error
}

func (e Error) Error() string {
	return e.ErrorMessage
}

func (e Error) Is(target error) bool {
	var err *Error

	if !errors.As(target, &err) {
		return false
	}

	return e.Code == err.Code
}

func (e *Error) SetErrorMessage(format string, params ...interface{}) {
	e.ErrorMessage = fmt.Sprintf(format, params...)
}

// NewErrorWrap wraps an external error.
func NewErrorWrap(err error, code ErrCode, errorMessageFormat string, params ...interface{}) *Error {
	e := NewError(code, errorMessageFormat, params...)
	e.parent = err

	return e
}

// NewError creates a new custom error of our format.
func NewError(code ErrCode, errorMessageFormat string, params ...interface{}) *Error {
	err := &Error{
		Code: code,
	}
	err.SetErrorMessage(errorMessageFormat, params...)

	return err
}

func NewDBErrorWrap(err error) error {
	switch err.(type) { //nolint:errorlint
	case *net.OpError:
		return NewErrorWrap(err, ErrCodeDatabaseFailure, "Database failure: %v", err)
	case *pq.Error:
		return NewErrorWrap(err, ErrCodeDatabaseError, "Database produced error: %v", err)
	}

	if errors.Is(err, dbr.ErrNotFound) {
		return NewErrorWrap(err, ErrCodeNotFound, "Not found")
	}

	return err
}

func NewUnauthorizedError() *Error {
	return NewError(ErrCodeNotAuthorized, "Unauthorized")
}

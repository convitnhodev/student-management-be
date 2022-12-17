package solveError

import (
	"github.com/pkg/errors"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"_"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"key"`
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootErr.Error()
}

func NewFullErrorResponse(statusCode int, rootErr error, message string, log string, key string) *AppError {
	return &AppError{
		statusCode,
		rootErr,
		message,
		log,
		key,
	}
}

func NewErrorResponse(rootErr error, message string, log string, key string) *AppError {
	return &AppError{
		http.StatusBadRequest,
		rootErr,
		message,
		log,
		key,
	}
}

func ErrDB(err error) *AppError {
	return NewErrorResponse(err, "some thing went wrong with DB", err.Error(), "DB_ERR")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "Invalid_request")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "internal error", err.Error(), "Internal_ERR")
}

// 401, Unauthorrize
func NewUnauthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}

func NewCustomError(root error, msg, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}
	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

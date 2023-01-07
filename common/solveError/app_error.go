package solveError

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"strings"
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

func ErrWaitingAdminAcp(err error) *AppError {
	return NewErrorResponse(err, "waiting admin acp", err.Error(), "AD_ACP")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "Invalid_request")
}

func ErrNoPermission(err error) *AppError {
	return NewErrorResponse(err, "no permission", err.Error(), "Invalid_permission")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "internal error", err.Error(), "Internal_ERR")
}

func ErrInvalidLogin(err error) *AppError {
	return NewFullErrorResponse(http.StatusBadRequest, err, "user or password invalid", err.Error(), "Internal_ERR")
}

func ErrInvalidCurrentPassword() *AppError {
	return NewFullErrorResponse(http.StatusBadRequest, errors.New("invalid password"), "current password invalid", "", "Internal_ERR")
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

func ErrEntityExisted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s existed", strings.ToLower(entity)),
		fmt.Sprintf("%s Existed", entity),
	)
}

func ErrEntityNotExisted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s is not exist", strings.ToLower(entity)),
		fmt.Sprintf("%s is not Exist", entity),
	)
}

package tokenProvider

import (
	"errors"
	_const "managerstudent/common/const"
	"managerstudent/common/solveError"
	"time"
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

var (
	ErrNotFound = solveError.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ErrNotFound",
	)

	ErrEncodingToken = solveError.NewCustomError(
		errors.New("error encoding the token"),
		"error encoding the token",
		"ErrEncodingToken",
	)

	ErrInvalidToken = solveError.NewCustomError(
		errors.New("invalid token provider"),
		"invalid token provider",
		"ErrInvalidToken",
	)
	ErrInvalidToken1 = solveError.NewCustomError(
		errors.New("invalid token provider"),
		"invalid token provider (access token expired)",
		"ErrInvalidToken1",
	)

	ErrInvalidToken2 = solveError.NewCustomError(
		errors.New("invalid token provider"),
		"invalid token provider (refresh token expired)",
		"ErrInvalidToken2",
	)
)

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type Account struct {
	AccessToken *Token
}

type TokenPayload struct {
	UserName string      `json:"user_name" bson:"user_name,omitempty"`
	Role     _const.Role `json:"role" bson:"role"`
}

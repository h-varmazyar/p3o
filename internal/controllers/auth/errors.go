package auth

import (
	"github.com/h-varmazyar/p3o/pkg/errors"
	"net/http"
)

// error code format: 110ab
var (
	ErrLoginFailed             = errors.NewWithHttp("login_failed", 11000, http.StatusBadRequest)
	ErrInvalidUsernamePassword = errors.NewWithHttp("invalid_username", 11001, http.StatusBadRequest)
	ErrPasswordHashingFailed   = errors.NewWithHttp("password_hashing_failed", 11002, http.StatusInternalServerError)
)

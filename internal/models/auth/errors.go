package auth

import "github.com/h-varmazyar/p3o/pkg/errors"

// error code format: 221ab
var (
	ErrFailedToCreateUser = errors.NewWithCode("create_link_failed", 21100)
	ErrUserNotFound       = errors.NewWithCode("user_not_found", 21101)
	ErrFailedToFetchUser  = errors.NewWithCode("failed_to_fetch_user", 21102)
)

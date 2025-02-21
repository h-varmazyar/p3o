package user

import "github.com/h-varmazyar/p3o/pkg/errors"

// error code format: 221ab
var (
	ErrFailedToCreateUser = errors.NewWithCode("create_link_failed", 22100)
	ErrUserNotFound       = errors.NewWithCode("user_not_found", 22101)
	ErrFailedToFetchUser  = errors.NewWithCode("failed_to_fetch_user", 22102)
)

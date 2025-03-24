package user

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/errors"
)

func (s Service) Logout(ctx context.Context, username string) error {
	return errors.ErrUnimplemented
}

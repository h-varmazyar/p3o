package user

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/errors"
)

func (s Service) Verify(ctx context.Context, username, verificationCode string) error {
	return errors.ErrUnimplemented
}

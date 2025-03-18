package link

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/errors"
)

func (s Service) Delete(ctx context.Context, userId uint, key string) error {
	link, err := s.linkRepo.ReturnByKey(ctx, key)
	if err != nil {
		return err
	}

	if link.OwnerId != userId {
		return errors.ErrLinkOwnerMismatch
	}

	if err = s.linkRepo.Delete(ctx, key); err != nil {
		return err
	}
	return nil
}

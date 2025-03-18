package link

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/errors"
)

func (s Service) Status(ctx context.Context, userId uint, key string) (string, error) {
	link, err := s.linkRepo.ReturnByKey(ctx, key)
	if err != nil {
		return "", err
	}

	if link.OwnerId != userId {
		return "", errors.ErrLinkOwnerMismatch
	}

	return string(link.Status), nil
}

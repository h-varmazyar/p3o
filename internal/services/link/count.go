package link

import (
	"context"

	"github.com/h-varmazyar/p3o/internal/errors"
)

func (s Service)TotalLinkCount(ctx context.Context, userId uint) (uint, error) {
	if count, err := s.linkRepo.TotalLinkCount(ctx, userId); err != nil {
		return 0, errors.ErrUnexpected.AddOriginalError(err)
	} else {
		return count, nil
	}
}
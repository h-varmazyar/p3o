package link

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/errors"
)

func (s Service) TotalLinkCount(ctx context.Context, userId uint) (domain.DashboardInfoItem, error) {
	count, err := s.linkRepo.TotalLinkCount(ctx, userId)
	if err != nil {
		return domain.DashboardInfoItem{}, errors.ErrUnexpected.AddOriginalError(err)
	}

	resp := domain.DashboardInfoItem{
		Count:       uint(count),
		Growth:      "44.3%",
		GrowthTrend: "+",
	}

	return resp, nil
}

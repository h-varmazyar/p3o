package link

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/domain"
)

func (s Service) TotalVisit(ctx context.Context, userId uint) (domain.DashboardInfoItem, error) {
	totalVisits, err := s.linkRepo.TotalVisits(ctx, userId)
	if err != nil {
		return domain.DashboardInfoItem{}, err
	}

	resp := domain.DashboardInfoItem{
		Count:       uint(totalVisits),
		Growth:      "5%",
		GrowthTrend: "-",
	}

	return resp, nil
}

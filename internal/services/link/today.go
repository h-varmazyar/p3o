package link

import (
	"context"
	"fmt"
	"github.com/h-varmazyar/p3o/internal/domain"
	"time"
)

func (s Service) TodayInfo(ctx context.Context, userId uint) (domain.DashboardInfoItem, error) {
	visitCount, err := s.visitRepo.VisitCount(ctx, userId, time.Now().Add(-1*time.Hour*24), time.Now())
	if err != nil {
		return domain.DashboardInfoItem{}, err
	}

	fmt.Println("today:", visitCount)

	resp := domain.DashboardInfoItem{
		Count:       uint(visitCount),
		Growth:      "23.4%",
		GrowthTrend: "+",
	}

	return resp, nil
}

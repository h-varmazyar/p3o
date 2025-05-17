package link

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/domain"
)

var days = map[string]string{
	"Sunday":    "یکشنبه",
	"Monday":    "دوشنبه",
	"Tuesday":   "سه شنبه",
	"Wednesday": "چهارشنبه",
	"Thursday":  "پنجشنبه",
	"Friday":    "جمعه",
	"Saturday":  "شنبه",
}

func (s Service) VisitChart(ctx context.Context, userId uint, key string) ([]domain.ChartItem, error) {
	linkId := uint(0)
	if key != "" {
		link, err := s.linkRepo.ReturnByKey(ctx, key)
		if err != nil {
			return nil, err
		}

		linkId = link.ID
	}
	visits, err := s.visitRepo.DailyVisitCount(ctx, userId, linkId, 7)
	if err != nil {
		return nil, err
	}

	resp := make([]domain.ChartItem, len(visits))
	for i, visit := range visits {
		resp[i] = domain.ChartItem{
			Count:     uint(visit.VisitCount),
			TimeLabel: days[visit.VisitDate.Weekday().String()],
		}
	}

	return resp, nil
}

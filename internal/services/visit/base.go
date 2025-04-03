package visit

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/repositories/visit"
	log "github.com/sirupsen/logrus"
	"time"
)

type visitRepository interface {
	VisitCount(ctx context.Context, userId uint, from, to time.Time) (int64, error)
	DailyVisitCount(ctx context.Context, userId uint, count uint) ([]visit.DailyCount, error)
}

type linkRepository interface {
	TotalVisits(ctx context.Context, userId uint) (int64, error)
}

type Service struct {
	log       *log.Logger
	visitRepo visitRepository
	linkRepo  linkRepository
}

func New(log *log.Logger, visitRepo visitRepository, linkRepo linkRepository) Service {
	return Service{
		log:       log,
		visitRepo: visitRepo,
		linkRepo:  linkRepo,
	}
}

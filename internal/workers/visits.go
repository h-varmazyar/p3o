package workers

import (
	"context"
	"database/sql"
	"github.com/h-varmazyar/p3o/internal/entities"
	log "github.com/sirupsen/logrus"
	"time"
)

type linkRepository interface {
	Update(context.Context, entities.Link) error
	ReturnById(ctx context.Context, id uint) (entities.Link, error)
}

type visitRepository interface {
	GetUnhandled(ctx context.Context) ([]entities.Visit, error)
	Update(ctx context.Context, visit entities.Visit) error
}

type VisitsWorker struct {
	log       *log.Logger
	linkRepo  linkRepository
	visitRepo visitRepository
}

func NewVisitWorker(log *log.Logger, visitRepo visitRepository, linkRepo linkRepository) *VisitsWorker {
	worker := &VisitsWorker{
		log:       log,
		linkRepo:  linkRepo,
		visitRepo: visitRepo,
	}

	go worker.start()

	return worker
}

func (w VisitsWorker) start() {
	w.log.Infof("************** visit worker started")
	ticker := time.NewTicker(time.Minute)

	for {
		select {
		case <-ticker.C:
			visits, err := w.visitRepo.GetUnhandled(context.Background())
			if err != nil {
				w.log.WithError(err).Error("failed to get unhandled visits")
				continue
			}

			for _, visit := range visits {
				err = w.handleVisit(context.Background(), visit)
				if err != nil {
					w.log.WithError(err).Error("failed to handle visit")
				}
			}

		}
	}
}

func (w VisitsWorker) handleVisit(ctx context.Context, visit entities.Visit) error {
	link, err := w.linkRepo.ReturnById(ctx, visit.LinkId)
	if err != nil {
		return err
	}

	if visit.Status == entities.VisitStatusCompleted {
		link.TotalVisit++
		if err = w.linkRepo.Update(ctx, link); err != nil {
			return err
		}
	}

	visit.HandledAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	err = w.visitRepo.Update(ctx, visit)
	if err != nil {
		return err
	}

	return nil
}

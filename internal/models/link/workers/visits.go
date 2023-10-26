package workers

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/models/link/repository"
	log "github.com/sirupsen/logrus"
)

type VisitRecord struct {
	LinkId    uint
	IpAddress string
}

type VisitsWorker struct {
	log       *log.Logger
	visitChan chan VisitRecord
	persistDB repository.Repository
}

var worker *VisitsWorker

func StartVisitWorker(log *log.Logger, visitChan chan VisitRecord, persistDB repository.Repository) error {
	if worker != nil {
		return ErrVisitWorkerStartedBefore
	}

	worker = &VisitsWorker{
		log:       log,
		visitChan: visitChan,
		persistDB: persistDB,
	}

	if err := worker.start(); err != nil {
		return err
	}

	return nil
}

func (w VisitsWorker) start() error {
	if w.visitChan == nil {
		return ErrNilVisitChannel
	}

	go func() {
		for record := range w.visitChan {
			if err := w.persistDB.Visit(context.Background(), record.LinkId); err != nil {
				w.log.WithError(err).Errorf("failed to increase visit of %v", record.LinkId)
			}
		}
	}()

	return nil
}

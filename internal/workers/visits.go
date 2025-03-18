package workers

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/repositories/link"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type VisitRecord struct {
	LinkId    uint
	IpAddress string
}

type VisitsWorker struct {
	log       *log.Logger
	visitChan chan VisitRecord
	persistDB link.Model
}

type Params struct {
	fx.In

	Log       *log.Logger
	VisitChan chan VisitRecord
	PersistDB link.Model
}

type Result struct {
	fx.Out

	Worker *VisitsWorker
}

func NewVisitWorker(p Params) (*VisitsWorker, error) {
	worker := &VisitsWorker{
		log:       p.Log,
		visitChan: p.VisitChan,
		persistDB: p.PersistDB,
	}

	if err := worker.start(); err != nil {
		return nil, err
	}

	//result := Result{
	//	Worker: worker,
	//}
	return worker, nil
}

func (w VisitsWorker) start() error {
	w.log.Infof("************** visit worker started")
	go func() {
		for record := range w.visitChan {
			if err := w.persistDB.Visit(context.Background(), record.LinkId); err != nil {
				w.log.WithError(err).Errorf("failed to increase visit of %v", record.LinkId)
			}
		}
	}()

	return nil
}

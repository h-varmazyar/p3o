package workers

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/h-varmazyar/p3o/internal/entities"
)

type linkRepository interface {
	Visit(ctx context.Context, linkId uint) error
	ReturnById(ctx context.Context, id uint) (entities.Link, error)
}

type visitRepository interface {
	Create(ctx context.Context, visit entities.Visit) (entities.Visit, error)
}

type VisitRecord struct {
	LinkId    uint
	IpAddress string
	OS entities.OS
	Browser entities.Browser
}

type VisitsWorker struct {
	log       *log.Logger
	visitChan chan VisitRecord
	linkRepo  linkRepository
	visitRepo visitRepository
}

func NewVisitWorker(log *log.Logger, visitChan chan VisitRecord) (*VisitsWorker, error) {
	worker := &VisitsWorker{
		log:       log,
		visitChan: visitChan,
		//linkRepo:  p.LinkRepo,
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
			if err := w.linkRepo.Visit(context.Background(), record.LinkId); err != nil {
				w.log.WithError(err).Errorf("failed to increase visit of %v", record.LinkId)
			}
		}
	}()

	return nil
}

func (w VisitsWorker) visit(ctx context.Context, record VisitRecord) error {
	link, err := w.linkRepo.ReturnById(ctx, record.LinkId)
	if err != nil {
		return err
	}

	visit := entities.Visit{
		LinkId:link.ID,
		UserId: link.OwnerId,
		OS      : record.OS,
		Browser :record.Browser,
		IP      : record.IpAddress,
	}

	if _, err = w.visitRepo.Create(ctx, visit); err!=nil {
		return err
	}

	return nil
}

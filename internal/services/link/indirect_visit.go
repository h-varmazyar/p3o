package link

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/entities"
	"github.com/h-varmazyar/p3o/internal/errors"
	"github.com/oklog/ulid/v2"
	"time"
)

func (s Service) IndirectVisit(ctx context.Context, key string, id string) (domain.Link, error) {
	link, err := s.linkRepo.ReturnByKey(ctx, key)
	if err != nil {
		return domain.Link{}, err
	}

	visitId, err := ulid.Parse(id)
	if err != nil {
		return domain.Link{}, err
	}
	visit, err := s.visitRepo.ReturnByID(ctx, visitId)
	if err != nil {
		return domain.Link{}, err
	}

	if visit.LinkId != link.ID {
		return domain.Link{}, errors.ErrLinkVisitMismatch
	}

	visit.Status = entities.VisitStatusCompleted
	visit.RedirectedAt = time.Now()

	if err = s.visitRepo.Update(ctx, visit); err != nil {
		return domain.Link{}, err
	}

	return domain.Link{Url: link.RealLink}, nil
}

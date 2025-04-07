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

	if _, err = ulid.Parse(id); err != nil {
		return domain.Link{}, err
	}
	visit, err := s.visitRepo.ReturnByID(ctx, id)
	if err != nil {
		return domain.Link{}, err
	}

	if visit.LinkId != link.ID {
		return domain.Link{}, errors.ErrLinkVisitMismatch
	}

	if visit.RedirectedAt.IsZero() {
		visit.RedirectedAt = time.Now()
	}
	visit.Status = entities.VisitStatusCompleted

	if err = s.visitRepo.Update(ctx, visit); err != nil {
		return domain.Link{}, err
	}

	return domain.Link{Url: link.RealLink}, nil
}

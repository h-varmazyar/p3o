package link

import (
	"context"
	"fmt"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/entities"
	"github.com/oklog/ulid/v2"
	"time"
)

func (s Service) ReturnByKey(ctx context.Context, key string) (domain.Link, error) {
	link, err := s.linkRepo.ReturnByKey(ctx, key)
	if err != nil {
		return domain.Link{}, err
	}

	visit := entities.Visit{
		ID :ulid.Make(),
		LinkId  :link.ID,
		UserId  : link.OwnerId,
		OS      :"",
		Browser :"",
		IP      :"",
	}
	url := ""
	if link.Immediate {
		url = link.RealLink
		visit.RedirectedAt = time.Now()
		visit.Status = entities.VisitStatusCompleted
	} else {
		url = fmt.Sprintf("%v/%v/%v", s.cfg.IndirectBaseURL, link.ID, visit.ID)
		visit.Status = entities.VisitStatusAdsPending
	}

	_, err = s.visitRepo.Create(ctx, visit)
	if err != nil {
		return domain.Link{}, err
	}

	return domain.Link{
		ShortLink: fmt.Sprintf("https://p3o.ir/%v", link.Key),
		Url:       url,
	}, nil
}

package link

import (
	"context"
	"fmt"
	"time"

	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/entities"
	"github.com/h-varmazyar/p3o/internal/errors"
	"github.com/oklog/ulid/v2"
)

func (s Service) ReturnByKey(ctx context.Context, key string) (domain.Link, error) {
	cachedLink, available, err:=s.linksCache.Get(key)
	if err != nil {
		return domain.Link{}, err
	}

	if available {
		if cachedLink.SystemLink{
			return domain.Link{}, errors.ErrLinkNotFound
		}
		//todo: handle visit
		return domain.Link{
			ShortLink: fmt.Sprintf("https://p3o.ir/%v", key),
			Url:       cachedLink.URL,
		}, nil
	}

	link, err := s.linkRepo.ReturnByKey(ctx, key)
	if err != nil {
		return domain.Link{}, err
	}

	visit := entities.Visit{
		ID:      ulid.Make().String(),
		LinkId:  link.ID,
		UserId:  link.OwnerId,
		OS:      "",
		Browser: "",
		IP:      "",
	}
	
	if link.Immediate {
		url = link.RealLink
		visit.RedirectedAt = time.Now()
		visit.Status = entities.VisitStatusCompleted
	} else {
		url = fmt.Sprintf("%v/%v/%v", s.cfg.IndirectBaseURL, link.Key, visit.ID)
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

package link

import (
	"context"
	sysErrors "errors"
	"fmt"
	"github.com/h-varmazyar/p3o/internal/errors"
	"time"

	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/entities"
	"github.com/oklog/ulid/v2"
)

func (s Service) ReturnByKey(ctx context.Context, key string) (domain.Link, error) {
	link, err := s.loadLink(ctx, key)
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

	url := ""
	if link.Immediate {
		url = link.RealLink
		visit.RedirectedAt = time.Now()
		visit.Status = entities.VisitStatusCompleted
	} else {
		url = fmt.Sprintf("%v/%v/%v", s.cfg.IndirectBaseURL, key, visit.ID)
		visit.Status = entities.VisitStatusAdsPending
	}

	_, err = s.visitRepo.Create(ctx, visit)
	if err != nil {
		return domain.Link{}, err
	}

	return domain.Link{
		ShortLink: fmt.Sprintf("https://p3o.ir/%v", key),
		Url:       url,
	}, nil
}

func (s Service) loadLink(ctx context.Context, key string) (entities.Link, error) {
	cachedLink, available, err := s.linksCache.Get(key)
	if err != nil {
		return entities.Link{}, err
	}

	link := entities.Link{}

	if available {
		if cachedLink.Error != nil {
			return entities.Link{}, cachedLink.Error
		}

		link.RealLink = cachedLink.URL
		link.OwnerId = cachedLink.OwnerID
		link.ID = cachedLink.ID
	} else {
		link, err = s.linkRepo.ReturnByKey(ctx, key)
		if err != nil {
			if sysErrors.Is(err, errors.ErrLinkNotFound) {
				cachedLink.Error = errors.ErrLinkNotFound
				s.linksCache.Set(key, cachedLink)
			}
			return entities.Link{}, err
		}

		cachedLink.URL = link.RealLink
		cachedLink.OwnerID = link.OwnerId
		cachedLink.ID = link.ID

		s.linksCache.Set(key, cachedLink)
	}

	return link, nil
}

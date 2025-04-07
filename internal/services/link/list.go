package link

import (
	"context"
	"fmt"
	"github.com/h-varmazyar/p3o/internal/domain"
)

func (s Service) List(ctx context.Context, userId uint) (domain.LinkList, error) {
	links, err := s.linkRepo.List(ctx, userId)
	if err != nil {
		s.log.WithError(err).Error("user link list")
		return domain.LinkList{}, err
	}

	all := domain.LinkList{
		Links: make([]domain.Link, len(links)),
	}

	for i, link := range links {
		all.Links[i] = domain.Link{
			ShortLink: fmt.Sprintf("https://p3o.ir/%v", link.Key),
			Url:       link.RealLink,
			Visits:    uint(link.TotalVisit),
			CreatedAt: link.CreatedAt,
		}
	}

	return all, err
}

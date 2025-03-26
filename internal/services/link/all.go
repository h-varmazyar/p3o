package link

import (
	"context"
	"fmt"
	"github.com/h-varmazyar/p3o/internal/domain"
)

func (s Service) All(ctx context.Context, userId uint) (domain.All, error) {
	s.log.Info("user:", userId)
	links, err := s.linkRepo.List(ctx, userId)
	if err != nil {
		s.log.WithError(err).Error("user link list")
		return domain.All{}, err
	}

	all := domain.All{
		Links: make([]domain.Link, len(links)),
	}

	for i, link := range links {
		all.Links[i] = domain.Link{
			ID:        link.ID,
			ShortLink: fmt.Sprintf("https://p3o.ir/%v", link.Key),
			Url:       link.RealLink,
			Visits:    uint(link.TotalVisit),
			CreatedAt: link.CreatedAt,
		}
	}

	return all, err
}

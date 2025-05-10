package link

import (
	"context"
	"fmt"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/pkg/pagination"
)

func (s Service) List(ctx context.Context, userId uint, paginator pagination.Paginator) (domain.LinkList, error) {
	links, linkPagination, err := s.linkRepo.List(ctx, userId, paginator)
	if err != nil {
		s.log.WithError(err).Error("user link list")
		return domain.LinkList{}, err
	}

	all := domain.LinkList{
		Links:      make([]domain.Link, len(links)),
		Pagination: linkPagination,
	}

	for i, link := range links {
		all.Links[i] = domain.Link{
			ShortLink:   fmt.Sprintf("https://p3o.ir/%v", link.Key),
			Url:         link.RealLink,
			Status:      link.Status.ToShowableString(),
			IsImmediate: link.Immediate,
			Visits:      uint(link.TotalVisit),
			CreatedAt:   link.CreatedAt,
		}
	}

	return all, err
}

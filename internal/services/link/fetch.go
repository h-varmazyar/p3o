package link

import (
	"context"
	"fmt"
	"github.com/h-varmazyar/p3o/internal/domain"
	"github.com/h-varmazyar/p3o/internal/errors"
)

func (s Service) Fetch(ctx context.Context, userId uint, key string) (domain.LinkDetails, error) {
	link, err := s.linkRepo.ReturnByKey(ctx, key)
	if err != nil {
		return domain.LinkDetails{}, err
	}

	if link.OwnerId != userId {
		return domain.LinkDetails{}, errors.ErrLinkOwnerMismatch
	}

	resp := domain.LinkDetails{
		ShortLink:           fmt.Sprintf("https://p3o.ir/%v", link.Key),
		Url:                 link.RealLink,
		Status:              link.Status.ToShowableString(),
		IsImmediate:         link.Immediate,
		Visits:              uint(link.TotalVisit),
		CreatedAt:           link.CreatedAt,
		ProtectedByPassword: link.Password != "",
		MaxVisit:            link.MaxVisit,
	}

	if link.ExpireAt.Valid {
		resp.ExpireAt = link.ExpireAt.Time
	}
	return resp, nil
}

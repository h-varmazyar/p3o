package link

import (
	"context"
	"fmt"
	"github.com/h-varmazyar/p3o/internal/domain"
)

func (s Service) ReturnByKey(ctx context.Context, key string) (domain.Link, error) {
	link, err := s.linkRepo.ReturnByKey(ctx, key)
	if err != nil {
		return domain.Link{}, err
	}

	return domain.Link{
		ID:        link.ID,
		ShortLink: fmt.Sprintf("https://p3o.ir/%v", link.Key),
		Url:       link.RealLink,
	}, nil
}

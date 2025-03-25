package link

import (
	"context"
	"fmt"
	"github.com/h-varmazyar/p3o/internal/domain"
)

func (s Service) ReturnByKey(ctx context.Context, key string) (domain.Link, error) {
	fmt.Println("key:::::::::", key)
	link, err := s.linkRepo.ReturnByKey(ctx, key)
	if err != nil {
		return domain.Link{}, err
	}

	return domain.Link{
		ID:  link.ID,
		Key: link.Key,
		Url: link.RealLink,
	}, nil
}

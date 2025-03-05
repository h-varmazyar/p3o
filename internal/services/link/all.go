package link

import (
	"context"

	"github.com/h-varmazyar/p3o/internal/entities"
)

func (s Service)All(ctx context.Context, userId uint) ([]entities.Link, error) {
	links, err:=s.linkRepo.List(ctx, userId)
	if err!=nil{
		return nil, err
	}

	return links, err
}
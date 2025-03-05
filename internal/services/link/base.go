package link

import (
	"context"

	"github.com/h-varmazyar/p3o/internal/entities"
)

type linkRepository interface{
	ReturnByKey(ctx context.Context, key string) (entities.Link, error)
	List(ctx context.Context, userId uint) ([]entities.Link, error)
	Update(ctx context.Context, link entities.Link) error
}

type Service struct{
	linkRepo linkRepository
}
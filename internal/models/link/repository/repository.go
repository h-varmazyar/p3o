package repository

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/entities"
)

type Repository interface {
	Create(ctx context.Context, link *entities.Link) error
	ReturnByKey(ctx context.Context, key string) (*entities.Link, error)
	Visit(ctx context.Context, id uint) error
}

package link

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/entities"
)

type Model interface {
	Create(ctx context.Context, link *entities.Link) error
	TotalCounts(ctx context.Context) (int64, error)
	TotalVisits(ctx context.Context) (int64, error)
	ReturnByKey(ctx context.Context, key string) (*entities.Link, error)
	Visit(ctx context.Context, id uint) error
}

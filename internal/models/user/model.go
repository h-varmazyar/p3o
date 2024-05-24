package user

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/entities"
)

type Repository interface {
	Create(ctx context.Context, link *entities.User) error
	ReturnByMobile(ctx context.Context, mobile string) (*entities.User, error)
	ReturnByEmail(ctx context.Context, email string) (*entities.User, error)
}

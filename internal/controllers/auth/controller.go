package auth

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/domain"
)

var ()

type UserService interface {
	Login(ctx context.Context, username, password string) (domain.Tokens, error)
	Logout(ctx context.Context, username string) error
	Register(ctx context.Context, req domain.RegisterUserReq) (domain.RegisterResp, error)
	Verify(ctx context.Context, req domain.VerifyUserReq) (domain.Tokens, error)
}

type Controller struct {
	userService UserService
}

func New(userSrv UserService) Controller {
	return Controller{
		userService: userSrv,
	}
}

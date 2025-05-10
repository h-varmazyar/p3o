package user

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/domain"
)

var ()

type UserService interface {
	ChangePassword(ctx context.Context, req domain.ChangePasswordReq) error
	EditUser(ctx context.Context, req domain.EditUserReq) error
	Get(ctx context.Context, userId uint) (domain.UserInfo, error)
}

type Controller struct {
	userService UserService
}

func New(userSrv UserService) Controller {
	return Controller{
		userService: userSrv,
	}
}

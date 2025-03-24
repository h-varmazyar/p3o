package auth

import (
	"context"
	"github.com/h-varmazyar/p3o/internal/domain"
	"go.uber.org/fx"
)

var ()

type userService interface {
	Login(ctx context.Context, username, password string) (domain.Tokens, error)
	Logout(ctx context.Context, username string) error
	Register(ctx context.Context, username, password string) error
	Verify(ctx context.Context, username, verificationCode string) error
}

type Controller struct {
	userService userService
}

type Params struct {
	fx.In

	UserService userService
}

type Result struct {
	fx.Out

	Controller *Controller
}

func New(p Params) Result {
	controller := &Controller{
		userService: p.UserService,
	}
	return Result{Controller: controller}
}

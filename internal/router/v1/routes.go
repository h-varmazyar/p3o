package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/internal/controllers/auth"
	"go.uber.org/fx"
)

type Router struct {
	ginRouter      *gin.RouterGroup
	authController *auth.Controller
}

type Params struct {
	fx.In

	AuthController *auth.Controller
}

type Result struct {
	fx.Out

	Router *Router
}

func New(params Params) Result {
	router := &Router{
		authController: params.AuthController,
	}

	return Result{Router: router}
}

func (r *Router) RegisterRoutes(ginRouter *gin.RouterGroup) {
	v1Router := ginRouter.Group("/v1")

	{
		authRouter := v1Router.Group("/auth")
		authRouter.POST("/login", r.authController.Login)
		authRouter.PUT("/verify", r.authController.Verify)
		authRouter.DELETE("/logout", r.authController.Logout)
	}
}

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/internal/controllers/auth"
	"github.com/h-varmazyar/p3o/internal/controllers/link"
	"go.uber.org/fx"
)

type Router struct {
	ginRouter      *gin.RouterGroup
	authController *auth.Controller
	linkController *link.Controller
}

type Params struct {
	fx.In

	AuthController *auth.Controller
	LinkController *link.Controller
}

type Result struct {
	fx.Out

	Router *Router
}

func New(params Params) Result {
	router := &Router{
		authController: params.AuthController,
		linkController: params.LinkController,
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
	{
		linkRouter := v1Router.Group("/link")
		linkRouter.POST("/", r.linkController.Create)
	}
}

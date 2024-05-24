package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/internal/controllers/auth"
)

type Router struct {
	ginRouter      *gin.RouterGroup
	authController *auth.Controller
}

func NewV1Router() *Router {
	return &Router{}
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

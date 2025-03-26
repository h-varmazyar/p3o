package v1

import (
	"github.com/gin-gonic/gin"
	visit "github.com/h-varmazyar/p3o/api/v1/visits"
	"github.com/h-varmazyar/p3o/internal/controllers/auth"
	"github.com/h-varmazyar/p3o/internal/controllers/link"
	"github.com/h-varmazyar/p3o/internal/controllers/user"
	"github.com/h-varmazyar/p3o/internal/router/middlewares"
)

type Router struct {
	ginRouter            *gin.RouterGroup
	authController       auth.Controller
	linksController      link.Controller
	usersController      user.Controller
	visitsController     visit.Controller
	publicAuthMiddleware middlewares.PublicAuthMiddleware
}

func New(authController auth.Controller, linkController link.Controller, publicAuthMiddleware middlewares.PublicAuthMiddleware) Router {
	return Router{
		authController:       authController,
		linksController:      linkController,
		publicAuthMiddleware: publicAuthMiddleware,
	}
}

func (r *Router) RegisterRoutes(ginRouter *gin.RouterGroup) {
	v1Router := ginRouter.Group("/v1")

	{
		authRouter := v1Router.Group("/auth")
		authRouter.POST("/login", r.authController.Login)
		authRouter.PUT("/verify", r.authController.Verify)
		authRouter.DELETE("/logout", r.authController.Logout)
		authRouter.DELETE("/register", r.authController.Register)
	}
	{
		linkRouter := v1Router.Group("/links").Use(r.publicAuthMiddleware.Handle)
		linkRouter.POST("", r.linksController.Create)
		linkRouter.GET("", r.linksController.All)
		linkRouter.GET("/counts", r.linksController.Counts)
		linkRouter.GET("/visits", r.linksController.Visits)
		linkRouter.GET("/:key", r.linksController.Status)
		linkRouter.DELETE("/:key", r.linksController.Delete)
		linkRouter.PATCH("/:key/activate", r.linksController.Activate)
		linkRouter.PATCH("/:key/deactivate", r.linksController.Deactivate)
	}
	//{
	//	userRouter := v1Router.Group("/users")
	//	userRouter.GET("/info", r.usersController.GetInfo)
	//	userRouter.PATCH("/change-password", r.usersController.ChangePassword)
	//}
	//{
	//	visitsController := v1Router.Group("/visits")
	//	visitsController.GET("/history", r.visitsController.History)
	//}
	//{
	//	visitsController := v1Router.Group("/visits")
	//	visitsController.GET("/history", r.visitsController.History)
	//}
}

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/internal/controllers/auth"
	"github.com/h-varmazyar/p3o/internal/controllers/dashboard"
	"github.com/h-varmazyar/p3o/internal/controllers/link"
	"github.com/h-varmazyar/p3o/internal/controllers/user"
	"github.com/h-varmazyar/p3o/internal/router/middlewares"
)

type Router struct {
	ginRouter            *gin.RouterGroup
	authController       auth.Controller
	dashboardController  dashboard.Controller
	linksController      link.Controller
	usersController      user.Controller
	publicAuthMiddleware middlewares.PublicAuthMiddleware
}

func New(authController auth.Controller, linkController link.Controller, dashboardController dashboard.Controller, usersController user.Controller, publicAuthMiddleware middlewares.PublicAuthMiddleware) Router {
	return Router{
		authController:       authController,
		linksController:      linkController,
		dashboardController:  dashboardController,
		usersController:      usersController,
		publicAuthMiddleware: publicAuthMiddleware,
	}
}

func (r *Router) RegisterRoutes(ginRouter *gin.RouterGroup) {
	v1Router := ginRouter.Group("/v1")

	{
		authRouter := v1Router.Group("/auth")
		authRouter.POST("/login", r.authController.Login)
		authRouter.DELETE("/logout", r.authController.Logout)
		authRouter.POST("/register", r.authController.Register)
		authRouter.POST("/verify", r.authController.Verify)
	}
	{
		linkRouter := v1Router.Group("/links").Use(r.publicAuthMiddleware.Handle)
		linkRouter.POST("", r.linksController.Create)
		linkRouter.GET("", r.linksController.List)
		linkRouter.GET("/:key", r.linksController.Fetch)
		linkRouter.PATCH("/:key", r.linksController.Edit)
		linkRouter.DELETE("/:key", r.linksController.Delete)
		linkRouter.GET("/:key/visit-chart", r.linksController.VisitChart)

		//todo: check usage
		linkRouter.GET("/:key/status", r.linksController.Status)
		linkRouter.PATCH("/:key/activate", r.linksController.Activate)
		linkRouter.PATCH("/:key/deactivate", r.linksController.Deactivate)

		//public route
		v1Router.GET("/links/:key/:id", r.linksController.IndirectVisit)
	}
	{
		dashboardRouter := v1Router.Group("/dashboard").Use(r.publicAuthMiddleware.Handle)
		dashboardRouter.GET("/recent", r.dashboardController.Recent)
		dashboardRouter.GET("/visit-chart", r.dashboardController.VisitChart)
		dashboardRouter.GET("/info", r.dashboardController.Info)
	}
	{
		userRouter := v1Router.Group("/users").Use(r.publicAuthMiddleware.Handle)
		userRouter.GET("/info", r.usersController.GetInfo)
		userRouter.PATCH("/edit", r.usersController.Edit)
		userRouter.PATCH("/change-password", r.usersController.ChangePassword)
	}
	//{
	//	visitsController := v1Router.Group("/visits")
	//	visitsController.GET("/history", r.visitsController.History)
	//}
	//{
	//	visitsController := v1Router.Group("/visits")
	//	visitsController.GET("/history", r.visitsController.History)
	//}
}

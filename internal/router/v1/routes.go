package v1

import (
	"github.com/gin-gonic/gin"
	visit "github.com/h-varmazyar/p3o/api/v1/visits"
	"github.com/h-varmazyar/p3o/internal/controllers/auth"
	"github.com/h-varmazyar/p3o/internal/controllers/link"
	"github.com/h-varmazyar/p3o/internal/controllers/user"
	"go.uber.org/fx"
)

type Router struct {
	ginRouter        *gin.RouterGroup
	authController   *auth.Controller
	linksController  *link.Controller
	usersController  *user.Controller
	visitsController *visit.Controller
}

type Params struct {
	fx.In

	AuthController   *auth.Controller
	LinksController  *link.Controller
	usersController  *user.Controller
	visitsController *visit.Controller
}

type Result struct {
	fx.Out

	Router *Router
}

func New(params Params) Result {
	router := &Router{
		authController:  params.AuthController,
		linksController: params.LinksController,
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
		authRouter.DELETE("/register", r.authController.Register)
	}
	{
		linkRouter := v1Router.Group("/links")
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

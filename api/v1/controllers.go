package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/api/v1/auth"
	"github.com/h-varmazyar/p3o/api/v1/links"
	visit "github.com/h-varmazyar/p3o/api/v1/visits"
	linkService "github.com/h-varmazyar/p3o/internal/models/link/service"
	log "github.com/sirupsen/logrus"
)

type Controller struct {
	log              *log.Logger
	authController   *auth.Controller
	LinksController  *links.Controller
	VisitsController *visit.Controller
}

func NewController(log *log.Logger, linkService linkService.Service) *Controller {
	return &Controller{
		log:              log,
		authController:   auth.NewController(log),
		LinksController:  links.NewController(log, linkService),
		VisitsController: visit.NewController(log),
	}
}

func (c *Controller) RegisterRoutes(router *gin.RouterGroup) {
	v1Router := router.Group("/v1")

	c.authController.RegisterRoutes(v1Router)
	c.LinksController.RegisterRoutes(v1Router)
	c.VisitsController.RegisterRoutes(v1Router)
}

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/p3o/internal/controllers/panel/v1/link"
	linkService "github.com/h-varmazyar/p3o/internal/models/link/service"
	log "github.com/sirupsen/logrus"
)

type Controller struct {
	link *link.Controller
	log  *log.Logger
}

func NewV1(log *log.Logger, linkService linkService.Service) *Controller {
	return &Controller{
		link: link.NewController(log, linkService),
	}
}

func (c *Controller) RegisterRoutes(router *gin.RouterGroup) {
	v1Router := router.Group("/v1")
	c.link.RegisterRoutes(v1Router)
}

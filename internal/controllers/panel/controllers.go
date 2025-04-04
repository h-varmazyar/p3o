package panel

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/h-varmazyar/p3o/internal/controllers/api/links"
	linkService "github.com/h-varmazyar/p3o/internal/repositories/link/service"
	log "github.com/sirupsen/logrus"
)

type Controller struct {
	v1  *v1.Controller
	log *log.Logger
}

func NewController(log *log.Logger, linkService linkService.Service) *Controller {
	c := &Controller{
		v1: v1.NewV1(log, linkService),
	}
	return c
}

func (c *Controller) RegisterRoutes(router *gin.Engine) {
	panelRouter := router.Group("/")
	c.v1.RegisterRoutes(panelRouter)
}
